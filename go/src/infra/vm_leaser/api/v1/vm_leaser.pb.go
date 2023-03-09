// Copyright 2022 The ChromiumOS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: infra/vm_leaser/api/v1/vm_leaser.proto

package vmleaserpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

// VMType specifies the type of VM that is requested and will be created.
type VMType int32

const (
	VMType_VM_TYPE_UNSPECIFIED VMType = 0
	VMType_VM_TYPE_DUT         VMType = 1
	VMType_VM_TYPE_DRONE       VMType = 2
)

// Enum value maps for VMType.
var (
	VMType_name = map[int32]string{
		0: "VM_TYPE_UNSPECIFIED",
		1: "VM_TYPE_DUT",
		2: "VM_TYPE_DRONE",
	}
	VMType_value = map[string]int32{
		"VM_TYPE_UNSPECIFIED": 0,
		"VM_TYPE_DUT":         1,
		"VM_TYPE_DRONE":       2,
	}
)

func (x VMType) Enum() *VMType {
	p := new(VMType)
	*p = x
	return p
}

func (x VMType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VMType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_enumTypes[0].Descriptor()
}

func (VMType) Type() protoreflect.EnumType {
	return &file_infra_vm_leaser_api_v1_vm_leaser_proto_enumTypes[0]
}

func (x VMType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VMType.Descriptor instead.
func (VMType) EnumDescriptor() ([]byte, []int) {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP(), []int{0}
}

type LeaseVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Generated on the client side, shared across retries but pseudo-unique
	// across different logical requests. Requests with the same key will be
	// treated as duplicate of original request, return the same response.
	IdempotencyKey string `protobuf:"bytes,1,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
	// This is the final end user (can be human or robot). Useful for both
	// debugging and analytics. For example for a tests triggered for a CL, this
	// field could indicate the CL author as they are the end user.
	//
	// For direct invocations like CLI, this is enforced at first entry point but
	// trusted from there.
	//
	// Not to be confused with LUCI auth which is done by the caller assuming the
	// appropriate identity from a permissions perspective — like LUCI project.
	OnBehalfOf string `protobuf:"bytes,2,opt,name=on_behalf_of,json=onBehalfOf,proto3" json:"on_behalf_of,omitempty"`
	// This is the quota the end user is requesting to use. One user can have
	// access to multiple quotas. For example, release, CQ, performance testing,
	// etc.
	QuotaId string `protobuf:"bytes,3,opt,name=quota_id,json=quotaId,proto3" json:"quota_id,omitempty"`
	// Optional with a good default.
	// Important to configure a good max (e.g. 10 min). This will put a ceiling on
	// time wasted if the client dies.
	LeaseDuration *durationpb.Duration `protobuf:"bytes,4,opt,name=lease_duration,json=leaseDuration,proto3" json:"lease_duration,omitempty"`
	// The populated requirements will specify the requirements for a VM host.
	HostReqs *VMRequirements `protobuf:"bytes,5,opt,name=host_reqs,json=hostReqs,proto3" json:"host_reqs,omitempty"`
}

func (x *LeaseVMRequest) Reset() {
	*x = LeaseVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaseVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaseVMRequest) ProtoMessage() {}

func (x *LeaseVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaseVMRequest.ProtoReflect.Descriptor instead.
func (*LeaseVMRequest) Descriptor() ([]byte, []int) {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP(), []int{0}
}

func (x *LeaseVMRequest) GetIdempotencyKey() string {
	if x != nil {
		return x.IdempotencyKey
	}
	return ""
}

func (x *LeaseVMRequest) GetOnBehalfOf() string {
	if x != nil {
		return x.OnBehalfOf
	}
	return ""
}

func (x *LeaseVMRequest) GetQuotaId() string {
	if x != nil {
		return x.QuotaId
	}
	return ""
}

func (x *LeaseVMRequest) GetLeaseDuration() *durationpb.Duration {
	if x != nil {
		return x.LeaseDuration
	}
	return nil
}

func (x *LeaseVMRequest) GetHostReqs() *VMRequirements {
	if x != nil {
		return x.HostReqs
	}
	return nil
}

// NEXT TAG: 11
type VMRequirements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GceImage       string `protobuf:"bytes,1,opt,name=gce_image,json=gceImage,proto3" json:"gce_image,omitempty"`
	GceRegion      string `protobuf:"bytes,2,opt,name=gce_region,json=gceRegion,proto3" json:"gce_region,omitempty"`
	GceProject     string `protobuf:"bytes,3,opt,name=gce_project,json=gceProject,proto3" json:"gce_project,omitempty"`
	GceNetwork     string `protobuf:"bytes,4,opt,name=gce_network,json=gceNetwork,proto3" json:"gce_network,omitempty"`
	GceSubnet      string `protobuf:"bytes,5,opt,name=gce_subnet,json=gceSubnet,proto3" json:"gce_subnet,omitempty"`
	GceMachineType string `protobuf:"bytes,6,opt,name=gce_machine_type,json=gceMachineType,proto3" json:"gce_machine_type,omitempty"`
	GceScope       string `protobuf:"bytes,7,opt,name=gce_scope,json=gceScope,proto3" json:"gce_scope,omitempty"`
	// Specified IP address for VM. If not, use ephemeral IP.
	GceIpAddress string `protobuf:"bytes,8,opt,name=gce_ip_address,json=gceIpAddress,proto3" json:"gce_ip_address,omitempty"`
	GceDiskSize  int64  `protobuf:"varint,10,opt,name=gce_disk_size,json=gceDiskSize,proto3" json:"gce_disk_size,omitempty"`
	Type         VMType `protobuf:"varint,9,opt,name=type,proto3,enum=vm_leaser.api.v1.VMType" json:"type,omitempty"`
}

func (x *VMRequirements) Reset() {
	*x = VMRequirements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMRequirements) ProtoMessage() {}

func (x *VMRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMRequirements.ProtoReflect.Descriptor instead.
func (*VMRequirements) Descriptor() ([]byte, []int) {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP(), []int{1}
}

func (x *VMRequirements) GetGceImage() string {
	if x != nil {
		return x.GceImage
	}
	return ""
}

func (x *VMRequirements) GetGceRegion() string {
	if x != nil {
		return x.GceRegion
	}
	return ""
}

func (x *VMRequirements) GetGceProject() string {
	if x != nil {
		return x.GceProject
	}
	return ""
}

func (x *VMRequirements) GetGceNetwork() string {
	if x != nil {
		return x.GceNetwork
	}
	return ""
}

func (x *VMRequirements) GetGceSubnet() string {
	if x != nil {
		return x.GceSubnet
	}
	return ""
}

func (x *VMRequirements) GetGceMachineType() string {
	if x != nil {
		return x.GceMachineType
	}
	return ""
}

func (x *VMRequirements) GetGceScope() string {
	if x != nil {
		return x.GceScope
	}
	return ""
}

func (x *VMRequirements) GetGceIpAddress() string {
	if x != nil {
		return x.GceIpAddress
	}
	return ""
}

func (x *VMRequirements) GetGceDiskSize() int64 {
	if x != nil {
		return x.GceDiskSize
	}
	return 0
}

func (x *VMRequirements) GetType() VMType {
	if x != nil {
		return x.Type
	}
	return VMType_VM_TYPE_UNSPECIFIED
}

type VM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address *VMAddress `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Type    VMType     `protobuf:"varint,3,opt,name=type,proto3,enum=vm_leaser.api.v1.VMType" json:"type,omitempty"`
}

