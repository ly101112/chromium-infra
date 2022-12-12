# Copyright 2016 The Chromium Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Utility functions and classes for dealing with saved queries.

Saved queries can be part of the project issue config, where they are
called "canned queries".  Or, they can be personal saved queries that
may appear in the search scope drop-down, on the user's dashboard, or
in the user's subscription.
"""
from __future__ import print_function
from __future__ import division
from __future__ import absolute_import

import logging
import re

from framework import template_helpers
from tracker import tracker_bizobj
from tracker import tracker_constants


MAX_QUERIES = 100


def ParseSavedQueries(cnxn, post_data, project_service, prefix=''):
  """Parse form data for the Saved Queries part of an admin form."""
  saved_queries = []
  for i in range(1, MAX_QUERIES + 1):
    if ('%ssavedquery_name_%s' % (prefix, i)) not in post_data:
      continue  # skip any entries that are blank or have no predicate.

    name = post_data['%ssavedquery_name_%s' % (prefix, i)].strip()
    if not name:
      continue  # skip any blank entries

    if '%ssavedquery_id_%s' % (prefix, i) in post_data:
      query_id = int(post_data['%ssavedquery_id_%s' % (prefix, i)])
    else:
      query_id = None  # a new query_id will be generated by the DB.

    project_names_str = post_data.get(
        '%ssavedquery_projects_%s' % (prefix, i), '')
    project_names = [pn.strip().lower()
                     for pn in re.split('[],;\s]+', project_names_str)
                     if pn.strip()]
    project_ids = list(project_service.LookupProjectIDs(
        cnxn, project_names).values())

    base_id = int(post_data['%ssavedquery_base_%s' % (prefix, i)])
    query = post_data['%ssavedquery_query_%s' % (prefix, i)].strip()

    subscription_mode_field = '%ssavedquery_sub_mode_%s' % (prefix, i)
    if subscription_mode_field in post_data:
      subscription_mode = post_data[subscription_mode_field].strip()
    else:
      subscription_mode = None

    saved_queries.append(tracker_bizobj.MakeSavedQuery(
        query_id, name, base_id, query, subscription_mode=subscription_mode,
        executes_in_project_ids=project_ids))

  return saved_queries


class SavedQueryView(template_helpers.PBProxy):
  """Wrapper class that makes it easier to display SavedQuery via EZT."""

  def __init__(self, sq, idx, cnxn, project_service):
    """Store relevant values for later display by EZT.

    Args:
      sq: A SavedQuery protocol buffer.
      idx: Int index of this saved query in the list.
      cnxn: connection to SQL database.
      project_service: persistence layer for project data.
    """
    super(SavedQueryView, self).__init__(sq)

    self.idx = idx
    base_query_name = 'All issues'
    for canned in tracker_constants.DEFAULT_CANNED_QUERIES:
      qid, name, _base_id, _query = canned
      if qid == sq.base_query_id:
        base_query_name = name

    if cnxn:
      project_names = sorted(project_service.LookupProjectNames(
          cnxn, sq.executes_in_project_ids).values())
      self.projects = ', '.join(project_names)
    else:
      self.projects = ''

    self.docstring = '[%s] %s' % (base_query_name, sq.query)


def SavedQueryToCond(saved_query):
  """Convert a SavedQuery PB to a user query condition string."""
  if saved_query is None:
    return ''

  base_cond = tracker_bizobj.GetBuiltInQuery(saved_query.base_query_id)
  cond = '%s %s' % (base_cond, saved_query.query)
  return cond.strip()


def SavedQueryIDToCond(cnxn, features_service, query_id):
  """Convert a can/query ID to a user query condition string."""
  built_in = tracker_bizobj.GetBuiltInQuery(query_id)
  if built_in:
    return built_in

  saved_query = features_service.GetSavedQuery(cnxn, query_id)
  return SavedQueryToCond(saved_query)
