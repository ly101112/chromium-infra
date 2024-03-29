# Copyright 2016 The Chromium Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""A set of functions that provide persistence for stars.

Stars can be on users, projects, or issues.
"""
from __future__ import print_function
from __future__ import division
from __future__ import absolute_import

import logging

import time
from features import filterrules_helpers
from framework import sql
from services import caches


USERSTAR_TABLE_NAME = 'UserStar'
PROJECTSTAR_TABLE_NAME = 'ProjectStar'
ISSUESTAR_TABLE_NAME = 'IssueStar'
HOTLISTSTAR_TABLE_NAME = 'HotlistStar'

# TODO(jrobbins): Consider adding memcache here if performance testing shows
# that stars are a bottleneck.  Keep in mind that issue star counts are
# already denormalized and stored in the Issue, which is cached in memcache.


class AbstractStarService(object):
  """The persistence layer for any kind of star data."""

  def __init__(self, cache_manager, tbl, item_col, user_col, cache_kind):
    """Constructor.

    Args:
      cache_manager: local cache with distributed invalidation.
      tbl: SQL table that stores star data.
      item_col: string SQL column name that holds int item IDs.
      user_col: string SQL column name that holds int user IDs
          of the user who starred the item.
      cache_kind: string saying the kind of RAM cache.
    """
    self.tbl = tbl
    self.item_col = item_col
    self.user_col = user_col

    # Items starred by users, keyed by user who did the starring.
    self.star_cache = caches.RamCache(cache_manager, 'user')
    # Users that starred an item, keyed by item ID.
    self.starrer_cache = caches.RamCache(cache_manager, cache_kind)
    # Counts of the users that starred an item, keyed by item ID.
    self.star_count_cache = caches.RamCache(cache_manager, cache_kind)

  def ExpungeStars(self, cnxn, item_id, commit=True, limit=None):
    """Wipes an item's stars from the system.

    Args:
      cnxn: connection to SQL database.
      item_id: ID of the item that's starred. ie: an issue, project, etc
      commit: whether to commit the change.
      limit: max stars to delete for performance reasons.
    """
    self.tbl.Delete(
        cnxn, commit=commit, limit=limit, **{self.item_col: item_id})

  def ExpungeStarsByUsers(self, cnxn, user_ids, limit=None):
    """Wipes a user's stars from the system.
    This method will not commit the operation. This method will
    not make changes to in-memory data.
    """
    self.tbl.Delete(cnxn, user_id=user_ids, commit=False, limit=limit)

  def LookupItemStarrers(self, cnxn, item_id):
    """Returns list of users having stars on the specified item."""
    starrer_list_dict = self.LookupItemsStarrers(cnxn, [item_id])
    return starrer_list_dict[item_id]

  def LookupItemsStarrers(self, cnxn, items_ids):
    """Returns {item_id: [uid, ...]} of users who starred these items."""
    starrer_list_dict, missed_ids = self.starrer_cache.GetAll(items_ids)

    if missed_ids:
      rows = self.tbl.Select(
          cnxn, cols=[self.item_col, self.user_col],
          **{self.item_col: missed_ids})
      # Ensure that every requested item_id has an entry so that even
      # zero-star items get cached.
      retrieved_starrers = {item_id: [] for item_id in missed_ids}
      for item_id, starrer_id in rows:
        retrieved_starrers[item_id].append(starrer_id)
      starrer_list_dict.update(retrieved_starrers)
      self.starrer_cache.CacheAll(retrieved_starrers)

    return starrer_list_dict

  def LookupStarredItemIDs(self, cnxn, starrer_user_id):
    """Returns list of item IDs that were starred by the specified user."""
    if not starrer_user_id:
      return []  # Anon user cannot star anything.

    cached_item_ids = self.star_cache.GetItem(starrer_user_id)
    if cached_item_ids is not None:
      return cached_item_ids

    rows = self.tbl.Select(cnxn, cols=[self.item_col], user_id=starrer_user_id)
    starred_ids = [row[0] for row in rows]
    self.star_cache.CacheItem(starrer_user_id, starred_ids)
    return starred_ids

  def IsItemStarredBy(self, cnxn, item_id, starrer_user_id):
    """Return True if the given issue is starred by the given user."""
    starred_ids = self.LookupStarredItemIDs(cnxn, starrer_user_id)
    return item_id in starred_ids

  def CountItemStars(self, cnxn, item_id):
    """Returns the number of stars on the specified item."""
    count_dict = self.CountItemsStars(cnxn, [item_id])
    return count_dict.get(item_id, 0)

  def CountItemsStars(self, cnxn, item_ids):
    """Get a dict {item_id: count} for the given items."""
    item_count_dict, missed_ids = self.star_count_cache.GetAll(item_ids)

    if missed_ids:
      rows = self.tbl.Select(
          cnxn, cols=[self.item_col, 'COUNT(%s)' % self.user_col],
          group_by=[self.item_col],
          **{self.item_col: missed_ids})
      # Ensure that every requested item_id has an entry so that even
      # zero-star items get cached.
      retrieved_counts = {item_id: 0 for item_id in missed_ids}
      retrieved_counts.update(rows)
      item_count_dict.update(retrieved_counts)
      self.star_count_cache.CacheAll(retrieved_counts)

    return item_count_dict

  def _SetStarsBatch(
      self, cnxn, item_id, starrer_user_ids, starred, commit=True):
    """Sets or unsets stars for the specified item and users."""
    if starred:
      rows = [(item_id, user_id) for user_id in starrer_user_ids]
      self.tbl.InsertRows(
          cnxn, [self.item_col, self.user_col], rows, ignore=True,
          commit=commit)
    else:
      self.tbl.Delete(
          cnxn, commit=commit,
          **{self.item_col: item_id, self.user_col: starrer_user_ids})

    self.star_cache.InvalidateKeys(cnxn, starrer_user_ids)
    self.starrer_cache.Invalidate(cnxn, item_id)
    self.star_count_cache.Invalidate(cnxn, item_id)

  def SetStarsBatch(
      self, cnxn, item_id, starrer_user_ids, starred, commit=True):
    """Sets or unsets stars for the specified item and users."""
    self._SetStarsBatch(
        cnxn, item_id, starrer_user_ids, starred, commit=commit)

  def SetStar(self, cnxn, item_id, starrer_user_id, starred):
    """Sets or unsets a star for the specified item and user."""
    self._SetStarsBatch(cnxn, item_id, [starrer_user_id], starred)


class UserStarService(AbstractStarService):
  """Star service for stars on users."""

  def __init__(self, cache_manager):
    tbl = sql.SQLTableManager(USERSTAR_TABLE_NAME)
    super(UserStarService, self).__init__(
        cache_manager, tbl, 'starred_user_id', 'user_id', 'user')


class ProjectStarService(AbstractStarService):
  """Star service for stars on projects."""

  def __init__(self, cache_manager):
    tbl = sql.SQLTableManager(PROJECTSTAR_TABLE_NAME)
    super(ProjectStarService, self).__init__(
        cache_manager, tbl, 'project_id', 'user_id', 'project')


class HotlistStarService(AbstractStarService):
  """Star service for stars on hotlists."""

  def __init__(self, cache_manager):
    tbl = sql.SQLTableManager(HOTLISTSTAR_TABLE_NAME)
    super(HotlistStarService, self).__init__(
        cache_manager, tbl, 'hotlist_id', 'user_id', 'hotlist')


class IssueStarService(AbstractStarService):
  """Star service for stars on issues."""

  def __init__(self, cache_manager):
    tbl = sql.SQLTableManager(ISSUESTAR_TABLE_NAME)
    super(IssueStarService, self).__init__(
        cache_manager, tbl, 'issue_id', 'user_id', 'issue')

    # HACK. Usually Monorail SQL table references should stay in their
    # respective service layer class. But for performance reasons, it's better
    # for us to directly query the Issue table here.
    self.issue_tbl = sql.SQLTableManager('Issue')

  def ExpungeStarsByUsers(self, cnxn, user_ids, limit=None):
    """Wipes a user's stars from the system.

    Ensure that issue metadata is updated on expunging.

    Args:
      cnxn: connection to SQL database.
      services:  connections to persistence layer.
      user_ids: users to delete stars for.
      limit: max stars to delete for performance reasons.
    """
    # TODO(zhangtiff): update star_count for updated issues. This is tricky
    # because star_count needs to be recomputd for each issue, so this likely
    # requires a task queue.

    timestamp = int(time.time())

    shard_id = sql.RandomShardID()
    issue_id_rows = self.tbl.Select(
        cnxn,
        cols=['IssueStar.issue_id'],
        user_id=user_ids,
        shard_id=shard_id,
        limit=limit)

    super(IssueStarService, self).ExpungeStarsByUsers(
        cnxn, user_ids, limit=limit)
    issue_ids = [row[0] for row in issue_id_rows]
    if issue_ids:
      self.issue_tbl.Update(
          cnxn, {'migration_modified': timestamp},
          id=issue_ids,
          commit=False,
          limit=limit)

  # pylint: disable=arguments-differ
  def SetStar(
      self, cnxn, services, config, issue_id, starrer_user_id, starred):
    """Add or remove a star on the given issue for the given user.

    Args:
      cnxn: connection to SQL database.
      services: connections to persistence layer.
      config: ProjectIssueConfig PB for the project containing the issue.
      issue_id: integer global ID of an issue.
      starrer_user_id: user ID of the user who starred the issue.
      starred: boolean True for adding a star, False when removing one.
    """
    self.SetStarsBatch(
        cnxn, services, config, issue_id, [starrer_user_id], starred)

  # pylint: disable=arguments-differ
  def SetStarsBatch(
      self, cnxn, services, config, issue_id, starrer_user_ids, starred):
    """Add or remove a star on the given issue for the given users.

    Args:
      cnxn: connection to SQL database.
      services: connections to persistence layer.
      config: ProjectIssueConfig PB for the project containing the issue.
      issue_id: integer global ID of an issue.
      starrer_user_id: user ID of the user who starred the issue.
      starred: boolean True for adding a star, False when removing one.
    """
    logging.info(
        'SetStarsBatch:%r, %r, %r', issue_id, starrer_user_ids, starred)
    super(IssueStarService, self).SetStarsBatch(
        cnxn, issue_id, starrer_user_ids, starred)

    # Because we will modify issues, load from DB rather than cache.
    issue = services.issue.GetIssue(cnxn, issue_id, use_cache=False)
    issue.star_count = self.CountItemStars(cnxn, issue_id)
    issue.migration_modified_timestamp = int(time.time())
    filterrules_helpers.ApplyFilterRules(cnxn, services, issue, config)
    # Note: only star_count could change due to the starring, but any
    # field could have changed as a result of filter rules.
    services.issue.UpdateIssue(cnxn, issue)

    self.star_cache.InvalidateKeys(cnxn, starrer_user_ids)
    self.starrer_cache.Invalidate(cnxn, issue_id)

  # TODO(crbug.com/monorail/8098): This method should replace SetStarsBatch.
  # New code should be calling SetStarsBatch_SkipIssueUpdate.
  # SetStarsBatch, does issue.star_count updating that should be done
  # in the business logic layer instead. E.g. We can create a
  # WorkEnv.BatchSetStars() that includes the star_count updating work.
  def SetStarsBatch_SkipIssueUpdate(
      self, cnxn, issue_id, starrer_user_ids, starred, commit=True):
    # type: (MonorailConnection, int, Sequence[int], bool, Optional[bool])
    #   -> None
    """Add or remove a star on the given issue for the given users.

    Note: unlike SetStarsBatch above, does not make any updates to the
      the issue itself e.g. updating issue.star_count.

    """
    logging.info(
        'SetStarsBatch:%r, %r, %r', issue_id, starrer_user_ids, starred)
    super(IssueStarService, self).SetStarsBatch(
        cnxn, issue_id, starrer_user_ids, starred, commit=commit)

    self.star_cache.InvalidateKeys(cnxn, starrer_user_ids)
    self.starrer_cache.Invalidate(cnxn, issue_id)
