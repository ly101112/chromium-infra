# Copyright 2015 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from infra_libs.ts_mon.common import metrics


# Extending HTTP status codes to client-side errors and timeouts.
STATUS_OK = 200
STATUS_ERROR = 901
STATUS_TIMEOUT = 902
STATUS_EXCEPTION = 909


request_bytes = metrics.CumulativeDistributionMetric(
    'http/request_bytes',
    description='Bytes sent per http request (body only).')
response_bytes = metrics.CumulativeDistributionMetric(
    'http/response_bytes',
    description='Bytes received per http requests (content only).')
durations = metrics.CumulativeDistributionMetric(
    'http/durations',
    description='Time elapsed between sending a request and getting a'
                ' response (including parsing) in milliseconds.')
response_status = metrics.CounterMetric('http/response_status')
