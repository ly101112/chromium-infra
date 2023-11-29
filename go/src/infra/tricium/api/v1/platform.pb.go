// Copyright 2017 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/tricium/api/v1/platform.proto

package tricium

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

// Enum names of supported platforms.
//
// The original idea was that these platforms could be used to describe either
// data dependencies or runtime platforms; this is not used, because the
// runtime platform is actually determined by builder config.
//
// In practice, the only platforms used are LINUX (and sometimes UBUNTU).
//
// Names must not contain "_", since this is used as a separator character
// in worker names.
type Platform_Name int32

const (
	// Use for platform-independent data types.
	Platform_ANY Platform_Name = 0
	// Generic Linux.
	Platform_LINUX  Platform_Name = 1
	Platform_UBUNTU Platform_Name = 2 // reserved 3,4,5,6,7,8,9,10
	// Generic Android.
	//
	// Deprecated: Marked as deprecated in infra/tricium/api/v1/platform.proto.
	Platform_ANDROID Platform_Name = 11 // reserved 12,13,14,15
	// Generic Mac
	//
	// Deprecated: Marked as deprecated in infra/tricium/api/v1/platform.proto.
	Platform_MAC Platform_Name = 16
	// Deprecated: Marked as deprecated in infra/tricium/api/v1/platform.proto.
	Platform_OSX Platform_Name = 17
	// Deprecated: Marked as deprecated in infra/tricium/api/v1/platform.proto.
	Platform_IOS Platform_Name = 18 // reserved 19,20
	// Generic Windows.
	//
	// Deprecated: Marked as deprecated in infra/tricium/api/v1/platform.proto.
	Platform_WINDOWS Platform_Name = 21 // reserved 22,23,24,25
	// Deprecated: Marked as deprecated in infra/tricium/api/v1/platform.proto.
	Platform_CHROMEOS Platform_Name = 26 // reserved 27,28,29,30
	// Deprecated: Marked as deprecated in infra/tricium/api/v1/platform.proto.
	Platform_FUCHSIA Platform_Name = 31 // reserved 32,33,34,35
)

// Enum value maps for Platform_Name.
var (
	Platform_Name_name = map[int32]string{
		0:  "ANY",
		1:  "LINUX",
		2:  "UBUNTU",
		11: "ANDROID",
		16: "MAC",
		17: "OSX",
		18: "IOS",
		21: "WINDOWS",
		26: "CHROMEOS",
		31: "FUCHSIA",
	}
	Platform_Name_value = map[string]int32{
		"ANY":      0,
		"LINUX":    1,
		"UBUNTU":   2,
		"ANDROID":  11,
		"MAC":      16,
		"OSX":      17,
		"IOS":      18,
		"WINDOWS":  21,
		"CHROMEOS": 26,
		"FUCHSIA":  31,
	}
)

func (x Platform_Name) Enum() *Platform_Name {
	p := new(Platform_Name)
	*p = x
	return p
}

func (x Platform_Name) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform_Name) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_tricium_api_v1_platform_proto_enumTypes[0].Descriptor()
}

func (Platform_Name) Type() protoreflect.EnumType {
	return &file_infra_tricium_api_v1_platform_proto_enumTypes[0]
}

func (x Platform_Name) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform_Name.Descriptor instead.
func (Platform_Name) EnumDescriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_platform_proto_rawDescGZIP(), []int{0, 0}
}

// Platforms supported by Tricium.
type Platform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Platform) Reset() {
	*x = Platform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform) ProtoMessage() {}

func (x *Platform) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform.ProtoReflect.Descriptor instead.
func (*Platform) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_platform_proto_rawDescGZIP(), []int{0}
}

// Platform details for supported platforms.
type Platform_Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name Platform_Name `protobuf:"varint,1,opt,name=name,proto3,enum=tricium.Platform_Name" json:"name,omitempty"`
	// Deprecated, ignored.
	//
	// Deprecated: Marked as deprecated in infra/tricium/api/v1/platform.proto.
	Dimensions []string `protobuf:"bytes,2,rep,name=dimensions,proto3" json:"dimensions,omitempty"`
	// Whether this platform can be used as a runtime platform.
	//
	// Not used, should always be true.
	HasRuntime bool `protobuf:"varint,3,opt,name=has_runtime,json=hasRuntime,proto3" json:"has_runtime,omitempty"`
}

