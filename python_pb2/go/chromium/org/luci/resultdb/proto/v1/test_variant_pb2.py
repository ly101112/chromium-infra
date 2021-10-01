# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: go.chromium.org/luci/resultdb/proto/v1/test_variant.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from go.chromium.org.luci.resultdb.proto.v1 import common_pb2 as go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_common__pb2
from go.chromium.org.luci.resultdb.proto.v1 import test_metadata_pb2 as go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_test__metadata__pb2
from go.chromium.org.luci.resultdb.proto.v1 import test_result_pb2 as go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_test__result__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='go.chromium.org/luci/resultdb/proto/v1/test_variant.proto',
  package='luci.resultdb.v1',
  syntax='proto3',
  serialized_options=b'Z/go.chromium.org/luci/resultdb/proto/v1;resultpb',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n9go.chromium.org/luci/resultdb/proto/v1/test_variant.proto\x12\x10luci.resultdb.v1\x1a\x33go.chromium.org/luci/resultdb/proto/v1/common.proto\x1a:go.chromium.org/luci/resultdb/proto/v1/test_metadata.proto\x1a\x38go.chromium.org/luci/resultdb/proto/v1/test_result.proto\"\xba\x02\n\x0bTestVariant\x12\x0f\n\x07test_id\x18\x01 \x01(\t\x12*\n\x07variant\x18\x02 \x01(\x0b\x32\x19.luci.resultdb.v1.Variant\x12\x14\n\x0cvariant_hash\x18\x03 \x01(\t\x12\x33\n\x06status\x18\x04 \x01(\x0e\x32#.luci.resultdb.v1.TestVariantStatus\x12\x33\n\x07results\x18\x05 \x03(\x0b\x32\".luci.resultdb.v1.TestResultBundle\x12\x37\n\x0c\x65xonerations\x18\x06 \x03(\x0b\x32!.luci.resultdb.v1.TestExoneration\x12\x35\n\rtest_metadata\x18\x07 \x01(\x0b\x32\x1e.luci.resultdb.v1.TestMetadata\"@\n\x10TestResultBundle\x12,\n\x06result\x18\x01 \x01(\x0b\x32\x1c.luci.resultdb.v1.TestResult\"K\n\x14TestVariantPredicate\x12\x33\n\x06status\x18\x01 \x01(\x0e\x32#.luci.resultdb.v1.TestVariantStatus*\xa0\x01\n\x11TestVariantStatus\x12#\n\x1fTEST_VARIANT_STATUS_UNSPECIFIED\x10\x00\x12\x0e\n\nUNEXPECTED\x10\n\x12\x18\n\x14UNEXPECTEDLY_SKIPPED\x10\x14\x12\t\n\x05\x46LAKY\x10\x1e\x12\x0e\n\nEXONERATED\x10(\x12\x13\n\x0fUNEXPECTED_MASK\x10-\x12\x0c\n\x08\x45XPECTED\x10\x32\x42\x31Z/go.chromium.org/luci/resultdb/proto/v1;resultpbb\x06proto3'
  ,
  dependencies=[go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_common__pb2.DESCRIPTOR,go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_test__metadata__pb2.DESCRIPTOR,go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_test__result__pb2.DESCRIPTOR,])

_TESTVARIANTSTATUS = _descriptor.EnumDescriptor(
  name='TestVariantStatus',
  full_name='luci.resultdb.v1.TestVariantStatus',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='TEST_VARIANT_STATUS_UNSPECIFIED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='UNEXPECTED', index=1, number=10,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='UNEXPECTEDLY_SKIPPED', index=2, number=20,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='FLAKY', index=3, number=30,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EXONERATED', index=4, number=40,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='UNEXPECTED_MASK', index=5, number=45,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EXPECTED', index=6, number=50,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=711,
  serialized_end=871,
)
_sym_db.RegisterEnumDescriptor(_TESTVARIANTSTATUS)

