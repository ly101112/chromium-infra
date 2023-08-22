# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from flask import Flask

import gae_ts_mon
import pipeline

from backend.handlers import rerun_analyses
from backend.handlers import rerun_analysis
from backend.handlers import update_component_config
from backend.handlers import update_inverted_index
from backend.handlers import update_repo_to_dep_path


backend_handler_mappings = [
    ('/process/rerun-analyses', 'rerun_analyses',
     rerun_analyses.RerunAnalyses().Handle, ['GET']),
    ('/process/rerun-analysis', 'rerun_analysis',
     rerun_analysis.RerunAnalysis().Handle, ['GET']),
    ('/process/update-component-config', 'update_component_config',
     update_component_config.UpdateComponentConfig().Handle, ['GET']),
    ('/process/update-inverted-index', 'update_inverted_index',
     update_inverted_index.UpdateInvertedIndex().Handle, ['GET']),
    ('/process/update-repo-to-dep-path', 'update_repo_to_dep_path',
     update_repo_to_dep_path.UpdateRepoToDepPath().Handle, ['GET']),
]

backend_app = Flask(__name__)
pipeline.create_handlers_map(backend_app)
for url, endpoint, view_func, methods in backend_handler_mappings:
  backend_app.add_url_rule(
      url, endpoint=endpoint, view_func=view_func, methods=methods)

# TODO(crbug.com/1322775) Migrate away from the shared prodx-mon-chrome-infra
# service account and change to gae_ts_mon.initialize_prod()
gae_ts_mon.initialize_adhoc(backend_app)
