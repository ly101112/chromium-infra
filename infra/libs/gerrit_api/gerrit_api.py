# Copyright 2014 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Yet another wrapper around Gerrit REST API."""

import base64
import cookielib
import functools
import json
import logging
import requests
import requests_cache
import time
import urllib

from requests.packages import urllib3


LOGGER = logging.getLogger(__name__)

def _not_read_only(f):
  @functools.wraps(f)
  def wrapper(self, *args, **kwargs):
    if self._read_only:
      raise AccessViolationException(
          'Method call of method not accessible for read_only Gerrit instance.')
    return f(self, *args, **kwargs)
  return wrapper


class AccessViolationException(Exception):
  """A method was called which would require write access to Gerrit."""

class UnexpectedResponseException(Exception):
  """Gerrit returned something unexpected."""

  def __init__(self, http_code, body):  # pragma: no cover
    super(UnexpectedResponseException, self).__init__()
    self.http_code = http_code
    self.body = body

  def __str__(self):  # pragma: no cover
    return 'Unexpected response (HTTP %d): %s' % (self.http_code, self.body)


class BlockCookiesPolicy(cookielib.DefaultCookiePolicy):
  def set_ok(self, cookie, request):
    return False # pragma: no cover


class Gerrit(object):
  """Wrapper around a single Gerrit host.

  Args:
    host (str): gerrit host name.
    creds (Credentials): provides credentials for the Gerrit host.
    throttle_delay_sec (int): minimal time delay between two requests, to
      avoid hammering the Gerrit server.
  """

  def __init__(self, host, creds, throttle_delay_sec=0, read_only=False):
    self._auth_header = 'Basic %s' % (
        base64.b64encode('%s:%s' % creds[host]))
    self._url_base = 'https://%s/a' % host.rstrip('/')
    self._throttle = throttle_delay_sec
    self._read_only = read_only
    self._last_call_ts = None
    self.session = requests.Session()
    # Do not use cookies with Gerrit. This breaks interaction with Google's
    # Gerrit instances. Do not use cookies as advised by the Gerrit team.
    self.session.cookies.set_policy(BlockCookiesPolicy())
    retry_config = urllib3.util.Retry(total=4, backoff_factor=2,
                                      status_forcelist=[500, 503])
    self.session.mount(self._url_base, requests.adapters.HTTPAdapter(
        max_retries=retry_config))


  def _request(self, method, request_path, params=None, body=None):
    """Sends HTTP request to Gerrit.

    Args:
      method: HTTP method (e.g 'GET', 'POST', ...).
      request_path: URL of the endpoint, relative to host (e.g. '/accounts/id').
      params: dict with query parameters.
      body: optional request body, will be serialized to JSON.

    Returns:
      Tuple (response code, deserialized JSON response).
    """
    if not request_path.startswith('/'):
      request_path = '/' + request_path

    full_url = '%s%s' % (self._url_base, request_path)

    # Wait to avoid Gerrit quota, don't wait if a response is in the cache.
    if self._throttle and not _is_response_cached(method, full_url):
      now = time.time()
      if self._last_call_ts and now - self._last_call_ts < self._throttle:
        time.sleep(self._throttle - (now - self._last_call_ts))
      self._last_call_ts = time.time()

    headers = {
        # This makes the server return compact JSON.
        'Accept': 'application/json',
        # This means responses will be gzip compressed.
        'Accept-encoding': 'gzip',
        'Authorization': self._auth_header,
    }

    if body is not None:
      body = json.dumps(body)
      headers['Content-Type'] = 'application/json;charset=UTF-8'
      headers['Content-Length'] = str(len(body))

    LOGGER.debug('%s %s', method, full_url)
    response = self.session.request(
        method=method,
        url=full_url,
        params=params,
        data=body,
        headers=headers)

    # Gerrit prepends )]}' to response.
    prefix = ')]}\'\n'
    body = response.text
    if body and body.startswith(prefix):
      body = json.loads(body[len(prefix):])

    return response.status_code, body

  def get_account(self, account_id):
    """Returns a dict describing a Gerrit account or None if no such account.
    Documentation:
    https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-account

    Args:
      account_id: email, numeric account id, or 'self'.

    Returns:
      None if no such account, AccountInfo dict otherwise.
    """
    assert '/' not in account_id
    code, body = self._request('GET', '/accounts/%s' % account_id)
    if code == 200:
      return body
    if code == 404:
      return None
    raise UnexpectedResponseException(code, body)

  @_not_read_only
  def add_group_members(self, group, members):
    """Adds a bunch of members to a group.
    Documentation:
    https://gerrit-review.googlesource.com/Documentation/rest-api-groups.html#_add_group_members

    Args:
      group: name of a group to add members to.
      members: iterable with emails of accounts to add to the group.

    Returns:
      None

    Raises:
      UnexpectedResponseException: if call failed.
    """
    if '/' in group:
      raise ValueError('Invalid group name: %s' % group)
    code, body = self._request(
        method='POST',
        request_path='/groups/%s/members.add' % group,
        body={'members': list(members)})
    if code != 200:
      raise UnexpectedResponseException(code, body)
    return body

  def is_account_active(self, account_id): # pragma: no cover
    if '/' in account_id:
      raise ValueError('Invalid account id: %s' % account_id)
    code, body = self._request(
        method='GET',
        request_path='/accounts/%s/active' % account_id)
    if code == 200:
      return True
    if code == 204:
      return False
    raise UnexpectedResponseException(code, body)

  @_not_read_only
  def activate_account(self, account_id): # pragma: no cover
    """Sets account state to 'active'.

    Args:
      account_id (str): account to update.

    Raises:
      UnexpectedResponseException: if gerrit does not answer as expected.
    """
    if '/' in account_id:
      raise ValueError('Invalid account id: %s' % account_id)
    code, body = self._request(
        method='PUT',
        request_path='/accounts/%s/active' % account_id)
    if code not in (200, 201):
      raise UnexpectedResponseException(code, body)

  def get_projects(self, prefix=''): # pragma: no cover
    """Returns list of projects with names starting with a prefix.

    Args:
      prefix (str): optional project name prefix to limit the listing to.

    Returns:
      Dict <project name> -> {'state': 'ACTIVE', 'parent': 'All-Projects'}

    Raises:
      UnexpectedResponseException: if gerrit does not answer as expected.
    """
    code, body = self._request(
        method='GET',
        request_path='/projects/?p=%s&t' % urllib.quote(prefix, safe=''))
    if code not in (200, 201):
      raise UnexpectedResponseException(code, body)
    return body

  def get_project_parent(self, project): # pragma: no cover
    """Retrieves the name of a project's parent project.

    Returns None If |project| is not registered on Gerrit or doesn't have
    a parent (like 'All-Projects').

    Args:
      project (str): project to query.

    Raises:
      UnexpectedResponseException: if gerrit does not answer as expected.
    """
    code, body = self._request(
        method='GET',
        request_path='/projects/%s/parent' % urllib.quote(project, safe=''))
    if code == 404:
      return None
    if code not in (200, 201):
      raise UnexpectedResponseException(code, body)
    assert isinstance(body, unicode)
    return body if body else None

  @_not_read_only
  def set_project_parent(self, project, parent, commit_message=None):
    """Changes project's parent project.
    Documentation:
    https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#set-project-parent

    Args:
      project (str): project to change.
      parent (str): parent to set.
      commit_message (str): message for corresponding refs/meta/config commit.

    Raises:
      UnexpectedResponseException: if gerrit does not answer as expected.
    """
    commit_message = (
        commit_message or ('Changing parent project to %s' % parent))
    code, body = self._request(
        method='PUT',
        request_path='/projects/%s/parent' % urllib.quote(project, safe=''),
        body={'parent': parent, 'commit_message': commit_message})
    if code not in (200, 201):
      raise UnexpectedResponseException(code, body)
    return body

  def _query(
      self,
      project,
      query_name=None,
      with_messages=True,
      with_labels=True,
      with_revisions=True,
      **kwargs):
    """Queries the Gerrit API changes endpoint. Returns a list of ChangeInfo
    dictionaries.
    Documentation:
    https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-changes

    Args:
      project: (str) The project name.
      query_name: (str) The name of the named query stored for the CQ user.
      with_messages: (bool) If True, adds the o=MESSAGES option.
      with_labels: (bool) If True, adds the o=LABELS option.
      with_revisions: (bool) If True, adds the o=ALL_REVISIONS option.
      kwargs: Allows to specify additional query parameters.
    """

    # We always restrict queries with the project name.
    request_params = 'project:%s' % project

    if query_name:
      request_params += '+query:%s' % query_name
    for operator,value in kwargs.iteritems():
      request_params += '+%s:%s' % (operator, value)

    if with_messages:
      request_params += '&o=MESSAGES'
    if with_labels:
      request_params += '&o=LABELS'
    if with_revisions:
      request_params += '&o=ALL_REVISIONS'

    request_path = '/changes/?q=%s' % request_params
    code, body = self._request(method='GET', request_path=request_path)
    if code != 200:
      raise UnexpectedResponseException(code, body)
    return body

  def get_pending_issues(self, project):
    """Returns a list of ChangeInfo dictionaries of all the pending issues on
    Gerrit.

    Args:
      project: The Gerrit project.
    """
    return self._query(project, query_name='pending_cls')

  def get_commit_ready_issues(self, project):
    """Returns a lit of ChangeInfo dictionaries of all issues on Gerrit ready to
    be committed

    Args:
      project: The Gerrit project.
    """
    return self._query(project, query_name='commit_ready_cls',
                       with_messages=False)


def _is_response_cached(method, full_url):
  """Returns True if response to GET request is in requests_cache.

  Args:
    method (str): http verb ('GET', 'POST', etc.)
    full_url (str): url, including the protocol
  Returns:
    is_cached (bool):
"""
  if method != 'GET':
    return False # pragma: no cover
  try:
    cache = requests_cache.get_cache()
  except AttributeError: # pragma: no cover
    cache = None
  return cache.has_url(full_url) if cache else False
