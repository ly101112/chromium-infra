// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package testmetrics

import (
	"context"
	"os"
	"strconv"
	"strings"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"go.chromium.org/luci/common/errors"
	"google.golang.org/api/iterator"

	"infra/appengine/chrome-test-health/api"
)

// Base queries to build from
const (
	testSingleDayQuery string = `
SELECT
	m.date,
	m.test_id,
	ANY_VALUE(m.test_name) AS test_name,
	ANY_VALUE(m.file_name) AS file_name,
	{metricAggregations},
	ARRAY_AGG(STRUCT(
		builder AS builder,
		project AS project,
		bucket AS bucket,
		test_suite AS test_suite,
		{metricNames}
		)
	) AS variants
FROM
	{table} AS m
WHERE
	DATE(date) IN UNNEST(@dates)
	AND component IN UNNEST(@components){fileNameClause}{filterClause}
GROUP BY date, test_id
ORDER BY {sortMetric} {sortDirection}
LIMIT @page_size OFFSET @page_offset`

	testMultiDayQuery string = `
WITH tests AS (
	SELECT
		m.date,
		m.test_id,
		ANY_VALUE(m.test_name) AS test_name,
		ANY_VALUE(m.file_name) AS file_name,
		{metricAggregations},
		ARRAY_AGG(STRUCT(
			builder AS builder,
			bucket AS bucket,
			project AS project,
			test_suite AS test_suite,
			{metricNames}
			)
		) AS variants
	FROM
		{table} AS m
	WHERE
		DATE(date) IN UNNEST(@dates)
		AND component IN UNNEST(@components){fileNameClause}{filterClause}
	GROUP BY m.date, m.test_id
), sorted_day AS (
	SELECT
		test_id,
		{sortMetric} AS rank
	FROM tests
	WHERE date = @sort_date
	ORDER BY {sortMetric} {sortDirection}
	LIMIT @page_size OFFSET @page_offset
)
SELECT t.*
FROM sorted_day AS s LEFT JOIN tests AS t USING(test_id)
ORDER BY rank {sortDirection}`

	unfilteredDirectorySingleDayQuery string = `
SELECT
	date,
	node_name,
	ARRAY_REVERSE(SPLIT(node_name, '/'))[SAFE_OFFSET(0)] AS display_name,
	ANY_VALUE(is_file) AS is_file,
	{metricAggregations},
FROM {fileTable}, UNNEST(@parents) AS parent
WHERE
	STARTS_WITH(node_name, parent || "/")
	-- The child folders and files can't have a / after the parent's name
	AND REGEXP_CONTAINS(SUBSTR(node_name, LENGTH(parent) + 2), "^[^/]*$")
	AND DATE(date) IN UNNEST(@dates)
	AND component IN UNNEST(@components)
GROUP BY date, node_name
ORDER BY is_file, {sortMetric} {sortDirection}`

	unfilteredDirectoryMultiDayQuery string = `
WITH nodes AS(
	SELECT
		date,
		node_name,
		ARRAY_REVERSE(SPLIT(node_name, '/'))[SAFE_OFFSET(0)] AS display_name,
		ANY_VALUE(is_file) AS is_file,
		{metricAggregations},
	FROM {fileTable}, UNNEST(@parents) AS parent
	WHERE
		STARTS_WITH(node_name, parent || "/")
		-- The child folders and files can't have a / after the parent's name
		AND REGEXP_CONTAINS(SUBSTR(node_name, LENGTH(parent) + 2), "^[^/]*$")
		AND DATE(date) IN UNNEST(@dates)
		AND component IN UNNEST(@components)
	GROUP BY date, node_name
), sorted_day AS (
	SELECT
		node_name,
		{sortMetric} AS rank
	FROM nodes
	WHERE date = @sort_date
)
SELECT t.*
FROM nodes AS t LEFT JOIN sorted_day AS s USING(node_name)
ORDER BY is_file, s.rank {sortDirection}`

	filteredDirectorySingleDayQuery string = `
WITH
test_summaries AS (
	SELECT
		file_name AS node_name,
		date,
		--metrics
		{metricAggregations},
	FROM {testTable}
	WHERE
		date IN UNNEST(@dates)
		AND file_name IS NOT NULL
		AND component IN UNNEST(@components)
		-- Apply the requested filter{filterClause}
	GROUP BY file_name, date, test_id
)
SELECT
	f.date,
	f.node_name,
	ARRAY_REVERSE(SPLIT(f.node_name, '/'))[SAFE_OFFSET(0)] AS display_name,
	ANY_VALUE(is_file) AS is_file,
	-- metrics
	{fileAggMetricTerms},
FROM {fileTable} AS f, UNNEST(@parents) AS parent
JOIN test_summaries t ON
	f.date = t.date
	AND STARTS_WITH(t.node_name, f.node_name)
WHERE
	STARTS_WITH(f.node_name, parent || "/")
	-- The child folders and files can't have a / after the parent's name
	AND REGEXP_CONTAINS(SUBSTR(f.node_name, LENGTH(parent) + 2), "^[^/]*$")
	AND DATE(f.date) IN UNNEST(@dates)
	AND component IN UNNEST(@components)
GROUP BY date, node_name
ORDER BY is_file, {sortMetric} {sortDirection}`

	filteredDirectoryMultiDayQuery string = `
WITH
test_summaries AS (
	SELECT
		file_name AS node_name,
		date,
		--metrics
		{metricAggregations},
	FROM {testTable}
	WHERE
		date IN UNNEST(@dates)
		AND file_name IS NOT NULL
		AND component IN UNNEST(@components)
		-- Apply the requested filter{filterClause}
	GROUP BY file_name, date, test_id
), node_summaries AS (
	SELECT
		f.date,
		f.node_name,
		ARRAY_REVERSE(SPLIT(f.node_name, '/'))[SAFE_OFFSET(0)] AS display_name,
		ANY_VALUE(is_file) AS is_file,
		-- metrics
		{fileAggMetricTerms},
	FROM {fileTable} AS f, UNNEST(@parents) AS parent
	JOIN test_summaries t ON
		f.date = t.date
		AND STARTS_WITH(t.node_name, f.node_name)
	WHERE
		STARTS_WITH(f.node_name, parent || "/")
		-- The child folders and files can't have a / after the parent's name
		AND REGEXP_CONTAINS(SUBSTR(f.node_name, LENGTH(parent) + 2), "^[^/]*$")
		AND DATE(f.date) IN UNNEST(@dates)
		AND component IN UNNEST(@components)
	GROUP BY date, node_name
), sorted_day AS (
	SELECT
		node_name,
		{sortMetric} AS rank
	FROM node_summaries
	WHERE date = @sort_date
)

SELECT node_summaries.*
FROM node_summaries LEFT JOIN sorted_day USING(node_name)
ORDER BY is_file, rank {sortDirection}`
)

