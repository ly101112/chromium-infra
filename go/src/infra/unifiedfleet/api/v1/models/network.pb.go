// Copyright 2020 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/unifiedfleet/api/v1/models/network.proto

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

type Nic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique serial_number or asset tag
	// The format will be nics/XXX
	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MacAddress string `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// Record the last update timestamp of this machine (In UTC timezone)
	UpdateTime      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	SwitchInterface *SwitchInterface       `protobuf:"bytes,4,opt,name=switch_interface,json=switchInterface,proto3" json:"switch_interface,omitempty"`
	// Refers to Machine name
	Machine string `protobuf:"bytes,5,opt,name=machine,proto3" json:"machine,omitempty"`
	// Refers to Rack name
	Rack string `protobuf:"bytes,7,opt,name=rack,proto3" json:"rack,omitempty"`
	// tags user can attach for easy querying/searching
	Tags []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	// Refers to Zone
	Zone string `protobuf:"bytes,10,opt,name=zone,proto3" json:"zone,omitempty"`
	// State of the nic
	ResourceState State `protobuf:"varint,11,opt,name=resource_state,json=resourceState,proto3,enum=unifiedfleet.api.v1.models.State" json:"resource_state,omitempty"`
}

func (x *Nic) Reset() {
	*x = Nic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nic) ProtoMessage() {}

func (x *Nic) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nic.ProtoReflect.Descriptor instead.
func (*Nic) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_network_proto_rawDescGZIP(), []int{0}
}

func (x *Nic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Nic) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *Nic) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Nic) GetSwitchInterface() *SwitchInterface {
	if x != nil {
		return x.SwitchInterface
	}
	return nil
}

func (x *Nic) GetMachine() string {
	if x != nil {
		return x.Machine
	}
	return ""
}

func (x *Nic) GetRack() string {
	if x != nil {
		return x.Rack
	}
	return ""
}

func (x *Nic) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Nic) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *Nic) GetResourceState() State {
	if x != nil {
		return x.ResourceState
	}
	return State_STATE_UNSPECIFIED
}

type Vlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique name for the Vlan
	// The format will be vlans/XXX
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The format will be subnet/CIDR.
	VlanAddress string `protobuf:"bytes,2,opt,name=vlan_address,json=vlanAddress,proto3" json:"vlan_address,omitempty"`
	// The number of IPs that in this vlan
	CapacityIp int32 `protobuf:"varint,3,opt,name=capacity_ip,json=capacityIp,proto3" json:"capacity_ip,omitempty"`
	// Record the last update timestamp of this Vlan (In UTC timezone)
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The description of the vlan.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// The ips that cannot be used by normal host in this vlan.
	// By default, the first 10 and last 1 ip in the vlan is automatically reserved.
	// No need to specify them here.
	ReservedIps []string `protobuf:"bytes,7,rep,name=reserved_ips,json=reservedIps,proto3" json:"reserved_ips,omitempty"`
	// Define the free ip range of this vlan that can be automatically assigned.
	// Any ips between the two ips are regarded as available ips.
	// Any ips that not in this range will be marked as reserved.
	// If they're empty, it means no ips in this vlan can be used as dhcp configs.
	// The format of the ip ranges are ipv4 string format, e.g. 192.168.1.23.
	// User can update reserved_ips to remove some of the available ips in the free range.
	FreeStartIpv4Str string `protobuf:"bytes,11,opt,name=free_start_ipv4Str,json=freeStartIpv4Str,proto3" json:"free_start_ipv4Str,omitempty"`
	FreeEndIpv4Str   string `protobuf:"bytes,12,opt,name=free_end_ipv4Str,json=freeEndIpv4Str,proto3" json:"free_end_ipv4Str,omitempty"`
	// State of the vlan
	ResourceState State `protobuf:"varint,8,opt,name=resource_state,json=resourceState,proto3,enum=unifiedfleet.api.v1.models.State" json:"resource_state,omitempty"`
	// vlan_number is same as the vlan name without prefix(browser/atl/cros)
	VlanNumber string `protobuf:"bytes,9,opt,name=vlan_number,json=vlanNumber,proto3" json:"vlan_number,omitempty"`
	// The zone that this vlan serves. One vlan can serve multiple zones.
	Zones []Zone `protobuf:"varint,10,rep,packed,name=zones,proto3,enum=unifiedfleet.api.v1.models.Zone" json:"zones,omitempty"`
	// Record the ACL info of the vlan
	Realm string `protobuf:"bytes,13,opt,name=realm,proto3" json:"realm,omitempty"`
	// Tags user can attach for easy querying/searching
	Tags []string `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Vlan) Reset() {
	*x = Vlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vlan) ProtoMessage() {}

func (x *Vlan) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vlan.ProtoReflect.Descriptor instead.
func (*Vlan) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_network_proto_rawDescGZIP(), []int{1}
}

