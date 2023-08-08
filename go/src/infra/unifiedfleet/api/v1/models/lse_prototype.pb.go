// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/unifiedfleet/api/v1/models/lse_prototype.proto

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

// The supported peripheral type in LSE definition. The list is not completed
// as we have many special setups in ChromeOS high-touch labs. Will add them later
// when it comes to use.
type PeripheralType int32

const (
	PeripheralType_PERIPHERAL_TYPE_UNSPECIFIED       PeripheralType = 0
	PeripheralType_PERIPHERAL_TYPE_SERVO             PeripheralType = 1
	PeripheralType_PERIPHERAL_TYPE_LABSTATION        PeripheralType = 2
	PeripheralType_PERIPHERAL_TYPE_RPM               PeripheralType = 3
	PeripheralType_PERIPHERAL_TYPE_KVM               PeripheralType = 4
	PeripheralType_PERIPHERAL_TYPE_SWITCH            PeripheralType = 5
	PeripheralType_PERIPHERAL_TYPE_BLUETOOTH_BTPEERS PeripheralType = 6
	PeripheralType_PERIPHERAL_TYPE_WIFICELL          PeripheralType = 7
	PeripheralType_PERIPHERAL_TYPE_CAMERA            PeripheralType = 8
)

// Enum value maps for PeripheralType.
var (
	PeripheralType_name = map[int32]string{
		0: "PERIPHERAL_TYPE_UNSPECIFIED",
		1: "PERIPHERAL_TYPE_SERVO",
		2: "PERIPHERAL_TYPE_LABSTATION",
		3: "PERIPHERAL_TYPE_RPM",
		4: "PERIPHERAL_TYPE_KVM",
		5: "PERIPHERAL_TYPE_SWITCH",
		6: "PERIPHERAL_TYPE_BLUETOOTH_BTPEERS",
		7: "PERIPHERAL_TYPE_WIFICELL",
		8: "PERIPHERAL_TYPE_CAMERA",
	}
	PeripheralType_value = map[string]int32{
		"PERIPHERAL_TYPE_UNSPECIFIED":       0,
		"PERIPHERAL_TYPE_SERVO":             1,
		"PERIPHERAL_TYPE_LABSTATION":        2,
		"PERIPHERAL_TYPE_RPM":               3,
		"PERIPHERAL_TYPE_KVM":               4,
		"PERIPHERAL_TYPE_SWITCH":            5,
		"PERIPHERAL_TYPE_BLUETOOTH_BTPEERS": 6,
		"PERIPHERAL_TYPE_WIFICELL":          7,
		"PERIPHERAL_TYPE_CAMERA":            8,
	}
)

func (x PeripheralType) Enum() *PeripheralType {
	p := new(PeripheralType)
	*p = x
	return p
}

func (x PeripheralType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeripheralType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_enumTypes[0].Descriptor()
}

func (PeripheralType) Type() protoreflect.EnumType {
	return &file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_enumTypes[0]
}

func (x PeripheralType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeripheralType.Descriptor instead.
func (PeripheralType) EnumDescriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescGZIP(), []int{0}
}

// The supported virtual type in LSE definition.
type VirtualType int32

const (
	VirtualType_VIRTUAL_TYPE_UNSPECIFIED VirtualType = 0
	VirtualType_VIRTUAL_TYPE_VM          VirtualType = 1
)

// Enum value maps for VirtualType.
var (
	VirtualType_name = map[int32]string{
		0: "VIRTUAL_TYPE_UNSPECIFIED",
		1: "VIRTUAL_TYPE_VM",
	}
	VirtualType_value = map[string]int32{
		"VIRTUAL_TYPE_UNSPECIFIED": 0,
		"VIRTUAL_TYPE_VM":          1,
	}
)

func (x VirtualType) Enum() *VirtualType {
	p := new(VirtualType)
	*p = x
	return p
}

func (x VirtualType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VirtualType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_enumTypes[1].Descriptor()
}

func (VirtualType) Type() protoreflect.EnumType {
	return &file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_enumTypes[1]
}

func (x VirtualType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VirtualType.Descriptor instead.
func (VirtualType) EnumDescriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescGZIP(), []int{1}
}

type RackLSEPrototype struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique name for the RackLSEPrototype.
	// The format will be rackLSEPrototypes/XXX
	Name                   string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PeripheralRequirements []*PeripheralRequirement `protobuf:"bytes,2,rep,name=peripheral_requirements,json=peripheralRequirements,proto3" json:"peripheral_requirements,omitempty"`
	// Record the last update timestamp of this RackLSEPrototype (In UTC timezone)
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// tags user can attach for easy querying/search
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *RackLSEPrototype) Reset() {
	*x = RackLSEPrototype{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RackLSEPrototype) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RackLSEPrototype) ProtoMessage() {}

func (x *RackLSEPrototype) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RackLSEPrototype.ProtoReflect.Descriptor instead.
func (*RackLSEPrototype) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescGZIP(), []int{0}
}