var (
	// Period to the table in the dataset to use
	periodToTestMetricTable = map[api.Period]string{
		api.Period_DAY:  "daily_test_metrics",
		api.Period_WEEK: "weekly_test_metrics",
	}
	// Period to table in the dataset to use
	periodToFileMetricTable = map[api.Period]string{
		api.Period_DAY:  "daily_file_metrics",
		api.Period_WEEK: "weekly_file_metrics",
	}
	// Lookup table for converting sort type to it's sql column name
	sortTypeSqlLookup = map[api.SortType]string{
		api.SortType_SORT_NUM_RUNS:      "num_runs",
		api.SortType_SORT_NUM_FAILURES:  "num_failures",
		api.SortType_SORT_AVG_RUNTIME:   "avg_runtime",
		api.SortType_SORT_TOTAL_RUNTIME: "total_runtime",
		api.SortType_SORT_AVG_CORES:     "avg_cores",
		api.SortType_SORT_P50_RUNTIME:   "p50_runtime",
		api.SortType_SORT_P90_RUNTIME:   "p90_runtime",
	}
	// Metrics that aren't summed but averaged when aggregated
	weightedAverageMetrics = map[api.MetricType]struct{}{
		api.MetricType_AVG_RUNTIME: {},
		api.MetricType_P50_RUNTIME: {},
		api.MetricType_P90_RUNTIME: {},
	}
	// Queries run in order to update the db
	updateQueries = []string{
		"sql/update_raw_metrics.sql",
		"sql/update_daily_test_metrics.sql",
		"sql/update_weekly_test_metrics.sql",
		"sql/update_daily_file_metrics.sql",
		"sql/update_weekly_file_metrics.sql",
		"sql/update_average_cores.sql",
	}
)

