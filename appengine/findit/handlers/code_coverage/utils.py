# Copyright 2021 The Chromium Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import re
import six
from google.cloud import storage
from google.cloud.storage.retry import DEFAULT_RETRY
import logging

from common.findit_http_client import FinditHttpClient
from gae_libs.gitiles.cached_gitiles_repository import CachedGitilesRepository
from model.code_coverage import CoverageReportModifier
from model.code_coverage import DependencyRepository
from waterfall import waterfall_config

# Cloud storage bucket used to store the source files fetched from gitile.
_SOURCE_FILE_GS_BUCKET = 'source-files-for-coverage'

_GCS_BUCKET_BLOB_REGEX = re.compile(r'/(.*?)/(.*)')

# Regex to identify all test related files
# It includes all those files
# - with 'tests/' or 'testing/' in their path
# - with _test_ or /test_ in their name
# - with filename ending with 'Test', 'Tests', 'test' or 'tests'
TEST_FILE_REGEX = r'(^|.+\/)test(s|ing)?\/.+|.+(_|\/)test_.+|.+(T|t)ests?\..*'

# Regex to identify main files
MAIN_FILE_REGEX = r'.*(m|M)ain\..*'

# Regex to identify example files
EXAMPLE_FILE_REGEX = r'.*example.*'


def GetPostsubmitPlatformInfoMap(luci_project):
  """Returns a map of postsubmit platform information.

  The map contains per-luci_project platform information, and following is
  an example config:
  {
    'postsubmit_platform_info_map': {
      'chromium': {
        'linux': {
          'bucket': 'ci',
          'buider': 'linux-code-coverage',
          'coverage_tool': 'clang',
          'ui_name': 'Linux (C/C++)',
        }
      }
    }
  }
  """
  return waterfall_config.GetCodeCoverageSettings().get(
      'postsubmit_platform_info_map', {}).get(luci_project, {})


def GetAllowedGitilesConfigs():
  """Returns the set of valid gitiles configurations.

  The returned structure contains the tree of valid hosts, projects, and refs.

  Please note that the hosts in the config are gitiles hosts instead of gerrit
  hosts, such as: 'chromium.googlesource.com'.

  Example config:
  {
    'allowed_gitiles_configs': {
      'chromium.googlesource.com': {
        'chromium/src': [
          'refs/heads/main',
        ]
      }
    }
  }
  """
  return waterfall_config.GetCodeCoverageSettings().get(
      'allowed_gitiles_configs', {})


def GetAllowedGerritConfigs():
  """Returns the set of valid gerrit configurations.

  The returned structure contains the map of valid gerrit_hosts to projects.

  Example config:
  {
    'allowed_gerrit_configs': {
      'chromium-review.googlesource.com': [
        'chromium/src'
      ]
    }
  }
  """
  return waterfall_config.GetCodeCoverageSettings().get(
      'allowed_gerrit_configs', {})


def GetMatchedDependencyRepository(manifest, file_path):  # pragma: no cover.
  """Gets the matched dependency in the manifest of the report.

  Args:
    manifest (DependencyRepository): Entity containing mapping from path prefix
                                     to corresponding repo.
    file_path (str): Source absolute path to the file.

  Returns:
    A DependencyRepository if a matched one is found and it is allowed,
    otherwise None.
  """
  assert file_path.startswith('//'), 'All file path should start with "//".'

  for dep in manifest:
    if file_path.startswith(
        dep.path) and dep.server_host in GetAllowedGitilesConfigs():
      return dep

  return None


def GetActiveReferenceCommits(server_host, project):
  """Returns commits against which coverage is to be generated.

  Returns ids of the CoverageReportModifier corresponding to the active
  reference commits.
  """
  query = CoverageReportModifier.query(
      CoverageReportModifier.server_host == server_host,
      CoverageReportModifier.project == project,
      CoverageReportModifier.is_active == True)
  modifier_ids = []
  for x in query.fetch():
    if x.reference_commit:
      modifier_ids.append(x.key.id())
  return modifier_ids


def GetLastActiveReferenceCommitForYear(server_host, project, year):
  """Returns last active commit for a given year

  Args:
    server_host (str): Gitiles hostname, e.g. "chromium.googlesource.com".
    project (str): Gitiles project name, e.g. "chromium/src.git".
    year (int): Year whose last commit is desired.
  """
  modifier_ids = GetActiveReferenceCommits(server_host, project)
  reference_commits = []
  logging.info("year = %d", year)
  logging.info("modifier_ids = %r", modifier_ids)
  for modifier_id in modifier_ids:
    modifier = CoverageReportModifier.Get(modifier_id)
    logging.info("modifier_year = %d", modifier.reference_commit_timestamp.year)
    if modifier.reference_commit_timestamp.year == year:
      reference_commits.append(
          (modifier.reference_commit_timestamp, modifier_id))
  reference_commits.sort()
  return reference_commits[-1][1]