func (x *RackLSEPrototype) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RackLSEPrototype) GetPeripheralRequirements() []*PeripheralRequirement {
	if x != nil {
		return x.PeripheralRequirements
	}
	return nil
}

func (x *RackLSEPrototype) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *RackLSEPrototype) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type MachineLSEPrototype struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique name for the MachineLSEPrototype.
	// The format will be machineLSEPrototypes/XXX
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// peripheral_requirements.peripheral_type must be unique.
	PeripheralRequirements []*PeripheralRequirement `protobuf:"bytes,2,rep,name=peripheral_requirements,json=peripheralRequirements,proto3" json:"peripheral_requirements,omitempty"`
	// Indicates the Rack Unit capacity of this setup, corresponding
	// to a Rack’s Rack Unit capacity.
	OccupiedCapacityRu int32 `protobuf:"varint,3,opt,name=occupied_capacity_ru,json=occupiedCapacityRu,proto3" json:"occupied_capacity_ru,omitempty"`
	// Record the last update timestamp of this MachineLSEPrototype (In UTC timezone)
	UpdateTime          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	VirtualRequirements []*VirtualRequirement  `protobuf:"bytes,5,rep,name=virtual_requirements,json=virtualRequirements,proto3" json:"virtual_requirements,omitempty"`
	// tags user can attach for easy querying/search
	Tags []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *MachineLSEPrototype) Reset() {
	*x = MachineLSEPrototype{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineLSEPrototype) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineLSEPrototype) ProtoMessage() {}

func (x *MachineLSEPrototype) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineLSEPrototype.ProtoReflect.Descriptor instead.
func (*MachineLSEPrototype) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescGZIP(), []int{1}
}

func (x *MachineLSEPrototype) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MachineLSEPrototype) GetPeripheralRequirements() []*PeripheralRequirement {
	if x != nil {
		return x.PeripheralRequirements
	}
	return nil
}

func (x *MachineLSEPrototype) GetOccupiedCapacityRu() int32 {
	if x != nil {
		return x.OccupiedCapacityRu
	}
	return 0
}

func (x *MachineLSEPrototype) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *MachineLSEPrototype) GetVirtualRequirements() []*VirtualRequirement {
	if x != nil {
		return x.VirtualRequirements
	}
	return nil
}

func (x *MachineLSEPrototype) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// The requirement for peripherals of a LSE. Usually it’s predefined
// by the designer of the test and lab, e.g. a test needs 2 cameras, 1 rpm,
// 1 servo, and a labstation.
// We probably also record cables as ChromeOS ACS lab wants to track the cable
// usage also.
type PeripheralRequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// It refers to the peripheral type that a LSE needs. The common use cases
	// include: kvm, switch, servo, rpm, labstation, camera, ...
	PeripheralType PeripheralType `protobuf:"varint,1,opt,name=peripheral_type,json=peripheralType,proto3,enum=unifiedfleet.api.v1.models.PeripheralType" json:"peripheral_type,omitempty"`
	// The minimum/maximum number of the peripherals that needed by a LSE, e.g.
	// A test needs 1-3 bluetooth bt peers to be set up.
	Min int32 `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max int32 `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *PeripheralRequirement) Reset() {
	*x = PeripheralRequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeripheralRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeripheralRequirement) ProtoMessage() {}

func (x *PeripheralRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeripheralRequirement.ProtoReflect.Descriptor instead.
func (*PeripheralRequirement) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescGZIP(), []int{2}
}

func (x *PeripheralRequirement) GetPeripheralType() PeripheralType {
	if x != nil {
		return x.PeripheralType
	}
	return PeripheralType_PERIPHERAL_TYPE_UNSPECIFIED
}

func (x *PeripheralRequirement) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *PeripheralRequirement) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

type VirtualRequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtualType VirtualType `protobuf:"varint,1,opt,name=virtual_type,json=virtualType,proto3,enum=unifiedfleet.api.v1.models.VirtualType" json:"virtual_type,omitempty"`
	// The minimum/maximum number of the vms that can be setup.
	Min int32 `protobuf:"varint,2,opt,name=min,proto3" json:"min,omitempty"`
	Max int32 `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *VirtualRequirement) Reset() {
	*x = VirtualRequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualRequirement) ProtoMessage() {}

func (x *VirtualRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualRequirement.ProtoReflect.Descriptor instead.
func (*VirtualRequirement) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescGZIP(), []int{3}
}

func (x *VirtualRequirement) GetVirtualType() VirtualType {
	if x != nil {
		return x.VirtualType
	}
	return VirtualType_VIRTUAL_TYPE_UNSPECIFIED
}

func (x *VirtualRequirement) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *VirtualRequirement) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

var File_infra_unifiedfleet_api_v1_models_lse_prototype_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDesc = []byte{
	0x0a, 0x34, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6c, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc6, 0x02, 0x0a, 0x10, 0x52, 0x61, 0x63, 0x6b, 0x4c, 0x53, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x6a, 0x0a, 0x17, 0x70, 0x65, 0x72, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x75, 0x6e, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x16, 0x70, 0x65,
	0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x3a, 0x5c, 0xea, 0x41, 0x59, 0x0a,
	0x31, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x52, 0x61, 0x63, 0x6b, 0x4c, 0x53, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x24, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x53, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x53, 0x45, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x22, 0xe7, 0x03, 0x0a, 0x13, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x4c, 0x53, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x6a, 0x0a, 0x17, 0x70, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x16, 0x70, 0x65, 0x72, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x30, 0x0a, 0x14, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12,
	0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x52, 0x75, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x61, 0x0a, 0x14, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x13, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x3a, 0x65, 0xea, 0x41, 0x62,
	0x0a, 0x34, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4c, 0x53, 0x45, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4c,
	0x53, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x4c, 0x53, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79, 0x70,
	0x65, 0x7d, 0x22, 0x90, 0x01, 0x0a, 0x15, 0x50, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x53, 0x0a, 0x0f,
	0x70, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x84, 0x01, 0x0a, 0x12, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x0c,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x27, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x2a, 0x9b, 0x02, 0x0a,
	0x0e, 0x50, 0x65, 0x72, 0x69, 0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x1b, 0x50, 0x45, 0x52, 0x49, 0x50, 0x48, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x19, 0x0a, 0x15, 0x50, 0x45, 0x52, 0x49, 0x50, 0x48, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x4f, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x50,
	0x45, 0x52, 0x49, 0x50, 0x48, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c,
	0x41, 0x42, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x50,
	0x45, 0x52, 0x49, 0x50, 0x48, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x50, 0x4d, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x45, 0x52, 0x49, 0x50, 0x48, 0x45, 0x52,
	0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x56, 0x4d, 0x10, 0x04, 0x12, 0x1a, 0x0a,
	0x16, 0x50, 0x45, 0x52, 0x49, 0x50, 0x48, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21, 0x50, 0x45, 0x52,
	0x49, 0x50, 0x48, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4c, 0x55,
	0x45, 0x54, 0x4f, 0x4f, 0x54, 0x48, 0x5f, 0x42, 0x54, 0x50, 0x45, 0x45, 0x52, 0x53, 0x10, 0x06,
	0x12, 0x1c, 0x0a, 0x18, 0x50, 0x45, 0x52, 0x49, 0x50, 0x48, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x57, 0x49, 0x46, 0x49, 0x43, 0x45, 0x4c, 0x4c, 0x10, 0x07, 0x12, 0x1a,
	0x0a, 0x16, 0x50, 0x45, 0x52, 0x49, 0x50, 0x48, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x41, 0x4d, 0x45, 0x52, 0x41, 0x10, 0x08, 0x2a, 0x40, 0x0a, 0x0b, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x56, 0x49, 0x52,
	0x54, 0x55, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x49, 0x52, 0x54, 0x55,
	0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x4d, 0x10, 0x01, 0x42, 0x28, 0x5a, 0x26,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65,
	0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x3b, 0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_goTypes = []interface{}{
	(PeripheralType)(0),           // 0: unifiedfleet.api.v1.models.PeripheralType
	(VirtualType)(0),              // 1: unifiedfleet.api.v1.models.VirtualType
	(*RackLSEPrototype)(nil),      // 2: unifiedfleet.api.v1.models.RackLSEPrototype
	(*MachineLSEPrototype)(nil),   // 3: unifiedfleet.api.v1.models.MachineLSEPrototype
	(*PeripheralRequirement)(nil), // 4: unifiedfleet.api.v1.models.PeripheralRequirement
	(*VirtualRequirement)(nil),    // 5: unifiedfleet.api.v1.models.VirtualRequirement
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_depIdxs = []int32{
	4, // 0: unifiedfleet.api.v1.models.RackLSEPrototype.peripheral_requirements:type_name -> unifiedfleet.api.v1.models.PeripheralRequirement
	6, // 1: unifiedfleet.api.v1.models.RackLSEPrototype.update_time:type_name -> google.protobuf.Timestamp
	4, // 2: unifiedfleet.api.v1.models.MachineLSEPrototype.peripheral_requirements:type_name -> unifiedfleet.api.v1.models.PeripheralRequirement
	6, // 3: unifiedfleet.api.v1.models.MachineLSEPrototype.update_time:type_name -> google.protobuf.Timestamp
	5, // 4: unifiedfleet.api.v1.models.MachineLSEPrototype.virtual_requirements:type_name -> unifiedfleet.api.v1.models.VirtualRequirement
	0, // 5: unifiedfleet.api.v1.models.PeripheralRequirement.peripheral_type:type_name -> unifiedfleet.api.v1.models.PeripheralType
	1, // 6: unifiedfleet.api.v1.models.VirtualRequirement.virtual_type:type_name -> unifiedfleet.api.v1.models.VirtualType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_init() }
func file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_lse_prototype_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RackLSEPrototype); i {
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
		file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineLSEPrototype); i {
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
		file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeripheralRequirement); i {
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
		file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualRequirement); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_depIdxs,
		EnumInfos:         file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_enumTypes,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_lse_prototype_proto = out.File
	file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_lse_prototype_proto_depIdxs = nil
}
