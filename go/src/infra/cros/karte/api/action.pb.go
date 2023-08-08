// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: infra/cros/karte/api/action.proto

package kartepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The status of an action is whether the action succeeded or failed.
type Action_Status int32

const (
	Action_STATUS_UNSPECIFIED Action_Status = 0
	Action_SUCCESS            Action_Status = 1
	Action_FAIL               Action_Status = 2
	Action_SKIP               Action_Status = 3
)

// Enum value maps for Action_Status.
var (
	Action_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "SUCCESS",
		2: "FAIL",
		3: "SKIP",
	}
	Action_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"SUCCESS":            1,
		"FAIL":               2,
		"SKIP":               3,
	}
)

func (x Action_Status) Enum() *Action_Status {
	p := new(Action_Status)
	*p = x
	return p
}

func (x Action_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_cros_karte_api_action_proto_enumTypes[0].Descriptor()
}

func (Action_Status) Type() protoreflect.EnumType {
	return &file_infra_cros_karte_api_action_proto_enumTypes[0]
}

func (x Action_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action_Status.Descriptor instead.
func (Action_Status) EnumDescriptor() ([]byte, []int) {
	return file_infra_cros_karte_api_action_proto_rawDescGZIP(), []int{0, 0}
}

// AllowFail records whether failure is allowed or not for the action in
// question. There are two features of actions commonly referred to as
// "criticality". The first is whether the action is a "critical action" or
// not. A critical action is one that is considered an entrypoint from the
// perspective of the Paris engine. The second is one where failures are NOT
// forgiven. This captures the second notion.
type Action_AllowFail int32

const (
	Action_ALLOW_FAIL_UNSPECIFIED Action_AllowFail = 0
	Action_ALLOW_FAIL             Action_AllowFail = 1
	Action_NO_ALLOW_FAIL          Action_AllowFail = 2
)

// Enum value maps for Action_AllowFail.
var (
	Action_AllowFail_name = map[int32]string{
		0: "ALLOW_FAIL_UNSPECIFIED",
		1: "ALLOW_FAIL",
		2: "NO_ALLOW_FAIL",
	}
	Action_AllowFail_value = map[string]int32{
		"ALLOW_FAIL_UNSPECIFIED": 0,
		"ALLOW_FAIL":             1,
		"NO_ALLOW_FAIL":          2,
	}
)

func (x Action_AllowFail) Enum() *Action_AllowFail {
	p := new(Action_AllowFail)
	*p = x
	return p
}

func (x Action_AllowFail) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action_AllowFail) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_cros_karte_api_action_proto_enumTypes[1].Descriptor()
}

func (Action_AllowFail) Type() protoreflect.EnumType {
	return &file_infra_cros_karte_api_action_proto_enumTypes[1]
}

func (x Action_AllowFail) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action_AllowFail.Descriptor instead.
func (Action_AllowFail) EnumDescriptor() ([]byte, []int) {
	return file_infra_cros_karte_api_action_proto_rawDescGZIP(), []int{0, 1}
}

// ActionType is the type of the action: verifier, condition, or recovery.
type Action_ActionType int32

const (
	// ACTION_TYPE_UNSPECIFIED is the default action type.
	Action_ACTION_TYPE_UNSPECIFIED Action_ActionType = 0
	// ACTION_TYPE_VERIFIER actions are critical actions and their dependencies.
	Action_ACTION_TYPE_VERIFIER Action_ActionType = 1
	// ACTION_TYPE_CONDITION actions are condition actions and their
	// dependencies.
	Action_ACTION_TYPE_CONDITION Action_ActionType = 2
	// ACTION_TYPE_RECOVERY actions are recovery actions and their dependencies.
	Action_ACTION_TYPE_RECOVERY Action_ActionType = 3
)

// Enum value maps for Action_ActionType.
var (
	Action_ActionType_name = map[int32]string{
		0: "ACTION_TYPE_UNSPECIFIED",
		1: "ACTION_TYPE_VERIFIER",
		2: "ACTION_TYPE_CONDITION",
		3: "ACTION_TYPE_RECOVERY",
	}
	Action_ActionType_value = map[string]int32{
		"ACTION_TYPE_UNSPECIFIED": 0,
		"ACTION_TYPE_VERIFIER":    1,
		"ACTION_TYPE_CONDITION":   2,
		"ACTION_TYPE_RECOVERY":    3,
	}
)

func (x Action_ActionType) Enum() *Action_ActionType {
	p := new(Action_ActionType)
	*p = x
	return p
}

func (x Action_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_cros_karte_api_action_proto_enumTypes[2].Descriptor()
}

func (Action_ActionType) Type() protoreflect.EnumType {
	return &file_infra_cros_karte_api_action_proto_enumTypes[2]
}

