// Copyright 2019 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/appengine/cros/lab_inventory/api/bigquery/changehistory.proto

package apibq

import (
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

type ChangeHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hostname    string                 `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Label       string                 `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"` // What is changed, e.g. servo_port.
	OldValue    string                 `protobuf:"bytes,4,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	NewValue    string                 `protobuf:"bytes,5,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	ByWhom      *ChangeHistory_User    `protobuf:"bytes,7,opt,name=by_whom,json=byWhom,proto3" json:"by_whom,omitempty"`
	Comment     string                 `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *ChangeHistory) Reset() {
	*x = ChangeHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeHistory) ProtoMessage() {}

func (x *ChangeHistory) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeHistory.ProtoReflect.Descriptor instead.
func (*ChangeHistory) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeHistory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangeHistory) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ChangeHistory) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ChangeHistory) GetOldValue() string {
	if x != nil {
		return x.OldValue
	}
	return ""
}

func (x *ChangeHistory) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

func (x *ChangeHistory) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *ChangeHistory) GetByWhom() *ChangeHistory_User {
	if x != nil {
		return x.ByWhom
	}
	return nil
}

func (x *ChangeHistory) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type ChangeHistory_User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ChangeHistory_User) Reset() {
	*x = ChangeHistory_User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeHistory_User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeHistory_User) ProtoMessage() {}

func (x *ChangeHistory_User) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeHistory_User.ProtoReflect.Descriptor instead.
func (*ChangeHistory_User) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ChangeHistory_User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChangeHistory_User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto protoreflect.FileDescriptor

var file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDesc = []byte{
	0x0a, 0x43, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x70, 0x69, 0x62, 0x71, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x02,
	0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x62, 0x79,
	0x5f, 0x77, 0x68, 0x6f, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x62, 0x71, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x62, 0x79, 0x57, 0x68, 0x6f, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x30, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x37, 0x5a, 0x35, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x63, 0x72,
	0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3b, 0x61, 0x70,
	0x69, 0x62, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDescOnce sync.Once
	file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDescData = file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDesc
)

func file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDescGZIP() []byte {
	file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDescOnce.Do(func() {
		file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDescData)
	})
	return file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDescData
}

var file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_goTypes = []interface{}{
	(*ChangeHistory)(nil),         // 0: apibq.ChangeHistory
	(*ChangeHistory_User)(nil),    // 1: apibq.ChangeHistory.User
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_depIdxs = []int32{
	2, // 0: apibq.ChangeHistory.updated_time:type_name -> google.protobuf.Timestamp
	1, // 1: apibq.ChangeHistory.by_whom:type_name -> apibq.ChangeHistory.User
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_init() }
func file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_init() {
	if File_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeHistory); i {
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
		file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeHistory_User); i {
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
			RawDescriptor: file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_goTypes,
		DependencyIndexes: file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_depIdxs,
		MessageInfos:      file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_msgTypes,
	}.Build()
	File_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto = out.File
	file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_rawDesc = nil
	file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_goTypes = nil
	file_infra_appengine_cros_lab_inventory_api_bigquery_changehistory_proto_depIdxs = nil
}