func (x *Platform_Details) Reset() {
	*x = Platform_Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Platform_Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform_Details) ProtoMessage() {}

func (x *Platform_Details) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform_Details.ProtoReflect.Descriptor instead.
func (*Platform_Details) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_platform_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Platform_Details) GetName() Platform_Name {
	if x != nil {
		return x.Name
	}
	return Platform_ANY
}

// Deprecated: Marked as deprecated in infra/tricium/api/v1/platform.proto.
func (x *Platform_Details) GetDimensions() []string {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

func (x *Platform_Details) GetHasRuntime() bool {
	if x != nil {
		return x.HasRuntime
	}
	return false
}

var File_infra_tricium_api_v1_platform_proto protoreflect.FileDescriptor

var file_infra_tricium_api_v1_platform_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x22, 0x9b,
	0x02, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x1a, 0x7a, 0x0a, 0x07, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x64, 0x69, 0x6d, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x61, 0x73, 0x5f, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x68, 0x61, 0x73,
	0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e,
	0x55, 0x58, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x42, 0x55, 0x4e, 0x54, 0x55, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x0b, 0x1a, 0x02, 0x08,
	0x01, 0x12, 0x0b, 0x0a, 0x03, 0x4d, 0x41, 0x43, 0x10, 0x10, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0b,
	0x0a, 0x03, 0x4f, 0x53, 0x58, 0x10, 0x11, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0b, 0x0a, 0x03, 0x49,
	0x4f, 0x53, 0x10, 0x12, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0f, 0x0a, 0x07, 0x57, 0x49, 0x4e, 0x44,
	0x4f, 0x57, 0x53, 0x10, 0x15, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x10, 0x0a, 0x08, 0x43, 0x48, 0x52,
	0x4f, 0x4d, 0x45, 0x4f, 0x53, 0x10, 0x1a, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0f, 0x0a, 0x07, 0x46,
	0x55, 0x43, 0x48, 0x53, 0x49, 0x41, 0x10, 0x1f, 0x1a, 0x02, 0x08, 0x01, 0x42, 0x1e, 0x5a, 0x1c,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_tricium_api_v1_platform_proto_rawDescOnce sync.Once
	file_infra_tricium_api_v1_platform_proto_rawDescData = file_infra_tricium_api_v1_platform_proto_rawDesc
)

func file_infra_tricium_api_v1_platform_proto_rawDescGZIP() []byte {
	file_infra_tricium_api_v1_platform_proto_rawDescOnce.Do(func() {
		file_infra_tricium_api_v1_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_tricium_api_v1_platform_proto_rawDescData)
	})
	return file_infra_tricium_api_v1_platform_proto_rawDescData
}

var file_infra_tricium_api_v1_platform_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_tricium_api_v1_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_tricium_api_v1_platform_proto_goTypes = []interface{}{
	(Platform_Name)(0),       // 0: tricium.Platform.Name
	(*Platform)(nil),         // 1: tricium.Platform
	(*Platform_Details)(nil), // 2: tricium.Platform.Details
}
var file_infra_tricium_api_v1_platform_proto_depIdxs = []int32{
	0, // 0: tricium.Platform.Details.name:type_name -> tricium.Platform.Name
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_tricium_api_v1_platform_proto_init() }
func file_infra_tricium_api_v1_platform_proto_init() {
	if File_infra_tricium_api_v1_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_tricium_api_v1_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform); i {
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
		file_infra_tricium_api_v1_platform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Platform_Details); i {
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
			RawDescriptor: file_infra_tricium_api_v1_platform_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_tricium_api_v1_platform_proto_goTypes,
		DependencyIndexes: file_infra_tricium_api_v1_platform_proto_depIdxs,
		EnumInfos:         file_infra_tricium_api_v1_platform_proto_enumTypes,
		MessageInfos:      file_infra_tricium_api_v1_platform_proto_msgTypes,
	}.Build()
	File_infra_tricium_api_v1_platform_proto = out.File
	file_infra_tricium_api_v1_platform_proto_rawDesc = nil
	file_infra_tricium_api_v1_platform_proto_goTypes = nil
	file_infra_tricium_api_v1_platform_proto_depIdxs = nil
}
