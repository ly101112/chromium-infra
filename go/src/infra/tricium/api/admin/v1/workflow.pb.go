// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: infra/tricium/api/admin/v1/workflow.proto

package admin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "infra/tricium/api/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Tricium workflow configuration.
//
// A Workflow is generated from a merged service and project config,
// and contains information required for one workflow run.
type Workflow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workers []*Worker `protobuf:"bytes,2,rep,name=workers,proto3" json:"workers,omitempty"`
	// Function definitions used for this workflow; these contain the function
	// owner and component, to be used when filling out a bug filing template.
	Functions             []*v1.Function `protobuf:"bytes,5,rep,name=functions,proto3" json:"functions,omitempty"`
	BuildbucketServerHost string         `protobuf:"bytes,6,opt,name=buildbucket_server_host,json=buildbucketServerHost,proto3" json:"buildbucket_server_host,omitempty"`
}

func (x *Workflow) Reset() {
	*x = Workflow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_admin_v1_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workflow) ProtoMessage() {}

func (x *Workflow) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_admin_v1_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workflow.ProtoReflect.Descriptor instead.
func (*Workflow) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_admin_v1_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *Workflow) GetWorkers() []*Worker {
	if x != nil {
		return x.Workers
	}
	return nil
}

func (x *Workflow) GetFunctions() []*v1.Function {
	if x != nil {
		return x.Functions
	}
	return nil
}

func (x *Workflow) GetBuildbucketServerHost() string {
	if x != nil {
		return x.BuildbucketServerHost
	}
	return ""
}

// A Tricium worker includes the details needed to execute an analyzer.
type Worker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of worker is combination of the function and platform name
	// for which results are provided, e.g "GitFileIsolator_LINUX".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Includes data dependencies for runtime type checking.
	// Platform-specific details are provided when required by the corresponding
	// data type.
	Needs               v1.Data_Type     `protobuf:"varint,2,opt,name=needs,proto3,enum=tricium.Data_Type" json:"needs,omitempty"`
	NeedsForPlatform    v1.Platform_Name `protobuf:"varint,3,opt,name=needs_for_platform,json=needsForPlatform,proto3,enum=tricium.Platform_Name" json:"needs_for_platform,omitempty"`
	Provides            v1.Data_Type     `protobuf:"varint,4,opt,name=provides,proto3,enum=tricium.Data_Type" json:"provides,omitempty"`
	ProvidesForPlatform v1.Platform_Name `protobuf:"varint,5,opt,name=provides_for_platform,json=providesForPlatform,proto3,enum=tricium.Platform_Name" json:"provides_for_platform,omitempty"`
	// Workers to run after this one.
	//
	// For recipe-based analyzers this should be empty.
	Next []string `protobuf:"bytes,6,rep,name=next,proto3" json:"next,omitempty"`
	// Name of the runtime platform configuration.
	RuntimePlatform v1.Platform_Name `protobuf:"varint,7,opt,name=runtime_platform,json=runtimePlatform,proto3,enum=tricium.Platform_Name" json:"runtime_platform,omitempty"`
	// Types that are assignable to Impl:
	//	*Worker_Recipe
	Impl isWorker_Impl `protobuf_oneof:"impl"`
}

func (x *Worker) Reset() {
	*x = Worker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_admin_v1_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Worker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Worker) ProtoMessage() {}

func (x *Worker) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_admin_v1_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Worker.ProtoReflect.Descriptor instead.
func (*Worker) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_admin_v1_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *Worker) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Worker) GetNeeds() v1.Data_Type {
	if x != nil {
		return x.Needs
	}
	return v1.Data_NONE
}

func (x *Worker) GetNeedsForPlatform() v1.Platform_Name {
	if x != nil {
		return x.NeedsForPlatform
	}
	return v1.Platform_ANY
}

func (x *Worker) GetProvides() v1.Data_Type {
	if x != nil {
		return x.Provides
	}
	return v1.Data_NONE
}

func (x *Worker) GetProvidesForPlatform() v1.Platform_Name {
	if x != nil {
		return x.ProvidesForPlatform
	}
	return v1.Platform_ANY
}

func (x *Worker) GetNext() []string {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *Worker) GetRuntimePlatform() v1.Platform_Name {
	if x != nil {
		return x.RuntimePlatform
	}
	return v1.Platform_ANY
}

func (m *Worker) GetImpl() isWorker_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (x *Worker) GetRecipe() *v1.Recipe {
	if x, ok := x.GetImpl().(*Worker_Recipe); ok {
		return x.Recipe
	}
	return nil
}

