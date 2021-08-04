// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/protojson"
)

// writeTestFiles writes TestFile protobuf messages to w in JSON Lines format.
func writeTestFiles(ctx context.Context, bqClient *bigquery.Client, w io.Writer) error {
	// Grab all tests in the past 1 week.
	// Use ci_test_results table (not CQ) because it is smaller and it does not
	// include test file names that never made it to the repo.
	q := bqClient.Query(`
		SELECT
			tr.test_metadata.location.file_name as Path,
			ARRAY_AGG(DISTINCT tr.test_metadata.name) as TestNames,
			# Extract the test target. Examples:
			# - "ninja://chrome/test:browser_tests/foo/bar" -> "browser_tests"
			# - "ninja://:blink_web_tests/foo/bar" -> "blink_web_tests"
			ARRAY_AGG(DISTINCT REGEXP_EXTRACT(test_id, "ninja://[^:]*:([^/]+)/") IGNORE NULLS) TestTargets,
		FROM chrome-luci-data.chromium.ci_test_results tr
		WHERE partition_time > TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 1 DAY)
			AND tr.test_metadata.location.file_name != ''
			AND tr.test_metadata.name != ''
		GROUP BY Path
	`)
	it, err := q.Read(ctx)
	if err != nil {
		return err
	}
	return writeTestFilesFrom(ctx, w, it.Next)
}

// writeTestFilesFrom writes test files returned by source to w.
// If source returns iterator.Done, writeTestFilesFrom exits.
func writeTestFilesFrom(ctx context.Context, w io.Writer, source func(dest interface{}) error) error {
	testFile := &TestFile{}
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		// Read the next row.
		switch err := source(testFile); {
		case err == iterator.Done:
			return ctx.Err()
		case err != nil:
			return err

		case mustAlwaysRunTest(testFile.Path):
			continue
		}

		jsonBytes, err := protojson.Marshal(testFile)
		if err != nil {
			return err
		}
		if _, err := w.Write(jsonBytes); err != nil {
			return err
		}
		if _, err := fmt.Fprintln(w); err != nil {
			return err
		}
	}
}

// readTestFiles reads test files written by writeTestFilesFrom().
func readTestFiles(r io.Reader, callback func(*TestFile) error) error {
	scan := bufio.NewScanner(r)
	scan.Buffer(nil, 1e7) // 10 MB.
	for scan.Scan() {
		testFile := &TestFile{}
		if err := protojson.Unmarshal(scan.Bytes(), testFile); err != nil {
			return err
		}
		if err := callback(testFile); err != nil {
			return err
		}
	}
	return scan.Err()
}