func (x *VM) Reset() {
	*x = VM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VM) ProtoMessage() {}

func (x *VM) ProtoReflect() protoreflect.Message {
	mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VM.ProtoReflect.Descriptor instead.
func (*VM) Descriptor() ([]byte, []int) {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP(), []int{2}
}

func (x *VM) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VM) GetAddress() *VMAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *VM) GetType() VMType {
	if x != nil {
		return x.Type
	}
	return VMType_VM_TYPE_UNSPECIFIED
}

type VMAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IP address of the device.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *VMAddress) Reset() {
	*x = VMAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMAddress) ProtoMessage() {}

func (x *VMAddress) ProtoReflect() protoreflect.Message {
	mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMAddress.ProtoReflect.Descriptor instead.
func (*VMAddress) Descriptor() ([]byte, []int) {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP(), []int{3}
}

func (x *VMAddress) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *VMAddress) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type LeaseVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Relevant information for the lease.
	LeaseId string `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	Vm      *VM    `protobuf:"bytes,2,opt,name=vm,proto3" json:"vm,omitempty"`
	// Client is responsible for extending the lease as needed.
	ExpirationTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
}

func (x *LeaseVMResponse) Reset() {
	*x = LeaseVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaseVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaseVMResponse) ProtoMessage() {}

func (x *LeaseVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaseVMResponse.ProtoReflect.Descriptor instead.
func (*LeaseVMResponse) Descriptor() ([]byte, []int) {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP(), []int{4}
}

func (x *LeaseVMResponse) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

func (x *LeaseVMResponse) GetVm() *VM {
	if x != nil {
		return x.Vm
	}
	return nil
}

func (x *LeaseVMResponse) GetExpirationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

type ExtendLeaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Original lease_id obtained when the lease was created.
	LeaseId string `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	// Important to configure a good lifetime max for leases can not be extended
	// indefinitely. Ideally, long running tasks should not exists and failing
	// tasks should just fail without extension.
	//
	// It is also important to ensure same limits as those for
	// LeaseVMRequest.duration.
	ExtendDuration *durationpb.Duration `protobuf:"bytes,2,opt,name=extend_duration,json=extendDuration,proto3" json:"extend_duration,omitempty"`
}

