# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is govered by a BSD-style
# license that can be found in the LICENSE file or at
# https://developers.google.com/open-source/licenses/bsd

"""Tests for converting internal protorpc to external protoc."""

import unittest

from api import converters
from api.api_proto import common_pb2
from api.api_proto import issue_objects_pb2
from proto import tracker_pb2
from testing import fake
from testing import testing_helpers
from tracker import tracker_bizobj


class ConverterFunctionsTest(unittest.TestCase):

  def testConvertUserRef(self):
    """We can convert user IDs to a UserRef."""
    users_by_id = {
        111L: testing_helpers.Blank(email='user@example.com'),
        }
    # No specified user
    actual = converters.ConvertUserRef(None, None, users_by_id)
    expected = common_pb2.UserRef(
        user_id=0, is_derived=False, email='----')
    self.assertEqual(expected, actual)

    # Explicitly specified user
    actual = converters.ConvertUserRef(111L, None, users_by_id)
    expected = common_pb2.UserRef(
        user_id=111L, is_derived=False, email='user@example.com')
    self.assertEqual(expected, actual)

    # Derived user
    actual = converters.ConvertUserRef(None, 111L, users_by_id)
    expected = common_pb2.UserRef(
        user_id=111L, is_derived=True, email='user@example.com')
    self.assertEqual(expected, actual)

  def testConvertUserRefs(self):
    """We can convert lists of user_ids into UserRefs."""
    users_by_id = {
        111L: testing_helpers.Blank(email='one@example.com'),
        222L: testing_helpers.Blank(email='two@example.com'),
        }
    # No specified users
    actual = converters.ConvertUserRefs(
        [], [], users_by_id)
    expected = []
    self.assertEqual(expected, actual)

    # A mix of explicit and derived users
    actual = converters.ConvertUserRefs(
        [111L], [222L], users_by_id)
    expected = [
      common_pb2.UserRef(
          user_id=111L, is_derived=False, email='one@example.com'),
      common_pb2.UserRef(
          user_id=222L, is_derived=True, email='two@example.com'),
      ]
    self.assertEqual(expected, actual)

  def testConvertLabels(self):
    """We can convert labels."""
    # No labels specified
    actual = converters.ConvertLabels([], [])
    self.assertEqual([], actual)

    # A mix of explicit and derived labels
    actual = converters.ConvertLabels(
        ['Milestone-66'], ['Restrict-View-CoreTeam'])
    expected = [
        common_pb2.LabelRef(label='Milestone-66', is_derived=False),
        common_pb2.LabelRef(label='Restrict-View-CoreTeam', is_derived=True),
        ]
    self.assertEqual(expected, actual)

  def testConvertComponents(self):
    """We can convert a list of components."""
    config = tracker_bizobj.MakeDefaultProjectIssueConfig(789)
    config.component_defs = [
      tracker_pb2.ComponentDef(component_id=1, path='UI'),
      tracker_pb2.ComponentDef(component_id=2, path='DB'),
      ]

    # No components specified
    actual = converters.ConvertComponents([], [], config)
    self.assertEqual([], actual)

    # A mix of explicit and derived components
    actual = converters.ConvertComponents([1], [2], config)
    expected = [
        common_pb2.ComponentRef(path='UI', is_derived=False),
        common_pb2.ComponentRef(path='DB', is_derived=True),
        ]
    self.assertEqual(expected, actual)

  def testConvertIssueRef(self):
    """We can convert a pair (project_name, local_id) to an IssueRef."""
    actual = converters.ConvertIssueRef(('proj', 1))
    self.assertEqual(
        common_pb2.IssueRef(project_name='proj', local_id=1),
        actual)

  def testConvertIssueRefs(self):
    """We can convert issue_ids to IssueRefs."""
    related_refs_dict = {
        78901: ('proj', 1),
        78902: ('proj', 2),
        }
    actual = converters.ConvertIssueRefs([78901, 78902], related_refs_dict)
    self.assertEqual(
        [common_pb2.IssueRef(project_name='proj', local_id=1),
         common_pb2.IssueRef(project_name='proj', local_id=2)],
        actual)

  def testConvertIssue(self):
    """We can convert a protorpc Issue to a protoc Issue."""
    users_by_id = {
        111L: testing_helpers.Blank(email='one@example.com'),
        222L: testing_helpers.Blank(email='two@example.com'),
        }
    related_refs_dict = {
        78901: ('proj', 1),
        78902: ('proj', 2),
        }
    now = 12345678
    config = tracker_bizobj.MakeDefaultProjectIssueConfig(789)
    config.component_defs = [
      tracker_pb2.ComponentDef(component_id=1, path='UI'),
      tracker_pb2.ComponentDef(component_id=2, path='DB'),
      ]
    issue = fake.MakeTestIssue(
      789, 3, 'sum', 'New', 111L, labels=['Hot'],
      derived_labels=['Scalability'], star_count=12, reporter_id=222L,
      opened_timestamp=now, component_ids=[1], project_name='proj',
      cc_ids=[111L], derived_cc_ids=[222L])

    actual = converters.ConvertIssue(
        issue, users_by_id, related_refs_dict, config)

    expected = issue_objects_pb2.Issue(
        project_name='proj', local_id=3, summary='sum',
        status_ref=common_pb2.StatusRef(
            status='New', is_derived=False, means_open=True),
        owner_ref=common_pb2.UserRef(
            user_id=111L, email='one@example.com', is_derived=False),
        cc_refs=[common_pb2.UserRef(
                     user_id=111L, email='one@example.com', is_derived=False),
                 common_pb2.UserRef(
                     user_id=222L, email='two@example.com', is_derived=True)],
        label_refs=[common_pb2.LabelRef(label='Hot', is_derived=False),
                    common_pb2.LabelRef(label='Scalability', is_derived=True)],
        component_refs=[common_pb2.ComponentRef(path='UI', is_derived=False)],
        is_deleted=False,
        reporter_ref=common_pb2.UserRef(
            user_id=222L, email='two@example.com', is_derived=False),
        opened_timestamp=now, star_count=12, is_spam=False, attachment_count=0
        )
    self.assertEqual(expected, actual)

