// Copyright 2018 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/libs/skylab/inventory/lab.proto

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

// NEXT TAG: 6
type Lab struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duts                 []*DeviceUnderTest     `protobuf:"bytes,1,rep,name=duts" json:"duts,omitempty"`
	ServoHosts           []*ServoHostDevice     `protobuf:"bytes,2,rep,name=servo_hosts,json=servoHosts" json:"servo_hosts,omitempty"`
	Chamelons            []*ChameleonDevice     `protobuf:"bytes,3,rep,name=chamelons" json:"chamelons,omitempty"`
	ServoHostConnections []*ServoHostConnection `protobuf:"bytes,4,rep,name=servo_host_connections,json=servoHostConnections" json:"servo_host_connections,omitempty"`
	ChameleonConnections []*ChameleonConnection `protobuf:"bytes,5,rep,name=chameleon_connections,json=chameleonConnections" json:"chameleon_connections,omitempty"`
}

func (x *Lab) Reset() {
	*x = Lab{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_libs_skylab_inventory_lab_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lab) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lab) ProtoMessage() {}

func (x *Lab) ProtoReflect() protoreflect.Message {
	mi := &file_infra_libs_skylab_inventory_lab_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lab.ProtoReflect.Descriptor instead.
func (*Lab) Descriptor() ([]byte, []int) {
	return file_infra_libs_skylab_inventory_lab_proto_rawDescGZIP(), []int{0}
}

func (x *Lab) GetDuts() []*DeviceUnderTest {
	if x != nil {
		return x.Duts
	}
	return nil
}

func (x *Lab) GetServoHosts() []*ServoHostDevice {
	if x != nil {
		return x.ServoHosts
	}
	return nil
}

func (x *Lab) GetChamelons() []*ChameleonDevice {
	if x != nil {
		return x.Chamelons
	}
	return nil
}

func (x *Lab) GetServoHostConnections() []*ServoHostConnection {
	if x != nil {
		return x.ServoHostConnections
	}
	return nil
}

func (x *Lab) GetChameleonConnections() []*ChameleonConnection {
	if x != nil {
		return x.ChameleonConnections
	}
	return nil
}

type Infrastructure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*Server `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty"`
}

func (x *Infrastructure) Reset() {
	*x = Infrastructure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_libs_skylab_inventory_lab_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Infrastructure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Infrastructure) ProtoMessage() {}

func (x *Infrastructure) ProtoReflect() protoreflect.Message {
	mi := &file_infra_libs_skylab_inventory_lab_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Infrastructure.ProtoReflect.Descriptor instead.
func (*Infrastructure) Descriptor() ([]byte, []int) {
	return file_infra_libs_skylab_inventory_lab_proto_rawDescGZIP(), []int{1}
}

func (x *Infrastructure) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

var File_infra_libs_skylab_inventory_lab_proto protoreflect.FileDescriptor

var file_infra_libs_skylab_inventory_lab_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x73, 0x6b, 0x79,
	0x6c, 0x61, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6c, 0x61,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73,
	0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x2c, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x69, 0x62,
	0x73, 0x2f, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f,
	0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x04, 0x0a, 0x03, 0x4c, 0x61, 0x62, 0x12,
	0x51, 0x0a, 0x04, 0x64, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x75,
	0x74, 0x73, 0x12, 0x5e, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x48, 0x6f, 0x73, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x48, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x5b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6d, 0x65, 0x6c, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b,
	0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6d, 0x65, 0x6c, 0x65, 0x6f, 0x6e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6d, 0x65, 0x6c, 0x6f, 0x6e, 0x73, 0x12,
	0x77, 0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x41, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f,
	0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x14, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x76, 0x0a, 0x15, 0x63, 0x68, 0x61, 0x6d,
	0x65, 0x6c, 0x65, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e,
	0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6d, 0x65, 0x6c, 0x65, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x63, 0x68, 0x61, 0x6d,
	0x65, 0x6c, 0x65, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x60, 0x0a, 0x0e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x4e, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x73, 0x6b, 0x79, 0x6c,
	0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6c, 0x69, 0x62, 0x73,
	0x2f, 0x73, 0x6b, 0x79, 0x6c, 0x61, 0x62, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x3b, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
}