func (x *ExtendLeaseRequest) Reset() {
	*x = ExtendLeaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendLeaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendLeaseRequest) ProtoMessage() {}

func (x *ExtendLeaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendLeaseRequest.ProtoReflect.Descriptor instead.
func (*ExtendLeaseRequest) Descriptor() ([]byte, []int) {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP(), []int{5}
}

func (x *ExtendLeaseRequest) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

func (x *ExtendLeaseRequest) GetExtendDuration() *durationpb.Duration {
	if x != nil {
		return x.ExtendDuration
	}
	return nil
}

type ExtendLeaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaseId        string                 `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	ExpirationTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
}

func (x *ExtendLeaseResponse) Reset() {
	*x = ExtendLeaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendLeaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendLeaseResponse) ProtoMessage() {}

func (x *ExtendLeaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendLeaseResponse.ProtoReflect.Descriptor instead.
func (*ExtendLeaseResponse) Descriptor() ([]byte, []int) {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP(), []int{6}
}

func (x *ExtendLeaseResponse) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

func (x *ExtendLeaseResponse) GetExpirationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

type ReleaseVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaseId    string `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
	GceProject string `protobuf:"bytes,2,opt,name=gce_project,json=gceProject,proto3" json:"gce_project,omitempty"`
	GceRegion  string `protobuf:"bytes,3,opt,name=gce_region,json=gceRegion,proto3" json:"gce_region,omitempty"`
}

func (x *ReleaseVMRequest) Reset() {
	*x = ReleaseVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseVMRequest) ProtoMessage() {}