// Client is used to fetch metrics from a given data source.
type Client struct {
	BqClient      *bigquery.Client
	ProjectId     string
	DataSet       string
	updateQueries []string
}

func bqToDateArray(dates []string) ([]civil.Date, error) {
	ret := make([]civil.Date, len(dates))
	for i, date := range dates {
		d, err := civil.ParseDate(date)
		if err != nil {
			return nil, err
		}
		ret[i] = d
	}
	return ret, nil
}

func parseDatasetQuery(r *strings.Replacer, fileName string) (string, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return r.Replace(string(bytes)), nil
}

// Initializes the testmetric client
func (c *Client) Init() error {
	if c.ProjectId == "" {
		c.ProjectId = "chrome-resources-staging"
	}
	if c.DataSet == "" {
		c.DataSet = "test_results"
	}

	r := strings.NewReplacer("{project}", c.ProjectId, "{dataset}", c.DataSet)

	for _, filename := range updateQueries {
		query, err := parseDatasetQuery(r, filename)
		if err != nil {
			return err
		}
		c.updateQueries = append(c.updateQueries, query)
	}
	return nil
}

// Lists the available monorail components
func (c *Client) ListComponents(ctx context.Context, req *api.ListComponentsRequest) (*api.ListComponentsResponse, error) {
	query := "SELECT DISTINCT component FROM chrome-metadata.chrome.monorail_component_owners ORDER BY component"
	q := c.BqClient.Query(query)

	job, err := q.Run(ctx)
	if err != nil {
		return nil, err
	}
	it, err := job.Read(ctx)
	if err != nil {
		return nil, err
	}
	response := &api.ListComponentsResponse{
		// Add the "Unknown" component for tests that are detached from their
		// component
		Components: []string{"Unknown"},
	}
	type row struct {
		Component string `bigquery:"component"`
	}
	for {
		var rowVals row
		err = it.Next(&rowVals)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.Annotate(err, "obtain next component row").Err()
		}
		response.Components = append(response.Components, rowVals.Component)
	}
	return response, nil
}

// Fetches requested metrics for the provided days and filters
func (c *Client) FetchMetrics(ctx context.Context, req *api.FetchTestMetricsRequest) (*api.FetchTestMetricsResponse, error) {
	q, err := c.createFetchMetricsQuery(req)
	if err != nil {
		return nil, errors.Annotate(err, "failed to parse the request into a query").Err()
	}

	job, err := q.Run(ctx)
	if err != nil {
		return nil, errors.Annotate(err, "failed to run start the query").Err()
	}
	it, err := job.Read(ctx)
	if err != nil {
		return nil, errors.Annotate(err, "the query failed to complete").Err()
	}

	return c.readFetchTestMetricsResponse(it, req)
}

