vars = {
  "chromium_git": "https://chromium.googlesource.com",
  "external_github": "https://chromium.googlesource.com/external/github.com",

  # This can be used to override the python used for generating the ENV python
  # environment. Unfortunately, some bots need this as they are attempting to
  # build the "infra/infra_python/${platform}" package which includes
  # a virtualenv which needs to point to a fixed, pre-deployed, python
  # interpreter on the system.
  #
  # This package is an awful way to distribute software, so if you see an
  # opportunity to kill it, please do so.
  "infra_env_python": "disabled",
  #
  # This is used during the transition phase of moving infra repos to git
  # submodules. To add new deps here check with the Chrome Source Team.
  "infra_superproject_checkout": False,
}

deps = {
  "build": {
    "url": "{chromium_git}/chromium/tools/build.git",
    "condition": "not infra_superproject_checkout",
  },

  # Used to initiate bootstrapping.
  #
  # This commit resolves to tag "16.7.12".
  "infra/bootstrap/virtualenv-ext":
     "{external_github}/pypa/virtualenv@" +
     "fdfec65ff031997503fb409f365ee3aeb4c2c89f",

  "infra/luci":
     "{chromium_git}/infra/luci/luci-py@" +
     "3a2ef1d0b965aa810fdd0a171f9033404d617e32",

  # This unpinned dependency is present because it is used by the trybots for
  # the recipes-py repo; They check out infra with this at HEAD, and then apply
  # the patch to it and run verifications within that copy of the repo. They
  # piggyback on top of infra in order to take advantage of it's precompiled
  # version of python-coverage.
  "infra/recipes-py":
     "{chromium_git}/infra/luci/recipes-py@" +
     "refs/heads/main",

  "infra/go/src/go.chromium.org/luci":
     "{chromium_git}/infra/luci/luci-go@" +
     "dfbd01ad9eaa62f953e990420c746ffbb0560ba6",

  "infra/go/src/go.chromium.org/chromiumos/config":
     "{chromium_git}/chromiumos/config@" +
     "832fc1b96461a68028f5eaf408de738bc79514a5",

  "infra/go/src/go.chromium.org/chromiumos/infra/proto":
     "{chromium_git}/chromiumos/infra/proto@" +
     "42f196c4e8e1c85812b568362ab967b4032c91d2",

  # Appengine third_party DEPS
  "infra/appengine/third_party/bootstrap":
     "{external_github}/twbs/bootstrap.git@" +
     "b4895a0d6dc493f17fe9092db4debe44182d42ac",

  "infra/appengine/third_party/cloudstorage":
     "{external_github}/GoogleCloudPlatform/appengine-gcs-client.git@" +
     "76162a98044f2a481e2ef34d32b7e8196e534b78",

  "infra/appengine/third_party/six":
     "{external_github}/benjaminp/six.git@" +
     "65486e4383f9f411da95937451205d3c7b61b9e1",

  "infra/appengine/third_party/oauth2client":
     "{external_github}/google/oauth2client.git@" +
     "e8b1e794d28f2117dd3e2b8feeb506b4c199c533",

  "infra/appengine/third_party/uritemplate":
     "{external_github}/uri-templates/uritemplate-py.git@" +
     "1e780a49412cdbb273e9421974cb91845c124f3f",

  "infra/appengine/third_party/httplib2":
     "{external_github}/jcgregorio/httplib2.git@" +
     "058a1f9448d5c27c23772796f83a596caf9188e6",

  "infra/appengine/third_party/endpoints-proto-datastore":
     "{external_github}/GoogleCloudPlatform/endpoints-proto-datastore.git@" +
     "971bca8e31a4ab0ec78b823add5a47394d78965a",

  "infra/appengine/third_party/difflibjs":
     "{external_github}/qiao/difflib.js.git@"
     "e11553ba3e303e2db206d04c95f8e51c5692ca28",

  "infra/appengine/third_party/pipeline":
     "{external_github}/GoogleCloudPlatform/appengine-pipelines.git@" +
     "58cf59907f67db359fe626ee06b6d3ac448c9e15",

  "infra/appengine/third_party/google-api-python-client":
     "{external_github}/google/google-api-python-client.git@" +
     "49d45a6c3318b75e551c3022020f46c78655f365",

  "infra/appengine/third_party/gae-pytz":
     "{chromium_git}/external/code.google.com/p/gae-pytz/@" +
     "4d72fd095c91f874aaafb892859acbe3f927b3cd",

  "infra/appengine/third_party/dateutil":
     "{chromium_git}/external/code.launchpad.net/dateutil/@" +
     "8c6026ba09716a4e164f5420120bfe2ebb2d9d82",

  ## For ease of development. These are pulled in as wheels for run.py/test.py
  "infra/packages/expect_tests":
     "{chromium_git}/infra/testing/expect_tests.git@" +
     "823c363d418baab7dc2dbdf184136dcfac532f64",
  "testing_support": {
     "url": "{chromium_git}/infra/testing/testing_support.git",
     "condition": "not infra_superproject_checkout",
  },

  "infra/appengine/third_party/npm_modules": {
     "url":
        "{chromium_git}/infra/third_party/npm_modules.git@" +
        "f83fafaa22f5ff396cf5306285ca3806d1b2cf1b",
     "condition": "checkout_linux or checkout_mac",
  },

  "gcloud": {
    'packages': [
      {
        'package': 'infra/3pp/tools/gcloud/${{os=mac,linux}}-${{arch=amd64}}',
        'version': 'version:2@427.0.0.chromium.3',
      }
    ],
    'dep_type': 'cipd',
    'condition': 'not infra_superproject_checkout',
  },

  "infra/cipd": {
    'packages': [
      {
        'package': 'infra/3pp/tools/protoc/${{os}}-${{arch=amd64}}',
        'version': 'version:2@21.7',
      },

      {
        'package': 'infra/3pp/tools/nodejs/${{os=linux,mac}}-${{arch}}',
        'version': 'version:2@16.13.0',
      },

      {
        'package': 'infra/3pp/tools/cloud-tasks-emulator/${{platform=linux-amd64}}',
        'version': 'version:2@1.1.1',
      },

      # TODO: Remove the arch pinning for the following two packages below when
      # they're rolled forward to a version that includes linux-arm64.
      {
        'package': 'infra/tools/luci/logdog/logdog/${{os}}-${{arch=amd64}}',
        'version': 'git_revision:fe9985447e6b95f4907774f05e9774f031700775',
      },

      {
        'package': 'infra/tools/cloudbuildhelper/${{os=mac,linux}}-${{arch=amd64}}',
        'version': 'git_revision:7088a6d730d06d978dfa0e9c5dd07e51aeb5916e',
      },

      # TODO: These should be built via 3pp instead.
      # See: docs/legacy/build_adb.md
      {
        'package': 'infra/adb/${{platform=linux-amd64}}',
        'version': 'adb_version:1.0.36',
      },
      # See: docs/legacy/build_fastboot.md
      {
        'package': 'infra/fastboot/${{platform=linux-amd64}}',
        'version': 'fastboot_version:5943271ace17',
      },

    ],
    'dep_type': 'cipd',
  },

  # Hosts legacy packages needed by the infra environment
  "infra/cipd/legacy": {
    'packages': [
      {
        'package': 'infra/3pp/tools/protoc/${{os}}-${{arch=amd64}}',
        'version': 'version:2@21.7',
      },
    ],
    'dep_type': 'cipd',
  },

  "infra/cipd/result_adapter": {
    'packages': [
      {
        'package': 'infra/tools/result_adapter/${{platform}}',
        'version': 'git_revision:fe9985447e6b95f4907774f05e9774f031700775',
      },
    ],
    'dep_type': 'cipd',
  },
}

hooks = [
  {
    "pattern": ".",
    "action": [
      "python3", "-u", "./infra/bootstrap/remove_orphaned_pycs.py",
    ],
  },

  {
    "pattern": ".",
    "condition": "'{infra_env_python}' != 'disabled'",
    "action": [
      Var("infra_env_python"), "-u", "./infra/bootstrap/bootstrap.py",
      "--deps_file", "infra/bootstrap/deps.pyl", "infra/ENV"
    ],
  },
  {
    "pattern": ".",
    "action": [
      "python3", "./infra/migration_warning.py"
    ],
    "condition": "not infra_superproject_checkout",
  },
]

recursedeps = ['build', 'infra/luci']
