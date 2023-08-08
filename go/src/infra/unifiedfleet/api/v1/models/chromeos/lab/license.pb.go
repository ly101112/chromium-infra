// Copyright 2020 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/unifiedfleet/api/v1/models/chromeos/lab/license.proto

package ufspb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Next Tag: 3
type LicenseType int32

const (
	LicenseType_LICENSE_TYPE_UNSPECIFIED LicenseType = 0
	// Microsoft Windows 10 Professional Desktop Operating System.
	// Contact the Chrome OS Parallels team for license specifics.
	LicenseType_LICENSE_TYPE_WINDOWS_10_PRO LicenseType = 1
	// Microsoft Office Standard.
	// Contact the Chrome OS Parallels team for license specifics.
	LicenseType_LICENSE_TYPE_MS_OFFICE_STANDARD LicenseType = 2
)

// Enum value maps for LicenseType.
var (
	LicenseType_name = map[int32]string{
		0: "LICENSE_TYPE_UNSPECIFIED",
		1: "LICENSE_TYPE_WINDOWS_10_PRO",
		2: "LICENSE_TYPE_MS_OFFICE_STANDARD",
	}
	LicenseType_value = map[string]int32{
		"LICENSE_TYPE_UNSPECIFIED":        0,
		"LICENSE_TYPE_WINDOWS_10_PRO":     1,
		"LICENSE_TYPE_MS_OFFICE_STANDARD": 2,
	}
)

func (x LicenseType) Enum() *LicenseType {
	p := new(LicenseType)
	*p = x
	return p
}

func (x LicenseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LicenseType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_enumTypes[0].Descriptor()
}

func (LicenseType) Type() protoreflect.EnumType {
	return &file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_enumTypes[0]
}

func (x LicenseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LicenseType.Descriptor instead.
func (LicenseType) EnumDescriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDescGZIP(), []int{0}
}

// Represents a Software License assigned to a device.
// Next Tag: 3
type License struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type LicenseType `protobuf:"varint,1,opt,name=type,proto3,enum=unifiedfleet.api.v1.models.chromeos.lab.LicenseType" json:"type,omitempty"`
	// An optional string to uniquely identify the license that was assigned,
	// for tracking purposes.
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *License) Reset() {
	*x = License{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *License) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*License) ProtoMessage() {}

func (x *License) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use License.ProtoReflect.Descriptor instead.
func (*License) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDescGZIP(), []int{0}
}

func (x *License) GetType() LicenseType {
	if x != nil {
		return x.Type
	}
	return LicenseType_LICENSE_TYPE_UNSPECIFIED
}

func (x *License) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

var File_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x2f,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65,
	0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x62, 0x22, 0x73, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x12, 0x48, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x34, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x62, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2a, 0x71, 0x0a, 0x0b, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x49,
	0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x4c, 0x49, 0x43, 0x45,
	0x4e, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53,
	0x5f, 0x31, 0x30, 0x5f, 0x50, 0x52, 0x4f, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x4c, 0x49, 0x43,
	0x45, 0x4e, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x53, 0x5f, 0x4f, 0x46, 0x46,
	0x49, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x02, 0x42, 0x35,
	0x5a, 0x33, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x3b,
	0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_goTypes = []interface{}{
	(LicenseType)(0), // 0: unifiedfleet.api.v1.models.chromeos.lab.LicenseType
	(*License)(nil),  // 1: unifiedfleet.api.v1.models.chromeos.lab.License
}
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_depIdxs = []int32{
	0, // 0: unifiedfleet.api.v1.models.chromeos.lab.License.type:type_name -> unifiedfleet.api.v1.models.chromeos.lab.LicenseType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_init() }
func file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*License); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_depIdxs,
		EnumInfos:         file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_enumTypes,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto = out.File
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_license_proto_depIdxs = nil
}
