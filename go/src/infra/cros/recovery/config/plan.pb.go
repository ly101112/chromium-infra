// Copyright 2021 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.7
// source: infra/cros/recovery/config/planpb/plan.proto

package config

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RunControl describe when and how often an action runs per plan execution.
type RunControl int32

const (
	// The action is run once per plan and rerun again after each successful
	// recovery action.
	RunControl_RERUN_AFTER_RECOVERY RunControl = 0
	// The action runs every time.
	RunControl_ALWAYS_RUN RunControl = 1
	// The action is run only once per plan.
	RunControl_RUN_ONCE RunControl = 2
)

// Enum value maps for RunControl.
var (
	RunControl_name = map[int32]string{
		0: "RERUN_AFTER_RECOVERY",
		1: "ALWAYS_RUN",
		2: "RUN_ONCE",
	}
	RunControl_value = map[string]int32{
		"RERUN_AFTER_RECOVERY": 0,
		"ALWAYS_RUN":           1,
		"RUN_ONCE":             2,
	}
)

func (x RunControl) Enum() *RunControl {
	p := new(RunControl)
	*p = x
	return p
}

func (x RunControl) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunControl) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_cros_recovery_config_planpb_plan_proto_enumTypes[0].Descriptor()
}

func (RunControl) Type() protoreflect.EnumType {
	return &file_infra_cros_recovery_config_planpb_plan_proto_enumTypes[0]
}

func (x RunControl) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunControl.Descriptor instead.
func (RunControl) EnumDescriptor() ([]byte, []int) {
	return file_infra_cros_recovery_config_planpb_plan_proto_rawDescGZIP(), []int{0}
}

// Configuration provides the plans to be used by the recovery engine.
type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map of plans provided by configuration.
	Plans map[string]*Plan `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// List of plan names in order to execute.
	PlanNames []string `protobuf:"bytes,2,rep,name=plan_names,json=planNames,proto3" json:"plan_names,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_cros_recovery_config_planpb_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_infra_cros_recovery_config_planpb_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_infra_cros_recovery_config_planpb_plan_proto_rawDescGZIP(), []int{0}
}

func (x *Configuration) GetPlans() map[string]*Plan {
	if x != nil {
		return x.Plans
	}
	return nil
}

func (x *Configuration) GetPlanNames() []string {
	if x != nil {
		return x.PlanNames
	}
	return nil
}

