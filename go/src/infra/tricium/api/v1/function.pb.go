// Copyright 2017 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/tricium/api/v1/function.proto

package tricium

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

// Originally, there were two types of functions; isolators and analyzers.
// For analyzer functions, the output must be of type Data.Results. After
// transition to recipe-based analyzers only, all functions should be analyzers
// and input type is ignored.
type Function_Type int32

const (
	Function_NONE Function_Type = 0
	// Deprecated: Marked as deprecated in infra/tricium/api/v1/function.proto.
	Function_ISOLATOR Function_Type = 1
	Function_ANALYZER Function_Type = 2
)

// Enum value maps for Function_Type.
var (
	Function_Type_name = map[int32]string{
		0: "NONE",
		1: "ISOLATOR",
		2: "ANALYZER",
	}
	Function_Type_value = map[string]int32{
		"NONE":     0,
		"ISOLATOR": 1,
		"ANALYZER": 2,
	}
)

func (x Function_Type) Enum() *Function_Type {
	p := new(Function_Type)
	*p = x
	return p
}

func (x Function_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Function_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_tricium_api_v1_function_proto_enumTypes[0].Descriptor()
}

func (Function_Type) Type() protoreflect.EnumType {
	return &file_infra_tricium_api_v1_function_proto_enumTypes[0]
}

func (x Function_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Function_Type.Descriptor instead.
func (Function_Type) EnumDescriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_function_proto_rawDescGZIP(), []int{0, 0}
}

// Tricium analyzer definition.
type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of this function.
	//
	// This field is required, but should always be ANALYZER.
	Type Function_Type `protobuf:"varint,1,opt,name=type,proto3,enum=tricium.Function_Type" json:"type,omitempty"`
	// The name of the analyzer.
	//
	// The name must be unique among Tricium functions within a Tricium instance.
	// The name is expected to be CamelCase; no spaces, underscores or dashes are
	// allowed.
	//
	// This field is required.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Data needed by this analyzer.
	//
	// After migration to recipe-based analyzers, this should always be
	// GIT_FILE_DETAILS.
	Needs Data_Type `protobuf:"varint,3,opt,name=needs,proto3,enum=tricium.Data_Type" json:"needs,omitempty"`
	// Data provided by this analyzer.
	//
	// This field is required.
	// After migration to recipe-based analyzers, this should always be
	// RESULTS.
	Provides Data_Type `protobuf:"varint,4,opt,name=provides,proto3,enum=tricium.Data_Type" json:"provides,omitempty"`
	// Path filters for this analyzer.
	//
	// Defined as a glob. The path filters only apply to the last part of the
	// path.
	PathFilters []string `protobuf:"bytes,5,rep,name=path_filters,json=pathFilters,proto3" json:"path_filters,omitempty"` // Default: "*"
	// Email address of the owner of this analyzer. Used for bug filing.
	//
	// This field is required.
	Owner string `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	// Monorail bug component for bug filing.
	//
	// This field is required.
	MonorailComponent string `protobuf:"bytes,7,opt,name=monorail_component,json=monorailComponent,proto3" json:"monorail_component,omitempty"`
	// Function implementations.
	//
	// Originally the idea was that an analyzer may run on many different platforms
	// and the comments from different platforms may be merged.
	//
	// This was not done in practice, so the number of impls should always be one.
	Impls []*Impl `protobuf:"bytes,9,rep,name=impls,proto3" json:"impls,omitempty"`
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_function_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_function_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_function_proto_rawDescGZIP(), []int{0}
}

func (x *Function) GetType() Function_Type {
	if x != nil {
		return x.Type
	}
	return Function_NONE
}

func (x *Function) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Function) GetNeeds() Data_Type {
	if x != nil {
		return x.Needs
	}
	return Data_NONE
}

func (x *Function) GetProvides() Data_Type {
	if x != nil {
		return x.Provides
	}
	return Data_NONE
}

func (x *Function) GetPathFilters() []string {
	if x != nil {
		return x.PathFilters
	}
	return nil
}

func (x *Function) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Function) GetMonorailComponent() string {
	if x != nil {
		return x.MonorailComponent
	}
	return ""
}

func (x *Function) GetImpls() []*Impl {
	if x != nil {
		return x.Impls
	}
	return nil
}

// Analyzer implementation.
//
// Implementation must be recipe-based; command-based (legacy) analyzers
// are no longer supported.
type Impl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data-dependency details specific to this implementation.
	//
	// This particular value of this field isn't significant, because
	// the platform is determined by the builder.
	NeedsForPlatform    Platform_Name `protobuf:"varint,1,opt,name=needs_for_platform,json=needsForPlatform,proto3,enum=tricium.Platform_Name" json:"needs_for_platform,omitempty"`
	ProvidesForPlatform Platform_Name `protobuf:"varint,2,opt,name=provides_for_platform,json=providesForPlatform,proto3,enum=tricium.Platform_Name" json:"provides_for_platform,omitempty"`
	// The platform to run this implementation on.
	//
	// This particular value of this field isn't significant, because
	// the platform is determined by the builder.
	RuntimePlatform Platform_Name `protobuf:"varint,3,opt,name=runtime_platform,json=runtimePlatform,proto3,enum=tricium.Platform_Name" json:"runtime_platform,omitempty"`
	// Types that are assignable to Impl:
	//
	//	*Impl_Recipe
	Impl isImpl_Impl `protobuf_oneof:"impl"`
	// Deadline for execution of corresponding workers.
	//
	// Note that this deadline includes the launch of a swarming task for the
	// corresponding worker, and collection of results from that worker.
	// Deadline should be given in seconds.
	Deadline int32 `protobuf:"varint,7,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *Impl) Reset() {
	*x = Impl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_function_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Impl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Impl) ProtoMessage() {}

