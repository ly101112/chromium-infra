// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: recovery_version.proto

package models

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

// RecoveryVersion represents stable version information to be used for repair
type RecoveryVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Board     string `protobuf:"bytes,1,opt,name=board,proto3" json:"board,omitempty"`
	Model     string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	OsImage   string `protobuf:"bytes,3,opt,name=os_image,json=osImage,proto3" json:"os_image,omitempty"`
	FwVersion string `protobuf:"bytes,4,opt,name=fw_version,json=fwVersion,proto3" json:"fw_version,omitempty"`
	FwImage   string `protobuf:"bytes,5,opt,name=fw_image,json=fwImage,proto3" json:"fw_image,omitempty"`
}

func (x *RecoveryVersion) Reset() {
	*x = RecoveryVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recovery_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoveryVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoveryVersion) ProtoMessage() {}

func (x *RecoveryVersion) ProtoReflect() protoreflect.Message {
	mi := &file_recovery_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoveryVersion.ProtoReflect.Descriptor instead.
func (*RecoveryVersion) Descriptor() ([]byte, []int) {
	return file_recovery_version_proto_rawDescGZIP(), []int{0}
}

func (x *RecoveryVersion) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

func (x *RecoveryVersion) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *RecoveryVersion) GetOsImage() string {
	if x != nil {
		return x.OsImage
	}
	return ""
}

func (x *RecoveryVersion) GetFwVersion() string {
	if x != nil {
		return x.FwVersion
	}
	return ""
}

func (x *RecoveryVersion) GetFwImage() string {
	if x != nil {
		return x.FwImage
	}
	return ""
}

var File_recovery_version_proto protoreflect.FileDescriptor

var file_recovery_version_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65,
	0x6f, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x22, 0x92, 0x01, 0x0a, 0x0f,
	0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x73, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x77, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x77, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x77, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x77, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x42, 0x1d, 0x5a, 0x1b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recovery_version_proto_rawDescOnce sync.Once
	file_recovery_version_proto_rawDescData = file_recovery_version_proto_rawDesc
)

func file_recovery_version_proto_rawDescGZIP() []byte {
	file_recovery_version_proto_rawDescOnce.Do(func() {
		file_recovery_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_recovery_version_proto_rawDescData)
	})
	return file_recovery_version_proto_rawDescData
}

var file_recovery_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_recovery_version_proto_goTypes = []interface{}{
	(*RecoveryVersion)(nil), // 0: chromeos.recovery.RecoveryVersion
}
var file_recovery_version_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_recovery_version_proto_init() }
func file_recovery_version_proto_init() {
	if File_recovery_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recovery_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoveryVersion); i {
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
			RawDescriptor: file_recovery_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_recovery_version_proto_goTypes,
		DependencyIndexes: file_recovery_version_proto_depIdxs,
		MessageInfos:      file_recovery_version_proto_msgTypes,
	}.Build()
	File_recovery_version_proto = out.File
	file_recovery_version_proto_rawDesc = nil
	file_recovery_version_proto_goTypes = nil
	file_recovery_version_proto_depIdxs = nil
}