func (x *Vlan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Vlan) GetVlanAddress() string {
	if x != nil {
		return x.VlanAddress
	}
	return ""
}

func (x *Vlan) GetCapacityIp() int32 {
	if x != nil {
		return x.CapacityIp
	}
	return 0
}

func (x *Vlan) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Vlan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Vlan) GetReservedIps() []string {
	if x != nil {
		return x.ReservedIps
	}
	return nil
}

func (x *Vlan) GetFreeStartIpv4Str() string {
	if x != nil {
		return x.FreeStartIpv4Str
	}
	return ""
}

func (x *Vlan) GetFreeEndIpv4Str() string {
	if x != nil {
		return x.FreeEndIpv4Str
	}
	return ""
}

func (x *Vlan) GetResourceState() State {
	if x != nil {
		return x.ResourceState
	}
	return State_STATE_UNSPECIFIED
}

func (x *Vlan) GetVlanNumber() string {
	if x != nil {
		return x.VlanNumber
	}
	return ""
}

func (x *Vlan) GetZones() []Zone {
	if x != nil {
		return x.Zones
	}
	return nil
}

func (x *Vlan) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *Vlan) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type DHCPConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MacAddress string                 `protobuf:"bytes,1,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Hostname   string                 `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Ip         string                 `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Vlan       string                 `protobuf:"bytes,5,opt,name=vlan,proto3" json:"vlan,omitempty"`
}

func (x *DHCPConfig) Reset() {
	*x = DHCPConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DHCPConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DHCPConfig) ProtoMessage() {}

func (x *DHCPConfig) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DHCPConfig.ProtoReflect.Descriptor instead.
func (*DHCPConfig) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_network_proto_rawDescGZIP(), []int{2}
}

func (x *DHCPConfig) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *DHCPConfig) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *DHCPConfig) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *DHCPConfig) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *DHCPConfig) GetVlan() string {
	if x != nil {
		return x.Vlan
	}
	return ""
}

// Message contains all dhcp configs.
type AllDHCPConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configs []*DHCPConfig `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *AllDHCPConfigs) Reset() {
	*x = AllDHCPConfigs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllDHCPConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllDHCPConfigs) ProtoMessage() {}

func (x *AllDHCPConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllDHCPConfigs.ProtoReflect.Descriptor instead.
func (*AllDHCPConfigs) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_network_proto_rawDescGZIP(), []int{3}
}

func (x *AllDHCPConfigs) GetConfigs() []*DHCPConfig {
	if x != nil {
		return x.Configs
	}
	return nil
}

