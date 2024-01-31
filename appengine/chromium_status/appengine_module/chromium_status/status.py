# coding=utf-8
# Copyright (c) 2012 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Status management pages."""

import datetime
import json
import logging
import re
import urllib2
import contextlib

from google.appengine.api import app_identity
from google.appengine.api import memcache
from google.appengine.ext import db

from appengine_module.chromium_status.base_page import BasePage
from appengine_module.chromium_status import utils


ALLOWED_ORIGINS = [
    'https://gerrit-int.chromium.org',
    'https://gerrit.chromium.org',
    'https://chrome-internal-review.googlesource.com',
    'https://chromium-review.googlesource.com',
]


DEFAULT_USERNAME = "user"

EMAIL_SCOPE = 'https://www.googleapis.com/auth/userinfo.email'

JSON_PREFIX = ")]}'"

ISO_FORMAT = '%Y-%m-%dT%H:%M:%S.%fZ'

TREE_STATUS_SERVICE = ('https://luci-tree-status.appspot.com/' +
                       'prpc/luci.tree_status.v1.TreeStatus/')

class TextFragment(object):
  """Simple object to hold text that might be linked"""

  def __init__(self, text, target=None, is_email=False):
    self.text = text
    self.target = target
    self.is_email = is_email

  def __repr__(self):
    return 'TextFragment({%s->%s})' % (self.text, self.target)


