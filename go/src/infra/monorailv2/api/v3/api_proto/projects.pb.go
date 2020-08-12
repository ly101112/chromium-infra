// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: api/v3/api_proto/projects.proto

package api_proto

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for CreateFieldDef method.
// Next available tag: 3
type CreateFieldDefRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The project resource where this field will be created.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The field to create.
	// It must have a display_name and a type with its corresponding settings.
	Fielddef *FieldDef `protobuf:"bytes,2,opt,name=fielddef,proto3" json:"fielddef,omitempty"`
}

func (x *CreateFieldDefRequest) Reset() {
	*x = CreateFieldDefRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v3_api_proto_projects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFieldDefRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFieldDefRequest) ProtoMessage() {}

func (x *CreateFieldDefRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v3_api_proto_projects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFieldDefRequest.ProtoReflect.Descriptor instead.
func (*CreateFieldDefRequest) Descriptor() ([]byte, []int) {
	return file_api_v3_api_proto_projects_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFieldDefRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateFieldDefRequest) GetFielddef() *FieldDef {
	if x != nil {
		return x.Fielddef
	}
	return nil
}

// Request message for ListIssueTemplates
// Next available tag: 4
type ListIssueTemplatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the project these templates belong to.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return. The service may return fewer than
	// this value.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListIssueTemplates` call.
	// Provide this to retrieve the subsequent page.
	// When paginating, all other parameters provided to
	// `ListIssueTemplatesRequest` must match the call that provided the token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListIssueTemplatesRequest) Reset() {
	*x = ListIssueTemplatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v3_api_proto_projects_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIssueTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIssueTemplatesRequest) ProtoMessage() {}

func (x *ListIssueTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v3_api_proto_projects_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIssueTemplatesRequest.ProtoReflect.Descriptor instead.
func (*ListIssueTemplatesRequest) Descriptor() ([]byte, []int) {
	return file_api_v3_api_proto_projects_proto_rawDescGZIP(), []int{1}
}

func (x *ListIssueTemplatesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListIssueTemplatesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListIssueTemplatesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message for ListIssueTemplates
// Next available tag: 3
type ListIssueTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Templates matching the given request.
	Templates []*IssueTemplate `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListIssueTemplatesResponse) Reset() {
	*x = ListIssueTemplatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v3_api_proto_projects_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIssueTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIssueTemplatesResponse) ProtoMessage() {}

func (x *ListIssueTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v3_api_proto_projects_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIssueTemplatesResponse.ProtoReflect.Descriptor instead.
func (*ListIssueTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_api_v3_api_proto_projects_proto_rawDescGZIP(), []int{2}
}

func (x *ListIssueTemplatesResponse) GetTemplates() []*IssueTemplate {
	if x != nil {
		return x.Templates
	}
	return nil
}

func (x *ListIssueTemplatesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request message for ListProjects
// Next available tag: 3
type ListProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of items to return. The service may return fewer than
	// this value.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListProjects` call.
	// Provide this to retrieve the subsequent page.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListProjectsRequest) Reset() {
	*x = ListProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v3_api_proto_projects_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsRequest) ProtoMessage() {}

func (x *ListProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v3_api_proto_projects_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return file_api_v3_api_proto_projects_proto_rawDescGZIP(), []int{3}
}

func (x *ListProjectsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListProjectsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message for ListProjects
// Next available tag: 3
type ListProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Projects matching the given request.
	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListProjectsResponse) Reset() {
	*x = ListProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v3_api_proto_projects_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse) ProtoMessage() {}

func (x *ListProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v3_api_proto_projects_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return file_api_v3_api_proto_projects_proto_rawDescGZIP(), []int{4}
}

func (x *ListProjectsResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *ListProjectsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_api_v3_api_proto_projects_proto protoreflect.FileDescriptor

var file_api_v3_api_proto_projects_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x33, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x44, 0x65, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x17, 0x0a, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x62, 0x75, 0x67, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x64, 0x65, 0x66, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x64, 0x65, 0x66, 0x22, 0x8e, 0x01, 0x0a, 0x19,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xfa, 0x41, 0x17, 0x0a, 0x15,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x62, 0x75, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7e, 0x0a, 0x1a,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x51, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x70, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f,
	0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x32, 0x99, 0x02, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x4d,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66,
	0x12, 0x22, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x22, 0x00, 0x12, 0x67, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x6f,
	0x6e, 0x6f, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x61, 0x69,
	0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x6f, 0x6e, 0x6f, 0x72,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a,
	0x10, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v3_api_proto_projects_proto_rawDescOnce sync.Once
	file_api_v3_api_proto_projects_proto_rawDescData = file_api_v3_api_proto_projects_proto_rawDesc
)

func file_api_v3_api_proto_projects_proto_rawDescGZIP() []byte {
	file_api_v3_api_proto_projects_proto_rawDescOnce.Do(func() {
		file_api_v3_api_proto_projects_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v3_api_proto_projects_proto_rawDescData)
	})
	return file_api_v3_api_proto_projects_proto_rawDescData
}

