// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Copied from https://source.corp.google.com/piper///depot/google3/java/com/google/chrome/crosbuilds/backend/hartapi/proto/pubsub/assetinfo/assetinfo.proto
// and edited

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: infra/unifiedfleet/api/v1/models/assetinfo.proto

package ufspb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

//TODO(anushruth): Use 0 as Undefined inline with https://google.aip.dev/126
type RequestStatus int32

const (
	RequestStatus_OK              RequestStatus = 0
	RequestStatus_INVALID_MESSAGE RequestStatus = 1
	RequestStatus_SERVER_ERROR    RequestStatus = 2
	RequestStatus_NO_ASSET_TAG    RequestStatus = 3
	RequestStatus_NO_GPN          RequestStatus = 4
)

// Enum value maps for RequestStatus.
var (
	RequestStatus_name = map[int32]string{
		0: "OK",
		1: "INVALID_MESSAGE",
		2: "SERVER_ERROR",
		3: "NO_ASSET_TAG",
		4: "NO_GPN",
	}
	RequestStatus_value = map[string]int32{
		"OK":              0,
		"INVALID_MESSAGE": 1,
		"SERVER_ERROR":    2,
		"NO_ASSET_TAG":    3,
		"NO_GPN":          4,
	}
)

func (x RequestStatus) Enum() *RequestStatus {
	p := new(RequestStatus)
	*p = x
	return p
}

func (x RequestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_unifiedfleet_api_v1_models_assetinfo_proto_enumTypes[0].Descriptor()
}

func (RequestStatus) Type() protoreflect.EnumType {
	return &file_infra_unifiedfleet_api_v1_models_assetinfo_proto_enumTypes[0]
}

func (x RequestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestStatus.Descriptor instead.
func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescGZIP(), []int{0}
}

type AssetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetTag string `protobuf:"bytes,1,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	// User not allowed to update this field. SSW will update this field.
	SerialNumber       string `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	CostCenter         string `protobuf:"bytes,3,opt,name=cost_center,json=costCenter,proto3" json:"cost_center,omitempty"`
	GoogleCodeName     string `protobuf:"bytes,4,opt,name=google_code_name,json=googleCodeName,proto3" json:"google_code_name,omitempty"`
	Model              string `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
	BuildTarget        string `protobuf:"bytes,6,opt,name=build_target,json=buildTarget,proto3" json:"build_target,omitempty"`
	ReferenceBoard     string `protobuf:"bytes,7,opt,name=reference_board,json=referenceBoard,proto3" json:"reference_board,omitempty"`
	EthernetMacAddress string `protobuf:"bytes,8,opt,name=ethernet_mac_address,json=ethernetMacAddress,proto3" json:"ethernet_mac_address,omitempty"`
	// User not allowed to update this field. SSW will update this field.
	Sku   string `protobuf:"bytes,9,opt,name=sku,proto3" json:"sku,omitempty"`
	Phase string `protobuf:"bytes,10,opt,name=phase,proto3" json:"phase,omitempty"`
	// User not allowed to update this field. SSW will update this field.
	Hwid string `protobuf:"bytes,11,opt,name=hwid,proto3" json:"hwid,omitempty"`
}

func (x *AssetInfo) Reset() {
	*x = AssetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetInfo) ProtoMessage() {}

func (x *AssetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetInfo.ProtoReflect.Descriptor instead.
func (*AssetInfo) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescGZIP(), []int{0}
}

func (x *AssetInfo) GetAssetTag() string {
	if x != nil {
		return x.AssetTag
	}
	return ""
}

func (x *AssetInfo) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *AssetInfo) GetCostCenter() string {
	if x != nil {
		return x.CostCenter
	}
	return ""
}

func (x *AssetInfo) GetGoogleCodeName() string {
	if x != nil {
		return x.GoogleCodeName
	}
	return ""
}

func (x *AssetInfo) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *AssetInfo) GetBuildTarget() string {
	if x != nil {
		return x.BuildTarget
	}
	return ""
}

func (x *AssetInfo) GetReferenceBoard() string {
	if x != nil {
		return x.ReferenceBoard
	}
	return ""
}

func (x *AssetInfo) GetEthernetMacAddress() string {
	if x != nil {
		return x.EthernetMacAddress
	}
	return ""
}

func (x *AssetInfo) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *AssetInfo) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

func (x *AssetInfo) GetHwid() string {
	if x != nil {
		return x.Hwid
	}
	return ""
}

type AssetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Asset tag(s) to lookup.
	AssetTags []string `protobuf:"bytes,1,rep,name=asset_tags,json=assetTags,proto3" json:"asset_tags,omitempty"`
}

func (x *AssetInfoRequest) Reset() {
	*x = AssetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetInfoRequest) ProtoMessage() {}

func (x *AssetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetInfoRequest.ProtoReflect.Descriptor instead.
func (*AssetInfoRequest) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescGZIP(), []int{1}
}

func (x *AssetInfoRequest) GetAssetTags() []string {
	if x != nil {
		return x.AssetTags
	}
	return nil
}

type AssetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicator of whether the request call succeed or not.
	RequestStatus RequestStatus `protobuf:"varint,1,opt,name=request_status,json=requestStatus,proto3,enum=unifiedfleet.api.v1.models.RequestStatus" json:"request_status,omitempty"`
	// A list of asset info retrieved from Hart.
	Assets []*AssetInfo `protobuf:"bytes,2,rep,name=assets,proto3" json:"assets,omitempty"`
	// A list of asset tags not found in Corporate Asset Management platform.
	MissingAssetTags []string `protobuf:"bytes,3,rep,name=missing_asset_tags,json=missingAssetTags,proto3" json:"missing_asset_tags,omitempty"`
	// A list of asset tags whose part number not found in Hart.
	FailedAssetTags []string `protobuf:"bytes,4,rep,name=failed_asset_tags,json=failedAssetTags,proto3" json:"failed_asset_tags,omitempty"`
}

func (x *AssetInfoResponse) Reset() {
	*x = AssetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetInfoResponse) ProtoMessage() {}

func (x *AssetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetInfoResponse.ProtoReflect.Descriptor instead.
func (*AssetInfoResponse) Descriptor() ([]byte, []int) {
	return file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescGZIP(), []int{2}
}

func (x *AssetInfoResponse) GetRequestStatus() RequestStatus {
	if x != nil {
		return x.RequestStatus
	}
	return RequestStatus_OK
}

func (x *AssetInfoResponse) GetAssets() []*AssetInfo {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *AssetInfoResponse) GetMissingAssetTags() []string {
	if x != nil {
		return x.MissingAssetTags
	}
	return nil
}

func (x *AssetInfoResponse) GetFailedAssetTags() []string {
	if x != nil {
		return x.FailedAssetTags
	}
	return nil
}

var File_infra_unifiedfleet_api_v1_models_assetinfo_proto protoreflect.FileDescriptor

var file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66,
	0x6c, 0x65, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf7, 0x02, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x6d, 0x61,
	0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x15, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65,
	0x12, 0x17, 0x0a, 0x04, 0x68, 0x77, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x04, 0x68, 0x77, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x10, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x22, 0xfe, 0x01, 0x0a,
	0x11, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x50, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x61, 0x67,
	0x73, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x2a, 0x5c, 0x0a,
	0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06,
	0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x10, 0x0a,
	0x0c, 0x4e, 0x4f, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x5f, 0x47, 0x50, 0x4e, 0x10, 0x04, 0x42, 0x28, 0x5a, 0x26, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b,
	0x75, 0x66, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescOnce sync.Once
	file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescData = file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDesc
)

func file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescGZIP() []byte {
	file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescOnce.Do(func() {
		file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescData)
	})
	return file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDescData
}

var file_infra_unifiedfleet_api_v1_models_assetinfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_infra_unifiedfleet_api_v1_models_assetinfo_proto_goTypes = []interface{}{
	(RequestStatus)(0),        // 0: unifiedfleet.api.v1.models.RequestStatus
	(*AssetInfo)(nil),         // 1: unifiedfleet.api.v1.models.AssetInfo
	(*AssetInfoRequest)(nil),  // 2: unifiedfleet.api.v1.models.AssetInfoRequest
	(*AssetInfoResponse)(nil), // 3: unifiedfleet.api.v1.models.AssetInfoResponse
}
var file_infra_unifiedfleet_api_v1_models_assetinfo_proto_depIdxs = []int32{
	0, // 0: unifiedfleet.api.v1.models.AssetInfoResponse.request_status:type_name -> unifiedfleet.api.v1.models.RequestStatus
	1, // 1: unifiedfleet.api.v1.models.AssetInfoResponse.assets:type_name -> unifiedfleet.api.v1.models.AssetInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infra_unifiedfleet_api_v1_models_assetinfo_proto_init() }
func file_infra_unifiedfleet_api_v1_models_assetinfo_proto_init() {
	if File_infra_unifiedfleet_api_v1_models_assetinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetInfo); i {
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
		file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetInfoRequest); i {
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
		file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetInfoResponse); i {
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
			RawDescriptor: file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_unifiedfleet_api_v1_models_assetinfo_proto_goTypes,
		DependencyIndexes: file_infra_unifiedfleet_api_v1_models_assetinfo_proto_depIdxs,
		EnumInfos:         file_infra_unifiedfleet_api_v1_models_assetinfo_proto_enumTypes,
		MessageInfos:      file_infra_unifiedfleet_api_v1_models_assetinfo_proto_msgTypes,
	}.Build()
	File_infra_unifiedfleet_api_v1_models_assetinfo_proto = out.File
	file_infra_unifiedfleet_api_v1_models_assetinfo_proto_rawDesc = nil
	file_infra_unifiedfleet_api_v1_models_assetinfo_proto_goTypes = nil
	file_infra_unifiedfleet_api_v1_models_assetinfo_proto_depIdxs = nil
}