func (x *Impl) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_function_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Impl.ProtoReflect.Descriptor instead.
func (*Impl) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_function_proto_rawDescGZIP(), []int{1}
}

func (x *Impl) GetNeedsForPlatform() Platform_Name {
	if x != nil {
		return x.NeedsForPlatform
	}
	return Platform_ANY
}

func (x *Impl) GetProvidesForPlatform() Platform_Name {
	if x != nil {
		return x.ProvidesForPlatform
	}
	return Platform_ANY
}

func (x *Impl) GetRuntimePlatform() Platform_Name {
	if x != nil {
		return x.RuntimePlatform
	}
	return Platform_ANY
}

func (m *Impl) GetImpl() isImpl_Impl {
	if m != nil {
		return m.Impl
	}
	return nil
}

func (x *Impl) GetRecipe() *Recipe {
	if x, ok := x.GetImpl().(*Impl_Recipe); ok {
		return x.Recipe
	}
	return nil
}

func (x *Impl) GetDeadline() int32 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

type isImpl_Impl interface {
	isImpl_Impl()
}

type Impl_Recipe struct {
	// Recipe for recipe-based implementation.
	Recipe *Recipe `protobuf:"bytes,5,opt,name=recipe,proto3,oneof"`
}

func (*Impl_Recipe) isImpl_Impl() {}