func (x Action_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action_ActionType.Descriptor instead.
func (Action_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_infra_cros_karte_api_action_proto_rawDescGZIP(), []int{0, 2}
}

// An action represents an event that was intentionally performed on a DUT.
// Examples include running a command on a DUT or resetting the servo
// attached to a DUT.
type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the action. Names are generated
	// automatically when a new action is created.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A kind is a coarse-grained type of an action, such as
	// ssh-attempt. New action_kinds will be created frequently so this field
	// is a string; see https://google.aip.dev/126 for details.
	Kind string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	// A swarming task ID is the ID of a single swarming task.
	// The swarming task of an action is the swarming task that invoked the
	// action.
	// For example, "4f6c0ba2ef3fc610" is a swarming task ID.
	SwarmingTaskId string `protobuf:"bytes,4,opt,name=swarming_task_id,json=swarmingTaskId,proto3" json:"swarming_task_id,omitempty"`
	// An asset tag is the tag of a given asset in UFS.
	// An asset tag may be a short number such as C444444 printed on a device,
	// or it may be a UUID in some circumstances.
	AssetTag string `protobuf:"bytes,5,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	// The start time is the time that an action started.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The stop time is the time that an action finished.
	StopTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=stop_time,json=stopTime,proto3" json:"stop_time,omitempty"`
	// The create time is the time that an action was created by Karte.
	// This is the time that the event was first received, since events are
	// immutable outside of rare cases.
	// This field is managed by Karte itself.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Status     Action_Status          `protobuf:"varint,9,opt,name=status,proto3,enum=chromeos.karte.Action_Status" json:"status,omitempty"`
	// The fail reason of an event is a diagnostic message that is emitted when
	// the action in question has failed.
	FailReason string `protobuf:"bytes,10,opt,name=fail_reason,json=failReason,proto3" json:"fail_reason,omitempty"`
	// The seal time is when the particular Karte record is sealed and no further
	// changes can be made.
	SealTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=seal_time,json=sealTime,proto3" json:"seal_time,omitempty"`
	// The client is the name of the entity creating the Action entry, e.g.
	// "paris".
	ClientName string `protobuf:"bytes,12,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	// The client version is the version of the entity creating the Action entry,
	// e.g. "0.0.1".
	ClientVersion string `protobuf:"bytes,13,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	// The buildbucket ID is the ID of the buildbucket build associated with the
	// event in question.
	BuildbucketId string `protobuf:"bytes,14,opt,name=buildbucket_id,json=buildbucketId,proto3" json:"buildbucket_id,omitempty"`
	// The hostname is the hostname of the DUT in question.
	Hostname string `protobuf:"bytes,15,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Modification_count returns the number of times that the Karte record has
	// been modified. This field is managed by Karte.
	ModificationCount int32 `protobuf:"varint,16,opt,name=modification_count,json=modificationCount,proto3" json:"modification_count,omitempty"`
	// model is the model of the DUT this event applies to.
	Model string `protobuf:"bytes,17,opt,name=model,proto3" json:"model,omitempty"`
	// board is the board of the DUT this event applies to.
	Board string `protobuf:"bytes,18,opt,name=board,proto3" json:"board,omitempty"`
	// recovered by is the name of the task that recovered the current action, if
	// one exists.
	RecoveredBy string `protobuf:"bytes,19,opt,name=recovered_by,json=recoveredBy,proto3" json:"recovered_by,omitempty"`
	// restarts is the number of times that the current plan was restarted.
	Restarts int32 `protobuf:"varint,20,opt,name=restarts,proto3" json:"restarts,omitempty"`
	// plan_name is the name of the plan that we're currently executing.
	PlanName   string            `protobuf:"bytes,21,opt,name=plan_name,json=planName,proto3" json:"plan_name,omitempty"`
	AllowFail  Action_AllowFail  `protobuf:"varint,22,opt,name=allow_fail,json=allowFail,proto3,enum=chromeos.karte.Action_AllowFail" json:"allow_fail,omitempty"`
	ActionType Action_ActionType `protobuf:"varint,23,opt,name=action_type,json=actionType,proto3,enum=chromeos.karte.Action_ActionType" json:"action_type,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_cros_karte_api_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_infra_cros_karte_api_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_infra_cros_karte_api_action_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Action) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Action) GetSwarmingTaskId() string {
	if x != nil {
		return x.SwarmingTaskId
	}
	return ""
}

func (x *Action) GetAssetTag() string {
	if x != nil {
		return x.AssetTag
	}
	return ""
}

func (x *Action) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Action) GetStopTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StopTime
	}
	return nil
}

func (x *Action) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Action) GetStatus() Action_Status {
	if x != nil {
		return x.Status
	}
	return Action_STATUS_UNSPECIFIED
}

