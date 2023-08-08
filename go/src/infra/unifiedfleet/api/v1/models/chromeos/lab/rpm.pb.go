// Copyright 2022 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/unifiedfleet/api/v1/models/chromeos/lab/rpm.proto

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

// Remote power management info.
// Next Tag: 3
type OSRPM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PowerunitName   string `protobuf:"bytes,1,opt,name=powerunit_name,json=powerunitName,proto3" json:"powerunit_name,omitempty"`
	PowerunitOutlet string `protobuf:"bytes,2,opt,name=powerunit_outlet,json=powerunitOutlet,proto3" json:"powerunit_outlet,omitempty"`
}

func (x *OSRPM) Reset() {
	*x = OSRPM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OSRPM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSRPM) ProtoMessage() {}

func (x *OSRPM) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSRPM.ProtoReflect.Descriptor instead.
func (*OSRPM) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDescGZIP(), []int{0}
}

func (x *OSRPM) GetPowerunitName() string {
	if x != nil {
		return x.PowerunitName
	}
	return ""
}

func (x *OSRPM) GetPowerunitOutlet() string {
	if x != nil {
		return x.PowerunitOutlet
	}
	return ""
}

var File_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDesc = []byte{
	0x0a, 0x37, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x2f,
	0x72, 0x70, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x75, 0x6e, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6c,
	0x61, 0x62, 0x22, 0x59, 0x0a, 0x05, 0x4f, 0x53, 0x52, 0x50, 0x4d, 0x12, 0x25, 0x0a, 0x0e, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x6f, 0x75, 0x74, 0x6c, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6f,
	0x77, 0x65, 0x72, 0x75, 0x6e, 0x69, 0x74, 0x4f, 0x75, 0x74, 0x6c, 0x65, 0x74, 0x42, 0x35, 0x5a,
	0x33, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x3b, 0x75,
	0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_goTypes = []interface{}{
	(*OSRPM)(nil), // 0: unifiedfleet.api.v1.models.chromeos.lab.OSRPM
}
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_init() }
func file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OSRPM); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_depIdxs,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto = out.File
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_rpm_proto_depIdxs = nil
}