// Specification of a recipe for a recipe-based analyzer.
type Recipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Project ID, e.g. "chromium".
	Project string `protobuf:"bytes,5,opt,name=project,proto3" json:"project,omitempty"`
	// Bucket name, e.g. "try".
	Bucket string `protobuf:"bytes,6,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Builder name, e.g. "linux-rel".
	Builder string `protobuf:"bytes,7,opt,name=builder,proto3" json:"builder,omitempty"`
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_v1_function_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_v1_function_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_v1_function_proto_rawDescGZIP(), []int{2}
}

func (x *Recipe) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Recipe) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Recipe) GetBuilder() string {
	if x != nil {
		return x.Builder
	}
	return ""
}

var File_infra_tricium_api_v1_function_proto protoreflect.FileDescriptor

var file_infra_tricium_api_v1_function_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x1a, 0x1f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x02, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x05, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x6d, 0x70, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x49, 0x6d, 0x70,
	0x6c, 0x52, 0x05, 0x69, 0x6d, 0x70, 0x6c, 0x73, 0x22, 0x30, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x08, 0x49, 0x53,
	0x4f, 0x4c, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x01, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x4e, 0x41, 0x4c, 0x59, 0x5a, 0x45, 0x52, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09,
	0x22, 0xb6, 0x02, 0x0a, 0x04, 0x49, 0x6d, 0x70, 0x6c, 0x12, 0x44, 0x0a, 0x12, 0x6e, 0x65, 0x65,
	0x64, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x10, 0x6e,
	0x65, 0x65, 0x64, 0x73, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x4a, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73,
	0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x41, 0x0a, 0x10, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0f, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x29,
	0x0a, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x48,
	0x00, 0x52, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x69, 0x6d, 0x70, 0x6c, 0x4a, 0x04, 0x08,
	0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0x6c, 0x0a, 0x06, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10,
	0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b,
	0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_tricium_api_v1_function_proto_rawDescOnce sync.Once
	file_infra_tricium_api_v1_function_proto_rawDescData = file_infra_tricium_api_v1_function_proto_rawDesc
)

func file_infra_tricium_api_v1_function_proto_rawDescGZIP() []byte {
	file_infra_tricium_api_v1_function_proto_rawDescOnce.Do(func() {
		file_infra_tricium_api_v1_function_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_tricium_api_v1_function_proto_rawDescData)
	})
	return file_infra_tricium_api_v1_function_proto_rawDescData
}

var file_infra_tricium_api_v1_function_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_tricium_api_v1_function_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_infra_tricium_api_v1_function_proto_goTypes = []interface{}{
	(Function_Type)(0), // 0: tricium.Function.Type
	(*Function)(nil),   // 1: tricium.Function
	(*Impl)(nil),       // 2: tricium.Impl
	(*Recipe)(nil),     // 3: tricium.Recipe
	(Data_Type)(0),     // 4: tricium.Data.Type
	(Platform_Name)(0), // 5: tricium.Platform.Name
}
var file_infra_tricium_api_v1_function_proto_depIdxs = []int32{
	0, // 0: tricium.Function.type:type_name -> tricium.Function.Type
	4, // 1: tricium.Function.needs:type_name -> tricium.Data.Type
	4, // 2: tricium.Function.provides:type_name -> tricium.Data.Type
	2, // 3: tricium.Function.impls:type_name -> tricium.Impl
	5, // 4: tricium.Impl.needs_for_platform:type_name -> tricium.Platform.Name
	5, // 5: tricium.Impl.provides_for_platform:type_name -> tricium.Platform.Name
	5, // 6: tricium.Impl.runtime_platform:type_name -> tricium.Platform.Name
	3, // 7: tricium.Impl.recipe:type_name -> tricium.Recipe
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_infra_tricium_api_v1_function_proto_init() }
func file_infra_tricium_api_v1_function_proto_init() {
	if File_infra_tricium_api_v1_function_proto != nil {
		return
	}
	file_infra_tricium_api_v1_data_proto_init()
	file_infra_tricium_api_v1_platform_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infra_tricium_api_v1_function_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Function); i {
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
		file_infra_tricium_api_v1_function_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Impl); i {
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
		file_infra_tricium_api_v1_function_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipe); i {
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
	file_infra_tricium_api_v1_function_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Impl_Recipe)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_tricium_api_v1_function_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_tricium_api_v1_function_proto_goTypes,
		DependencyIndexes: file_infra_tricium_api_v1_function_proto_depIdxs,
		EnumInfos:         file_infra_tricium_api_v1_function_proto_enumTypes,
		MessageInfos:      file_infra_tricium_api_v1_function_proto_msgTypes,
	}.Build()
	File_infra_tricium_api_v1_function_proto = out.File
	file_infra_tricium_api_v1_function_proto_rawDesc = nil
	file_infra_tricium_api_v1_function_proto_goTypes = nil
	file_infra_tricium_api_v1_function_proto_depIdxs = nil
}
