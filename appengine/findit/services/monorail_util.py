# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
"""Functions for interfacing with Mororail bugs."""
import logging

from gae_libs import appengine_util
from monorail_api import Issue
from monorail_api import IssueTrackerAPI
from services import issue_constants


def GetOpenIssues(query, monorail_project):
  """Searches for open bugs that match the query.

  This method wraps a call IssueTrackerAPI.getIssues(), it is needed because
  it's unclear from the API name and the documentation whether the returned
  issues are all open issues or not, so check the property to make sure that
  only open bugs are considered.

  Args:
    query: A query to search for bugs on Monorail.
    monorail_project: The monorail project to search for.

  Returns:
    A list of open bugs that match the query.
  """
  issue_tracker_api = IssueTrackerAPI(
      monorail_project, use_staging=appengine_util.IsStaging())
  issues = issue_tracker_api.getIssues(query)
  if not issues:
    return []

  return [issue for issue in issues if issue.open]


def OpenBugAlreadyExistsForId(bug_id, project_id='chromium'):
  """Returns True if the bug exists and is open on monorail."""
  existing_bug = GetMergedDestinationIssueForId(bug_id, project_id)
  return existing_bug and existing_bug.open


def GetMergedDestinationIssueForId(issue_id, monorail_project='chromium'):
  """Given an id, traverse the merge chain to get the destination issue.

  Args:
    issue_id: The id to get merged destination issue for.
    monorail_project: The Monorail project the issue is on.

  Returns:
    The destination issue if the original issue was merged, otherwise itself,
    and returns None if there is an exception while communicating with
    Monorail.

    NOTE: If there is a cycle in the merge chain, the first visited issue in
    the cycle will be returned.
  """
  if issue_id is None:
    return None

  issue_tracker_api = IssueTrackerAPI(
      monorail_project, use_staging=appengine_util.IsStaging())
  issue = issue_tracker_api.getIssue(issue_id)
  visited_issues = set()

  while issue and issue.merged_into:
    logging.info('Issue %s was merged into %s on project: %s.', issue.id,
                 issue.merged_into, monorail_project)
    visited_issues.add(issue)
    issue = issue_tracker_api.getIssue(issue.merged_into)
    if issue in visited_issues:
      # There is a cycle, bails out.
      break

  return issue


def CreateBug(issue, project_id='chromium'):
  """Creates a bug with the given information.

  Returns:
    (int) id of the bug that was filed.
  """
  assert issue

  issue_tracker_api = IssueTrackerAPI(
      project_id, use_staging=appengine_util.IsStaging())
  issue_tracker_api.create(issue)

  return issue.id


def UpdateBug(issue, comment, project_id='chromium'):
  """Creates a bug with the given information."""
  assert issue

  issue_tracker_api = IssueTrackerAPI(
      project_id, use_staging=appengine_util.IsStaging())
  issue_tracker_api.update(issue, comment, send_email=True)

  return issue.id


def CreateIssueWithIssueGenerator(issue_generator):
  """Creates a new issue with a given issue generator.

  Args:
    issue_generator: A FlakyTestIssueGenerator object.

  Returns:
    The id of the newly created issue.
  """
  labels = issue_generator.GetLabels()
  labels.append(issue_generator.GetPriority())
  issue = Issue({
      'status': issue_generator.GetStatus(),
      'summary': issue_generator.GetSummary(),
      'description': issue_generator.GetDescription(),
      'projectId': issue_generator.GetMonorailProject(),
      'labels': labels,
      'fieldValues': [issue_generator.GetFlakyTestCustomizedField()]
  })

  issue_id = CreateBug(issue, issue_generator.GetMonorailProject())
  if issue_id:
    issue_generator.OnIssueCreated()

  return issue_id


def UpdateIssueWithIssueGenerator(issue_id, issue_generator):
  """Updates an existing issue with a given issue generator.

  Args:
    issue_id: Id of the issue to be updated.
    issue_generator: A FlakyTestIssueGenerator object.
  """
  issue = GetMergedDestinationIssueForId(issue_id,
                                         issue_generator.GetMonorailProject())
  for label in issue_generator.GetLabels():
    # It is most likely that existing issues already have their priorities set
    # by developers, so it would be annoy if FindIt tries to overwrite it.
    if label.startswith('Pri-'):
      continue

    if (label == issue_constants.SHERIFF_CHROMIUM_LABEL and
        not issue_generator.ShouldRestoreChromiumSheriffLabel()):
      continue

    if label not in issue.labels:
      issue.labels.append(label)

  issue.field_values.append(issue_generator.GetFlakyTestCustomizedField())
  UpdateBug(issue, issue_generator.GetComment(),
            issue_generator.GetMonorailProject())
  issue_generator.OnIssueUpdated()
