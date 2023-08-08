// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: infra/appengine/cros/lab_inventory/api/bigquery/hwid_server.proto

package apibq

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Next Tag: 4
type HWIDInventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hwid        string                 `protobuf:"bytes,1,opt,name=hwid,proto3" json:"hwid,omitempty"`
	Sku         string                 `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
}

func (x *HWIDInventory) Reset() {
	*x = HWIDInventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HWIDInventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HWIDInventory) ProtoMessage() {}

func (x *HWIDInventory) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HWIDInventory.ProtoReflect.Descriptor instead.
func (*HWIDInventory) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDescGZIP(), []int{0}
}

func (x *HWIDInventory) GetHwid() string {
	if x != nil {
		return x.Hwid
	}
	return ""
}

func (x *HWIDInventory) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *HWIDInventory) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

var File_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto protoreflect.FileDescriptor

var file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDesc = []byte{
	0x0a, 0x41, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2f, 0x68, 0x77, 0x69, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x70, 0x69, 0x62, 0x71, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x0d, 0x48,
	0x57, 0x49, 0x44, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x77, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x77, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x6b, 0x75, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x37, 0x5a, 0x35, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDescOnce sync.Once
	file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDescData = file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDesc
)

func file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDescGZIP() []byte {
	file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDescOnce.Do(func() {
		file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDescData)
	})
	return file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDescData
}

var file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_goTypes = []interface{}{
	(*HWIDInventory)(nil),         // 0: apibq.HWIDInventory
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_depIdxs = []int32{
	1, // 0: apibq.HWIDInventory.updated_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_init() }
func file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_init() {
	if File_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HWIDInventory); i {
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
			RawDescriptor: file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_goTypes,
		DependencyIndexes: file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_depIdxs,
		MessageInfos:      file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_msgTypes,
	}.Build()
	File_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto = out.File
	file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_rawDesc = nil
	file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_goTypes = nil
	file_infra_appengine_cros_lab_inventory_api_bigquery_hwid_server_proto_depIdxs = nil
}
