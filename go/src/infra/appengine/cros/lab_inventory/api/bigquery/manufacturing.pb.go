// Copyright 2019 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: infra/appengine/cros/lab_inventory/api/bigquery/manufacturing.proto

package apibq

import (
	manufacturing "go.chromium.org/chromiumos/infra/proto/go/manufacturing"
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

type ManufacturingInventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManufacturingId string                 `protobuf:"bytes,1,opt,name=manufacturing_id,json=manufacturingId,proto3" json:"manufacturing_id,omitempty"`
	Config          *manufacturing.Config  `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	UpdatedTime     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
}

func (x *ManufacturingInventory) Reset() {
	*x = ManufacturingInventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManufacturingInventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManufacturingInventory) ProtoMessage() {}

func (x *ManufacturingInventory) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManufacturingInventory.ProtoReflect.Descriptor instead.
func (*ManufacturingInventory) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDescGZIP(), []int{0}
}

func (x *ManufacturingInventory) GetManufacturingId() string {
	if x != nil {
		return x.ManufacturingId
	}
	return ""
}

func (x *ManufacturingInventory) GetConfig() *manufacturing.Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ManufacturingInventory) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

var File_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto protoreflect.FileDescriptor

var file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDesc = []byte{
	0x0a, 0x43, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2f, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x70, 0x69, 0x62, 0x71, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61,
	0x63, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x16, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63,
	0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x29, 0x0a, 0x10, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x61, 0x6e,
	0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x63, 0x72, 0x6f, 0x73,
	0x2f, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3b, 0x61, 0x70, 0x69, 0x62,
	0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDescOnce sync.Once
	file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDescData = file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDesc
)

func file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDescGZIP() []byte {
	file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDescOnce.Do(func() {
		file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDescData)
	})
	return file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDescData
}

var file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_goTypes = []interface{}{
	(*ManufacturingInventory)(nil), // 0: apibq.ManufacturingInventory
	(*manufacturing.Config)(nil),   // 1: manufacturing.Config
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_depIdxs = []int32{
	1, // 0: apibq.ManufacturingInventory.config:type_name -> manufacturing.Config
	2, // 1: apibq.ManufacturingInventory.updated_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_init() }
func file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_init() {
	if File_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManufacturingInventory); i {
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
			RawDescriptor: file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_goTypes,
		DependencyIndexes: file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_depIdxs,
		MessageInfos:      file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_msgTypes,
	}.Build()
	File_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto = out.File
	file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_rawDesc = nil
	file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_goTypes = nil
	file_infra_appengine_cros_lab_inventory_api_bigquery_manufacturing_proto_depIdxs = nil
}