func (x *ReleaseVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseVMRequest.ProtoReflect.Descriptor instead.
func (*ReleaseVMRequest) Descriptor() ([]byte, []int) {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP(), []int{7}
}

func (x *ReleaseVMRequest) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

func (x *ReleaseVMRequest) GetGceProject() string {
	if x != nil {
		return x.GceProject
	}
	return ""
}

func (x *ReleaseVMRequest) GetGceRegion() string {
	if x != nil {
		return x.GceRegion
	}
	return ""
}

type ReleaseVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaseId string `protobuf:"bytes,1,opt,name=lease_id,json=leaseId,proto3" json:"lease_id,omitempty"`
}

func (x *ReleaseVMResponse) Reset() {
	*x = ReleaseVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseVMResponse) ProtoMessage() {}

func (x *ReleaseVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseVMResponse.ProtoReflect.Descriptor instead.
func (*ReleaseVMResponse) Descriptor() ([]byte, []int) {
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP(), []int{8}
}

func (x *ReleaseVMResponse) GetLeaseId() string {
	if x != nil {
		return x.LeaseId
	}
	return ""
}

var File_infra_vm_leaser_api_v1_vm_leaser_proto protoreflect.FileDescriptor

var file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x0e,
	0x4c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x6f, 0x6e, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x6c, 0x66, 0x5f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f,
	0x6e, 0x42, 0x65, 0x68, 0x61, 0x6c, 0x66, 0x4f, 0x66, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x72,
	0x65, 0x71, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x6d, 0x5f, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x4d, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x73, 0x22, 0xec, 0x02, 0x0a, 0x0e, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x63, 0x65, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x63, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x63, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x63, 0x65, 0x5f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x63, 0x65, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x63, 0x65, 0x5f, 0x73, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x63, 0x65, 0x53,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x67, 0x63, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x63, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x67, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e,
	0x67, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x63, 0x65, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x67, 0x63, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x67, 0x63, 0x65, 0x44, 0x69,
	0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x79, 0x0a, 0x02, 0x56, 0x4d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x6d,
	0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x4d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x33, 0x0a, 0x09, 0x56, 0x4d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x4d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x02, 0x76, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x4d, 0x52, 0x02, 0x76, 0x6d, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x73,
	0x0a, 0x12, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x42, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x13, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x4c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6d, 0x0a, 0x10, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x63, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x67, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x63,
	0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x67, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x11, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x64, 0x2a, 0x45, 0x0a, 0x06, 0x56, 0x4d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x56, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x55, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x56, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x52, 0x4f, 0x4e, 0x45, 0x10, 0x02,
	0x32, 0x99, 0x02, 0x0a, 0x0f, 0x56, 0x4d, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x4d, 0x12,
	0x20, 0x2e, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x09, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x56, 0x4d, 0x12, 0x22, 0x2e, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x4d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c,
	0x0a, 0x0b, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x24, 0x2e,
	0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x4c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x76, 0x6d, 0x5f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x6d, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x72, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescOnce sync.Once
	file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescData = file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDesc
)

func file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescGZIP() []byte {
	file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescOnce.Do(func() {
		file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescData)
	})
	return file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDescData
}

