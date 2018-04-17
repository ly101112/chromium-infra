# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import datetime
import json
import logging
import mock

from google.appengine.api import taskqueue

from dto import swarming_task_error
from infra_api_clients.swarming import swarming_util
from libs import analysis_status
from libs.test_results import test_results_util
from model.flake.flake_swarming_task import FlakeSwarmingTask
from model.flake.master_flake_analysis import MasterFlakeAnalysis
from model.wf_swarming_task import WfSwarmingTask
from services import swarmed_test_util
from waterfall.process_base_swarming_task_result_pipeline import (
    ProcessBaseSwarmingTaskResultPipeline)
from waterfall.process_flake_swarming_task_result_pipeline import (
    ProcessFlakeSwarmingTaskResultPipeline)
from waterfall.process_swarming_task_result_pipeline import (
    ProcessSwarmingTaskResultPipeline)
from waterfall.test import wf_testcase
from waterfall.trigger_base_swarming_task_pipeline import NO_TASK
from waterfall.trigger_base_swarming_task_pipeline import NO_TASK_EXCEPTION

_ISOLATED_SERVER = 'https://isolateserver.appspot.com'
_ISOLATED_STORAGE_URL = 'isolateserver.storage.googleapis.com'
_SAMPLE_FAILURE_LOG = {
    'all_tests': [
        'TestSuite1.test1',
        'TestSuite1.test2',
        'TestSuite1.test3',
    ],
    'per_iteration_data': [{
        'TestSuite1.test1': [{
            'status': 'SUCCESS',
            'other_info': 'N/A'
        }],
        'TestSuite1.test2': [{
            'status': 'FAILURE',
            'other_info': 'N/A'
        }, {
            'status': 'FAILURE',
            'other_info': 'N/A'
        }, {
            'status': 'SUCCESS',
            'other_info': 'N/A'
        }],
        'TestSuite1.test3': [{
            'status': 'FAILURE',
            'other_info': 'N/A'
        }, {
            'status': 'FAILURE',
            'other_info': 'N/A'
        }, {
            'status': 'FAILURE',
            'other_info': 'N/A'
        }]
    }, {
        'TestSuite1.test1': [{
            'status': 'SUCCESS',
            'other_info': 'N/A'
        }],
        'TestSuite1.test2': [{
            'status': 'SUCCESS',
            'other_info': 'N/A'
        }],
        'TestSuite1.test3': [{
            'status': 'FAILURE',
            'other_info': 'N/A'
        }, {
            'status': 'FAILURE',
            'other_info': 'N/A'
        }, {
            'status': 'FAILURE',
            'other_info': 'N/A'
        }]
    }]
}

_SWARMING_TASK_RESULTS = {
    'task_id1': {
        'state': 'COMPLETED',
        'tags': ['priority:25', 'ref_name:abc_tests'],
        'outputs_ref': {
            'isolatedserver': _ISOLATED_SERVER,
            'namespace': 'default-gzip',
            'isolated': 'shard1_isolated'
        },
        'created_ts': '2016-02-10T18:32:06.538220',
        'started_ts': '2016-02-10T18:32:09.090550',
        'completed_ts': '2016-02-10T18:33:09'
    },
    'task_id2': {
        'state': 'TIMED_OUT',
        'outputs_ref': None
    },
    'task_id3': {
        'state': 'COMPLETED',
    },
    'task_id5': {
        'state': 'COMPLETED',
        'outputs_ref': {
            'isolatedserver': _ISOLATED_SERVER,
            'namespace': 'default-gzip',
            'isolated': 'shard5_isolated'
        },
    },
    'task_id6': {
        'state': 'BOT_DIED',
        'outputs_ref': {
            'isolatedserver': _ISOLATED_SERVER,
            'namespace': 'default-gzip',
            'isolated': 'shard5_isolated'
        },
    },
}