class LinkableText(object):
  """Turns arbitrary text into a set of links"""

  CR_REV_URL = 'https://crrev.com'

  GERRIT_URLS = {
      'chrome': 'https://chrome-internal-review.googlesource.com',
      'chromium': 'https://chromium-review.googlesource.com',
  }

  WATERFALL_URLS = {
      'chromeos': 'https://uberchromegw.corp.google.com/i/chromeos',
      # chromium-status.appspot.com:
      'chromium': 'https://build.chromium.org/p/chromium',
      'chromiumos': 'http://build.chromium.org/p/chromiumos',
      # infra-status.appspot.com:
      'infra': 'https://build.chromium.org/p/chromium.infra',
      # v8-status.appspot.com:
      'v8': 'https://build.chromium.org/p/client.v8',
  }

  APP_PREFIXES = (
      'dev~',
      's~',
  )

  # Automatically linkify known strings for the user.
  _CONVERTS = []

  @classmethod
  def register_converter(cls, regex, target, pretty, is_email, flags=re.I):
    """Register a new conversion for creating links from text"""
    cls._CONVERTS.append(
        (re.compile(regex, flags=flags), target, pretty, is_email))

  @classmethod
  def bootstrap(cls, is_chromiumos):
    """Add conversions (possibly specific to |app_name| instance)"""
    # Convert Chromium bug links.  Support the forms:
    # chromium:1234
    # The project list comes from: https://bugs.chromium.org/hosting/
    # $ curl -s https://bugs.chromium.org/hosting/ | \
    #   grep -o 'href="/p/[^/"]*/"' | \
    #   sed -e 's:.*/p/::' -e 's:/"::'
    chromium_projects = (
        'angleproject',
        'aomedia',
        'boringssl',
        'chromedriver',
        'chromium',
        'crashpad',
        'gerrit',
        'google-breakpad',
        'gyp',
        'libyuv',
        'linux-syscall-support',
        'monorail',
        'nativeclient',
        'oss-fuzz',
        'pdfium',
        'project-zero',
        'skia',
        'swiftshader',
        'v8',
        'webm',
        'webp',
        'webports',
        'webrtc',
    )
    cls.register_converter(
        r'\b((%s):([0-9]+))\b' % r'|'.join(chromium_projects),
        r'https://bugs.chromium.org/p/\2/issues/detail?id=\3', r'\1', False)

    # Convert <project>:<number> bug links.  Support the forms:
    # chromium:1234
    # chromium-os:1234
    cls.register_converter(
        #   1   2        3      4        5       6 7
        r'\b((https?://)?((crbug|crosbug)(\.com)?(/(p/)?[0-9]+)))\b',
        r'https://\4.com\6', r'\1', False)

    # Convert internal b/ & b: bug links.
    cls.register_converter(
        r'\b(https?://)?(b[:/]([0-9]+))\b',
        r'https://b.corp.google.com/\3', r'\2', False)

    # Convert e-mail addresses.
    cls.register_converter(
        r'(([-+.a-z0-9_!#$%&*/=?^_`{|}~]+)@[-a-z0-9.]+\.[a-z0-9]+)\b',
        r'\1', r'\2', True)

    # Convert SHA1's to links to the respective commit on crrev.com.
    cls.register_converter(
        r'\b([0-9a-f]{40})\b',
        r'%s/\1' % cls.CR_REV_URL, r'\1', False)

    # Convert public gerrit CL numbers which take the form:
    # CL:1234
    cls.register_converter(
        r'\b(CL[: ]([0-9]+))\b',
        r'%s/\2' % cls.GERRIT_URLS['chromium'], r'\1', False)
    # Convert internal gerrit CL numbers which take the form:
    # CL:*1234
    cls.register_converter(
        r'\b(CL[: ]\*([0-9]+))\b',
        r'%s/\2' % cls.GERRIT_URLS['chrome'], r'\1', False)

    # Match the string:
    #   Automatic: "cbuildbot" on "x86-generic ASAN" from.
    # Do this for everyone since "cbuildbot" is unique to CrOS.
    # Otherwise, we'd do it only for chromium |app_name| instances.
    cls.register_converter(
        r'("cbuildbot" on "([^"]+ (canary|master|launcher))")',
        r'%s/builders/\2' % cls.WATERFALL_URLS['chromeos'], r'\1', False)
    cls.register_converter(
        r'("cbuildbot" on "([^"]+)")',
        r'%s/builders/\2' % cls.WATERFALL_URLS['chromiumos'], r'\1', False)

    # Try to linkify /builder path names according to the tree.
    if BasePage.APP_NAME in cls.WATERFALL_URLS:
      cls.register_converter(
          r'([^/]*)/builders([^"]+) ',
          r'%s/builders\2' % cls.WATERFALL_URLS[BasePage.APP_NAME], r'\2',
          False)

    # Convert all other URLs into links. Regexp based on @stephenhay's idea from
    # https://mathiasbynens.be/demo/url-regex.
    cls.register_converter(
        r'\b(https?://[^\s/$.?#].[^\s]*)\b', r'\1', r'\1', False)

    if is_chromiumos:
      # Match the string '"builder name"-internal/public-buildnumber:'. E.g.,
      #   "Canary master"-i-120:
      # This applies only to the CrOS instance where the builders may update
      # the tree status directly.
      cls.register_converter(
          r'("([\w\s]+)"-i-(\d+):)',
          r'%s/builders/\2/builds/\3' % cls.WATERFALL_URLS['chromeos'],
          r'\1', False
      )
      cls.register_converter(
          r'("([\w\s]+)"-p-(\d+):)',
          r'%s/builders/\2/builds/\3' % cls.WATERFALL_URLS['chromiumos'],
          r'\1', False
      )

  @classmethod
  def parse(cls, text):
    """Creates a list of TextFragment objects based on |text|"""
    if not text:
      return []
    for prog, target, pretty_text, is_email in cls._CONVERTS:
      m = prog.search(text)
      if m:
        link = TextFragment(m.expand(pretty_text),
                            target=m.expand(target),
                            is_email=is_email)
        left_links = cls.parse(text[:m.start()].rstrip())
        right_links = cls.parse(text[m.end():].lstrip())
        return left_links + [link] + right_links
    return [TextFragment(text)]

  def __init__(self, text):
    self.raw_text = text
    self.links = self.parse(text.strip())

  def __str__(self):
    return self.raw_text


