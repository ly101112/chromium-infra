// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/unifiedfleet/api/v1/models/change_event.proto

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

// represents the ChangeEvent generated when there is any change for the asset
type ChangeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The string representation of the changed (multi-component) field paths,
	// See explanation of the field path here:
	// https://github.com/protocolbuffers/protobuf/blob/50e03cdde3ef1fc7e9674db0a98ee1dea93f6fb2/src/google/protobuf/field_mask.proto#L43
	// machine.serial_number, chromeos_machine.model, dut.config.peripherals.wifi.wificell,
	// peripheral_requirement.min
	EventLabel string `protobuf:"bytes,2,opt,name=event_label,json=eventLabel,proto3" json:"event_label,omitempty"`
	// The string representation of the changed item, e.g.
	// machine.serial_number: from "" => A
	// chromeos_machine.model: from modelA => modelB
	// dut.config.peripherals.wifi.wificell: from false => true
	// periphral_requirement.min: from 1 => 3
	OldValue string `protobuf:"bytes,3,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	NewValue string `protobuf:"bytes,4,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	// Record the last update timestamp of this Event (In UTC timezone)
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	UserEmail  string                 `protobuf:"bytes,6,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Comment    string                 `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *ChangeEvent) Reset() {
	*x = ChangeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_change_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeEvent) ProtoMessage() {}

func (x *ChangeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_change_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeEvent.ProtoReflect.Descriptor instead.
func (*ChangeEvent) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeEvent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChangeEvent) GetEventLabel() string {
	if x != nil {
		return x.EventLabel
	}
	return ""
}

func (x *ChangeEvent) GetOldValue() string {
	if x != nil {
		return x.OldValue
	}
	return ""
}

func (x *ChangeEvent) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

func (x *ChangeEvent) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *ChangeEvent) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *ChangeEvent) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

var File_infra_unifiedfleet_api_v1_models_change_event_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4,
	0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3b, 0xea, 0x41, 0x38, 0x0a, 0x26, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x7d, 0x42, 0x28, 0x5a, 0x26, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b, 0x75, 0x66, 0x73, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_change_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_unifiedfleet_api_v1_models_change_event_proto_goTypes = []interface{}{
	(*ChangeEvent)(nil),           // 0: unifiedfleet.api.v1.models.ChangeEvent
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_infra_unifiedfleet_api_v1_models_change_event_proto_depIdxs = []int32{
	1, // 0: unifiedfleet.api.v1.models.ChangeEvent.update_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_change_event_proto_init() }
func file_infra_unifiedfleet_api_v1_models_change_event_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_change_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_change_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeEvent); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_change_event_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_change_event_proto_depIdxs,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_change_event_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_change_event_proto = out.File
	file_infra_unifiedfleet_api_v1_models_change_event_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_change_event_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_change_event_proto_depIdxs = nil
}