var file_api_v3_api_proto_projects_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v3_api_proto_projects_proto_goTypes = []interface{}{
	(*CreateFieldDefRequest)(nil),      // 0: monorail.v3.CreateFieldDefRequest
	(*ListIssueTemplatesRequest)(nil),  // 1: monorail.v3.ListIssueTemplatesRequest
	(*ListIssueTemplatesResponse)(nil), // 2: monorail.v3.ListIssueTemplatesResponse
	(*ListProjectsRequest)(nil),        // 3: monorail.v3.ListProjectsRequest
	(*ListProjectsResponse)(nil),       // 4: monorail.v3.ListProjectsResponse
	(*FieldDef)(nil),                   // 5: monorail.v3.FieldDef
	(*IssueTemplate)(nil),              // 6: monorail.v3.IssueTemplate
	(*Project)(nil),                    // 7: monorail.v3.Project
}
var file_api_v3_api_proto_projects_proto_depIdxs = []int32{
	5, // 0: monorail.v3.CreateFieldDefRequest.fielddef:type_name -> monorail.v3.FieldDef
	6, // 1: monorail.v3.ListIssueTemplatesResponse.templates:type_name -> monorail.v3.IssueTemplate
	7, // 2: monorail.v3.ListProjectsResponse.projects:type_name -> monorail.v3.Project
	0, // 3: monorail.v3.Projects.CreateFieldDef:input_type -> monorail.v3.CreateFieldDefRequest
	1, // 4: monorail.v3.Projects.ListIssueTemplates:input_type -> monorail.v3.ListIssueTemplatesRequest
	3, // 5: monorail.v3.Projects.ListProjects:input_type -> monorail.v3.ListProjectsRequest
	5, // 6: monorail.v3.Projects.CreateFieldDef:output_type -> monorail.v3.FieldDef
	2, // 7: monorail.v3.Projects.ListIssueTemplates:output_type -> monorail.v3.ListIssueTemplatesResponse
	4, // 8: monorail.v3.Projects.ListProjects:output_type -> monorail.v3.ListProjectsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v3_api_proto_projects_proto_init() }
func file_api_v3_api_proto_projects_proto_init() {
	if File_api_v3_api_proto_projects_proto != nil {
		return
	}
	file_api_v3_api_proto_project_objects_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v3_api_proto_projects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFieldDefRequest); i {
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
		file_api_v3_api_proto_projects_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIssueTemplatesRequest); i {
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
		file_api_v3_api_proto_projects_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIssueTemplatesResponse); i {
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
		file_api_v3_api_proto_projects_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsRequest); i {
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
		file_api_v3_api_proto_projects_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsResponse); i {
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
			RawDescriptor: file_api_v3_api_proto_projects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v3_api_proto_projects_proto_goTypes,
		DependencyIndexes: file_api_v3_api_proto_projects_proto_depIdxs,
		MessageInfos:      file_api_v3_api_proto_projects_proto_msgTypes,
	}.Build()
	File_api_v3_api_proto_projects_proto = out.File
	file_api_v3_api_proto_projects_proto_rawDesc = nil
	file_api_v3_api_proto_projects_proto_goTypes = nil
	file_api_v3_api_proto_projects_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProjectsClient is the client API for Projects service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectsClient interface {
	// Creates a new fieldDef (custom field).
	//
	// Raises:
	//   NOT_FOUND if some given users do not exist.
	//   ALREADY_EXISTS if a field with the same name owned by the project
	//   already exists.
	//   INVALID_INPUT if there was a problem with the input.
	//   PERMISSION_DENIED if the user cannot edit the project.
	CreateFieldDef(ctx context.Context, in *CreateFieldDefRequest, opts ...grpc.CallOption) (*FieldDef, error)
	// Returns all templates for specified project.
	//
	// Raises:
	//   NOT_FOUND if the requested parent project is not found.
	//   INVALID_ARGUMENT if the given `parent` is not valid.
	ListIssueTemplates(ctx context.Context, in *ListIssueTemplatesRequest, opts ...grpc.CallOption) (*ListIssueTemplatesResponse, error)
	// Returns all projects hosted on Monorail.
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
}
type projectsPRPCClient struct {
	client *prpc.Client
}

func NewProjectsPRPCClient(client *prpc.Client) ProjectsClient {
	return &projectsPRPCClient{client}
}

func (c *projectsPRPCClient) CreateFieldDef(ctx context.Context, in *CreateFieldDefRequest, opts ...grpc.CallOption) (*FieldDef, error) {
	out := new(FieldDef)
	err := c.client.Call(ctx, "monorail.v3.Projects", "CreateFieldDef", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsPRPCClient) ListIssueTemplates(ctx context.Context, in *ListIssueTemplatesRequest, opts ...grpc.CallOption) (*ListIssueTemplatesResponse, error) {
	out := new(ListIssueTemplatesResponse)
	err := c.client.Call(ctx, "monorail.v3.Projects", "ListIssueTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsPRPCClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.client.Call(ctx, "monorail.v3.Projects", "ListProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type projectsClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectsClient(cc grpc.ClientConnInterface) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) CreateFieldDef(ctx context.Context, in *CreateFieldDefRequest, opts ...grpc.CallOption) (*FieldDef, error) {
	out := new(FieldDef)
	err := c.cc.Invoke(ctx, "/monorail.v3.Projects/CreateFieldDef", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) ListIssueTemplates(ctx context.Context, in *ListIssueTemplatesRequest, opts ...grpc.CallOption) (*ListIssueTemplatesResponse, error) {
	out := new(ListIssueTemplatesResponse)
	err := c.cc.Invoke(ctx, "/monorail.v3.Projects/ListIssueTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, "/monorail.v3.Projects/ListProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectsServer is the server API for Projects service.
type ProjectsServer interface {
	// Creates a new fieldDef (custom field).
	//
	// Raises:
	//   NOT_FOUND if some given users do not exist.
	//   ALREADY_EXISTS if a field with the same name owned by the project
	//   already exists.
	//   INVALID_INPUT if there was a problem with the input.
	//   PERMISSION_DENIED if the user cannot edit the project.
	CreateFieldDef(context.Context, *CreateFieldDefRequest) (*FieldDef, error)
	// Returns all templates for specified project.
	//
	// Raises:
	//   NOT_FOUND if the requested parent project is not found.
	//   INVALID_ARGUMENT if the given `parent` is not valid.
	ListIssueTemplates(context.Context, *ListIssueTemplatesRequest) (*ListIssueTemplatesResponse, error)
	// Returns all projects hosted on Monorail.
	ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
}

// UnimplementedProjectsServer can be embedded to have forward compatible implementations.
type UnimplementedProjectsServer struct {
}

func (*UnimplementedProjectsServer) CreateFieldDef(context.Context, *CreateFieldDefRequest) (*FieldDef, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFieldDef not implemented")
}
func (*UnimplementedProjectsServer) ListIssueTemplates(context.Context, *ListIssueTemplatesRequest) (*ListIssueTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIssueTemplates not implemented")
}
func (*UnimplementedProjectsServer) ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}

func RegisterProjectsServer(s prpc.Registrar, srv ProjectsServer) {
	s.RegisterService(&_Projects_serviceDesc, srv)
}

func _Projects_CreateFieldDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFieldDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).CreateFieldDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v3.Projects/CreateFieldDef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).CreateFieldDef(ctx, req.(*CreateFieldDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_ListIssueTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIssueTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).ListIssueTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v3.Projects/ListIssueTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).ListIssueTemplates(ctx, req.(*ListIssueTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.v3.Projects/ListProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).ListProjects(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Projects_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monorail.v3.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFieldDef",
			Handler:    _Projects_CreateFieldDef_Handler,
		},
		{
			MethodName: "ListIssueTemplates",
			Handler:    _Projects_ListIssueTemplates_Handler,
		},
		{
			MethodName: "ListProjects",
			Handler:    _Projects_ListProjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v3/api_proto/projects.proto",
}