// IP is an intermediate record object, not an object to be exposed by API.
type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// can be converted to and from the string ip address
	Ipv4     uint32 `protobuf:"varint,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Vlan     string `protobuf:"bytes,3,opt,name=vlan,proto3" json:"vlan,omitempty"`
	Occupied bool   `protobuf:"varint,4,opt,name=occupied,proto3" json:"occupied,omitempty"`
	// store the string ip address
	Ipv4Str string `protobuf:"bytes,5,opt,name=ipv4_str,json=ipv4Str,proto3" json:"ipv4_str,omitempty"`
	// Indicate if this IP is reserved (cannot be auto-assigned to any host)
	Reserve bool `protobuf:"varint,6,opt,name=reserve,proto3" json:"reserve,omitempty"`
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_network_proto_rawDescGZIP(), []int{4}
}

func (x *IP) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IP) GetIpv4() uint32 {
	if x != nil {
		return x.Ipv4
	}
	return 0
}

func (x *IP) GetVlan() string {
	if x != nil {
		return x.Vlan
	}
	return ""
}

func (x *IP) GetOccupied() bool {
	if x != nil {
		return x.Occupied
	}
	return false
}

func (x *IP) GetIpv4Str() string {
	if x != nil {
		return x.Ipv4Str
	}
	return ""
}

func (x *IP) GetReserve() bool {
	if x != nil {
		return x.Reserve
	}
	return false
}

var File_infra_unifiedfleet_api_v1_models_network_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_network_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x04, 0x0a,
	0x03, 0x4e, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x56, 0x0a, 0x10, 0x73,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x52, 0x0f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x75, 0x6e, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x41, 0x0a, 0x04,
	0x72, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x03, 0xfa,
	0x41, 0x27, 0x0a, 0x25, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x63, 0x6b, 0x52, 0x04, 0x72, 0x61, 0x63, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x48, 0x0a, 0x0e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x35, 0xea, 0x41, 0x32, 0x0a, 0x24, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x69,
	0x63, 0x12, 0x0a, 0x6e, 0x69, 0x63, 0x73, 0x2f, 0x7b, 0x6e, 0x69, 0x63, 0x7d, 0x4a, 0x04, 0x08,
	0x06, 0x10, 0x07, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a, 0x52, 0x03, 0x6c, 0x61, 0x62, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xd6, 0x04, 0x0a, 0x04, 0x56, 0x6c, 0x61, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x6c, 0x61, 0x6e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61,
	0x63, 0x69, 0x74, 0x79, 0x49, 0x70, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x49, 0x70, 0x73, 0x12, 0x2c, 0x0a,
	0x12, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x70, 0x76, 0x34,
	0x53, 0x74, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x72, 0x65, 0x65, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x49, 0x70, 0x76, 0x34, 0x53, 0x74, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x66,
	0x72, 0x65, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x53, 0x74, 0x72, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x72, 0x65, 0x65, 0x45, 0x6e, 0x64, 0x49, 0x70,
	0x76, 0x34, 0x53, 0x74, 0x72, 0x12, 0x48, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x24, 0x0a, 0x0b, 0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x76, 0x6c, 0x61, 0x6e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65,
	0x61, 0x6c, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x3a, 0x38, 0xea, 0x41, 0x35, 0x0a, 0x25, 0x75, 0x6e,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56,
	0x6c, 0x61, 0x6e, 0x12, 0x0c, 0x76, 0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x7b, 0x76, 0x6c, 0x61, 0x6e,
	0x7d, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xdb,
	0x01, 0x0a, 0x0a, 0x44, 0x48, 0x43, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x04,
	0x76, 0x6c, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xfa, 0x41, 0x27, 0x0a,
	0x25, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2d, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2d, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x56, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x22, 0x52, 0x0a, 0x0e,
	0x41, 0x6c, 0x6c, 0x44, 0x48, 0x43, 0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x40,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x48, 0x43,
	0x50, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x22, 0x8d, 0x01, 0x0a, 0x02, 0x49, 0x50, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x76,
	0x6c, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x70, 0x76, 0x34, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x70, 0x76, 0x34, 0x53, 0x74, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x42, 0x28, 0x5a, 0x26, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x3b, 0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_network_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_network_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_network_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_network_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_network_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_network_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_network_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_infra_unifiedfleet_api_v1_models_network_proto_goTypes = []interface{}{
	(*Nic)(nil),                   // 0: unifiedfleet.api.v1.models.Nic
	(*Vlan)(nil),                  // 1: unifiedfleet.api.v1.models.Vlan
	(*DHCPConfig)(nil),            // 2: unifiedfleet.api.v1.models.DHCPConfig
	(*AllDHCPConfigs)(nil),        // 3: unifiedfleet.api.v1.models.AllDHCPConfigs
	(*IP)(nil),                    // 4: unifiedfleet.api.v1.models.IP
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*SwitchInterface)(nil),       // 6: unifiedfleet.api.v1.models.SwitchInterface
	(State)(0),                    // 7: unifiedfleet.api.v1.models.State
	(Zone)(0),                     // 8: unifiedfleet.api.v1.models.Zone
}
var file_infra_unifiedfleet_api_v1_models_network_proto_depIdxs = []int32{
	5, // 0: unifiedfleet.api.v1.models.Nic.update_time:type_name -> google.protobuf.Timestamp
	6, // 1: unifiedfleet.api.v1.models.Nic.switch_interface:type_name -> unifiedfleet.api.v1.models.SwitchInterface
	7, // 2: unifiedfleet.api.v1.models.Nic.resource_state:type_name -> unifiedfleet.api.v1.models.State
	5, // 3: unifiedfleet.api.v1.models.Vlan.update_time:type_name -> google.protobuf.Timestamp
	7, // 4: unifiedfleet.api.v1.models.Vlan.resource_state:type_name -> unifiedfleet.api.v1.models.State
	8, // 5: unifiedfleet.api.v1.models.Vlan.zones:type_name -> unifiedfleet.api.v1.models.Zone
	5, // 6: unifiedfleet.api.v1.models.DHCPConfig.update_time:type_name -> google.protobuf.Timestamp
	2, // 7: unifiedfleet.api.v1.models.AllDHCPConfigs.configs:type_name -> unifiedfleet.api.v1.models.DHCPConfig
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_network_proto_init() }
func file_infra_unifiedfleet_api_v1_models_network_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_network_proto != nil {
		return
	}
	file_infra_unifiedfleet_api_v1_models_peripherals_proto_init()
	file_infra_unifiedfleet_api_v1_models_state_proto_init()
	file_infra_unifiedfleet_api_v1_models_location_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nic); i {
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
		file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vlan); i {
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
		file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DHCPConfig); i {
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
		file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllDHCPConfigs); i {
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
		file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IP); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_network_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_network_proto_depIdxs,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_network_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_network_proto = out.File
	file_infra_unifiedfleet_api_v1_models_network_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_network_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_network_proto_depIdxs = nil
}