var file_infra_vm_leaser_api_v1_vm_leaser_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_infra_vm_leaser_api_v1_vm_leaser_proto_goTypes = []interface{}{
	(VMType)(0),                   // 0: vm_leaser.api.v1.VMType
	(*LeaseVMRequest)(nil),        // 1: vm_leaser.api.v1.LeaseVMRequest
	(*VMRequirements)(nil),        // 2: vm_leaser.api.v1.VMRequirements
	(*VM)(nil),                    // 3: vm_leaser.api.v1.VM
	(*VMAddress)(nil),             // 4: vm_leaser.api.v1.VMAddress
	(*LeaseVMResponse)(nil),       // 5: vm_leaser.api.v1.LeaseVMResponse
	(*ExtendLeaseRequest)(nil),    // 6: vm_leaser.api.v1.ExtendLeaseRequest
	(*ExtendLeaseResponse)(nil),   // 7: vm_leaser.api.v1.ExtendLeaseResponse
	(*ReleaseVMRequest)(nil),      // 8: vm_leaser.api.v1.ReleaseVMRequest
	(*ReleaseVMResponse)(nil),     // 9: vm_leaser.api.v1.ReleaseVMResponse
	(*durationpb.Duration)(nil),   // 10: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_infra_vm_leaser_api_v1_vm_leaser_proto_depIdxs = []int32{
	10, // 0: vm_leaser.api.v1.LeaseVMRequest.lease_duration:type_name -> google.protobuf.Duration
	2,  // 1: vm_leaser.api.v1.LeaseVMRequest.host_reqs:type_name -> vm_leaser.api.v1.VMRequirements
	0,  // 2: vm_leaser.api.v1.VMRequirements.type:type_name -> vm_leaser.api.v1.VMType
	4,  // 3: vm_leaser.api.v1.VM.address:type_name -> vm_leaser.api.v1.VMAddress
	0,  // 4: vm_leaser.api.v1.VM.type:type_name -> vm_leaser.api.v1.VMType
	3,  // 5: vm_leaser.api.v1.LeaseVMResponse.vm:type_name -> vm_leaser.api.v1.VM
	11, // 6: vm_leaser.api.v1.LeaseVMResponse.expiration_time:type_name -> google.protobuf.Timestamp
	10, // 7: vm_leaser.api.v1.ExtendLeaseRequest.extend_duration:type_name -> google.protobuf.Duration
	11, // 8: vm_leaser.api.v1.ExtendLeaseResponse.expiration_time:type_name -> google.protobuf.Timestamp
	1,  // 9: vm_leaser.api.v1.VMLeaserService.LeaseVM:input_type -> vm_leaser.api.v1.LeaseVMRequest
	8,  // 10: vm_leaser.api.v1.VMLeaserService.ReleaseVM:input_type -> vm_leaser.api.v1.ReleaseVMRequest
	6,  // 11: vm_leaser.api.v1.VMLeaserService.ExtendLease:input_type -> vm_leaser.api.v1.ExtendLeaseRequest
	5,  // 12: vm_leaser.api.v1.VMLeaserService.LeaseVM:output_type -> vm_leaser.api.v1.LeaseVMResponse
	9,  // 13: vm_leaser.api.v1.VMLeaserService.ReleaseVM:output_type -> vm_leaser.api.v1.ReleaseVMResponse
	7,  // 14: vm_leaser.api.v1.VMLeaserService.ExtendLease:output_type -> vm_leaser.api.v1.ExtendLeaseResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_infra_vm_leaser_api_v1_vm_leaser_proto_init() }
func file_infra_vm_leaser_api_v1_vm_leaser_proto_init() {
	if File_infra_vm_leaser_api_v1_vm_leaser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaseVMRequest); i {
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
		file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMRequirements); i {
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
		file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VM); i {
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
		file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMAddress); i {
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
		file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaseVMResponse); i {
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
		file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendLeaseRequest); i {
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
		file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendLeaseResponse); i {
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
		file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseVMRequest); i {
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
		file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseVMResponse); i {
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
			RawDescriptor: file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_vm_leaser_api_v1_vm_leaser_proto_goTypes,
		DependencyIndexes: file_infra_vm_leaser_api_v1_vm_leaser_proto_depIdxs,
		EnumInfos:         file_infra_vm_leaser_api_v1_vm_leaser_proto_enumTypes,
		MessageInfos:      file_infra_vm_leaser_api_v1_vm_leaser_proto_msgTypes,
	}.Build()
	File_infra_vm_leaser_api_v1_vm_leaser_proto = out.File
	file_infra_vm_leaser_api_v1_vm_leaser_proto_rawDesc = nil
	file_infra_vm_leaser_api_v1_vm_leaser_proto_goTypes = nil
	file_infra_vm_leaser_api_v1_vm_leaser_proto_depIdxs = nil
}
