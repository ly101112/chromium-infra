# Generated by the pRPC protocol buffer compiler plugin.  DO NOT EDIT!
# source: go.chromium.org/luci/cv/api/migration/migration.proto

import base64
import zlib

from google.protobuf import descriptor_pb2

# Includes description of the go.chromium.org/luci/cv/api/migration/migration.proto and all of its transitive
# dependencies. Includes source code info.
FILE_DESCRIPTOR_SET = descriptor_pb2.FileDescriptorSet()
FILE_DESCRIPTOR_SET.ParseFromString(zlib.decompress(base64.b64decode(
    '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')))
_INDEX = {
    f.name: {
      'descriptor': f,
      'services': {s.name: s for s in f.service},
    }
    for f in FILE_DESCRIPTOR_SET.file
}


MigrationServiceDescription = {
  'file_descriptor_set': FILE_DESCRIPTOR_SET,
  'file_descriptor': _INDEX[u'go.chromium.org/luci/cv/api/migration/migration.proto']['descriptor'],
  'service_descriptor': _INDEX[u'go.chromium.org/luci/cv/api/migration/migration.proto']['services'][u'Migration'],
}