class Status(db.Model):
  """Description for the status table."""
  # The username who added this status.
  username = db.StringProperty(required=True)
  # The date when the status got added.
  date = db.DateTimeProperty(auto_now_add=True)
  # The date after which this entry will be deleted.
  expiry_date = db.DateTimeProperty()
  # The message. It can contain html code.
  message = db.StringProperty(required=True)

  _general_state = None

  def __init__(self, *args, **kwargs):
    # Normalize newlines otherwise the DB store barfs.  We don't really want to
    # make this field handle newlines as none of the places where we output the
    # content is designed to handle it, nor the clients that consume us.
    kwargs['message'] = kwargs.get('message', '').replace('\n', ' ')
    super(Status, self).__init__(*args, **kwargs)

  @property
  def username_links(self):
    return LinkableText(self.username)

  @property
  def message_links(self):
    return LinkableText(self.message)

  @property
  def general_state(self):
    """Returns a string representing the state that the status message
    describes.

    Note: Keep in sync with main.html help text.
    """
    if self._general_state is not None:
      return self._general_state
    message = self.message
    closed = re.search('close', message, re.IGNORECASE)
    if closed and re.search('maint', message, re.IGNORECASE):
      return 'maintenance'
    if re.search('throt', message, re.IGNORECASE):
      return 'throttled'
    if closed:
      return 'closed'
    return 'open'

  @property
  def can_commit_freely(self):
    return self.general_state == 'open'

  def AsDict(self):
    data = super(Status, self).AsDict()
    data['general_state'] = self.general_state
    data['can_commit_freely'] = self.can_commit_freely
    if 'key' not in data or not data['key']:
      data['key'] = -1
    return data

  @classmethod
  def FromDict(cls, value):
    """Creates a Status object from a dict with message, createUser and
    createTime keys (i.e. the Status proto from the new tree status app)."""
    status = Status(
        message=value['message'],
        username=value['createUser'],
        date=datetime.datetime.strptime(value['createTime'], ISO_FORMAT))
    status._general_state = value['generalState'].lower()
    return status


def tree_name_for_prpc():
  """Get the tree name for the current tree status instance in the new Tree
  Status service.

  This is equivalent to the app engine application id with the -status suffix
  removed (if present)
  """
  name = app_identity.get_application_id()
  if name.endswith('-hr'):
    name = name[:-len('-hr')]
  if name.endswith('-status'):
    name = name[:-len('-status')]
  return name


def tree_status_prpc(method, request_json):
  """Makes a PRPC call to the given method in the luci tree status server

    request_json should be a JSON serializable representation of the request
    object.

    returns a parsed JSON structure of the response.

    raises an Exception on any errors.
  """
  request = urllib2.Request(
      TREE_STATUS_SERVICE + method, json.dumps(request_json), {
          'Accept':
              'application/json',
          'Authorization':
              'Bearer ' + app_identity.get_access_token([EMAIL_SCOPE])[0],
          'Content-Type':
              'application/json',
      })
  with contextlib.closing(urllib2.urlopen(request)) as response:
    if response.getcode() != 200:
      logging.warning('server returned HTTP error: %s', response.getcode())
      raise Exception('tree status server returned HTTP error for ' + method +
                      ' ' + json.dumps(request_json) + ': ' +
                      response.getcode())
    text = response.read()
    if text.startswith(JSON_PREFIX):
      text = text[len(JSON_PREFIX):]
    return json.loads(text)

def get_status():
  """Returns the current Status, e.g. the most recent one."""
  status = memcache.get('last_status')
  if status is None:
    try:
      status = Status.FromDict(
          tree_status_prpc(
              'GetStatus',
              {'name': 'trees/' + tree_name_for_prpc() + '/status/latest'}))
      memcache.set('last_status', status, 30)  # cache for 30 seconds
      return status
    except Exception as e:
      logging.warning(
          'exception when making tree status PRPC request', exc_info=e)
    status = Status.all().order('-date').get()
    memcache.set('last_status', status, 30)  # cache for 30 seconds
  return status


def put_status(status):
  """Sets the current Status, e.g. append a new one."""
  status.put()
  memcache.set('last_status', status, 30)  # cache for 30 seconds
  memcache.delete('last_statuses')


def get_last_statuses(limit):
  """Returns the last |limit| statuses."""
  statuses = memcache.get('last_statuses')
  if not statuses or len(statuses) < limit:
    try:
      values_json = tree_status_prpc('ListStatus', {
          'parent': 'trees/' + tree_name_for_prpc() + '/status',
          'pageSize': limit
      })
      statuses = [Status.FromDict(value) for value in values_json['status']]
      memcache.set('last_statuses', statuses, 30)  # cache for 30 seconds
      return statuses[:limit]
    except Exception as e:
      logging.warning(
          'exception when making tree status PRPC request', exc_info=e)
    statuses = Status.all().order('-date').fetch(limit)
    memcache.add('last_statuses', statuses, 30)
  return statuses[:limit]