// Plan holds information about actions for recovery engine to execute.
type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Critical actions are actions which have to pass for plan to succeed.
	// Order matters.
	CriticalActions []string `protobuf:"bytes,1,rep,name=critical_actions,json=criticalActions,proto3" json:"critical_actions,omitempty"`
	// Map of all actions used by the plan.
	Actions map[string]*Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// When running multiple plans, whether to continue running the next plan
	// if this plan fails.
	AllowFail bool `protobuf:"varint,3,opt,name=allow_fail,json=allowFail,proto3" json:"allow_fail,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_cros_recovery_config_planpb_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_infra_cros_recovery_config_planpb_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_infra_cros_recovery_config_planpb_plan_proto_rawDescGZIP(), []int{1}
}

func (x *Plan) GetCriticalActions() []string {
	if x != nil {
		return x.CriticalActions
	}
	return nil
}

func (x *Plan) GetActions() map[string]*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Plan) GetAllowFail() bool {
	if x != nil {
		return x.AllowFail
	}
	return false
}

// Action describes how to run the action, including its dependencies,
// conditions, and other attributes.
type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of actions to determine if this action is applicable for the resource.
	// If any condition fails then this action will be skipped.
	Conditions []string `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// List of actions that must pass before executing this action's exec
	// function.
	Dependencies []string `protobuf:"bytes,2,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	// Name of the exec function to use.
	// The name of the action will be used if not provided.
	ExecName string `protobuf:"bytes,3,opt,name=exec_name,json=execName,proto3" json:"exec_name,omitempty"`
	// Allowed time to execute exec function.
	// If not specified, defaults to 60 seconds.
	// The default may change in the future.
	ExecTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=exec_timeout,json=execTimeout,proto3" json:"exec_timeout,omitempty"`
	// Extra arguments provided to the exec function.
	// What arguments are allowed depends on the exec function.
	ExecExtraArgs []string `protobuf:"bytes,5,rep,name=exec_extra_args,json=execExtraArgs,proto3" json:"exec_extra_args,omitempty"`
	// List of actions used to recover this action if exec function fails.
	RecoveryActions []string `protobuf:"bytes,6,rep,name=recovery_actions,json=recoveryActions,proto3" json:"recovery_actions,omitempty"`
	// If set to true, then the action is treated as if it passed even if it
	// and all its recovery actions failed.
	AllowFailAfterRecovery bool `protobuf:"varint,7,opt,name=allow_fail_after_recovery,json=allowFailAfterRecovery,proto3" json:"allow_fail_after_recovery,omitempty"`
	// Controls how and when the action can be rerun throughout the plan.
	RunControl RunControl `protobuf:"varint,8,opt,name=run_control,json=runControl,proto3,enum=chromeos.recovery.RunControl" json:"run_control,omitempty"`
	// Documentation to describe detail of the action.
	Docs []string `protobuf:"bytes,9,rep,name=docs,proto3" json:"docs,omitempty"`
	// The metrics config specifies how we handle metrics created by this action.
	MetricsConfig *MetricsConfig `protobuf:"bytes,10,opt,name=metrics_config,json=metricsConfig,proto3" json:"metrics_config,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_cros_recovery_config_planpb_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_infra_cros_recovery_config_planpb_plan_proto_msgTypes[2]
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
	return file_infra_cros_recovery_config_planpb_plan_proto_rawDescGZIP(), []int{2}
}

func (x *Action) GetConditions() []string {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *Action) GetDependencies() []string {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *Action) GetExecName() string {
	if x != nil {
		return x.ExecName
	}
	return ""
}

func (x *Action) GetExecTimeout() *durationpb.Duration {
	if x != nil {
		return x.ExecTimeout
	}
	return nil
}

func (x *Action) GetExecExtraArgs() []string {
	if x != nil {
		return x.ExecExtraArgs
	}
	return nil
}

func (x *Action) GetRecoveryActions() []string {
	if x != nil {
		return x.RecoveryActions
	}
	return nil
}

func (x *Action) GetAllowFailAfterRecovery() bool {
	if x != nil {
		return x.AllowFailAfterRecovery
	}
	return false
}

func (x *Action) GetRunControl() RunControl {
	if x != nil {
		return x.RunControl
	}
	return RunControl_RERUN_AFTER_RECOVERY
}

func (x *Action) GetDocs() []string {
	if x != nil {
		return x.Docs
	}
	return nil
}

func (x *Action) GetMetricsConfig() *MetricsConfig {
	if x != nil {
		return x.MetricsConfig
	}
	return nil
}

var File_infra_cros_recovery_config_planpb_plan_proto protoreflect.FileDescriptor

var file_infra_cros_recovery_config_planpb_plan_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x35, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x05, 0x70, 0x6c,
	0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x51, 0x0a, 0x0a,
	0x50, 0x6c, 0x61, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xe7, 0x01, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x72, 0x69, 0x74,
	0x69, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x66, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x46, 0x61,
	0x69, 0x6c, 0x1a, 0x55, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x72,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd2, 0x03, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x65,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78,
	0x65, 0x63, 0x45, 0x78, 0x74, 0x72, 0x61, 0x41, 0x72, 0x67, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x66, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x46, 0x61, 0x69, 0x6c, 0x41, 0x66, 0x74, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x12, 0x3e, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f,
	0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x63, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x6f, 0x63, 0x73, 0x12, 0x47, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0x44,
	0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x14,
	0x52, 0x45, 0x52, 0x55, 0x4e, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x43, 0x4f,
	0x56, 0x45, 0x52, 0x59, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4c, 0x57, 0x41, 0x59, 0x53,
	0x5f, 0x52, 0x55, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x55, 0x4e, 0x5f, 0x4f, 0x4e,
	0x43, 0x45, 0x10, 0x02, 0x42, 0x23, 0x5a, 0x21, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x63, 0x72,
	0x6f, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_infra_cros_recovery_config_planpb_plan_proto_rawDescOnce sync.Once
	file_infra_cros_recovery_config_planpb_plan_proto_rawDescData = file_infra_cros_recovery_config_planpb_plan_proto_rawDesc
)

func file_infra_cros_recovery_config_planpb_plan_proto_rawDescGZIP() []byte {
	file_infra_cros_recovery_config_planpb_plan_proto_rawDescOnce.Do(func() {
		file_infra_cros_recovery_config_planpb_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_cros_recovery_config_planpb_plan_proto_rawDescData)
	})
	return file_infra_cros_recovery_config_planpb_plan_proto_rawDescData
}

var file_infra_cros_recovery_config_planpb_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_cros_recovery_config_planpb_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_infra_cros_recovery_config_planpb_plan_proto_goTypes = []interface{}{
	(RunControl)(0),             // 0: chromeos.recovery.RunControl
	(*Configuration)(nil),       // 1: chromeos.recovery.Configuration
	(*Plan)(nil),                // 2: chromeos.recovery.Plan
	(*Action)(nil),              // 3: chromeos.recovery.Action
	nil,                         // 4: chromeos.recovery.Configuration.PlansEntry
	nil,                         // 5: chromeos.recovery.Plan.ActionsEntry
	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
	(*MetricsConfig)(nil),       // 7: chromeos.recovery.MetricsConfig
}
var file_infra_cros_recovery_config_planpb_plan_proto_depIdxs = []int32{
	4, // 0: chromeos.recovery.Configuration.plans:type_name -> chromeos.recovery.Configuration.PlansEntry
	5, // 1: chromeos.recovery.Plan.actions:type_name -> chromeos.recovery.Plan.ActionsEntry
	6, // 2: chromeos.recovery.Action.exec_timeout:type_name -> google.protobuf.Duration
	0, // 3: chromeos.recovery.Action.run_control:type_name -> chromeos.recovery.RunControl
	7, // 4: chromeos.recovery.Action.metrics_config:type_name -> chromeos.recovery.MetricsConfig
	2, // 5: chromeos.recovery.Configuration.PlansEntry.value:type_name -> chromeos.recovery.Plan
	3, // 6: chromeos.recovery.Plan.ActionsEntry.value:type_name -> chromeos.recovery.Action
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_infra_cros_recovery_config_planpb_plan_proto_init() }
func file_infra_cros_recovery_config_planpb_plan_proto_init() {
	if File_infra_cros_recovery_config_planpb_plan_proto != nil {
		return
	}
	file_infra_cros_recovery_config_planpb_metric_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infra_cros_recovery_config_planpb_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
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
		file_infra_cros_recovery_config_planpb_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
		file_infra_cros_recovery_config_planpb_plan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_infra_cros_recovery_config_planpb_plan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_cros_recovery_config_planpb_plan_proto_goTypes,
		DependencyIndexes: file_infra_cros_recovery_config_planpb_plan_proto_depIdxs,
		EnumInfos:         file_infra_cros_recovery_config_planpb_plan_proto_enumTypes,
		MessageInfos:      file_infra_cros_recovery_config_planpb_plan_proto_msgTypes,
	}.Build()
	File_infra_cros_recovery_config_planpb_plan_proto = out.File
	file_infra_cros_recovery_config_planpb_plan_proto_rawDesc = nil
	file_infra_cros_recovery_config_planpb_plan_proto_goTypes = nil
	file_infra_cros_recovery_config_planpb_plan_proto_depIdxs = nil
}
