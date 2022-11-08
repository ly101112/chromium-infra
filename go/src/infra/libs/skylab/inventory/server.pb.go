// Copyright 2018 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: infra/libs/skylab/inventory/server.proto

package inventory

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

// NEXT TAG: 4
type Server_Status int32

const (
	Server_STATUS_INVALID         Server_Status = 0
	Server_STATUS_PRIMARY         Server_Status = 1
	Server_STATUS_BACKUP          Server_Status = 2
	Server_STATUS_REPAIR_REQUIRED Server_Status = 3
)

// Enum value maps for Server_Status.
var (
	Server_Status_name = map[int32]string{
		0: "STATUS_INVALID",
		1: "STATUS_PRIMARY",
		2: "STATUS_BACKUP",
		3: "STATUS_REPAIR_REQUIRED",
	}
	Server_Status_value = map[string]int32{
		"STATUS_INVALID":         0,
		"STATUS_PRIMARY":         1,
		"STATUS_BACKUP":          2,
		"STATUS_REPAIR_REQUIRED": 3,
	}
)

func (x Server_Status) Enum() *Server_Status {
	p := new(Server_Status)
	*p = x
	return p
}

func (x Server_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Server_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_libs_skylab_inventory_server_proto_enumTypes[0].Descriptor()
}

func (Server_Status) Type() protoreflect.EnumType {
	return &file_infra_libs_skylab_inventory_server_proto_enumTypes[0]
}

func (x Server_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Server_Status) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Server_Status(num)
	return nil
}

// Deprecated: Use Server_Status.Descriptor instead.
func (Server_Status) EnumDescriptor() ([]byte, []int) {
	return file_infra_libs_skylab_inventory_server_proto_rawDescGZIP(), []int{0, 0}
}

// NEXT TAG: 14
// Note: Update ROLE_MAP in skylab_inventory/translation_utils.py accordingly
type Server_Role int32

const (
	Server_ROLE_INVALID            Server_Role = 0
	Server_ROLE_AFE                Server_Role = 1
	Server_ROLE_CRASH_SERVER       Server_Role = 2
	Server_ROLE_DATABASE           Server_Role = 3
	Server_ROLE_DATABASE_SLAVE     Server_Role = 4
	Server_ROLE_DEVSERVER          Server_Role = 5
	Server_ROLE_DRONE              Server_Role = 6
	Server_ROLE_GOLO_PROXY         Server_Role = 7
	Server_ROLE_HOST_SCHEDULER     Server_Role = 8
	Server_ROLE_SCHEDULER          Server_Role = 9
	Server_ROLE_SENTINEL           Server_Role = 10
	Server_ROLE_SHARD              Server_Role = 11
	Server_ROLE_SUITE_SCHEDULER    Server_Role = 12
	Server_ROLE_SKYLAB_DRONE       Server_Role = 13
	Server_ROLE_SKYLAB_SUITE_PROXY Server_Role = 14
	Server_ROLE_RPMSERVER          Server_Role = 15
)

// Enum value maps for Server_Role.
var (
	Server_Role_name = map[int32]string{
		0:  "ROLE_INVALID",
		1:  "ROLE_AFE",
		2:  "ROLE_CRASH_SERVER",
		3:  "ROLE_DATABASE",
		4:  "ROLE_DATABASE_SLAVE",
		5:  "ROLE_DEVSERVER",
		6:  "ROLE_DRONE",
		7:  "ROLE_GOLO_PROXY",
		8:  "ROLE_HOST_SCHEDULER",
		9:  "ROLE_SCHEDULER",
		10: "ROLE_SENTINEL",
		11: "ROLE_SHARD",
		12: "ROLE_SUITE_SCHEDULER",
		13: "ROLE_SKYLAB_DRONE",
		14: "ROLE_SKYLAB_SUITE_PROXY",
		15: "ROLE_RPMSERVER",
	}
	Server_Role_value = map[string]int32{
		"ROLE_INVALID":            0,
		"ROLE_AFE":                1,
		"ROLE_CRASH_SERVER":       2,
		"ROLE_DATABASE":           3,
		"ROLE_DATABASE_SLAVE":     4,
		"ROLE_DEVSERVER":          5,
		"ROLE_DRONE":              6,
		"ROLE_GOLO_PROXY":         7,
		"ROLE_HOST_SCHEDULER":     8,
		"ROLE_SCHEDULER":          9,
		"ROLE_SENTINEL":           10,
		"ROLE_SHARD":              11,
		"ROLE_SUITE_SCHEDULER":    12,
		"ROLE_SKYLAB_DRONE":       13,
		"ROLE_SKYLAB_SUITE_PROXY": 14,
		"ROLE_RPMSERVER":          15,
	}
)

func (x Server_Role) Enum() *Server_Role {
	p := new(Server_Role)
	*p = x
	return p
}

func (x Server_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Server_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_libs_skylab_inventory_server_proto_enumTypes[1].Descriptor()
}

func (Server_Role) Type() protoreflect.EnumType {
	return &file_infra_libs_skylab_inventory_server_proto_enumTypes[1]
}

func (x Server_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Server_Role) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Server_Role(num)
	return nil
}

