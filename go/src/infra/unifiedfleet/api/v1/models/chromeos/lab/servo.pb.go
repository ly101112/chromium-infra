// Copyright 2020 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.7
// source: infra/unifiedfleet/api/v1/models/chromeos/lab/servo.proto

package ufspb

import (
	api "go.chromium.org/chromiumos/config/go/test/lab/api"
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

// Servo Setup Type describes the capabilities of servos.
// Next Tag : 3
type ServoSetupType int32

const (
	ServoSetupType_SERVO_SETUP_REGULAR ServoSetupType = 0
	ServoSetupType_SERVO_SETUP_DUAL_V4 ServoSetupType = 1
	// SERVO_SETUP_INVALID explicitly marks errors in servo setup.
	ServoSetupType_SERVO_SETUP_INVALID ServoSetupType = 2
)

// Enum value maps for ServoSetupType.
var (
	ServoSetupType_name = map[int32]string{
		0: "SERVO_SETUP_REGULAR",
		1: "SERVO_SETUP_DUAL_V4",
		2: "SERVO_SETUP_INVALID",
	}
	ServoSetupType_value = map[string]int32{
		"SERVO_SETUP_REGULAR": 0,
		"SERVO_SETUP_DUAL_V4": 1,
		"SERVO_SETUP_INVALID": 2,
	}
)

func (x ServoSetupType) Enum() *ServoSetupType {
	p := new(ServoSetupType)
	*p = x
	return p
}

func (x ServoSetupType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServoSetupType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_enumTypes[0].Descriptor()
}

func (ServoSetupType) Type() protoreflect.EnumType {
	return &file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_enumTypes[0]
}

func (x ServoSetupType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServoSetupType.Descriptor instead.
func (ServoSetupType) EnumDescriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescGZIP(), []int{0}
}

// Servo Firmware Channel describes the firmware expected to have on servos.
// Next Tag : 4
type ServoFwChannel int32

const (
	// Servo firmware from Stable channel.
	ServoFwChannel_SERVO_FW_STABLE ServoFwChannel = 0
	// The previous Servo firmware from Stable channel.
	ServoFwChannel_SERVO_FW_PREV ServoFwChannel = 1
	// Servo firmware from Dev channel.
	ServoFwChannel_SERVO_FW_DEV ServoFwChannel = 2
	// Servo firmware from Alpha channel.
	ServoFwChannel_SERVO_FW_ALPHA ServoFwChannel = 3
)

// Enum value maps for ServoFwChannel.
var (
	ServoFwChannel_name = map[int32]string{
		0: "SERVO_FW_STABLE",
		1: "SERVO_FW_PREV",
		2: "SERVO_FW_DEV",
		3: "SERVO_FW_ALPHA",
	}
	ServoFwChannel_value = map[string]int32{
		"SERVO_FW_STABLE": 0,
		"SERVO_FW_PREV":   1,
		"SERVO_FW_DEV":    2,
		"SERVO_FW_ALPHA":  3,
	}
)

func (x ServoFwChannel) Enum() *ServoFwChannel {
	p := new(ServoFwChannel)
	*p = x
	return p
}

func (x ServoFwChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServoFwChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_enumTypes[1].Descriptor()
}

func (ServoFwChannel) Type() protoreflect.EnumType {
	return &file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_enumTypes[1]
}

func (x ServoFwChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServoFwChannel.Descriptor instead.
func (ServoFwChannel) EnumDescriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescGZIP(), []int{1}
}

// NEXT TAG: 13
type Servo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Servo-specific configs
	ServoHostname string `protobuf:"bytes,2,opt,name=servo_hostname,json=servoHostname,proto3" json:"servo_hostname,omitempty"`
	ServoPort     int32  `protobuf:"varint,3,opt,name=servo_port,json=servoPort,proto3" json:"servo_port,omitempty"`
	ServoSerial   string `protobuf:"bytes,4,opt,name=servo_serial,json=servoSerial,proto3" json:"servo_serial,omitempty"`
	// Based on https://docs.google.com/document/d/1TPp7yp-uwFUh5xOnBLI4jPYtYD7IcdyQ1dgqFqtcJEU/edit?ts=5d8eafb7#heading=h.csdfk1i6g0l
	// servo_type will contain different setup of servos. So string is recommended than enum.
	ServoType  string         `protobuf:"bytes,5,opt,name=servo_type,json=servoType,proto3" json:"servo_type,omitempty"`
	ServoSetup ServoSetupType `protobuf:"varint,7,opt,name=servo_setup,json=servoSetup,proto3,enum=unifiedfleet.api.v1.models.chromeos.lab.ServoSetupType" json:"servo_setup,omitempty"`
	// Based on http://go/fleet-servo-topology
	ServoTopology  *ServoTopology `protobuf:"bytes,8,opt,name=servo_topology,json=servoTopology,proto3" json:"servo_topology,omitempty"`
	ServoFwChannel ServoFwChannel `protobuf:"varint,9,opt,name=servo_fw_channel,json=servoFwChannel,proto3,enum=unifiedfleet.api.v1.models.chromeos.lab.ServoFwChannel" json:"servo_fw_channel,omitempty"`
	ServoComponent []string       `protobuf:"bytes,11,rep,name=servo_component,json=servoComponent,proto3" json:"servo_component,omitempty"`
	// b/190538710 optional docker container name if servod is running in docker
	DockerContainerName string `protobuf:"bytes,10,opt,name=docker_container_name,json=dockerContainerName,proto3" json:"docker_container_name,omitempty"`
	// UsbDrive contains details of the servo's plugged USB drive.
	UsbDrive *api.UsbDrive `protobuf:"bytes,12,opt,name=usb_drive,json=usbDrive,proto3" json:"usb_drive,omitempty"`
}

