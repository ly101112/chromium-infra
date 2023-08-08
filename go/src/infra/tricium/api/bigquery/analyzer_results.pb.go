// Copyright 2019 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: infra/tricium/api/bigquery/analyzer_results.proto

package apibq

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type AnalysisRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The revision information for the Gerrit change being analyzed by this
	// analysis run.
	GerritRevision *v1.GerritRevision `protobuf:"bytes,1,opt,name=gerrit_revision,json=gerritRevision,proto3" json:"gerrit_revision,omitempty"`
	// The revision number. In Gerrit this is the change revision and is
	// displayed as the patchset number in PolyGerrit.
	RevisionNumber int32 `protobuf:"varint,2,opt,name=revision_number,json=revisionNumber,proto3" json:"revision_number,omitempty"`
	// All files in the change revision analyzed by the run.
	Files []*v1.Data_File `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
	// Time when the request was received.
	RequestedTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=requested_time,json=requestedTime,proto3" json:"requested_time,omitempty"`
	// Platform for which the result applies.
	ResultPlatform v1.Platform_Name `protobuf:"varint,5,opt,name=result_platform,json=resultPlatform,proto3,enum=tricium.Platform_Name" json:"result_platform,omitempty"`
	// Overall state for the run result. As results are only sent after
	// completion PENDING and RUNNING would never be used.
	ResultState v1.State `protobuf:"varint,6,opt,name=result_state,json=resultState,proto3,enum=tricium.State" json:"result_state,omitempty"`
	// Comments added to Gerrit during this analysis run.
	Comments []*AnalysisRun_GerritComment `protobuf:"bytes,7,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *AnalysisRun) Reset() {
	*x = AnalysisRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_bigquery_analyzer_results_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisRun) ProtoMessage() {}

func (x *AnalysisRun) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_bigquery_analyzer_results_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisRun.ProtoReflect.Descriptor instead.
func (*AnalysisRun) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_bigquery_analyzer_results_proto_rawDescGZIP(), []int{0}
}

func (x *AnalysisRun) GetGerritRevision() *v1.GerritRevision {
	if x != nil {
		return x.GerritRevision
	}
	return nil
}

func (x *AnalysisRun) GetRevisionNumber() int32 {
	if x != nil {
		return x.RevisionNumber
	}
	return 0
}

func (x *AnalysisRun) GetFiles() []*v1.Data_File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *AnalysisRun) GetRequestedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RequestedTime
	}
	return nil
}

func (x *AnalysisRun) GetResultPlatform() v1.Platform_Name {
	if x != nil {
		return x.ResultPlatform
	}
	return v1.Platform_ANY
}

func (x *AnalysisRun) GetResultState() v1.State {
	if x != nil {
		return x.ResultState
	}
	return v1.State_PENDING
}

func (x *AnalysisRun) GetComments() []*AnalysisRun_GerritComment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type AnalysisRun_GerritComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The comment generated by the analysis run.
	Comment *v1.Data_Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	// Comment creation time.
	//
	// Comment creation time in terms of when it is tracked in the service not
	// when it is created by the analyzer.
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// Analyzer function name.
	//
	// e.g., "ClangTidy".
	Analyzer string `protobuf:"bytes,3,opt,name=analyzer,proto3" json:"analyzer,omitempty"`
	// Platforms this comment applies to.
	Platforms []v1.Platform_Name `protobuf:"varint,4,rep,packed,name=platforms,proto3,enum=tricium.Platform_Name" json:"platforms,omitempty"`
	// Has this comment been selected to be displayed on the review?
	// For example, comments outside of the changed lines are not included
	// (i.e. selected) in the results posted to the Gerrit change.
	Selected bool `protobuf:"varint,5,opt,name=selected,proto3" json:"selected,omitempty"`
}

func (x *AnalysisRun_GerritComment) Reset() {
	*x = AnalysisRun_GerritComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_tricium_api_bigquery_analyzer_results_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisRun_GerritComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisRun_GerritComment) ProtoMessage() {}

func (x *AnalysisRun_GerritComment) ProtoReflect() protoreflect.Message {
	mi := &file_infra_tricium_api_bigquery_analyzer_results_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisRun_GerritComment.ProtoReflect.Descriptor instead.
func (*AnalysisRun_GerritComment) Descriptor() ([]byte, []int) {
	return file_infra_tricium_api_bigquery_analyzer_results_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AnalysisRun_GerritComment) GetComment() *v1.Data_Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *AnalysisRun_GerritComment) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *AnalysisRun_GerritComment) GetAnalyzer() string {
	if x != nil {
		return x.Analyzer
	}
	return ""
}

func (x *AnalysisRun_GerritComment) GetPlatforms() []v1.Platform_Name {
	if x != nil {
		return x.Platforms
	}
	return nil
}

func (x *AnalysisRun_GerritComment) GetSelected() bool {
	if x != nil {
		return x.Selected
	}
	return false
}

var File_infra_tricium_api_bigquery_analyzer_results_proto protoreflect.FileDescriptor

var file_infra_tricium_api_bigquery_analyzer_results_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x70, 0x69, 0x62, 0x71, 0x1a, 0x1f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x05, 0x0a, 0x0b, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x52, 0x75, 0x6e, 0x12, 0x40, 0x0a, 0x0f, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x28, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3f, 0x0a,
	0x0f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d,
	0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0e,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x31,
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x62, 0x71, 0x2e, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0xed, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x12, 0x34, 0x0a,
	0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x75, 0x6d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_tricium_api_bigquery_analyzer_results_proto_rawDescOnce sync.Once
	file_infra_tricium_api_bigquery_analyzer_results_proto_rawDescData = file_infra_tricium_api_bigquery_analyzer_results_proto_rawDesc
)

func file_infra_tricium_api_bigquery_analyzer_results_proto_rawDescGZIP() []byte {
	file_infra_tricium_api_bigquery_analyzer_results_proto_rawDescOnce.Do(func() {
		file_infra_tricium_api_bigquery_analyzer_results_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_tricium_api_bigquery_analyzer_results_proto_rawDescData)
	})
	return file_infra_tricium_api_bigquery_analyzer_results_proto_rawDescData
}

var file_infra_tricium_api_bigquery_analyzer_results_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_tricium_api_bigquery_analyzer_results_proto_goTypes = []interface{}{
	(*AnalysisRun)(nil),               // 0: apibq.AnalysisRun
	(*AnalysisRun_GerritComment)(nil), // 1: apibq.AnalysisRun.GerritComment
	(*v1.GerritRevision)(nil),         // 2: tricium.GerritRevision
	(*v1.Data_File)(nil),              // 3: tricium.Data.File
	(*timestamppb.Timestamp)(nil),     // 4: google.protobuf.Timestamp
	(v1.Platform_Name)(0),             // 5: tricium.Platform.Name
	(v1.State)(0),                     // 6: tricium.State
	(*v1.Data_Comment)(nil),           // 7: tricium.Data.Comment
}
var file_infra_tricium_api_bigquery_analyzer_results_proto_depIdxs = []int32{
	2, // 0: apibq.AnalysisRun.gerrit_revision:type_name -> tricium.GerritRevision
	3, // 1: apibq.AnalysisRun.files:type_name -> tricium.Data.File
	4, // 2: apibq.AnalysisRun.requested_time:type_name -> google.protobuf.Timestamp
	5, // 3: apibq.AnalysisRun.result_platform:type_name -> tricium.Platform.Name
	6, // 4: apibq.AnalysisRun.result_state:type_name -> tricium.State
	1, // 5: apibq.AnalysisRun.comments:type_name -> apibq.AnalysisRun.GerritComment
	7, // 6: apibq.AnalysisRun.GerritComment.comment:type_name -> tricium.Data.Comment
	4, // 7: apibq.AnalysisRun.GerritComment.created_time:type_name -> google.protobuf.Timestamp
	5, // 8: apibq.AnalysisRun.GerritComment.platforms:type_name -> tricium.Platform.Name
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_infra_tricium_api_bigquery_analyzer_results_proto_init() }
func file_infra_tricium_api_bigquery_analyzer_results_proto_init() {
	if File_infra_tricium_api_bigquery_analyzer_results_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_tricium_api_bigquery_analyzer_results_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisRun); i {
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
		file_infra_tricium_api_bigquery_analyzer_results_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisRun_GerritComment); i {
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
			RawDescriptor: file_infra_tricium_api_bigquery_analyzer_results_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_tricium_api_bigquery_analyzer_results_proto_goTypes,
		DependencyIndexes: file_infra_tricium_api_bigquery_analyzer_results_proto_depIdxs,
		MessageInfos:      file_infra_tricium_api_bigquery_analyzer_results_proto_msgTypes,
	}.Build()
	File_infra_tricium_api_bigquery_analyzer_results_proto = out.File
	file_infra_tricium_api_bigquery_analyzer_results_proto_rawDesc = nil
	file_infra_tricium_api_bigquery_analyzer_results_proto_goTypes = nil
	file_infra_tricium_api_bigquery_analyzer_results_proto_depIdxs = nil
}