TestVariantStatus = enum_type_wrapper.EnumTypeWrapper(_TESTVARIANTSTATUS)
TEST_VARIANT_STATUS_UNSPECIFIED = 0
UNEXPECTED = 10
UNEXPECTEDLY_SKIPPED = 20
FLAKY = 30
EXONERATED = 40
UNEXPECTED_MASK = 45
EXPECTED = 50



_TESTVARIANT = _descriptor.Descriptor(
  name='TestVariant',
  full_name='luci.resultdb.v1.TestVariant',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='test_id', full_name='luci.resultdb.v1.TestVariant.test_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='variant', full_name='luci.resultdb.v1.TestVariant.variant', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='variant_hash', full_name='luci.resultdb.v1.TestVariant.variant_hash', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='luci.resultdb.v1.TestVariant.status', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='results', full_name='luci.resultdb.v1.TestVariant.results', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='exonerations', full_name='luci.resultdb.v1.TestVariant.exonerations', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='test_metadata', full_name='luci.resultdb.v1.TestVariant.test_metadata', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=251,
  serialized_end=565,
)


_TESTRESULTBUNDLE = _descriptor.Descriptor(
  name='TestResultBundle',
  full_name='luci.resultdb.v1.TestResultBundle',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='luci.resultdb.v1.TestResultBundle.result', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=567,
  serialized_end=631,
)


_TESTVARIANTPREDICATE = _descriptor.Descriptor(
  name='TestVariantPredicate',
  full_name='luci.resultdb.v1.TestVariantPredicate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='luci.resultdb.v1.TestVariantPredicate.status', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=633,
  serialized_end=708,
)

_TESTVARIANT.fields_by_name['variant'].message_type = go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_common__pb2._VARIANT
_TESTVARIANT.fields_by_name['status'].enum_type = _TESTVARIANTSTATUS
_TESTVARIANT.fields_by_name['results'].message_type = _TESTRESULTBUNDLE
_TESTVARIANT.fields_by_name['exonerations'].message_type = go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_test__result__pb2._TESTEXONERATION
_TESTVARIANT.fields_by_name['test_metadata'].message_type = go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_test__metadata__pb2._TESTMETADATA
_TESTRESULTBUNDLE.fields_by_name['result'].message_type = go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_test__result__pb2._TESTRESULT
_TESTVARIANTPREDICATE.fields_by_name['status'].enum_type = _TESTVARIANTSTATUS
DESCRIPTOR.message_types_by_name['TestVariant'] = _TESTVARIANT
DESCRIPTOR.message_types_by_name['TestResultBundle'] = _TESTRESULTBUNDLE
DESCRIPTOR.message_types_by_name['TestVariantPredicate'] = _TESTVARIANTPREDICATE
DESCRIPTOR.enum_types_by_name['TestVariantStatus'] = _TESTVARIANTSTATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TestVariant = _reflection.GeneratedProtocolMessageType('TestVariant', (_message.Message,), {
  'DESCRIPTOR' : _TESTVARIANT,
  '__module__' : 'go.chromium.org.luci.resultdb.proto.v1.test_variant_pb2'
  # @@protoc_insertion_point(class_scope:luci.resultdb.v1.TestVariant)
  })
_sym_db.RegisterMessage(TestVariant)

TestResultBundle = _reflection.GeneratedProtocolMessageType('TestResultBundle', (_message.Message,), {
  'DESCRIPTOR' : _TESTRESULTBUNDLE,
  '__module__' : 'go.chromium.org.luci.resultdb.proto.v1.test_variant_pb2'
  # @@protoc_insertion_point(class_scope:luci.resultdb.v1.TestResultBundle)
  })
_sym_db.RegisterMessage(TestResultBundle)

TestVariantPredicate = _reflection.GeneratedProtocolMessageType('TestVariantPredicate', (_message.Message,), {
  'DESCRIPTOR' : _TESTVARIANTPREDICATE,
  '__module__' : 'go.chromium.org.luci.resultdb.proto.v1.test_variant_pb2'
  # @@protoc_insertion_point(class_scope:luci.resultdb.v1.TestVariantPredicate)
  })
_sym_db.RegisterMessage(TestVariantPredicate)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