func (c *Client) createFetchMetricsQuery(req *api.FetchTestMetricsRequest) (*bigquery.Query, error) {
	dates, err := bqToDateArray(req.GetDates())
	if err != nil {
		return nil, err
	}

	metricNames := make([]string, len(req.Metrics))
	for i := 0; i < len(req.Metrics); i++ {
		metricNames[i] = MetricSqlName(req.Metrics[i])
	}

	// Terms for converting the rolling up the variants
	metricAggregations := make([]string, len(req.Metrics))
	for i := 0; i < len(req.Metrics); i++ {
		name := MetricSqlName(req.Metrics[i])
		if _, ok := weightedAverageMetrics[req.Metrics[i]]; ok {
			metricAggregations[i] = `SUM(` + name + ` * num_runs) / SUM(num_runs) AS ` + name
		} else {
			metricAggregations[i] = `SUM(` + name + `) AS ` + name
		}
	}

	table, ok := periodToTestMetricTable[req.Period]
	if !ok {
		return nil, errors.Reason("Received unsupported period request: '%s'", req.Period).Err()
	}

	sortMetric := "test_id"
	// A default value of 0 maps to the name which for test based fetches is
	// test_id. Other values map to metrics in both file and directory tables
	if req.Sort != nil && req.Sort.Metric != api.SortType_SORT_NAME {
		sortMetric = sortTypeSqlLookup[req.Sort.Metric]
	}
	sortDirection := "ASC"
	if req.Sort != nil && !req.Sort.Ascending {
		sortDirection = "DESC"
	}
	sortDate := req.Dates[0]
	if req.Sort != nil && req.Sort.SortDate != "" {
		sortDate = req.Sort.SortDate
	}

	fileNameClause := ""
	if len(req.FileNames) > 0 {
		fileNameClause = `
	AND file_name IN UNNEST(@file_names)`
	}

	filterClause := ""
	var filterParameters []bigquery.QueryParameter
	if req.Filter != "" {
		for i, filter := range strings.Split(req.Filter, " ") {
			filterClause += `
	AND REGEXP_CONTAINS(CONCAT(test_name, ' ', file_name, ' ', bucket, '/', builder, ' ', test_suite), @filter` + strconv.Itoa(i) + `)`
			filterParameters = append(filterParameters, bigquery.QueryParameter{
				Name:  "filter" + strconv.Itoa(i),
				Value: filter,
			})
		}
	}

	replacements := []string{
		"{metricAggregations}", strings.Join(metricAggregations, ",\n\t"),
		"{metricNames}", strings.Join(metricNames, ",\n\t\t"),
		"{table}", c.ProjectId + `.` + c.DataSet + `.` + table,
		"{filterClause}", filterClause,
		"{fileNameClause}", fileNameClause,
		"{sortMetric}", sortMetric,
		"{sortDirection}", sortDirection,
	}

	r := strings.NewReplacer(replacements...)
	// TODO(sshrimp): this query construction is pretty hard to read and should be refactored
	var query string
	if len(req.Dates) == 1 {
		query = testSingleDayQuery
	} else {
		query = testMultiDayQuery
	}

	q := c.BqClient.Query(r.Replace(query))

	q.Parameters = []bigquery.QueryParameter{
		{Name: "dates", Value: dates},
		{Name: "page_size", Value: req.PageSize + 1},
		{Name: "page_offset", Value: req.PageOffset},
		{Name: "components", Value: req.Components},
		{Name: "file_names", Value: req.FileNames},
		{Name: "sort_date", Value: sortDate},
	}
	q.Parameters = append(q.Parameters, filterParameters...)
	return q, nil
}

func (*Client) readFetchTestMetricsResponse(it *bigquery.RowIterator, req *api.FetchTestMetricsRequest) (*api.FetchTestMetricsResponse, error) {
	// maps for quick lookup of the test id
	testIdToTestDateMetricData := make(map[string]*api.TestDateMetricData)
	variantHashToTestDateMetricData := make(map[string]map[string]*api.TestVariantData)

	response := &api.FetchTestMetricsResponse{
		LastPage: int64(it.TotalRows) != req.PageSize+1,
	}
	for i := int64(0); i < req.PageSize; i++ {
		var rowVals rowLoader
		err := it.Next(&rowVals)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.Annotate(err, "obtain next test summary row").Err()
		}
		testId := rowVals.String("test_id")
		testIdData, ok := testIdToTestDateMetricData[testId]
		if !ok {
			testIdData = &api.TestDateMetricData{
				TestId:   testId,
				TestName: rowVals.NullString("test_name").StringVal,
				FileName: rowVals.NullString("file_name").StringVal,
				Metrics:  make(map[string]*api.TestMetricsArray),
			}
			testIdToTestDateMetricData[testId] = testIdData
			variantHashToTestDateMetricData[testId] = make(map[string]*api.TestVariantData)
			// Add to the actual struct being returned
			response.Tests = append(response.Tests, testIdData)
		}

		// Each row is the summary of a test id for a day
		date := rowVals.Date("date").String()

		// Handle the test's rollup metrics
		testIdData.Metrics[date] = &api.TestMetricsArray{
			Data: rowVals.Metrics(req.Metrics),
		}

		// Handle the indvidiual variant metrics
		repeated := true
		val, err := rowVals.valueWithType("variants", bigquery.RecordFieldType, repeated)
		if err != nil {
			return nil, err
		}

		i, _ := rowVals.fieldIndex("variants")
		rowSchema := rowVals.schema[i].Schema

		variantRows := val.([]bigquery.Value)
		for _, variantRow := range variantRows {
			var variantRowVals rowLoader
			if err := variantRowVals.Load(variantRow.([]bigquery.Value), rowSchema); err != nil {
				return nil, err
			}

			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, errors.Annotate(err, "obtain next variant summary row").Err()
			}

			builder := variantRowVals.NullString("builder").StringVal
			project := variantRowVals.NullString("project").StringVal
			bucket := variantRowVals.NullString("bucket").StringVal
			suite := variantRowVals.NullString("test_suite").StringVal
			builderSuite := builder + ":" + project + ":" + bucket + ":" + suite
			builderSuiteData, ok := variantHashToTestDateMetricData[testId][builderSuite]
			if !ok {
				fields := ""
				for _, field := range rowSchema {
					fields += field.Name
				}
				builderSuiteData = &api.TestVariantData{
					Builder: builder,
					Project: project,
					Bucket:  bucket,
					Suite:   suite,
					Metrics: make(map[string]*api.TestMetricsArray),
				}
				variantHashToTestDateMetricData[testId][builderSuite] = builderSuiteData
				testIdData.Variants = append(testIdData.Variants, builderSuiteData)
			}
			builderSuiteData.Metrics[date] = &api.TestMetricsArray{
				Data: variantRowVals.Metrics(req.Metrics),
			}

			if err := variantRowVals.Error(); err != nil {
				return nil, err
			}
		}
		if err := rowVals.Error(); err != nil {
			return nil, err
		}
	}
	return response, nil
}

