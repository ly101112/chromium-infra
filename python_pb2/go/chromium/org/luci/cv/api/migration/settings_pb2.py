# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: go.chromium.org/luci/cv/api/migration/settings.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='go.chromium.org/luci/cv/api/migration/settings.proto',
  package='migration',
  syntax='proto3',
  serialized_options=b'Z1go.chromium.org/luci/cv/api/migration;migrationpb',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n4go.chromium.org/luci/cv/api/migration/settings.proto\x12\tmigration\"\xc4\x02\n\x08Settings\x12.\n\tapi_hosts\x18\x01 \x03(\x0b\x32\x1b.migration.Settings.ApiHost\x12\x32\n\x0buse_cv_runs\x18\x03 \x01(\x0b\x32\x1d.migration.Settings.UseCVRuns\x12\x30\n\x0epssa_migration\x18\x02 \x01(\x0b\x32\x18.migration.PSSAMigration\x1a]\n\x07\x41piHost\x12\x0c\n\x04host\x18\x01 \x01(\t\x12\x16\n\x0eproject_regexp\x18\x02 \x03(\t\x12\x1e\n\x16project_regexp_exclude\x18\x04 \x03(\t\x12\x0c\n\x04prod\x18\x03 \x01(\x08\x1a\x43\n\tUseCVRuns\x12\x16\n\x0eproject_regexp\x18\x01 \x03(\t\x12\x1e\n\x16project_regexp_exclude\x18\x02 \x03(\t\"+\n\rPSSAMigration\x12\x1a\n\x12projects_blocklist\x18\x01 \x03(\tB3Z1go.chromium.org/luci/cv/api/migration;migrationpbb\x06proto3'
)




_SETTINGS_APIHOST = _descriptor.Descriptor(
  name='ApiHost',
  full_name='migration.Settings.ApiHost',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='host', full_name='migration.Settings.ApiHost.host', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='project_regexp', full_name='migration.Settings.ApiHost.project_regexp', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='project_regexp_exclude', full_name='migration.Settings.ApiHost.project_regexp_exclude', index=2,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='prod', full_name='migration.Settings.ApiHost.prod', index=3,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=230,
  serialized_end=323,
)

_SETTINGS_USECVRUNS = _descriptor.Descriptor(
  name='UseCVRuns',
  full_name='migration.Settings.UseCVRuns',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_regexp', full_name='migration.Settings.UseCVRuns.project_regexp', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='project_regexp_exclude', full_name='migration.Settings.UseCVRuns.project_regexp_exclude', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=325,
  serialized_end=392,
)

_SETTINGS = _descriptor.Descriptor(
  name='Settings',
  full_name='migration.Settings',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='api_hosts', full_name='migration.Settings.api_hosts', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='use_cv_runs', full_name='migration.Settings.use_cv_runs', index=1,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='pssa_migration', full_name='migration.Settings.pssa_migration', index=2,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_SETTINGS_APIHOST, _SETTINGS_USECVRUNS, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=68,
  serialized_end=392,
)


_PSSAMIGRATION = _descriptor.Descriptor(
  name='PSSAMigration',
  full_name='migration.PSSAMigration',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='projects_blocklist', full_name='migration.PSSAMigration.projects_blocklist', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=394,
  serialized_end=437,
)

_SETTINGS_APIHOST.containing_type = _SETTINGS
_SETTINGS_USECVRUNS.containing_type = _SETTINGS
_SETTINGS.fields_by_name['api_hosts'].message_type = _SETTINGS_APIHOST
_SETTINGS.fields_by_name['use_cv_runs'].message_type = _SETTINGS_USECVRUNS
_SETTINGS.fields_by_name['pssa_migration'].message_type = _PSSAMIGRATION
DESCRIPTOR.message_types_by_name['Settings'] = _SETTINGS
DESCRIPTOR.message_types_by_name['PSSAMigration'] = _PSSAMIGRATION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Settings = _reflection.GeneratedProtocolMessageType('Settings', (_message.Message,), {

  'ApiHost' : _reflection.GeneratedProtocolMessageType('ApiHost', (_message.Message,), {
    'DESCRIPTOR' : _SETTINGS_APIHOST,
    '__module__' : 'go.chromium.org.luci.cv.api.migration.settings_pb2'
    # @@protoc_insertion_point(class_scope:migration.Settings.ApiHost)
    })
  ,

  'UseCVRuns' : _reflection.GeneratedProtocolMessageType('UseCVRuns', (_message.Message,), {
    'DESCRIPTOR' : _SETTINGS_USECVRUNS,
    '__module__' : 'go.chromium.org.luci.cv.api.migration.settings_pb2'
    # @@protoc_insertion_point(class_scope:migration.Settings.UseCVRuns)
    })
  ,
  'DESCRIPTOR' : _SETTINGS,
  '__module__' : 'go.chromium.org.luci.cv.api.migration.settings_pb2'
  # @@protoc_insertion_point(class_scope:migration.Settings)
  })
_sym_db.RegisterMessage(Settings)
_sym_db.RegisterMessage(Settings.ApiHost)
_sym_db.RegisterMessage(Settings.UseCVRuns)

PSSAMigration = _reflection.GeneratedProtocolMessageType('PSSAMigration', (_message.Message,), {
  'DESCRIPTOR' : _PSSAMIGRATION,
  '__module__' : 'go.chromium.org.luci.cv.api.migration.settings_pb2'
  # @@protoc_insertion_point(class_scope:migration.PSSAMigration)
  })
_sym_db.RegisterMessage(PSSAMigration)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