type isWorker_Impl interface {
	isWorker_Impl()
}

type Worker_Recipe struct {
	// Recipe for recipe-based implementation.
	Recipe *v1.Recipe `protobuf:"bytes,12,opt,name=recipe,proto3,oneof"`
}

func (*Worker_Recipe) isWorker_Impl() {}

var File_infra_tricium_api_admin_v1_workflow_proto protoreflect.FileDescriptor

var file_infra_tricium_api_admin_v1_workflow_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x1a, 0x1f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75,
	0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69,
	0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f,
	0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01,
	0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x27, 0x0a, 0x07, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d,
	0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x02, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x22, 0xaa,
	0x03, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74,
	0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x05, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x12, 0x44, 0x0a, 0x12, 0x6e, 0x65, 0x65, 0x64, 0x73,
	0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x10, 0x6e, 0x65, 0x65,
	0x64, 0x73, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2e, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x12, 0x4a, 0x0a,
	0x15, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74,
	0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x46, 0x6f,
	0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78,
	0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x41, 0x0a,
	0x10, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75,
	0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x69,
	0x6d, 0x70, 0x6c, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a, 0x4a,
	0x04, 0x08, 0x0a, 0x10, 0x0b, 0x4a, 0x04, 0x08, 0x0b, 0x10, 0x0c, 0x42, 0x22, 0x5a, 0x20, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_tricium_api_admin_v1_workflow_proto_rawDescOnce sync.Once
	file_infra_tricium_api_admin_v1_workflow_proto_rawDescData = file_infra_tricium_api_admin_v1_workflow_proto_rawDesc
)

func file_infra_tricium_api_admin_v1_workflow_proto_rawDescGZIP() []byte {
	file_infra_tricium_api_admin_v1_workflow_proto_rawDescOnce.Do(func() {
		file_infra_tricium_api_admin_v1_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_tricium_api_admin_v1_workflow_proto_rawDescData)
	})
	return file_infra_tricium_api_admin_v1_workflow_proto_rawDescData
}

var file_infra_tricium_api_admin_v1_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_tricium_api_admin_v1_workflow_proto_goTypes = []interface{}{
	(*Workflow)(nil),      // 0: admin.Workflow
	(*Worker)(nil),        // 1: admin.Worker
	(*v1.Function)(nil),   // 2: tricium.Function
	(v1.Data_Type)(0),     // 3: tricium.Data.Type
	(v1.Platform_Name)(0), // 4: tricium.Platform.Name
	(*v1.Recipe)(nil),     // 5: tricium.Recipe
}
var file_infra_tricium_api_admin_v1_workflow_proto_depIdxs = []int32{
	1, // 0: admin.Workflow.workers:type_name -> admin.Worker
	2, // 1: admin.Workflow.functions:type_name -> tricium.Function
	3, // 2: admin.Worker.needs:type_name -> tricium.Data.Type
	4, // 3: admin.Worker.needs_for_platform:type_name -> tricium.Platform.Name
	3, // 4: admin.Worker.provides:type_name -> tricium.Data.Type
	4, // 5: admin.Worker.provides_for_platform:type_name -> tricium.Platform.Name
	4, // 6: admin.Worker.runtime_platform:type_name -> tricium.Platform.Name
	5, // 7: admin.Worker.recipe:type_name -> tricium.Recipe
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_infra_tricium_api_admin_v1_workflow_proto_init() }
func file_infra_tricium_api_admin_v1_workflow_proto_init() {
	if File_infra_tricium_api_admin_v1_workflow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_tricium_api_admin_v1_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workflow); i {
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
		file_infra_tricium_api_admin_v1_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Worker); i {
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
	file_infra_tricium_api_admin_v1_workflow_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Worker_Recipe)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_tricium_api_admin_v1_workflow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_tricium_api_admin_v1_workflow_proto_goTypes,
		DependencyIndexes: file_infra_tricium_api_admin_v1_workflow_proto_depIdxs,
		MessageInfos:      file_infra_tricium_api_admin_v1_workflow_proto_msgTypes,
	}.Build()
	File_infra_tricium_api_admin_v1_workflow_proto = out.File
	file_infra_tricium_api_admin_v1_workflow_proto_rawDesc = nil
	file_infra_tricium_api_admin_v1_workflow_proto_goTypes = nil
	file_infra_tricium_api_admin_v1_workflow_proto_depIdxs = nil
}