func (c *Client) createDirectoryQuery(req *api.FetchDirectoryMetricsRequest) (*bigquery.Query, error) {
	dates, err := bqToDateArray(req.GetDates())
	if err != nil {
		return nil, err
	}

	// Terms to aggregate the metric names between files. This is a sum even
	// for averaged runtimes
	fileAggMetricTerms := make([]string, len(req.Metrics))
	for i := 0; i < len(req.Metrics); i++ {
		fileAggMetricTerms[i] = `SUM(t.` + MetricSqlName(req.Metrics[i]) + `) AS ` + MetricSqlName(req.Metrics[i])
	}

	// Terms for converting the rolling up the variants
	metricAggregations := make([]string, len(req.Metrics))
	for i := 0; i < len(req.Metrics); i++ {
		name := MetricSqlName(req.Metrics[i])
		if _, ok := weightedAverageMetrics[req.Metrics[i]]; ok {
			metricAggregations[i] = `SUM(` + name + ` * num_runs) / SUM(num_runs) AS ` + name
		} else {
			metricAggregations[i] = `SUM(` + name + `) AS ` + name
		}
	}

	fileTable, ok := periodToFileMetricTable[req.Period]
	if !ok {
		return nil, errors.Reason("Received unsupported period request: '%s'", req.Period).Err()
	}
	testTable, ok := periodToTestMetricTable[req.Period]
	if !ok {
		return nil, errors.Reason("Received unsupported period request: '%s'", req.Period).Err()
	}

	sortMetric := "node_name"
	// A default value of 0 maps to the name which for file based fetches is
	// node_name. Other values map to metrics in both file and directory tables
	if req.Sort != nil && req.Sort.Metric != api.SortType_SORT_NAME {
		sortMetric = sortTypeSqlLookup[req.Sort.Metric]
	}
	sortDirection := "ASC"
	if req.Sort != nil && !req.Sort.Ascending {
		sortDirection = "DESC"
	}
	sortDate := req.Dates[0]
	if req.Sort != nil && req.Sort.SortDate != "" {
		sortDate = req.Sort.SortDate
	}

	filterClause := ""
	var filterParameters []bigquery.QueryParameter
	if req.Filter != "" {
		for i, filter := range strings.Split(req.Filter, " ") {
			filterClause += `
		AND REGEXP_CONTAINS(CONCAT(test_name, ' ', file_name, ' ', bucket, '/', builder, ' ', test_suite), @filter` + strconv.Itoa(i) + `)`
			filterParameters = append(filterParameters, bigquery.QueryParameter{
				Name:  "filter" + strconv.Itoa(i),
				Value: filter,
			})
		}
	}

	replacements := []string{
		"{metricAggregations}", strings.Join(metricAggregations, ",\n\t"),
		"{testTable}", c.ProjectId + `.` + c.DataSet + `.` + testTable,
		"{filterClause}", filterClause,
		"{fileAggMetricTerms}", strings.Join(fileAggMetricTerms, ",\n\t"),
		"{fileTable}", c.ProjectId + `.` + c.DataSet + `.` + fileTable,
		"{sortMetric}", sortMetric,
		"{sortDirection}", sortDirection,
	}
	r := strings.NewReplacer(replacements...)

	var query string
	if req.Filter == "" {
		if len(req.Dates) == 1 {
			query = unfilteredDirectorySingleDayQuery
		} else {
			query = unfilteredDirectoryMultiDayQuery
		}
	} else {
		if len(req.Dates) == 1 {
			query = filteredDirectorySingleDayQuery
		} else {
			query = filteredDirectoryMultiDayQuery
		}
	}

	q := c.BqClient.Query(r.Replace(query))

	q.Parameters = []bigquery.QueryParameter{
		{Name: "dates", Value: dates},
		{Name: "components", Value: req.Components},
		{Name: "parents", Value: req.ParentIds},
		{Name: "sort_date", Value: sortDate},
	}
	q.Parameters = append(q.Parameters, filterParameters...)
	return q, nil
}

