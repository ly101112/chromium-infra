# Copyright 2015 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import json

import webapp2

from handlers import config
from model import wf_config
from testing_utils import testing


_MOCK_MASTERS_TO_BLACKLISTED_STEPS = {
    'master1': ['unsupported_step1', 'unsupported_step2'],
    'master2': ['unsupported_step3', 'unsupported_step4'],
}

_MOCK_VERSION_NUMBER = 12


class ConfigTest(testing.AppengineTestCase):
  app_module = webapp2.WSGIApplication([
      ('/config', config.Configuration),
  ], debug=True)

  def testGetConfigurationSettings(self):
    class MockFinditConfig():
      masters_to_blacklisted_steps = _MOCK_MASTERS_TO_BLACKLISTED_STEPS

    def MockSettings():
      return MockFinditConfig

    def MockGetCurrentVersionNumber():
      return _MOCK_VERSION_NUMBER

    self.mock(wf_config, 'FinditConfig', MockFinditConfig)
    self.mock(wf_config, 'Settings', MockSettings)
    self.mock(wf_config, 'GetCurrentVersionNumber', MockGetCurrentVersionNumber)
    self.mock_current_user(user_email='test@chromium.org', is_admin=True)

    response = self.test_app.get('/config', params={'format': 'json'})
    self.assertEquals(response.status_int, 200)

    expected_response = {
        'masters': _MOCK_MASTERS_TO_BLACKLISTED_STEPS,
        'version': _MOCK_VERSION_NUMBER,
    }

    self.assertEquals(expected_response, response.json_body)

  def testValidateMastersDict(self):
    self.assertTrue(config._SupportedMastersConfigIsValid({
        'a': ['string1', 'string2', 'string3'],
        'b': ['string1', 'string2', 'string3'],
    }))
    self.assertFalse(config._SupportedMastersConfigIsValid({
        'a': {}
    }))
    self.assertFalse(config._SupportedMastersConfigIsValid({
        'a': [1, 2, 3]
    }))
    self.assertFalse(config._SupportedMastersConfigIsValid([]))
    self.assertFalse(config._SupportedMastersConfigIsValid([{
        'a': ['b', 'c', 'd']
    }]))

  def testConfigurationDictIsValid(self):
    self.assertTrue(config._ConfigurationDictIsValid({
        'masters_to_blacklisted_steps': {
            'a': []
        }
    }))
    self.assertFalse(config._ConfigurationDictIsValid([]))
    self.assertFalse(config._ConfigurationDictIsValid({
        'this_is_not_a_valid_property': []
    }))

  def testPostConfigurationSettings(self):

    class MockFinditConfig():
      masters_to_blacklisted_steps = {
          'a': [],
      }

      def modify(self, **kwargs):
        for k, v in kwargs.iteritems():
          setattr(self, k, v)

    mock_config = MockFinditConfig()

    def MockSettings():
      return mock_config

    def MockGetCurrentVersionNumber():
      return _MOCK_VERSION_NUMBER

    self.mock_current_user(user_email='test@chromium.org', is_admin=True)
    self.mock(wf_config, 'FinditConfig', MockFinditConfig)
    self.mock(wf_config, 'Settings', MockSettings)
    self.mock(wf_config, 'GetCurrentVersionNumber', MockGetCurrentVersionNumber)

    params = {
        'format': 'json',
        'data': json.dumps({
            'masters_to_blacklisted_steps': {
                'a': ['1', '2', '3'],
                'b': []
            }
        })
    }

    expected_response = {
        'masters': {
            'a': ['1', '2', '3'],
            'b': []
        },
        'version': _MOCK_VERSION_NUMBER,
    }

    response = self.test_app.post('/config', params=params)
    self.assertEquals(expected_response, response.json_body)
