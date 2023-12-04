// Copyright 2023 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: infra/tools/dirmd/proto/component_def.proto

package dirmdpb

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

// TODO(crbug.com/1505875) - Remove this once migration is complete.
// Note that these definitions are separate from the ones in dir_metadata.proto
// as these are for compatibility to the mapping definitions provided.
type ComponentsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComponentDef []*ComponentMapping `protobuf:"bytes,1,rep,name=component_def,json=componentDef,proto3" json:"component_def,omitempty"`
}

func (x *ComponentsConfig) Reset() {
	*x = ComponentsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tools_dirmd_proto_component_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentsConfig) ProtoMessage() {}

func (x *ComponentsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tools_dirmd_proto_component_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentsConfig.ProtoReflect.Descriptor instead.
func (*ComponentsConfig) Descriptor() ([]byte, []int) {
	return file_infra_tools_dirmd_proto_component_def_proto_rawDescGZIP(), []int{0}
}

func (x *ComponentsConfig) GetComponentDef() []*ComponentMapping {
	if x != nil {
		return x.ComponentDef
	}
	return nil
}

type ComponentMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Monorail component path. E.g. "Applications>repo".
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// ID for the corresponding component in Buganizer.
	BuganizerId int64 `protobuf:"varint,2,opt,name=buganizer_id,json=buganizerId,proto3" json:"buganizer_id,omitempty"`
}

func (x *ComponentMapping) Reset() {
	*x = ComponentMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tools_dirmd_proto_component_def_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentMapping) ProtoMessage() {}

func (x *ComponentMapping) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tools_dirmd_proto_component_def_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentMapping.ProtoReflect.Descriptor instead.
func (*ComponentMapping) Descriptor() ([]byte, []int) {
	return file_infra_tools_dirmd_proto_component_def_proto_rawDescGZIP(), []int{1}
}

func (x *ComponentMapping) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ComponentMapping) GetBuganizerId() int64 {
	if x != nil {
		return x.BuganizerId
	}
	return 0
}

var File_infra_tools_dirmd_proto_component_def_proto protoreflect.FileDescriptor

var file_infra_tools_dirmd_proto_component_def_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x64, 0x69,
	0x72, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x64, 0x69, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x5e, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4a, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x64, 0x69, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x66, 0x22, 0x49, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x75,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x62, 0x75, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x49, 0x64, 0x42, 0x21, 0x5a,
	0x1f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x64, 0x69, 0x72,
	0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x64, 0x69, 0x72, 0x6d, 0x64, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_tools_dirmd_proto_component_def_proto_rawDescOnce sync.Once
	file_infra_tools_dirmd_proto_component_def_proto_rawDescData = file_infra_tools_dirmd_proto_component_def_proto_rawDesc
)

func file_infra_tools_dirmd_proto_component_def_proto_rawDescGZIP() []byte {
	file_infra_tools_dirmd_proto_component_def_proto_rawDescOnce.Do(func() {
		file_infra_tools_dirmd_proto_component_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_tools_dirmd_proto_component_def_proto_rawDescData)
	})
	return file_infra_tools_dirmd_proto_component_def_proto_rawDescData
}

var file_infra_tools_dirmd_proto_component_def_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_tools_dirmd_proto_component_def_proto_goTypes = []interface{}{
	(*ComponentsConfig)(nil), // 0: chrome.dir_metadata.ComponentsConfig
	(*ComponentMapping)(nil), // 1: chrome.dir_metadata.ComponentMapping
}
var file_infra_tools_dirmd_proto_component_def_proto_depIdxs = []int32{
	1, // 0: chrome.dir_metadata.ComponentsConfig.component_def:type_name -> chrome.dir_metadata.ComponentMapping
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_tools_dirmd_proto_component_def_proto_init() }
func file_infra_tools_dirmd_proto_component_def_proto_init() {
	if File_infra_tools_dirmd_proto_component_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_tools_dirmd_proto_component_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentsConfig); i {
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
		file_infra_tools_dirmd_proto_component_def_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentMapping); i {
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
			RawDescriptor: file_infra_tools_dirmd_proto_component_def_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_tools_dirmd_proto_component_def_proto_goTypes,
		DependencyIndexes: file_infra_tools_dirmd_proto_component_def_proto_depIdxs,
		MessageInfos:      file_infra_tools_dirmd_proto_component_def_proto_msgTypes,
	}.Build()
	File_infra_tools_dirmd_proto_component_def_proto = out.File
	file_infra_tools_dirmd_proto_component_def_proto_rawDesc = nil
	file_infra_tools_dirmd_proto_component_def_proto_goTypes = nil
	file_infra_tools_dirmd_proto_component_def_proto_depIdxs = nil
}