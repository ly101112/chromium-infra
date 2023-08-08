// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Copied from google3/google/internal/chromeos/hwid/v2/hwid_service.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/unifiedfleet/api/v1/models/hwid.proto

package ufspb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The response message for `HwidService.GetDutLabel`.
type GetDutLabelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The DUT labels decoded from the HWID string.
	DutLabel *DutLabel `protobuf:"bytes,1,opt,name=dut_label,json=dutLabel,proto3" json:"dut_label,omitempty"`
}

func (x *GetDutLabelResponse) Reset() {
	*x = GetDutLabelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDutLabelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDutLabelResponse) ProtoMessage() {}

func (x *GetDutLabelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDutLabelResponse.ProtoReflect.Descriptor instead.
func (*GetDutLabelResponse) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescGZIP(), []int{0}
}

func (x *GetDutLabelResponse) GetDutLabel() *DutLabel {
	if x != nil {
		return x.DutLabel
	}
	return nil
}

// A set of labels representing the features of the device, can be revealed
// by decoding the HWID string.
type DutLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Possible labels in the labels field.
	PossibleLabels []string          `protobuf:"bytes,1,rep,name=possible_labels,json=possibleLabels,proto3" json:"possible_labels,omitempty"`
	Labels         []*DutLabel_Label `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"` // NOLINT
}

func (x *DutLabel) Reset() {
	*x = DutLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DutLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DutLabel) ProtoMessage() {}

func (x *DutLabel) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DutLabel.ProtoReflect.Descriptor instead.
func (*DutLabel) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescGZIP(), []int{1}
}

func (x *DutLabel) GetPossibleLabels() []string {
	if x != nil {
		return x.PossibleLabels
	}
	return nil
}

func (x *DutLabel) GetLabels() []*DutLabel_Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

type HwidData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku      string    `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Variant  string    `protobuf:"bytes,2,opt,name=variant,proto3" json:"variant,omitempty"`
	Hwid     string    `protobuf:"bytes,3,opt,name=hwid,proto3" json:"hwid,omitempty"`
	DutLabel *DutLabel `protobuf:"bytes,4,opt,name=dut_label,json=dutLabel,proto3" json:"dut_label,omitempty"`
}

func (x *HwidData) Reset() {
	*x = HwidData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HwidData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HwidData) ProtoMessage() {}

func (x *HwidData) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HwidData.ProtoReflect.Descriptor instead.
func (*HwidData) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescGZIP(), []int{2}
}

func (x *HwidData) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *HwidData) GetVariant() string {
	if x != nil {
		return x.Variant
	}
	return ""
}

func (x *HwidData) GetHwid() string {
	if x != nil {
		return x.Hwid
	}
	return ""
}

func (x *HwidData) GetDutLabel() *DutLabel {
	if x != nil {
		return x.DutLabel
	}
	return nil
}

// All labels extracted from the HWID string.
type DutLabel_Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the label, will always exist in the field of `possible_labels`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The value of this label.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DutLabel_Label) Reset() {
	*x = DutLabel_Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DutLabel_Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DutLabel_Label) ProtoMessage() {}

func (x *DutLabel_Label) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DutLabel_Label.ProtoReflect.Descriptor instead.
func (*DutLabel_Label) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DutLabel_Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DutLabel_Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_infra_unifiedfleet_api_v1_models_hwid_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x68, 0x77, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x44, 0x75, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x09, 0x64, 0x75, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x44, 0x75, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x08, 0x64, 0x75, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x22, 0xb4, 0x01, 0x0a, 0x08, 0x44, 0x75, 0x74, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x75, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x3b,
	0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x08,
	0x48, 0x77, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x77, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x77, 0x69, 0x64, 0x12, 0x41, 0x0a, 0x09, 0x64, 0x75, 0x74, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x75, 0x6e,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x75, 0x74, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x08, 0x64, 0x75, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x28, 0x5a, 0x26, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b,
	0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_infra_unifiedfleet_api_v1_models_hwid_proto_goTypes = []interface{}{
	(*GetDutLabelResponse)(nil), // 0: unifiedfleet.api.v1.models.GetDutLabelResponse
	(*DutLabel)(nil),            // 1: unifiedfleet.api.v1.models.DutLabel
	(*HwidData)(nil),            // 2: unifiedfleet.api.v1.models.HwidData
	(*DutLabel_Label)(nil),      // 3: unifiedfleet.api.v1.models.DutLabel.Label
}
var file_infra_unifiedfleet_api_v1_models_hwid_proto_depIdxs = []int32{
	1, // 0: unifiedfleet.api.v1.models.GetDutLabelResponse.dut_label:type_name -> unifiedfleet.api.v1.models.DutLabel
	3, // 1: unifiedfleet.api.v1.models.DutLabel.labels:type_name -> unifiedfleet.api.v1.models.DutLabel.Label
	1, // 2: unifiedfleet.api.v1.models.HwidData.dut_label:type_name -> unifiedfleet.api.v1.models.DutLabel
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_hwid_proto_init() }
func file_infra_unifiedfleet_api_v1_models_hwid_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_hwid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDutLabelResponse); i {
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
		file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DutLabel); i {
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
		file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HwidData); i {
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
		file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DutLabel_Label); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_hwid_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_hwid_proto_depIdxs,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_hwid_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_hwid_proto = out.File
	file_infra_unifiedfleet_api_v1_models_hwid_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_hwid_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_hwid_proto_depIdxs = nil
}
