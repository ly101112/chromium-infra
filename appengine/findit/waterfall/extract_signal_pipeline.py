# Copyright 2015 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from gae_libs.http.http_client_appengine import HttpClientAppengine
from gae_libs.pipeline_wrapper import BasePipeline
from services import extract_signal


class ExtractSignalPipeline(BasePipeline):
  """A pipeline to extract failure signals from each failed step."""

  HTTP_CLIENT = HttpClientAppengine()

  # Arguments number differs from overridden method - pylint: disable=W0221
  def run(self, failure_info):
    """Extracts failure signals from failed steps.

    Args:
      failure_info (dict): Output of pipeline DetectFirstFailurePipeline.run().

    Returns:
      A dict like below:
      {
        'step_name1': waterfall.failure_signal.FailureSignal.ToDict(),
        ...
      }
    """
    return extract_signal.ExtractSignals(failure_info, self.HTTP_CLIENT)
