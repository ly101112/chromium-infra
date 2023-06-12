-- This query updates the total test suite runetime metric in the
-- cq_builder_metrics_day table. It uses an interval of the last 2 days so that
-- there is some redundancy if the job fails This query is meant to be almost
-- identical to the one in cq_builder_metrics_week_total_test_runtime.sql.

-- The lines below are used by the deploy tool.
--name: Populate cq_builder_metrics_day total test runtime
--schedule: every 4 hours synchronized

DECLARE start_date DATE DEFAULT DATE_SUB(CURRENT_DATE('PST8PDT'), INTERVAL 2 DAY);
-- This isn't really needed, but useful to have around when doing backfills
-- The end_date is exclusive, which is why we add a day to the current date.
DECLARE end_date DATE DEFAULT DATE_ADD(CURRENT_DATE('PST8PDT'), INTERVAL 1 DAY);
DECLARE start_ts TIMESTAMP DEFAULT TIMESTAMP(start_date, 'PST8PDT');
DECLARE end_ts TIMESTAMP DEFAULT TIMESTAMP(end_date, 'PST8PDT');

-- Merge statement
MERGE INTO
  `chrome-trooper-analytics.metrics.cq_builder_metrics_day` AS T
USING
  (
  WITH builds AS (
    SELECT
      EXTRACT(DATE FROM b.start_time AT TIME ZONE "PST8PDT") AS `date`,
      b.id,
      b.builder.builder,
      b.start_time,
      b.infra.swarming.task_id
    FROM
      `cr-buildbucket.chromium.builds` b,
      `chrome-trooper-analytics.metrics.cq_builders` cq
    WHERE
      -- As we bucket the build using start_date, we need to include any builds
      -- that were created on the previous day.
      b.create_time >= TIMESTAMP_SUB(start_ts, INTERVAL 1 DAY)
      AND b.create_time < end_ts
      AND b.builder.bucket = 'try'
      AND b.builder.project = 'chromium'
      AND b.builder.builder = cq.builder
    ),
    -- This table is needed to dedupe swarming task IDs, as sometimes there are
    -- duplicate rows in task_results_summary
    deduped_tasks AS (
      SELECT
        b.date,
        b.id AS build_id,
        t.task_id,
        any_value(b.builder) builder,
        (
          SELECT
            SPLIT(tag, ':')[offset(1)]
          FROM UNNEST(t.request.tags) tag
          WHERE STARTS_WITH(tag, 'test_suite:')
        ) AS test_suite,
        MAX(b.start_time) max_start_time,
        MAX(timestamp_diff(t.end_time, t.start_time, SECOND)) task_duration_sec,
      FROM
        builds b,
        `chromium-swarm.swarming.task_results_summary` t
      WHERE
        t.end_time >= TIMESTAMP_SUB(start_ts, INTERVAL 1 DAY)
        AND t.end_time < TIMESTAMP_ADD(end_ts, INTERVAL 1 DAY)
        AND b.task_id = t.request.parent_task_id
        -- Exclude compilator tasks
        AND t.request.name not like '%-compilator-%'
      GROUP BY
        b.date, b.id, t.task_id, test_suite
    ),
    -- Combine the swarming task times by the test suite for each build
    test_tasks AS (
    SELECT
      build_id,
      test_suite,
      ANY_VALUE(b.date) date,
      ANY_VALUE(b.builder) builder,
      MAX(b.max_start_time) builder_start_time,
      SUM(task_duration_sec) total_suite_runtime
    FROM
      deduped_tasks b
    GROUP BY build_id, test_suite
    ),
    -- Calculate the quantiles for each builder/suite
    total_swarming_times AS (
    SELECT
      t.date,
      t.builder,
      MAX(t.builder_start_time) max_builder_start_time,
      t.test_suite,
      APPROX_QUANTILES(total_suite_runtime, 100) total_suite_time_quantiles,
      count(build_id) num_runs,
    FROM test_tasks t
    WHERE t.test_suite is not null
    GROUP BY t.date, t.builder, t.test_suite
    HAVING num_runs > 10
    )
  SELECT
    date,
    'P50 Total Suite Runtime' AS metric,
    builder,
    MAX(max_builder_start_time) max_builder_start_time,
    ARRAY_AGG(
      STRUCT(test_suite AS label, CAST(total_suite_time_quantiles[OFFSET(50)] AS NUMERIC) AS value)
      ORDER BY test_suite
    ) AS value_agg,
  FROM total_swarming_times
  GROUP BY date, builder
  UNION ALL
  SELECT
    date,
    'P90 Total Suite Runtime' AS metric,
    builder,
    MAX(max_builder_start_time) AS max_builder_start_time,
    ARRAY_AGG(
      STRUCT(test_suite AS label, CAST(total_suite_time_quantiles[OFFSET(90)] AS NUMERIC) AS value)
      ORDER BY test_suite
    ) AS value_agg,
  FROM total_swarming_times
  GROUP BY date, builder
  ) S
ON
  T.date = S.date AND T.metric = S.metric AND T.builder = S.builder
  WHEN MATCHED AND T.checkpoint != string(S.max_builder_start_time, "UTC") THEN
    UPDATE SET value_agg = S.value_agg, checkpoint = string(S.max_builder_start_time, "UTC"), last_updated = current_timestamp()
  WHEN NOT MATCHED THEN
    INSERT (date, metric, builder, value_agg, last_updated, checkpoint)
    VALUES (date, metric, builder, value_agg, current_timestamp(), string(max_builder_start_time, "UTC"));