_EXPECTED_TESTS_STATUS = {
    'TestSuite1.test1': {
        'total_run': 2,
        'SUCCESS': 2
    },
    'TestSuite1.test2': {
        'total_run': 4,
        'SUCCESS': 2,
        'FAILURE': 2
    },
    'TestSuite1.test3': {
        'total_run': 6,
        'FAILURE': 6
    },
}

_EXPECTED_CLASSIFIED_TESTS = {
    'flaky_tests': ['TestSuite1.test2', 'TestSuite1.test1'],
    'reliable_tests': ['TestSuite1.test3']
}


class ProcessBaseSwarmingTaskResultPipelineTest(wf_testcase.WaterfallTestCase):

  def _MockedGetSwarmingTaskResultById(self, _host, task_id, _):
    return _SWARMING_TASK_RESULTS[task_id], None

  def setUp(self):
    super(ProcessBaseSwarmingTaskResultPipelineTest, self).setUp()
    self.pipeline = ProcessBaseSwarmingTaskResultPipeline()
    self.master_name = 'm'
    self.builder_name = 'b'
    self.build_number = 121
    self.step_name = 'abc_tests on platform'
    self.test_name = 'test'
    self.mock(swarming_util, 'GetSwarmingTaskResultById',
              self._MockedGetSwarmingTaskResultById)

  def testConvertDateTime(self):
    fmt = '%Y-%m-%dT%H:%M:%S.%f'
    time_string = '2016-02-10T18:32:06.538220'
    test_time = self.pipeline._ConvertDateTime(time_string)
    time = datetime.datetime.strptime(time_string, fmt)
    self.assertEqual(test_time, time)

  def testConvertDateTimeNone(self):
    time_string = ''
    test_time = self.pipeline._ConvertDateTime(time_string)
    self.assertIsNone(test_time)

  def testConvertDateTimefailure(self):
    with self.assertRaises(ValueError):
      self.pipeline._ConvertDateTime('abc')

  def testCheckTestsRunStatusesNoOutPutJson(self):
    self.assertEqual(
        {},
        ProcessBaseSwarmingTaskResultPipeline._CheckTestsRunStatuses(
            self.pipeline, None, ()))

  def testCheckTestsRunStatuses(self):
    tests_statuses = (
        ProcessSwarmingTaskResultPipeline()
        ._CheckTestsRunStatuses(_SAMPLE_FAILURE_LOG))
    self.assertEqual(_EXPECTED_TESTS_STATUS, tests_statuses)

  @mock.patch.object(
      swarmed_test_util, 'GetOutputJsonByOutputsRef', return_value=(None, None))
  def testMonitorSwarmingTaskTimeOut(self, _):
    # Override swarming config settings to force a timeout.
    override_swarming_settings = {'task_timeout_hours': -1}
    self.UpdateUnitTestConfigSettings('swarming_settings',
                                      override_swarming_settings)

    task = FlakeSwarmingTask.Create(self.master_name, self.builder_name,
                                    self.build_number, self.step_name,
                                    self.test_name)
    task.task_id = 'task_id1'
    task.put()

    analysis = MasterFlakeAnalysis.Create(self.master_name, self.builder_name,
                                          self.build_number, self.step_name,
                                          self.test_name)
    analysis.Save()

    pipeline = ProcessFlakeSwarmingTaskResultPipeline(
        self.master_name,
        self.builder_name,
        self.build_number,
        self.step_name,
        self.build_number,
        self.test_name,
        1,
        task_id='task_id1')
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name, 'task_id1', self.build_number, self.test_name,
                 1)
    pipeline.callback(callback_params=pipeline.last_params)
    # Reload from ID to get all internal properties in sync.
    pipeline = ProcessFlakeSwarmingTaskResultPipeline.from_id(
        pipeline.pipeline_id)
    pipeline.finalized()
    step_name, task_info = pipeline.outputs.default.value
    self.assertEqual('abc_tests', task_info)
    self.assertEqual(self.step_name, step_name)

    task = FlakeSwarmingTask.Get(self.master_name, self.builder_name,
                                 self.build_number, self.step_name,
                                 self.test_name)
    self.assertEqual(analysis_status.ERROR, task.status)
    self.assertIsNone(task.tests_statuses)

  def testMonitorSwarmingTaskNotRunning(self):
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id2'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)
    # Reload from ID to get all internal properties in sync.
    pipeline = ProcessSwarmingTaskResultPipeline.from_id(pipeline.pipeline_id)
    pipeline.finalized()
    step_name, flaky_tests = pipeline.outputs.default.value

    self.assertEqual(self.step_name, step_name)
    self.assertEqual([], flaky_tests)

    task = WfSwarmingTask.Get(self.master_name, self.builder_name,
                              self.build_number, self.step_name)

    self.assertEqual(analysis_status.ERROR, task.status)
    self.assertEqual({}, task.tests_statuses)
    self.assertEqual({}, task.classified_tests)

  @mock.patch.object(
      swarming_util,
      'GetSwarmingTaskResultById',
      return_value=({}, {
          'code': 1,
          'message': 'error'
      }))
  def testMonitorSwarmingTaskGetSwarmingTaskResultIdError(self, _):
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id2'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)

    self.assertEqual(analysis_status.ERROR, task.status)
    self.assertEqual(task.error, {'code': 1, 'message': 'error'})

  @mock.patch.object(
      swarming_util,
      'GetSwarmingTaskResultById',
      return_value=(_SWARMING_TASK_RESULTS['task_id1'], {
          'code': 1,
          'message': 'error'
      }))
  @mock.patch.object(
      swarmed_test_util,
      'GetOutputJsonByOutputsRef',
      return_value=(_SAMPLE_FAILURE_LOG, None))
  def testMonitorSwarmingTaskGetSwarmingTaskResultIdErrorRecovered(self, *_):
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id1'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)

    self.assertEqual(analysis_status.COMPLETED, task.status)
    self.assertEqual(task.error, {'code': 1, 'message': 'error'})

  @mock.patch.object(
      swarmed_test_util,
      'GetOutputJsonByOutputsRef',
      return_value=(None, {
          'code': 1,
          'message': 'error'
      }))
  def testMonitorSwarmingTaskGetOutputJsonByOutputsRefError(self, _):
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id1'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)

    self.assertEqual(analysis_status.ERROR, task.status)
    self.assertEqual(task.error, {'code': 1, 'message': 'error'})

  def testMonitorSwarmingTaskWhereTaskFailed(self):
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id3'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)

    self.assertEqual(analysis_status.ERROR, task.status)
    self.assertIsNotNone(task.error)

  def testMonitorSwarmingTaskWhereNoTaskOutputs(self):
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id3'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)

    self.assertEqual(analysis_status.ERROR, task.status)
    self.assertEqual(
        task.error, {
            'code': swarming_task_error.NO_TASK_OUTPUTS,
            'message': 'outputs_ref is None'
        })

  def testMonitorSwarmingTaskNotCompletedWithExitCode(self):
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id6'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)

    self.assertEqual(analysis_status.ERROR, task.status)
    self.assertEqual(task.error, {
        'code': swarming_task_error.BOT_DIED,
        'message': 'BOT_DIED'
    })

  @mock.patch.object(
      swarmed_test_util,
      'GetOutputJsonByOutputsRef',
      return_value=(_SAMPLE_FAILURE_LOG, None))
  def testProcessSwarmingTaskResultPipeline(self, _):
    # End to end test.
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id1'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)
    pipeline.callback(callback_params=pipeline.last_params)
    # Reload from ID to get all internal properties in sync.
    pipeline = ProcessSwarmingTaskResultPipeline.from_id(pipeline.pipeline_id)
    pipeline.finalized()
    step_name, flaky_tests = pipeline.outputs.default.value

    self.assertEqual(self.step_name, step_name)
    self.assertEqual(_EXPECTED_CLASSIFIED_TESTS['flaky_tests'], flaky_tests)

    task = WfSwarmingTask.Get(self.master_name, self.builder_name,
                              self.build_number, self.step_name)

    self.assertEqual(analysis_status.COMPLETED, task.status)
    self.assertEqual(_EXPECTED_TESTS_STATUS, task.tests_statuses)
    self.assertEqual(_EXPECTED_CLASSIFIED_TESTS, task.classified_tests)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 32, 6, 538220), task.created_time)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 32, 9, 90550), task.started_time)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 33, 9), task.completed_time)
    self.assertEqual('abc_tests', task.canonical_step_name)

  @mock.patch.object(
      swarmed_test_util,
      'GetOutputJsonByOutputsRef',
      return_value=(_SAMPLE_FAILURE_LOG, None))
  def testProcessSwarmingTaskResultPipelineSerializedCallback(self, _):
    # End to end test.
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id1'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)
    pipeline.callback(callback_params=json.dumps(pipeline.last_params))
    # Reload from ID to get all internal properties in sync.
    pipeline = ProcessSwarmingTaskResultPipeline.from_id(pipeline.pipeline_id)
    pipeline.finalized()
    step_name, flaky_tests = pipeline.outputs.default.value

    self.assertEqual(self.step_name, step_name)
    self.assertEqual(_EXPECTED_CLASSIFIED_TESTS['flaky_tests'], flaky_tests)

    task = WfSwarmingTask.Get(self.master_name, self.builder_name,
                              self.build_number, self.step_name)

    self.assertEqual(analysis_status.COMPLETED, task.status)
    self.assertEqual(_EXPECTED_TESTS_STATUS, task.tests_statuses)
    self.assertEqual(_EXPECTED_CLASSIFIED_TESTS, task.classified_tests)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 32, 6, 538220), task.created_time)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 32, 9, 90550), task.started_time)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 33, 9), task.completed_time)
    self.assertEqual('abc_tests', task.canonical_step_name)

  def testMonitorSwarmingTaskStepNotExist(self):
    task_id = NO_TASK

    task = FlakeSwarmingTask.Create(self.master_name, self.builder_name,
                                    self.build_number, self.step_name,
                                    self.test_name)
    task.put()

    analysis = MasterFlakeAnalysis.Create(self.master_name, self.builder_name,
                                          self.build_number, self.step_name,
                                          self.test_name)
    analysis.Save()

    pipeline = ProcessFlakeSwarmingTaskResultPipeline(
        self.master_name, self.builder_name, self.build_number, self.step_name,
        task_id, self.build_number, self.test_name, 1)
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name, task_id, self.build_number, self.test_name, 1)
    # Reload from ID to get all internal properties in sync.
    pipeline = ProcessFlakeSwarmingTaskResultPipeline.from_id(
        pipeline.pipeline_id)
    pipeline.finalized()
    _, step_name_no_platform = pipeline.outputs.default.value

    self.assertIsNone(task.task_id)
    self.assertTrue(task.has_valid_artifact)
    self.assertEqual(analysis_status.SKIPPED, task.status)
    self.assertIsNone(step_name_no_platform)

  def testMonitorSwarmingTaskBuildException(self):
    task_id = NO_TASK_EXCEPTION

    task = FlakeSwarmingTask.Create(self.master_name, self.builder_name,
                                    self.build_number, self.step_name,
                                    self.test_name)
    task.task_id = 'task_id'
    task.put()

    analysis = MasterFlakeAnalysis.Create(self.master_name, self.builder_name,
                                          self.build_number, self.step_name,
                                          self.test_name)
    analysis.Save()

    pipeline = ProcessFlakeSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name, task_id, self.build_number, self.test_name, 1)

    self.assertIsNone(task.task_id)
    self.assertFalse(task.has_valid_artifact)

  def testMonitorSwarmingTaskFailedToTriggerUndetectedError(self):
    task = FlakeSwarmingTask.Create(self.master_name, self.builder_name,
                                    self.build_number, self.step_name,
                                    self.test_name)
    task.put()

    analysis = MasterFlakeAnalysis.Create(self.master_name, self.builder_name,
                                          self.build_number, self.step_name,
                                          self.test_name)
    analysis.Save()

    pipeline = ProcessFlakeSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name, None, self.build_number, self.test_name, 1)

    self.assertIsNone(task.task_id)
    self.assertIsNotNone(task.error)

  @mock.patch.object(
      swarmed_test_util,
      'GetOutputJsonByOutputsRef',
      return_value=(_SAMPLE_FAILURE_LOG, None))
  @mock.patch.object(test_results_util, 'IsTestResultsValid', return_value=True)
  def testProcessSwarmingTaskResultPipelineIdempotency(self, *_):
    # End to end test.
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id1'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    # NB, run is called twice.
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)
    pipeline.callback(callback_params=pipeline.last_params)
    # Reload from ID to get all internal properties in sync.
    pipeline = ProcessSwarmingTaskResultPipeline.from_id(pipeline.pipeline_id)
    pipeline.finalized()
    step_name, flaky_tests = pipeline.outputs.default.value

    self.assertEqual(self.step_name, step_name)
    self.assertEqual(_EXPECTED_CLASSIFIED_TESTS['flaky_tests'], flaky_tests)

    task = WfSwarmingTask.Get(self.master_name, self.builder_name,
                              self.build_number, self.step_name)

    self.assertEqual(analysis_status.COMPLETED, task.status)
    self.assertEqual(_EXPECTED_TESTS_STATUS, task.tests_statuses)
    self.assertEqual(_EXPECTED_CLASSIFIED_TESTS, task.classified_tests)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 32, 6, 538220), task.created_time)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 32, 9, 90550), task.started_time)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 33, 9), task.completed_time)
    self.assertEqual('abc_tests', task.canonical_step_name)

  @mock.patch.object(logging, 'warning')
  @mock.patch.object(
      ProcessBaseSwarmingTaskResultPipeline,
      'get_callback_task',
      side_effect=taskqueue.TombstonedTaskError)
  def testDelayCallbackException(self, _, mocked_logging):
    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.delay_callback(60, pipeline.last_params, name='name')
    self.assertTrue(mocked_logging.called)

  @mock.patch.object(
      swarmed_test_util,
      'GetOutputJsonByOutputsRef',
      return_value=(_SAMPLE_FAILURE_LOG, None))
  def testProcessSwarmingTaskResultPipelineBackwardCompatibleCallback(self, _):
    # End to end test.
    task = WfSwarmingTask.Create(self.master_name, self.builder_name,
                                 self.build_number, self.step_name)
    task.task_id = 'task_id1'
    task.put()

    pipeline = ProcessSwarmingTaskResultPipeline()
    pipeline.start_test()
    pipeline.run(self.master_name, self.builder_name, self.build_number,
                 self.step_name)
    pipeline.callback(**pipeline.last_params)
    # Reload from ID to get all internal properties in sync.
    pipeline = ProcessSwarmingTaskResultPipeline.from_id(pipeline.pipeline_id)
    pipeline.finalized()
    step_name, flaky_tests = pipeline.outputs.default.value

    self.assertEqual(self.step_name, step_name)
    self.assertEqual(_EXPECTED_CLASSIFIED_TESTS['flaky_tests'], flaky_tests)

    task = WfSwarmingTask.Get(self.master_name, self.builder_name,
                              self.build_number, self.step_name)

    self.assertEqual(analysis_status.COMPLETED, task.status)
    self.assertEqual(_EXPECTED_TESTS_STATUS, task.tests_statuses)
    self.assertEqual(_EXPECTED_CLASSIFIED_TESTS, task.classified_tests)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 32, 6, 538220), task.created_time)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 32, 9, 90550), task.started_time)
    self.assertEqual(
        datetime.datetime(2016, 2, 10, 18, 33, 9), task.completed_time)
    self.assertEqual('abc_tests', task.canonical_step_name)