func (x *Servo) Reset() {
	*x = Servo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Servo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Servo) ProtoMessage() {}

func (x *Servo) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Servo.ProtoReflect.Descriptor instead.
func (*Servo) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescGZIP(), []int{0}
}

func (x *Servo) GetServoHostname() string {
	if x != nil {
		return x.ServoHostname
	}
	return ""
}

func (x *Servo) GetServoPort() int32 {
	if x != nil {
		return x.ServoPort
	}
	return 0
}

func (x *Servo) GetServoSerial() string {
	if x != nil {
		return x.ServoSerial
	}
	return ""
}

func (x *Servo) GetServoType() string {
	if x != nil {
		return x.ServoType
	}
	return ""
}

func (x *Servo) GetServoSetup() ServoSetupType {
	if x != nil {
		return x.ServoSetup
	}
	return ServoSetupType_SERVO_SETUP_REGULAR
}

func (x *Servo) GetServoTopology() *ServoTopology {
	if x != nil {
		return x.ServoTopology
	}
	return nil
}

func (x *Servo) GetServoFwChannel() ServoFwChannel {
	if x != nil {
		return x.ServoFwChannel
	}
	return ServoFwChannel_SERVO_FW_STABLE
}

func (x *Servo) GetServoComponent() []string {
	if x != nil {
		return x.ServoComponent
	}
	return nil
}

func (x *Servo) GetDockerContainerName() string {
	if x != nil {
		return x.DockerContainerName
	}
	return ""
}

func (x *Servo) GetUsbDrive() *api.UsbDrive {
	if x != nil {
		return x.UsbDrive
	}
	return nil
}

