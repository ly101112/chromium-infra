// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: infra/cros/cmd/cros-tool-runner/api/templates.proto

package api

import (
	api1 "go.chromium.org/chromiumos/config/go/test/api"
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

// Wrapper message of one of container-specific templates
type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Container:
	//	*Template_CrosDut
	//	*Template_CrosProvision
	//	*Template_CrosTest
	Container isTemplate_Container `protobuf_oneof:"container"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescGZIP(), []int{0}
}

func (m *Template) GetContainer() isTemplate_Container {
	if m != nil {
		return m.Container
	}
	return nil
}

func (x *Template) GetCrosDut() *CrosDutTemplate {
	if x, ok := x.GetContainer().(*Template_CrosDut); ok {
		return x.CrosDut
	}
	return nil
}

func (x *Template) GetCrosProvision() *CrosProvisionTemplate {
	if x, ok := x.GetContainer().(*Template_CrosProvision); ok {
		return x.CrosProvision
	}
	return nil
}

func (x *Template) GetCrosTest() *CrosTestTemplate {
	if x, ok := x.GetContainer().(*Template_CrosTest); ok {
		return x.CrosTest
	}
	return nil
}

type isTemplate_Container interface {
	isTemplate_Container()
}

type Template_CrosDut struct {
	CrosDut *CrosDutTemplate `protobuf:"bytes,1,opt,name=cros_dut,json=crosDut,proto3,oneof"`
}

type Template_CrosProvision struct {
	CrosProvision *CrosProvisionTemplate `protobuf:"bytes,2,opt,name=cros_provision,json=crosProvision,proto3,oneof"`
}

type Template_CrosTest struct {
	CrosTest *CrosTestTemplate `protobuf:"bytes,3,opt,name=cros_test,json=crosTest,proto3,oneof"`
}

func (*Template_CrosDut) isTemplate_Container() {}

func (*Template_CrosProvision) isTemplate_Container() {}

func (*Template_CrosTest) isTemplate_Container() {}

// Plain template to demonstrate the usage of cros-dut container. All fields are
// required.
type CrosDutTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an existing network to join
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	// Host directory to be mounted into the container for logs and other artifacts
	ArtifactDir string `protobuf:"bytes,2,opt,name=artifact_dir,json=artifactDir,proto3" json:"artifact_dir,omitempty"`
	// Cache server address
	CacheServer *api.IpEndpoint `protobuf:"bytes,3,opt,name=cache_server,json=cacheServer,proto3" json:"cache_server,omitempty"`
	// Dut ssh address
	DutAddress *api.IpEndpoint `protobuf:"bytes,4,opt,name=dut_address,json=dutAddress,proto3" json:"dut_address,omitempty"`
}

func (x *CrosDutTemplate) Reset() {
	*x = CrosDutTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrosDutTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrosDutTemplate) ProtoMessage() {}

func (x *CrosDutTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrosDutTemplate.ProtoReflect.Descriptor instead.
func (*CrosDutTemplate) Descriptor() ([]byte, []int) {
	return file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescGZIP(), []int{1}
}

func (x *CrosDutTemplate) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *CrosDutTemplate) GetArtifactDir() string {
	if x != nil {
		return x.ArtifactDir
	}
	return ""
}

func (x *CrosDutTemplate) GetCacheServer() *api.IpEndpoint {
	if x != nil {
		return x.CacheServer
	}
	return nil
}

func (x *CrosDutTemplate) GetDutAddress() *api.IpEndpoint {
	if x != nil {
		return x.DutAddress
	}
	return nil
}

// Plain template to demonstrate the usage of cros-provision container. All
// fields are required.
type CrosProvisionTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an existing network to join
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	// Host directory to be mounted into the container for logs and other artifacts
	ArtifactDir string `protobuf:"bytes,2,opt,name=artifact_dir,json=artifactDir,proto3" json:"artifact_dir,omitempty"`
	// CrosProvisionRequest payload can have template placeholders, to be populated
	// and write to a json file inside the artifact directory and pass along to
	// cros-provision server. Note that inventory_server is no longer needed
	InputRequest *api1.CrosProvisionRequest `protobuf:"bytes,3,opt,name=input_request,json=inputRequest,proto3" json:"input_request,omitempty"`
}

func (x *CrosProvisionTemplate) Reset() {
	*x = CrosProvisionTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrosProvisionTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrosProvisionTemplate) ProtoMessage() {}

func (x *CrosProvisionTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrosProvisionTemplate.ProtoReflect.Descriptor instead.
func (*CrosProvisionTemplate) Descriptor() ([]byte, []int) {
	return file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescGZIP(), []int{2}
}

func (x *CrosProvisionTemplate) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *CrosProvisionTemplate) GetArtifactDir() string {
	if x != nil {
		return x.ArtifactDir
	}
	return ""
}

func (x *CrosProvisionTemplate) GetInputRequest() *api1.CrosProvisionRequest {
	if x != nil {
		return x.InputRequest
	}
	return nil
}

// Plain template to demonstrate the usage of cros-test container. All fields
// are required.
type CrosTestTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an existing network to join
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	// Host directory, two sub-directories will be created and mounted into the
	// container. `<artifact_dir>/cros-test/cros-test` for logs and
	// `<artifact_dir>/cros-test/results` for result artifacts to be published
	ArtifactDir string `protobuf:"bytes,2,opt,name=artifact_dir,json=artifactDir,proto3" json:"artifact_dir,omitempty"`
}

func (x *CrosTestTemplate) Reset() {
	*x = CrosTestTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrosTestTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrosTestTemplate) ProtoMessage() {}

func (x *CrosTestTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrosTestTemplate.ProtoReflect.Descriptor instead.
func (*CrosTestTemplate) Descriptor() ([]byte, []int) {
	return file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescGZIP(), []int{3}
}

func (x *CrosTestTemplate) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *CrosTestTemplate) GetArtifactDir() string {
	if x != nil {
		return x.ArtifactDir
	}
	return ""
}

var File_infra_cros_cmd_cros_tool_runner_api_templates_proto protoreflect.FileDescriptor

var file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x63, 0x6d, 0x64,
	0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x2d, 0x72, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x74, 0x72, 0x76, 0x32, 0x2e, 0x61, 0x70, 0x69,
	0x1a, 0x2c, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x6c, 0x61, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x08, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x63, 0x72, 0x6f, 0x73, 0x5f, 0x64,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x74, 0x72, 0x76, 0x32,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x6f, 0x73, 0x44, 0x75, 0x74, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x07, 0x63, 0x72, 0x6f, 0x73, 0x44, 0x75, 0x74, 0x12,
	0x49, 0x0a, 0x0e, 0x63, 0x72, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x74, 0x72, 0x76, 0x32, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x72, 0x6f,
	0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x09, 0x63, 0x72,
	0x6f, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x63, 0x74, 0x72, 0x76, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x6f, 0x73, 0x54, 0x65,
	0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x08, 0x63, 0x72,
	0x6f, 0x73, 0x54, 0x65, 0x73, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x22, 0xdc, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x6f, 0x73, 0x44, 0x75, 0x74, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x64, 0x69,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x44, 0x69, 0x72, 0x12, 0x46, 0x0a, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6c, 0x61, 0x62,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x70, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x0b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0b,
	0x64, 0x75, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x70, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x75, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0xa4, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x44, 0x69, 0x72, 0x12, 0x4e, 0x0a, 0x0d, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x6f, 0x73, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x10, 0x43, 0x72, 0x6f,
	0x73, 0x54, 0x65, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x44, 0x69, 0x72, 0x42, 0x25, 0x5a, 0x23, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x72, 0x6f,
	0x73, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x2d, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescOnce sync.Once
	file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescData = file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDesc
)

func file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescGZIP() []byte {
	file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescOnce.Do(func() {
		file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescData)
	})
	return file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDescData
}

var file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_infra_cros_cmd_cros_tool_runner_api_templates_proto_goTypes = []interface{}{
	(*Template)(nil),                  // 0: ctrv2.api.Template
	(*CrosDutTemplate)(nil),           // 1: ctrv2.api.CrosDutTemplate
	(*CrosProvisionTemplate)(nil),     // 2: ctrv2.api.CrosProvisionTemplate
	(*CrosTestTemplate)(nil),          // 3: ctrv2.api.CrosTestTemplate
	(*api.IpEndpoint)(nil),            // 4: chromiumos.test.lab.api.IpEndpoint
	(*api1.CrosProvisionRequest)(nil), // 5: chromiumos.test.api.CrosProvisionRequest
}
var file_infra_cros_cmd_cros_tool_runner_api_templates_proto_depIdxs = []int32{
	1, // 0: ctrv2.api.Template.cros_dut:type_name -> ctrv2.api.CrosDutTemplate
	2, // 1: ctrv2.api.Template.cros_provision:type_name -> ctrv2.api.CrosProvisionTemplate
	3, // 2: ctrv2.api.Template.cros_test:type_name -> ctrv2.api.CrosTestTemplate
	4, // 3: ctrv2.api.CrosDutTemplate.cache_server:type_name -> chromiumos.test.lab.api.IpEndpoint
	4, // 4: ctrv2.api.CrosDutTemplate.dut_address:type_name -> chromiumos.test.lab.api.IpEndpoint
	5, // 5: ctrv2.api.CrosProvisionTemplate.input_request:type_name -> chromiumos.test.api.CrosProvisionRequest
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_infra_cros_cmd_cros_tool_runner_api_templates_proto_init() }
func file_infra_cros_cmd_cros_tool_runner_api_templates_proto_init() {
	if File_infra_cros_cmd_cros_tool_runner_api_templates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
		file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrosDutTemplate); i {
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
		file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrosProvisionTemplate); i {
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
		file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrosTestTemplate); i {
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
	file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Template_CrosDut)(nil),
		(*Template_CrosProvision)(nil),
		(*Template_CrosTest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_cros_cmd_cros_tool_runner_api_templates_proto_goTypes,
		DependencyIndexes: file_infra_cros_cmd_cros_tool_runner_api_templates_proto_depIdxs,
		MessageInfos:      file_infra_cros_cmd_cros_tool_runner_api_templates_proto_msgTypes,
	}.Build()
	File_infra_cros_cmd_cros_tool_runner_api_templates_proto = out.File
	file_infra_cros_cmd_cros_tool_runner_api_templates_proto_rawDesc = nil
	file_infra_cros_cmd_cros_tool_runner_api_templates_proto_goTypes = nil
	file_infra_cros_cmd_cros_tool_runner_api_templates_proto_depIdxs = nil
}