def parse_date(date):
  """Parses a date."""
  match = re.match(r'^(\d\d\d\d)-(\d\d)-(\d\d)$', date)
  if match:
    return datetime.datetime(
        int(match.group(1)), int(match.group(2)), int(match.group(3)))
  if date.isdigit():
    return datetime.datetime.utcfromtimestamp(int(date))
  return None


def limit_length(string, length):
  """Limits the string |string| at length |length|.

  Inserts an ellipsis if it is cut.
  """
  string = unicode(string.strip())
  if len(string) > length:
    string = string[:length - 1] + u'…'
  return string


class AllStatusPage(BasePage):
  """Displays a big chunk, 1500, status values."""
  @utils.requires_read_access
  def get(self):
    query = db.Query(Status).order('-date')
    start_date = self.request.get('startTime')
    if start_date:
      query.filter('date <', parse_date(start_date))
    try:
      limit = int(self.request.get('limit'))
    except ValueError:
      limit = 1000
    end_date = self.request.get('endTime')
    beyond_end_of_range_status = None
    if end_date:
      query.filter('date >=', parse_date(end_date))
      # We also need to get the very next status in the range, otherwise
      # the caller can't tell what the effective tree status was at time
      # |end_date|.
      beyond_end_of_range_status = Status.all(
          ).filter('date <', end_date).order('-date').get()

    out_format = self.request.get('format', 'csv')
    if out_format in ['csv', 'json']:
      try:
        # Note that this disregards the startTime/endTime parameters, we can
        # re-introduce these if required, but I suspect they are only used
        # from the UI which is being retired.
        values = tree_status_prpc(
            'ListStatus', {
                'parent': 'trees/' + tree_name_for_prpc() + '/status',
                'pageSize': limit
            })
        statuses = [Status.FromDict(value) for value in values['status']]
      except Exception as e:
        logging.warning(
            "exception in prpc call to tree status server from AllStatusPage",
            exc_info=e)
        statuses = query.fetch(limit)
    if out_format == 'csv':
      # It's not really an html page.
      self.response.headers['Content-Type'] = 'text/plain'
      template_values = self.InitializeTemplate(self.APP_NAME + ' Tree Status')
      template_values['status'] = statuses
      template_values['beyond_end_of_range_status'] = beyond_end_of_range_status
      self.DisplayTemplate('allstatus.html', template_values)
    elif out_format == 'json':
      self.response.headers['Content-Type'] = 'application/json'
      self.response.headers['Access-Control-Allow-Origin'] = '*'
      statuses = [s.AsDict() for s in statuses]
      if beyond_end_of_range_status:
        statuses.append(beyond_end_of_range_status.AsDict())
      data = json.dumps(statuses)
      callback = self.request.get('callback')
      if callback:
        if re.match(r'^[a-zA-Z$_][a-zA-Z$0-9._]*$', callback):
          data = '%s(%s);' % (callback, data)
      self.response.out.write(data)
    else:
      self.response.headers['Content-Type'] = 'text/plain'
      self.response.out.write('Invalid format')


class CurrentPage(BasePage):
  """Displays the /current page."""

  def get(self):
    # Show login link on current status bar when login is required.
    out_format = self.request.get('format', 'html')
    if out_format == 'html' and not self.read_access and not self.user:
      template_values = self.InitializeTemplate(self.APP_NAME + ' Tree Status')
      template_values['show_login'] = True
      self.DisplayTemplate('current.html', template_values, use_cache=True)
    else:
      self._handle()

  @utils.requires_bot_login
  def post(self):
    """Handles the same get request from a backdoor.

    POST to receive the password plaintext without polluting the logs.
    """
    return self._handle()

  @utils.requires_read_access
  def _handle(self):
    """Displays the current message in various formats."""
    out_format = self.request.get('format', 'html')
    status = get_status()
    if out_format == 'raw':
      self.response.headers['Content-Type'] = 'text/plain'
      self.response.headers['Access-Control-Allow-Origin'] = '*'
      self.response.out.write(status.message)
    elif out_format == 'json':
      self.response.headers['Content-Type'] = 'application/json'
      origin = self.request.headers.get('Origin')
      if self.request.get('with_credentials') and origin in ALLOWED_ORIGINS:
        self.response.headers['Access-Control-Allow-Origin'] = origin
        self.response.headers['Access-Control-Allow-Credentials'] = 'true'
      else:
        self.response.headers['Access-Control-Allow-Origin'] = '*'
      data = json.dumps(status.AsDict())
      callback = self.request.get('callback')
      if callback:
        if re.match(r'^[a-zA-Z$_][a-zA-Z$0-9._]*$', callback):
          data = '%s(%s);' % (callback, data)
      self.response.out.write(data)
    elif out_format == 'html':
      template_values = self.InitializeTemplate(self.APP_NAME + ' Tree Status')
      template_values['show_login'] = False
      template_values['message'] = status.message
      template_values['state'] = status.general_state
      self.DisplayTemplate('current.html', template_values, use_cache=True)
    else:
      self.error(400)