// Servo Topology describe connected servo devices on DUT set-up to provide Servo functionality.
// Next Tag : 3
type ServoTopology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Main     *ServoTopologyItem   `protobuf:"bytes,1,opt,name=main,proto3" json:"main,omitempty"`
	Children []*ServoTopologyItem `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ServoTopology) Reset() {
	*x = ServoTopology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServoTopology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServoTopology) ProtoMessage() {}

func (x *ServoTopology) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServoTopology.ProtoReflect.Descriptor instead.
func (*ServoTopology) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescGZIP(), []int{1}
}

func (x *ServoTopology) GetMain() *ServoTopologyItem {
	if x != nil {
		return x.Main
	}
	return nil
}

func (x *ServoTopology) GetChildren() []*ServoTopologyItem {
	if x != nil {
		return x.Children
	}
	return nil
}

// Servo Topology Item describe details of one servo device on DUT set-up.
// Next Tag : 6
type ServoTopologyItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type provides the type of servo device. Keeping as String to avoid issue with introduce new type.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// sysfs_product provides the product name of the device recorded in File System.
	SysfsProduct string `protobuf:"bytes,2,opt,name=sysfs_product,json=sysfsProduct,proto3" json:"sysfs_product,omitempty"`
	// serial provides the serial number of the device.
	Serial string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	// usb_hub_port provides the port connection to the device.
	// e.g. '1-6.2.2' where
	//
	//	'1-6'  - port on the labstation
	//	'2'    - port on smart-hub connected to the labstation
	//	'2'    - port on servo hub (part of servo_v4 or servo_v4.1) connected to the smart-hub
	//
	// The same path will look '1-6.2' if connected servo_v4 directly to the labstation.
	UsbHubPort string `protobuf:"bytes,4,opt,name=usb_hub_port,json=usbHubPort,proto3" json:"usb_hub_port,omitempty"`
	// This is the firmware version of servo device.
	FwVersion string `protobuf:"bytes,5,opt,name=fw_version,json=fwVersion,proto3" json:"fw_version,omitempty"`
}

func (x *ServoTopologyItem) Reset() {
	*x = ServoTopologyItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServoTopologyItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServoTopologyItem) ProtoMessage() {}

func (x *ServoTopologyItem) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServoTopologyItem.ProtoReflect.Descriptor instead.
func (*ServoTopologyItem) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescGZIP(), []int{2}
}

func (x *ServoTopologyItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ServoTopologyItem) GetSysfsProduct() string {
	if x != nil {
		return x.SysfsProduct
	}
	return ""
}

func (x *ServoTopologyItem) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *ServoTopologyItem) GetUsbHubPort() string {
	if x != nil {
		return x.UsbHubPort
	}
	return ""
}

func (x *ServoTopologyItem) GetFwVersion() string {
	if x != nil {
		return x.FwVersion
	}
	return ""
}

var File_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDesc = []byte{
	0x0a, 0x39, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73,
	0x2e, 0x6c, 0x61, 0x62, 0x1a, 0x4f, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6c, 0x61,
	0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x62, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x04, 0x0a, 0x05, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x48, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x5f,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x6f, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x6f, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x6f,
	0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65,
	0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x53, 0x65, 0x74, 0x75,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x53, 0x65, 0x74, 0x75,
	0x70, 0x12, 0x5d, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x75, 0x6e, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e,
	0x6c, 0x61, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79,
	0x12, 0x61, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x5f, 0x66, 0x77, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73,
	0x2e, 0x6c, 0x61, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x46, 0x77, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x46, 0x77, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x6f, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x3e, 0x0a, 0x09, 0x75, 0x73, 0x62, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73,
	0x62, 0x44, 0x72, 0x69, 0x76, 0x65, 0x52, 0x08, 0x75, 0x73, 0x62, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0xb7, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x6f,
	0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x4e, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6c, 0x61, 0x62,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x56, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x72, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73,
	0x2e, 0x6c, 0x61, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x22, 0xa5, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x79,
	0x73, 0x66, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x66, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x62, 0x5f, 0x68,
	0x75, 0x62, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x73, 0x62, 0x48, 0x75, 0x62, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x77, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x5b, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76,
	0x6f, 0x53, 0x65, 0x74, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45,
	0x52, 0x56, 0x4f, 0x5f, 0x53, 0x45, 0x54, 0x55, 0x50, 0x5f, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41,
	0x52, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x4f, 0x5f, 0x53, 0x45, 0x54,
	0x55, 0x50, 0x5f, 0x44, 0x55, 0x41, 0x4c, 0x5f, 0x56, 0x34, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x53, 0x45, 0x52, 0x56, 0x4f, 0x5f, 0x53, 0x45, 0x54, 0x55, 0x50, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x02, 0x2a, 0x5e, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x46, 0x77,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x52, 0x56, 0x4f,
	0x5f, 0x46, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x53, 0x45, 0x52, 0x56, 0x4f, 0x5f, 0x46, 0x57, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56, 0x4f, 0x5f, 0x46, 0x57, 0x5f, 0x44, 0x45, 0x56, 0x10,
	0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x52, 0x56, 0x4f, 0x5f, 0x46, 0x57, 0x5f, 0x41, 0x4c,
	0x50, 0x48, 0x41, 0x10, 0x03, 0x42, 0x35, 0x5a, 0x33, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65,
	0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x3b, 0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_goTypes = []interface{}{
	(ServoSetupType)(0),       // 0: unifiedfleet.api.v1.models.chromeos.lab.ServoSetupType
	(ServoFwChannel)(0),       // 1: unifiedfleet.api.v1.models.chromeos.lab.ServoFwChannel
	(*Servo)(nil),             // 2: unifiedfleet.api.v1.models.chromeos.lab.Servo
	(*ServoTopology)(nil),     // 3: unifiedfleet.api.v1.models.chromeos.lab.ServoTopology
	(*ServoTopologyItem)(nil), // 4: unifiedfleet.api.v1.models.chromeos.lab.ServoTopologyItem
	(*api.UsbDrive)(nil),      // 5: chromiumos.test.lab.api.UsbDrive
}
var file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_depIdxs = []int32{
	0, // 0: unifiedfleet.api.v1.models.chromeos.lab.Servo.servo_setup:type_name -> unifiedfleet.api.v1.models.chromeos.lab.ServoSetupType
	3, // 1: unifiedfleet.api.v1.models.chromeos.lab.Servo.servo_topology:type_name -> unifiedfleet.api.v1.models.chromeos.lab.ServoTopology
	1, // 2: unifiedfleet.api.v1.models.chromeos.lab.Servo.servo_fw_channel:type_name -> unifiedfleet.api.v1.models.chromeos.lab.ServoFwChannel
	5, // 3: unifiedfleet.api.v1.models.chromeos.lab.Servo.usb_drive:type_name -> chromiumos.test.lab.api.UsbDrive
	4, // 4: unifiedfleet.api.v1.models.chromeos.lab.ServoTopology.main:type_name -> unifiedfleet.api.v1.models.chromeos.lab.ServoTopologyItem
	4, // 5: unifiedfleet.api.v1.models.chromeos.lab.ServoTopology.children:type_name -> unifiedfleet.api.v1.models.chromeos.lab.ServoTopologyItem
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_init() }
func file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Servo); i {
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
		file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServoTopology); i {
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
		file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServoTopologyItem); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_depIdxs,
		EnumInfos:         file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_enumTypes,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto = out.File
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_chromeos_lab_servo_proto_depIdxs = nil
}
