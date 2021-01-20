# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api_clients_config.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='api_clients_config.proto',
  package='monorail',
  syntax='proto2',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x18\x61pi_clients_config.proto\x12\x08monorail\"\xa4\x01\n\x11ProjectPermission\x12\x0f\n\x07project\x18\x01 \x01(\t\x12;\n\x04role\x18\x02 \x01(\x0e\x32 .monorail.ProjectPermission.Role:\x0b\x63ontributor\x12\x19\n\x11\x65xtra_permissions\x18\x03 \x03(\t\"&\n\x04Role\x12\r\n\tcommitter\x10\x01\x12\x0f\n\x0b\x63ontributor\x10\x02\"\x98\x02\n\x06\x43lient\x12\x14\n\x0c\x63lient_email\x18\x01 \x01(\t\x12\x14\n\x0c\x64isplay_name\x18\x02 \x01(\t\x12\x11\n\tclient_id\x18\x03 \x01(\t\x12\x17\n\x0f\x61llowed_origins\x18\n \x03(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12\x38\n\x13project_permissions\x18\x05 \x03(\x0b\x32\x1b.monorail.ProjectPermission\x12\x1c\n\x0cperiod_limit\x18\x06 \x01(\x05:\x06\x31\x30\x30\x30\x30\x30\x12\x1f\n\x0elifetime_limit\x18\x07 \x01(\x05:\x07\x31\x30\x30\x30\x30\x30\x30\x12\x10\n\x08\x63ontacts\x18\x08 \x03(\t\x12\x16\n\tqpm_limit\x18\t \x01(\x05:\x03\x31\x30\x30\".\n\tClientCfg\x12!\n\x07\x63lients\x18\x01 \x03(\x0b\x32\x10.monorail.Client'
)



_PROJECTPERMISSION_ROLE = _descriptor.EnumDescriptor(
  name='Role',
  full_name='monorail.ProjectPermission.Role',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='committer', index=0, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='contributor', index=1, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=165,
  serialized_end=203,
)
_sym_db.RegisterEnumDescriptor(_PROJECTPERMISSION_ROLE)


_PROJECTPERMISSION = _descriptor.Descriptor(
  name='ProjectPermission',
  full_name='monorail.ProjectPermission',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='monorail.ProjectPermission.project', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='role', full_name='monorail.ProjectPermission.role', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=True, default_value=2,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='extra_permissions', full_name='monorail.ProjectPermission.extra_permissions', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _PROJECTPERMISSION_ROLE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=39,
  serialized_end=203,
)


_CLIENT = _descriptor.Descriptor(
  name='Client',
  full_name='monorail.Client',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='client_email', full_name='monorail.Client.client_email', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='display_name', full_name='monorail.Client.display_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='client_id', full_name='monorail.Client.client_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='allowed_origins', full_name='monorail.Client.allowed_origins', index=3,
      number=10, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='description', full_name='monorail.Client.description', index=4,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='project_permissions', full_name='monorail.Client.project_permissions', index=5,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='period_limit', full_name='monorail.Client.period_limit', index=6,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=True, default_value=100000,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='lifetime_limit', full_name='monorail.Client.lifetime_limit', index=7,
      number=7, type=5, cpp_type=1, label=1,
      has_default_value=True, default_value=1000000,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='contacts', full_name='monorail.Client.contacts', index=8,
      number=8, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='qpm_limit', full_name='monorail.Client.qpm_limit', index=9,
      number=9, type=5, cpp_type=1, label=1,
      has_default_value=True, default_value=100,
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
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=206,
  serialized_end=486,
)


_CLIENTCFG = _descriptor.Descriptor(
  name='ClientCfg',
  full_name='monorail.ClientCfg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='clients', full_name='monorail.ClientCfg.clients', index=0,
      number=1, type=11, cpp_type=10, label=3,
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
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=488,
  serialized_end=534,
)

_PROJECTPERMISSION.fields_by_name['role'].enum_type = _PROJECTPERMISSION_ROLE
_PROJECTPERMISSION_ROLE.containing_type = _PROJECTPERMISSION
_CLIENT.fields_by_name['project_permissions'].message_type = _PROJECTPERMISSION
_CLIENTCFG.fields_by_name['clients'].message_type = _CLIENT
DESCRIPTOR.message_types_by_name['ProjectPermission'] = _PROJECTPERMISSION
DESCRIPTOR.message_types_by_name['Client'] = _CLIENT
DESCRIPTOR.message_types_by_name['ClientCfg'] = _CLIENTCFG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ProjectPermission = _reflection.GeneratedProtocolMessageType('ProjectPermission', (_message.Message,), {
  'DESCRIPTOR' : _PROJECTPERMISSION,
  '__module__' : 'api_clients_config_pb2'
  # @@protoc_insertion_point(class_scope:monorail.ProjectPermission)
  })
_sym_db.RegisterMessage(ProjectPermission)

Client = _reflection.GeneratedProtocolMessageType('Client', (_message.Message,), {
  'DESCRIPTOR' : _CLIENT,
  '__module__' : 'api_clients_config_pb2'
  # @@protoc_insertion_point(class_scope:monorail.Client)
  })
_sym_db.RegisterMessage(Client)

ClientCfg = _reflection.GeneratedProtocolMessageType('ClientCfg', (_message.Message,), {
  'DESCRIPTOR' : _CLIENTCFG,
  '__module__' : 'api_clients_config_pb2'
  # @@protoc_insertion_point(class_scope:monorail.ClientCfg)
  })
_sym_db.RegisterMessage(ClientCfg)


# @@protoc_insertion_point(module_scope)