// Deprecated: Use Server_Role.Descriptor instead.
func (Server_Role) EnumDescriptor() ([]byte, []int) {
	return file_infra_libs_skylab_inventory_server_proto_rawDescGZIP(), []int{0, 1}
}

// NEXT TAG: 9
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname    *string        `protobuf:"bytes,1,req,name=hostname" json:"hostname,omitempty"`
	Cname       *string        `protobuf:"bytes,2,opt,name=cname" json:"cname,omitempty"`
	Environment *Environment   `protobuf:"varint,3,req,name=environment,enum=chrome.chromeos_infra.skylab.proto.inventory.Environment" json:"environment,omitempty"`
	Status      *Server_Status `protobuf:"varint,4,req,name=status,enum=chrome.chromeos_infra.skylab.proto.inventory.Server_Status" json:"status,omitempty"`
	Roles       []Server_Role  `protobuf:"varint,5,rep,name=roles,enum=chrome.chromeos_infra.skylab.proto.inventory.Server_Role" json:"roles,omitempty"`
	Attributes  *Attributes    `protobuf:"bytes,6,opt,name=attributes" json:"attributes,omitempty"`
	Notes       *string        `protobuf:"bytes,7,opt,name=notes" json:"notes,omitempty"`
	// List of dut_uids serviced by this server.
	// This can mean different things for different roles.
	//
	//	skylab-drone: These are the DUTs owned by the drone.
	DutUids []string `protobuf:"bytes,8,rep,name=dut_uids,json=dutUids" json:"dut_uids,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_libs_skylab_inventory_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_infra_libs_skylab_inventory_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_infra_libs_skylab_inventory_server_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetHostname() string {
	if x != nil && x.Hostname != nil {
		return *x.Hostname
	}
	return ""
}

func (x *Server) GetCname() string {
	if x != nil && x.Cname != nil {
		return *x.Cname
	}
	return ""
}

func (x *Server) GetEnvironment() Environment {
	if x != nil && x.Environment != nil {
		return *x.Environment
	}
	return Environment_ENVIRONMENT_INVALID
}

func (x *Server) GetStatus() Server_Status {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Server_STATUS_INVALID
}

func (x *Server) GetRoles() []Server_Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Server) GetAttributes() *Attributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Server) GetNotes() string {
	if x != nil && x.Notes != nil {
		return *x.Notes
	}
	return ""
}

func (x *Server) GetDutUids() []string {
	if x != nil {
		return x.DutUids
	}
	return nil
}

// NOTE: Please update SERVER_ATTRIBUTE_TYPE_MAP in
// skylab_inventory/translation_utils.py accordingly.
// NEXT TAG: 6
type Attributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip                        *string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	MaxProcesses              *int32  `protobuf:"varint,2,opt,name=max_processes,json=maxProcesses" json:"max_processes,omitempty"`
	MysqlServerId             *string `protobuf:"bytes,3,opt,name=mysql_server_id,json=mysqlServerId" json:"mysql_server_id,omitempty"`
	DevserverRestrictedSubnet *string `protobuf:"bytes,4,opt,name=devserver_restricted_subnet,json=devserverRestrictedSubnet" json:"devserver_restricted_subnet,omitempty"`
	DevserverPort             *int32  `protobuf:"varint,5,opt,name=devserver_port,json=devserverPort" json:"devserver_port,omitempty"`
}

func (x *Attributes) Reset() {
	*x = Attributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_libs_skylab_inventory_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attributes) ProtoMessage() {}

func (x *Attributes) ProtoReflect() protoreflect.Message {
	mi := &file_infra_libs_skylab_inventory_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attributes.ProtoReflect.Descriptor instead.
func (*Attributes) Descriptor() ([]byte, []int) {
	return file_infra_libs_skylab_inventory_server_proto_rawDescGZIP(), []int{1}
}

func (x *Attributes) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *Attributes) GetMaxProcesses() int32 {
	if x != nil && x.MaxProcesses != nil {
		return *x.MaxProcesses
	}
	return 0
}

func (x *Attributes) GetMysqlServerId() string {
	if x != nil && x.MysqlServerId != nil {
		return *x.MysqlServerId
	}
	return ""
}

func (x *Attributes) GetDevserverRestrictedSubnet() string {
	if x != nil && x.DevserverRestrictedSubnet != nil {
		return *x.DevserverRestrictedSubnet
	}
	return ""
}

func (x *Attributes) GetDevserverPort() int32 {
	if x != nil && x.DevserverPort != nil {
		return *x.DevserverPort
	}
	return 0
}

var File_infra_libs_skylab_inventory_server_proto protoreflect.FileDescriptor

var file_infra_libs_skylab_inventory_server_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x73, 0x6b, 0x79,
	0x6c, 0x61, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x28, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f,
	0x6c, 0x69, 0x62, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x80, 0x07, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x5b, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79,
	0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x53, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4f, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x39, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65,
	0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x12, 0x58, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73,
	0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x74, 0x55, 0x69, 0x64, 0x73, 0x22, 0x5f, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x49, 0x4d, 0x41, 0x52, 0x59, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50,
	0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x50,
	0x41, 0x49, 0x52, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x22, 0xd4,
	0x02, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x4c, 0x45, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x4f, 0x4c,
	0x45, 0x5f, 0x41, 0x46, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x4f, 0x4c, 0x45, 0x5f,
	0x43, 0x52, 0x41, 0x53, 0x48, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x10,
	0x03, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41,
	0x53, 0x45, 0x5f, 0x53, 0x4c, 0x41, 0x56, 0x45, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f,
	0x4c, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x05, 0x12, 0x0e,
	0x0a, 0x0a, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x44, 0x52, 0x4f, 0x4e, 0x45, 0x10, 0x06, 0x12, 0x13,
	0x0a, 0x0f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x47, 0x4f, 0x4c, 0x4f, 0x5f, 0x50, 0x52, 0x4f, 0x58,
	0x59, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x48, 0x4f, 0x53, 0x54,
	0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x52, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x52, 0x10, 0x09,
	0x12, 0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x49, 0x4e, 0x45,
	0x4c, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x52,
	0x44, 0x10, 0x0b, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x55, 0x49, 0x54,
	0x45, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x52, 0x10, 0x0c, 0x12, 0x15, 0x0a,
	0x11, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x4b, 0x59, 0x4c, 0x41, 0x42, 0x5f, 0x44, 0x52, 0x4f,
	0x4e, 0x45, 0x10, 0x0d, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x4b, 0x59,
	0x4c, 0x41, 0x42, 0x5f, 0x53, 0x55, 0x49, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x58, 0x59, 0x10,
	0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x52, 0x50, 0x4d, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x10, 0x0f, 0x22, 0xd0, 0x01, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x79, 0x73,
	0x71, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x3e, 0x0a, 0x1b, 0x64, 0x65, 0x76, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x64, 0x65, 0x76, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x53, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64, 0x65, 0x76, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x27, 0x5a, 0x25, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x3b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79,
}

var (
	file_infra_libs_skylab_inventory_server_proto_rawDescOnce sync.Once
	file_infra_libs_skylab_inventory_server_proto_rawDescData = file_infra_libs_skylab_inventory_server_proto_rawDesc
)

func file_infra_libs_skylab_inventory_server_proto_rawDescGZIP() []byte {
	file_infra_libs_skylab_inventory_server_proto_rawDescOnce.Do(func() {
		file_infra_libs_skylab_inventory_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_libs_skylab_inventory_server_proto_rawDescData)
	})
	return file_infra_libs_skylab_inventory_server_proto_rawDescData
}

var file_infra_libs_skylab_inventory_server_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_infra_libs_skylab_inventory_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_libs_skylab_inventory_server_proto_goTypes = []interface{}{
	(Server_Status)(0), // 0: chrome.chromeos_infra.skylab.proto.inventory.Server.Status
	(Server_Role)(0),   // 1: chrome.chromeos_infra.skylab.proto.inventory.Server.Role
	(*Server)(nil),     // 2: chrome.chromeos_infra.skylab.proto.inventory.Server
	(*Attributes)(nil), // 3: chrome.chromeos_infra.skylab.proto.inventory.Attributes
	(Environment)(0),   // 4: chrome.chromeos_infra.skylab.proto.inventory.Environment
}
var file_infra_libs_skylab_inventory_server_proto_depIdxs = []int32{
	4, // 0: chrome.chromeos_infra.skylab.proto.inventory.Server.environment:type_name -> chrome.chromeos_infra.skylab.proto.inventory.Environment
	0, // 1: chrome.chromeos_infra.skylab.proto.inventory.Server.status:type_name -> chrome.chromeos_infra.skylab.proto.inventory.Server.Status
	1, // 2: chrome.chromeos_infra.skylab.proto.inventory.Server.roles:type_name -> chrome.chromeos_infra.skylab.proto.inventory.Server.Role
	3, // 3: chrome.chromeos_infra.skylab.proto.inventory.Server.attributes:type_name -> chrome.chromeos_infra.skylab.proto.inventory.Attributes
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_infra_libs_skylab_inventory_server_proto_init() }
func file_infra_libs_skylab_inventory_server_proto_init() {
	if File_infra_libs_skylab_inventory_server_proto != nil {
		return
	}
	file_infra_libs_skylab_inventory_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infra_libs_skylab_inventory_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_infra_libs_skylab_inventory_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attributes); i {
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
			RawDescriptor: file_infra_libs_skylab_inventory_server_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_libs_skylab_inventory_server_proto_goTypes,
		DependencyIndexes: file_infra_libs_skylab_inventory_server_proto_depIdxs,
		EnumInfos:         file_infra_libs_skylab_inventory_server_proto_enumTypes,
		MessageInfos:      file_infra_libs_skylab_inventory_server_proto_msgTypes,
	}.Build()
	File_infra_libs_skylab_inventory_server_proto = out.File
	file_infra_libs_skylab_inventory_server_proto_rawDesc = nil
	file_infra_libs_skylab_inventory_server_proto_goTypes = nil
	file_infra_libs_skylab_inventory_server_proto_depIdxs = nil
}
