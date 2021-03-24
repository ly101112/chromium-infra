// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: infra/unifiedfleet/api/v1/models/caching_service.proto

package ufspb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// CachingService refers to caching service information in the chromeos lab
//
// A lab can have multiple caching services running. Each caching service has a primary node
// and a secondary node and serve a particular subnet.
// UFS stores the information of all caching services available in the labs.
type CachingService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// caching service name
	// format will be 'cachingservices/{ipv4}'
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// port info of the caching service
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// particular subnet which the caching service serves/supports
	ServingSubnet string `protobuf:"bytes,3,opt,name=serving_subnet,json=servingSubnet,proto3" json:"serving_subnet,omitempty"`
	// ipv4 address of the primary node of the caching service
	PrimaryNode string `protobuf:"bytes,4,opt,name=primary_node,json=primaryNode,proto3" json:"primary_node,omitempty"`
	// ipv4 address of the secondary node of the caching service
	SecondaryNode string `protobuf:"bytes,5,opt,name=secondary_node,json=secondaryNode,proto3" json:"secondary_node,omitempty"`
	// state of the caching service
	State State `protobuf:"varint,6,opt,name=state,proto3,enum=unifiedfleet.api.v1.models.State" json:"state,omitempty"`
	// description of the caching service
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// record the last update timestamp of this caching service (In UTC timezone)
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *CachingService) Reset() {
	*x = CachingService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_caching_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CachingService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CachingService) ProtoMessage() {}

func (x *CachingService) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_caching_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CachingService.ProtoReflect.Descriptor instead.
func (*CachingService) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDescGZIP(), []int{0}
}

func (x *CachingService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CachingService) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *CachingService) GetServingSubnet() string {
	if x != nil {
		return x.ServingSubnet
	}
	return ""
}

func (x *CachingService) GetPrimaryNode() string {
	if x != nil {
		return x.PrimaryNode
	}
	return ""
}

func (x *CachingService) GetSecondaryNode() string {
	if x != nil {
		return x.SecondaryNode
	}
	return ""
}

func (x *CachingService) GetState() State {
	if x != nil {
		return x.State
	}
	return State_STATE_UNSPECIFIED
}

func (x *CachingService) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CachingService) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_infra_unifiedfleet_api_v1_models_caching_service_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDesc = []byte{
	0x0a, 0x36, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x94, 0x03, 0x0a, 0x0e, 0x43, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x4c, 0xea, 0x41, 0x49, 0x0a, 0x2f, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16,
	0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x70, 0x76, 0x34, 0x7d, 0x42, 0x28, 0x5a, 0x26, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f,
	0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b, 0x75, 0x66, 0x73, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_caching_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_unifiedfleet_api_v1_models_caching_service_proto_goTypes = []interface{}{
	(*CachingService)(nil),        // 0: unifiedfleet.api.v1.models.CachingService
	(State)(0),                    // 1: unifiedfleet.api.v1.models.State
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_infra_unifiedfleet_api_v1_models_caching_service_proto_depIdxs = []int32{
	1, // 0: unifiedfleet.api.v1.models.CachingService.state:type_name -> unifiedfleet.api.v1.models.State
	2, // 1: unifiedfleet.api.v1.models.CachingService.update_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_caching_service_proto_init() }
func file_infra_unifiedfleet_api_v1_models_caching_service_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_caching_service_proto != nil {
		return
	}
	file_infra_unifiedfleet_api_v1_models_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_caching_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CachingService); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_caching_service_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_caching_service_proto_depIdxs,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_caching_service_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_caching_service_proto = out.File
	file_infra_unifiedfleet_api_v1_models_caching_service_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_caching_service_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_caching_service_proto_depIdxs = nil
}
