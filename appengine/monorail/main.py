# Copyright 2016 The Chromium Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
"""Main program for Monorail.

Monorail is an issue tracking tool that is based on the code.google.com
issue tracker, but it has been ported to Google AppEngine and Google Cloud SQL.
"""
from __future__ import print_function
from __future__ import division
from __future__ import absolute_import

import flask
import six

# Fix imports before importing gae_ts_mon.
import import_utils

import_utils.FixImports()

import gae_ts_mon

import registerpages
from framework import sorting
from services import service_manager

if six.PY3:
  # https://github.com/GoogleCloudPlatform/appengine-python-standard/issues/70
  import functools
  from google.appengine.api import memcache
  unpickler = functools.partial(six.moves.cPickle.Unpickler, encoding='bytes')
  memcache.setup_client(memcache.Client(unpickler=unpickler))

services = service_manager.set_up_services()
sorting.InitializeArtValues(services)

app = flask.Flask(__name__)

registerpages.ServletRegistry().Register(services, app)
registerpages.RegisterEndpointsUrls(app)

gae_ts_mon.initialize_prod(app)