func (x *Action) GetFailReason() string {
	if x != nil {
		return x.FailReason
	}
	return ""
}

func (x *Action) GetSealTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SealTime
	}
	return nil
}

func (x *Action) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *Action) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

func (x *Action) GetBuildbucketId() string {
	if x != nil {
		return x.BuildbucketId
	}
	return ""
}

func (x *Action) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Action) GetModificationCount() int32 {
	if x != nil {
		return x.ModificationCount
	}
	return 0
}

func (x *Action) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Action) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

func (x *Action) GetRecoveredBy() string {
	if x != nil {
		return x.RecoveredBy
	}
	return ""
}

func (x *Action) GetRestarts() int32 {
	if x != nil {
		return x.Restarts
	}
	return 0
}

func (x *Action) GetPlanName() string {
	if x != nil {
		return x.PlanName
	}
	return ""
}

func (x *Action) GetAllowFail() Action_AllowFail {
	if x != nil {
		return x.AllowFail
	}
	return Action_ALLOW_FAIL_UNSPECIFIED
}

func (x *Action) GetActionType() Action_ActionType {
	if x != nil {
		return x.ActionType
	}
	return Action_ACTION_TYPE_UNSPECIFIED
}

var File_infra_cros_karte_api_action_proto protoreflect.FileDescriptor

var file_infra_cros_karte_api_action_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6b, 0x61, 0x72,
	0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6b, 0x61,
	0x72, 0x74, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd7, 0x09, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6b, 0x61, 0x72, 0x74,
	0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x5f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x6c,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x12,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6b, 0x61, 0x72, 0x74, 0x65, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x09,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x42, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x6b, 0x61, 0x72, 0x74, 0x65, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x41, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x03,
	0x22, 0x4a, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x16, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f,
	0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x22, 0x78, 0x0a, 0x0a,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f,
	0x56, 0x45, 0x52, 0x59, 0x10, 0x03, 0x3a, 0x2f, 0xea, 0x41, 0x2c, 0x0a, 0x18, 0x6b, 0x61, 0x72,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6b, 0x61, 0x72, 0x74, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x3b, 0x6b, 0x61, 0x72, 0x74, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_infra_cros_karte_api_action_proto_rawDescOnce sync.Once
	file_infra_cros_karte_api_action_proto_rawDescData = file_infra_cros_karte_api_action_proto_rawDesc
)

func file_infra_cros_karte_api_action_proto_rawDescGZIP() []byte {
	file_infra_cros_karte_api_action_proto_rawDescOnce.Do(func() {
		file_infra_cros_karte_api_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_cros_karte_api_action_proto_rawDescData)
	})
	return file_infra_cros_karte_api_action_proto_rawDescData
}

var file_infra_cros_karte_api_action_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_infra_cros_karte_api_action_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_cros_karte_api_action_proto_goTypes = []interface{}{
	(Action_Status)(0),            // 0: chromeos.karte.Action.Status
	(Action_AllowFail)(0),         // 1: chromeos.karte.Action.AllowFail
	(Action_ActionType)(0),        // 2: chromeos.karte.Action.ActionType
	(*Action)(nil),                // 3: chromeos.karte.Action
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_infra_cros_karte_api_action_proto_depIdxs = []int32{
	4, // 0: chromeos.karte.Action.start_time:type_name -> google.protobuf.Timestamp
	4, // 1: chromeos.karte.Action.stop_time:type_name -> google.protobuf.Timestamp
	4, // 2: chromeos.karte.Action.create_time:type_name -> google.protobuf.Timestamp
	0, // 3: chromeos.karte.Action.status:type_name -> chromeos.karte.Action.Status
	4, // 4: chromeos.karte.Action.seal_time:type_name -> google.protobuf.Timestamp
	1, // 5: chromeos.karte.Action.allow_fail:type_name -> chromeos.karte.Action.AllowFail
	2, // 6: chromeos.karte.Action.action_type:type_name -> chromeos.karte.Action.ActionType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_infra_cros_karte_api_action_proto_init() }
func file_infra_cros_karte_api_action_proto_init() {
	if File_infra_cros_karte_api_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_cros_karte_api_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
			RawDescriptor: file_infra_cros_karte_api_action_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_cros_karte_api_action_proto_goTypes,
		DependencyIndexes: file_infra_cros_karte_api_action_proto_depIdxs,
		EnumInfos:         file_infra_cros_karte_api_action_proto_enumTypes,
		MessageInfos:      file_infra_cros_karte_api_action_proto_msgTypes,
	}.Build()
	File_infra_cros_karte_api_action_proto = out.File
	file_infra_cros_karte_api_action_proto_rawDesc = nil
	file_infra_cros_karte_api_action_proto_goTypes = nil
	file_infra_cros_karte_api_action_proto_depIdxs = nil
}