class StatusPage(BasePage):
  """Displays the /status page."""

  def get(self):
    """Displays 1 if the tree is open, and 0 if the tree is closed."""
    # NOTE: This item is always public to allow waterfalls to check it.
    status = get_status()
    self.response.headers['Cache-Control'] = 'no-cache, private, max-age=0'
    self.response.headers['Content-Type'] = 'text/plain'
    self.response.out.write(str(int(status.can_commit_freely)))

  @utils.requires_bot_login
  @utils.requires_write_access
  def post(self):
    """Adds a new message from a backdoor.

    The main difference with MainPage.post() is that it doesn't look for
    conflicts and doesn't redirect to /.
    """
    # TODO(tandrii): switch to using service accounts.
    message = self.request.get('message')
    message = limit_length(message, 500)
    username = self.request.get('username')
    if message and username:
      expiry_date = datetime.datetime.utcnow() + datetime.timedelta(days=140)
      put_status(
          Status(message=message, username=username, expiry_date=expiry_date))
    self.response.out.write('OK')


class StatusViewerPage(BasePage):
  """Displays the /status_viewer page."""

  @utils.requires_read_access
  def get(self):
    """Displays status_viewer.html template."""
    template_values = self.InitializeTemplate(self.APP_NAME + ' Tree Status')
    self.DisplayTemplate('status_viewer.html', template_values)


class MainPage(BasePage):
  """Displays the main page containing the last 25 messages."""

  # NOTE: This is require_login in order to ensure that authentication doesn't
  # happen while changing the tree status.
  @utils.requires_login
  @utils.requires_read_access
  def get(self):
    self.redirect('https://ci.chromium.org/ui/labs/tree-status/' +
                  tree_name_for_prpc())

  @utils.requires_login
  @utils.requires_write_access
  def post(self):
    """Adds a new message."""
    # We pass these variables back into get(), prepare them.
    last_message = ''
    error_message = ''

    # Get the posted information.
    new_message = self.request.get('message')
    new_message = limit_length(new_message, 500)
    last_status_key = self.request.get('last_status_key')
    if not new_message:
      # A submission contained no data. It's a better experience to redirect
      # in this case.
      self.redirect("/")
      return

    current_status = get_status()
    if current_status and (last_status_key != str(current_status.key().id())):
      error_message = ('Message not saved, mid-air collision detected, '
                       'please resolve any conflicts and try again!')
      last_message = new_message
      return self._handle(error_message, last_message)
    else:
      expiry_date = datetime.datetime.utcnow() + datetime.timedelta(days=140)
      put_status(
          Status(
              message=new_message,
              username=self.user.email(),
              expiry_date=expiry_date))
      self.redirect("/")


class CleanupUsernamesCron(BasePage):
  """ Handler of the cron job that cleans up usernames. """

  @utils.requires_work_queue_login
  def get(self):
    """Cleans up usernames after 30 days by setting
       the username field to 'user'.
    """
    logging.info('Cleaning up Status usernames.')

    threshold_date = datetime.datetime.utcnow() - datetime.timedelta(days=30)
    statuses = Status.gql('WHERE date < :1', threshold_date).fetch(limit=None)
    updated_statuses = [s for s in statuses if s.username != DEFAULT_USERNAME]
    for status in updated_statuses:
      status.username = DEFAULT_USERNAME
    db.put(updated_statuses)

    logging.info('Status usernames clean up done, %d rows affected.',
                 len(updated_statuses))

    return self.response.out.write('Usernames cleared.')


def bootstrap():
  # Guarantee that at least one instance exists.
  if db.GqlQuery('SELECT __key__ FROM Status').get() is None:
    Status(username='none', message='welcome to status').put()
  LinkableText.bootstrap(BasePage.IS_CHROMIUMOS)