def GetFileContentFromGitiles(manifest, file_path,
                              revision):  # pragma: no cover.
  """Fetches the content of a specific revision of a file from gitiles.

  Args:
    manifest (DependencyRepository): Entity containing mapping from path prefix
                                     to corresponding repo.
    file_path (str): Source absolute path to the file.
    revision (str): The gitile revision of the file.

  Returns:
    The content of the source file."""
  assert file_path.startswith('//'), 'All file path should start with "//".'
  assert revision, 'A valid revision is required'

  dependency = GetMatchedDependencyRepository(manifest, file_path)
  assert dependency, ('%s file does not belong to any dependency repository' %
                      file_path)

  # Calculate the relative path to the root of the dependency repository itself.
  relative_file_path = file_path[len(dependency.path):]
  repo = CachedGitilesRepository(FinditHttpClient(), dependency.project_url)
  return repo.GetSource(relative_file_path, revision)


def ComposeSourceFileGsPath(manifest, file_path, revision):
  """Composes a cloud storage path for a specific revision of a source file.

  Args:
    manifest (DependencyRepository): Entity containing mapping from path prefix
                                     to corresponding repo.
    file_path (str): Source absolute path to the file.
    revision (str): The gitile revision of the file in its own repo.

  Returns:
    Cloud storage path to the file, in the format /bucket/object. For example,
    /source-files-for-coverage/chromium.googlesource.com/v8/v8/src/date.cc/1234.
  """
  assert file_path.startswith('//'), 'All file path should start with "//".'
  assert revision, 'A valid revision is required'

  dependency = GetMatchedDependencyRepository(manifest, file_path)
  assert dependency, ('%s file does not belong to any dependency repository' %
                      file_path)

  # Calculate the relative path to the root of the dependency repository itself.
  relative_file_path = file_path[len(dependency.path):]
  return '/%s/%s/%s/%s/%s' % (_SOURCE_FILE_GS_BUCKET, dependency.server_host,
                              dependency.project, relative_file_path, revision)


def IsFileAvailableInGs(gs_path):  # pragma: no cover.
  """Returns True if the specified object exists, otherwise False.

  Args:
    gs_path (str): Path to the file, in the format /bucket/object.

  Returns:
    True if the object exists, otherwise False.
  """
  try:
    # extract first word between slashes as bucket
    m = _GCS_BUCKET_BLOB_REGEX.match(gs_path)
    logging.error(gs_path)
    bucket, blob = m[1], m[2]
    storage_client = storage.Client()
    bucket = storage_client.bucket(bucket)
    blob = bucket.blob(blob)
    return blob.exists()
  except Exception as e:
    logging.error(e)

def WriteFileContentToGs(gs_path, content):  # pragma: no cover.
  """Writes the content of a file to cloud storage.

  Args:
    gs_path (str): Path to the file, in the format /bucket/object.
    content (str): Content of the file.
  """
  # extract first word between slashes as bucket
  m = _GCS_BUCKET_BLOB_REGEX.match(gs_path)
  bucket, blob = m[1], m[2]
  storage_client = storage.Client()
  bucket = storage_client.bucket(bucket)
  blob = bucket.blob(blob)
  modified_retry = DEFAULT_RETRY.with_delay(multiplier=2)
  with blob.open('w', content_type='text/plain', retry=modified_retry) as f:
    f.write(content)


def GetFileContentFromGs(gs_path):  # pragma: no cover.
  """Reads the content of a file in cloud storage.

  This method is more expensive than |_IsFileAvailableInGs|, so if the goal is
  to check if a file exists, |_IsFileAvailableInGs| is preferred.

  Args:
    gs_path (str): Path to the file, in the format /bucket/object.

  Returns:
    The content of the file if it exists, otherwise None."""
  try:
    # extract first word between slashes as bucket
    m = _GCS_BUCKET_BLOB_REGEX.match(gs_path)
    bucket, blob = m[1], m[2]
    storage_client = storage.Client()
    bucket = storage_client.bucket(bucket)
    blob = bucket.blob(blob)
    with blob.open('rb') as f:
      return f.read()
  except Exception as e:
    logging.error(e)
    return None
