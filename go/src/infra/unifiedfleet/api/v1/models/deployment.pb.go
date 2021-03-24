// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This proto definition describes some basic units in deployment.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: infra/unifiedfleet/api/v1/models/deployment.proto

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

// Next tag: 3
type PushConfigType int32

const (
	PushConfigType_PUSH_CONFIG_TYPE_UNSPECIFIED  PushConfigType = 0
	PushConfigType_PUSH_CONFIG_TYPE_MOBILECONFIG PushConfigType = 1
	PushConfigType_PUSH_CONFIG_TYPE_PACKAGE      PushConfigType = 2
)

// Enum value maps for PushConfigType.
var (
	PushConfigType_name = map[int32]string{
		0: "PUSH_CONFIG_TYPE_UNSPECIFIED",
		1: "PUSH_CONFIG_TYPE_MOBILECONFIG",
		2: "PUSH_CONFIG_TYPE_PACKAGE",
	}
	PushConfigType_value = map[string]int32{
		"PUSH_CONFIG_TYPE_UNSPECIFIED":  0,
		"PUSH_CONFIG_TYPE_MOBILECONFIG": 1,
		"PUSH_CONFIG_TYPE_PACKAGE":      2,
	}
)

func (x PushConfigType) Enum() *PushConfigType {
	p := new(PushConfigType)
	*p = x
	return p
}

func (x PushConfigType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushConfigType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_unifiedfleet_api_v1_models_deployment_proto_enumTypes[0].Descriptor()
}

func (PushConfigType) Type() protoreflect.EnumType {
	return &file_infra_unifiedfleet_api_v1_models_deployment_proto_enumTypes[0]
}

func (x PushConfigType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushConfigType.Descriptor instead.
func (PushConfigType) EnumDescriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDescGZIP(), []int{0}
}

// Payload describes the payload needed in deployment.
//
// Next tag: 3
type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// It refers to the unique payload name,
	// e.g. energy_saver, disable_gatekeeper, puppet3_package, ..., etc.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// e.g. PUSH_CONFIG_TYPE_MOBILECONFIG
	Config PushConfigType `protobuf:"varint,2,opt,name=config,proto3,enum=unifiedfleet.api.v1.models.PushConfigType" json:"config,omitempty"`
	// It refers to a g3 path for a profile, e.g.
	//      //depot/google3/ops/macops/mdm/mega/chrome/profiles/energy_saver.mobileconfig
	// or refers to a server link for a package, e.g.
	//      https://macos-server-1.golo.chromium.org/manifests/clpuppet3.plist
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Payload) GetConfig() PushConfigType {
	if x != nil {
		return x.Config
	}
	return PushConfigType_PUSH_CONFIG_TYPE_UNSPECIFIED
}

func (x *Payload) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_infra_unifiedfleet_api_v1_models_deployment_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22,
	0x75, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a,
	0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x2a, 0x73, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x55, 0x53, 0x48,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x55,
	0x53, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x4f, 0x42, 0x49, 0x4c, 0x45, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x01, 0x12, 0x1c, 0x0a,
	0x18, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x50, 0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x10, 0x02, 0x42, 0x28, 0x5a, 0x26, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b,
	0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_unifiedfleet_api_v1_models_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_unifiedfleet_api_v1_models_deployment_proto_goTypes = []interface{}{
	(PushConfigType)(0), // 0: unifiedfleet.api.v1.models.PushConfigType
	(*Payload)(nil),     // 1: unifiedfleet.api.v1.models.Payload
}
var file_infra_unifiedfleet_api_v1_models_deployment_proto_depIdxs = []int32{
	0, // 0: unifiedfleet.api.v1.models.Payload.config:type_name -> unifiedfleet.api.v1.models.PushConfigType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_deployment_proto_init() }
func file_infra_unifiedfleet_api_v1_models_deployment_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_deployment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_deployment_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_deployment_proto_depIdxs,
		EnumInfos:         file_infra_unifiedfleet_api_v1_models_deployment_proto_enumTypes,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_deployment_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_deployment_proto = out.File
	file_infra_unifiedfleet_api_v1_models_deployment_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_deployment_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_deployment_proto_depIdxs = nil
}
