# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

PYTHON_VERSION_COMPATIBILITY = 'PY2+3'

DEPS = [
    'recipe_engine/buildbucket',
    'recipe_engine/context',
    'recipe_engine/platform',
    'recipe_engine/step',
    'recipe_engine/properties',

    'infra_checkout',
]


def RunSteps(api):
  co = api.infra_checkout.checkout(
      gclient_config_name='infra',
      patch_root='infra',
      generate_env_with_system_python=api.properties.get('sys_env', False),
      go_version_variant=api.properties.get('go_version_variant'))
  co.gclient_runhooks()
  _ = co.bot_update_step  # coverage...
  with co.go_env():
    api.step('go test', ['go', 'test', 'infra/...'])
  with co.go_env():  # for coverage
    api.step('go test', ['go', 'test', 'infra/...'])
  with api.context(cwd=co.patch_root_path):
    api.step('python tests', ['python3', 'test.py', 'test', 'infra'])
  with api.context(cwd=co.path):
    api.step('dirs', ['python3', api.resource('dirs.py')])


def GenTests(api):
  for plat in ('linux', 'mac', 'win'):
    yield (
        api.test(plat) +
        api.platform(plat, 64) +
        api.buildbucket.ci_build(
            project='infra',
            bucket='ci',
            git_repo='https://chromium.googlesource.com/infra/infra',
        )
    )

  yield api.test(
      'sys_env',
      api.properties(sys_env=True),
      api.buildbucket.ci_build(
          project='infra',
          bucket='ci',
          git_repo='https://chromium.googlesource.com/infra/infra',
      )
  )

  yield api.test(
      'sys_env win',
      api.properties(sys_env=True),
      api.platform('win', 64),
      api.buildbucket.ci_build(
          project='infra',
          bucket='ci',
          git_repo='https://chromium.googlesource.com/infra/infra',
      )
  )

  yield api.test(
      'sys_env arm',
      api.properties(sys_env=True),
      api.platform('linux', 64, 'arm'),
      api.buildbucket.ci_build(
          project='infra',
          bucket='ci',
          git_repo='https://chromium.googlesource.com/infra/infra',
      )
  )

  yield api.test(
      'override go version',
      api.properties(go_version_variant='bleeding_edge'),
      api.buildbucket.ci_build(
          project='infra',
          bucket='ci',
          git_repo='https://chromium.googlesource.com/infra/infra',
      )
  )