var (
	file_infra_libs_skylab_inventory_lab_proto_rawDescOnce sync.Once
	file_infra_libs_skylab_inventory_lab_proto_rawDescData = file_infra_libs_skylab_inventory_lab_proto_rawDesc
)

func file_infra_libs_skylab_inventory_lab_proto_rawDescGZIP() []byte {
	file_infra_libs_skylab_inventory_lab_proto_rawDescOnce.Do(func() {
		file_infra_libs_skylab_inventory_lab_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_libs_skylab_inventory_lab_proto_rawDescData)
	})
	return file_infra_libs_skylab_inventory_lab_proto_rawDescData
}

var file_infra_libs_skylab_inventory_lab_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_libs_skylab_inventory_lab_proto_goTypes = []interface{}{
	(*Lab)(nil),                 // 0: chrome.chromeos_infra.skylab.proto.inventory.Lab
	(*Infrastructure)(nil),      // 1: chrome.chromeos_infra.skylab.proto.inventory.Infrastructure
	(*DeviceUnderTest)(nil),     // 2: chrome.chromeos_infra.skylab.proto.inventory.DeviceUnderTest
	(*ServoHostDevice)(nil),     // 3: chrome.chromeos_infra.skylab.proto.inventory.ServoHostDevice
	(*ChameleonDevice)(nil),     // 4: chrome.chromeos_infra.skylab.proto.inventory.ChameleonDevice
	(*ServoHostConnection)(nil), // 5: chrome.chromeos_infra.skylab.proto.inventory.ServoHostConnection
	(*ChameleonConnection)(nil), // 6: chrome.chromeos_infra.skylab.proto.inventory.ChameleonConnection
	(*Server)(nil),              // 7: chrome.chromeos_infra.skylab.proto.inventory.Server
}
var file_infra_libs_skylab_inventory_lab_proto_depIdxs = []int32{
	2, // 0: chrome.chromeos_infra.skylab.proto.inventory.Lab.duts:type_name -> chrome.chromeos_infra.skylab.proto.inventory.DeviceUnderTest
	3, // 1: chrome.chromeos_infra.skylab.proto.inventory.Lab.servo_hosts:type_name -> chrome.chromeos_infra.skylab.proto.inventory.ServoHostDevice
	4, // 2: chrome.chromeos_infra.skylab.proto.inventory.Lab.chamelons:type_name -> chrome.chromeos_infra.skylab.proto.inventory.ChameleonDevice
	5, // 3: chrome.chromeos_infra.skylab.proto.inventory.Lab.servo_host_connections:type_name -> chrome.chromeos_infra.skylab.proto.inventory.ServoHostConnection
	6, // 4: chrome.chromeos_infra.skylab.proto.inventory.Lab.chameleon_connections:type_name -> chrome.chromeos_infra.skylab.proto.inventory.ChameleonConnection
	7, // 5: chrome.chromeos_infra.skylab.proto.inventory.Infrastructure.servers:type_name -> chrome.chromeos_infra.skylab.proto.inventory.Server
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_infra_libs_skylab_inventory_lab_proto_init() }
func file_infra_libs_skylab_inventory_lab_proto_init() {
	if File_infra_libs_skylab_inventory_lab_proto != nil {
		return
	}
	file_infra_libs_skylab_inventory_connection_proto_init()
	file_infra_libs_skylab_inventory_device_proto_init()
	file_infra_libs_skylab_inventory_server_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infra_libs_skylab_inventory_lab_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lab); i {
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
		file_infra_libs_skylab_inventory_lab_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Infrastructure); i {
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
			RawDescriptor: file_infra_libs_skylab_inventory_lab_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_libs_skylab_inventory_lab_proto_goTypes,
		DependencyIndexes: file_infra_libs_skylab_inventory_lab_proto_depIdxs,
		MessageInfos:      file_infra_libs_skylab_inventory_lab_proto_msgTypes,
	}.Build()
	File_infra_libs_skylab_inventory_lab_proto = out.File
	file_infra_libs_skylab_inventory_lab_proto_rawDesc = nil
	file_infra_libs_skylab_inventory_lab_proto_goTypes = nil
	file_infra_libs_skylab_inventory_lab_proto_depIdxs = nil
}