func (*Client) readFetchDirectoryMetricsResponse(it *bigquery.RowIterator, req *api.FetchDirectoryMetricsRequest) (*api.FetchDirectoryMetricsResponse, error) {

	// maps for quick lookup of the node id
	filenameToTestDateMetricData := make(map[string]*api.DirectoryNode)

	response := &api.FetchDirectoryMetricsResponse{}
	for {
		var rowVals rowLoader
		err := it.Next(&rowVals)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.Annotate(err, "obtain next test summary row").Err()
		}
		nodeName := rowVals.String("node_name")
		dirNode, ok := filenameToTestDateMetricData[nodeName]
		if !ok {
			nodeType := api.DirectoryNodeType_DIRECTORY
			if rowVals.Bool("is_file") {
				nodeType = api.DirectoryNodeType_FILENAME
			}
			dirNode = &api.DirectoryNode{
				Id:      nodeName,
				Metrics: make(map[string]*api.TestMetricsArray),
				Name:    rowVals.String("display_name"),
				Type:    nodeType,
			}
			filenameToTestDateMetricData[nodeName] = dirNode
			// Add to the actual struct being returned
			response.Nodes = append(response.Nodes, dirNode)
		}

		// Each row is the summary of a node for a day
		date := rowVals.Date("date").String()

		// Handle the metric rollup metrics
		dirNode.Metrics[date] = &api.TestMetricsArray{
			Data: rowVals.Metrics(req.Metrics),
		}

		if err := rowVals.Error(); err != nil {
			return nil, err
		}
	}
	return response, nil
}

// Fetches requested metrics for the provided days and filters for a given
// directory node. A directory node represents a directory or file and the
// metrics are the combined metrics of the tests in these locations
func (c *Client) FetchDirectoryMetrics(ctx context.Context, req *api.FetchDirectoryMetricsRequest) (*api.FetchDirectoryMetricsResponse, error) {
	q, err := c.createDirectoryQuery(req)
	if err != nil {
		return nil, err
	}

	job, err := q.Run(ctx)
	if err != nil {
		return nil, err
	}
	it, err := job.Read(ctx)
	if err != nil {
		return nil, err
	}

	return c.readFetchDirectoryMetricsResponse(it, req)
}

// Updates the summary tables between the days in the provided
// UpdateMetricsTableRequest. All rollups (e.g. weekly/monthly) will be updated
// as well. The dates are inclusive
func (c *Client) UpdateSummary(ctx context.Context, fromDate civil.Date, toDate civil.Date) error {
	for _, query := range c.updateQueries {
		err := c.runUpdateSummary(ctx, fromDate, toDate, query)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) runUpdateSummary(ctx context.Context, fromDate civil.Date, toDate civil.Date, query string) error {
	q := c.BqClient.Query(query)

	q.Parameters = []bigquery.QueryParameter{
		{Name: "from_date", Value: fromDate},
		{Name: "to_date", Value: toDate},
	}

	job, err := q.Run(ctx)
	if err != nil {
		return errors.Annotate(err, "failed to start the job").Err()
	}

	job_status, err := job.Wait(ctx)
	if err != nil {
		return errors.Annotate(err, "failed to finish the query").Err()
	}
	return job_status.Err()
}