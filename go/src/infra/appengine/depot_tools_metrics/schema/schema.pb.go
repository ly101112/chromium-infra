// Copyright 2019 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: infra/appengine/depot_tools_metrics/schema/schema.proto

package schema

import (
	proto "go.chromium.org/luci/buildbucket/proto"
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

type SyncStatus int32

const (
	SyncStatus_UNSPECIFIED         SyncStatus = 0
	SyncStatus_SYNC_STATUS_FAILURE SyncStatus = 1
	SyncStatus_SYNC_STATUS_SUCCESS SyncStatus = 2
	SyncStatus_SYNC_STATUS_NO_OP   SyncStatus = 3
)

// Enum value maps for SyncStatus.
var (
	SyncStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "SYNC_STATUS_FAILURE",
		2: "SYNC_STATUS_SUCCESS",
		3: "SYNC_STATUS_NO_OP",
	}
	SyncStatus_value = map[string]int32{
		"UNSPECIFIED":         0,
		"SYNC_STATUS_FAILURE": 1,
		"SYNC_STATUS_SUCCESS": 2,
		"SYNC_STATUS_NO_OP":   3,
	}
)

func (x SyncStatus) Enum() *SyncStatus {
	p := new(SyncStatus)
	*p = x
	return p
}

func (x SyncStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_enumTypes[0].Descriptor()
}

func (SyncStatus) Type() protoreflect.EnumType {
	return &file_infra_appengine_depot_tools_metrics_schema_schema_proto_enumTypes[0]
}

func (x SyncStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SyncStatus.Descriptor instead.
func (SyncStatus) EnumDescriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{0}
}

// HttpRequest stores information on the HTTP requests made by the command.
type HttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The host the request was made to. Must be one of the |knownHTTPHosts| in
	// metrics/constants.go.
	// e.g. chromium-review.googlesource.com
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The HTTP method used to make the request (e.g. GET, POST).
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// The path and URL arguments of the request.
	// The path must be one of the |knownHTTPPaths| and the arguments must be
	// |knownHTTPArguments| as defined in metrics/constants.go.
	//
	// The URL is not recorded since it might contain PII. Similarly, in most
	// cases, only the name of the arguments (and not their values) are recorded.
	// When the possible values for an argument is a fixed set, as is the case for
	// "o-parameters" in Gerrit, they'll be recorded as arguments.
	// Each argument is recorded separately, so as to make it easier to query.
	//
	// e.g. If the request was to
	// '/changes/?q=owner:foo@example.com+is:open&n=3&o=LABELS&o=ALL_REVISIONS'
	// The path will be '/changes' and the arguments will be 'q', 'n', 'o',
	// 'LABELS' and 'ALL_REVISIONS'.
	Path      string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Arguments []string `protobuf:"bytes,4,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// The HTTP response status.
	Status int64 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	// The latency of the HTTP request in seconds.
	// TODO(ehmaldonado): Consider converting to google.protobuf.Duration.
	ResponseTime float64 `protobuf:"fixed64,6,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
}

func (x *HttpRequest) Reset() {
	*x = HttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRequest) ProtoMessage() {}

func (x *HttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRequest.ProtoReflect.Descriptor instead.
func (*HttpRequest) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{0}
}

func (x *HttpRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *HttpRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HttpRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *HttpRequest) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *HttpRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HttpRequest) GetResponseTime() float64 {
	if x != nil {
		return x.ResponseTime
	}
	return 0
}

// SubCommand stores information on the sub-commands executed by the command.
type SubCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The sub-command that was executed. Must be one of the |knownSubCommands| in
	// metrics/constans.go.
	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	// The arguments passed to the sub-command. All arguments must be
	// |knownSubCommandArguments| as defined in metrics/constants.go.
	Arguments []string `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// The runtime of the sub-command runtime in seconds.
	// TODO(ehmaldonado): Consider converting to google.protobuf.Duration.
	ExecutionTime float64 `protobuf:"fixed64,3,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	// The exit code of the sub-command.
	ExitCode int64 `protobuf:"varint,4,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *SubCommand) Reset() {
	*x = SubCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubCommand) ProtoMessage() {}

func (x *SubCommand) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubCommand.ProtoReflect.Descriptor instead.
func (*SubCommand) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{1}
}

func (x *SubCommand) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *SubCommand) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *SubCommand) GetExecutionTime() float64 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *SubCommand) GetExitCode() int64 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

// GitDependency stores information about the git dependencies for the current
// project.
type GitDependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path where this dependency was synced, relative to gclient root.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The url for the dependency.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// The revision the dependency was synced to.
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// Time to sync dependency to revision in seconds.
	ExecutionTime float64 `protobuf:"fixed64,4,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	// Whether we failed to sync the dependency to the desired revision.
	SyncStatus SyncStatus `protobuf:"varint,5,opt,name=sync_status,json=syncStatus,proto3,enum=schema.SyncStatus" json:"sync_status,omitempty"`
}

func (x *GitDependency) Reset() {
	*x = GitDependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitDependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitDependency) ProtoMessage() {}

func (x *GitDependency) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitDependency.ProtoReflect.Descriptor instead.
func (*GitDependency) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{2}
}

func (x *GitDependency) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GitDependency) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GitDependency) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *GitDependency) GetExecutionTime() float64 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *GitDependency) GetSyncStatus() SyncStatus {
	if x != nil {
		return x.SyncStatus
	}
	return SyncStatus_UNSPECIFIED
}

// Hook stores information about hooks run as part of gclient.
type Hook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The command that was executed.
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	// The name of the hook.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The directory where to execute the hook.
	Cwd string `protobuf:"bytes,3,opt,name=cwd,proto3" json:"cwd,omitempty"`
	// Conditional string. The hook is run if it evaluates to true.
	Condition string `protobuf:"bytes,4,opt,name=condition,proto3" json:"condition,omitempty"`
	// Time it took to execute the hook.
	ExecutionTime float64 `protobuf:"fixed64,5,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	// The exit code of the hook.
	ExitCode int64 `protobuf:"varint,6,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *Hook) Reset() {
	*x = Hook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hook) ProtoMessage() {}

func (x *Hook) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hook.ProtoReflect.Descriptor instead.
func (*Hook) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{3}
}

func (x *Hook) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Hook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hook) GetCwd() string {
	if x != nil {
		return x.Cwd
	}
	return ""
}

func (x *Hook) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *Hook) GetExecutionTime() float64 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *Hook) GetExitCode() int64 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

// EnvVar stores an environment variable.
type EnvVar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the environment variable.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The value of the environment variable.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EnvVar) Reset() {
	*x = EnvVar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvVar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvVar) ProtoMessage() {}

func (x *EnvVar) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvVar.ProtoReflect.Descriptor instead.
func (*EnvVar) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{4}
}

func (x *EnvVar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvVar) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// BotMetrics stores information about the bot environment from which the
// command was executed.
type BotMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The build from which this command was executed.
	BuildId int64 `protobuf:"varint,1,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	// The builder corresponding to the build.
	Builder *proto.BuilderID `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
}

func (x *BotMetrics) Reset() {
	*x = BotMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotMetrics) ProtoMessage() {}

func (x *BotMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotMetrics.ProtoReflect.Descriptor instead.
func (*BotMetrics) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{5}
}

func (x *BotMetrics) GetBuildId() int64 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

func (x *BotMetrics) GetBuilder() *proto.BuilderID {
	if x != nil {
		return x.Builder
	}
	return nil
}

// Metrics stores information for a depot_tools command's execution.
type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version of the format used to report the metrics.
	MetricsVersion int64 `protobuf:"varint,1,opt,name=metrics_version,json=metricsVersion,proto3" json:"metrics_version,omitempty"`
	// A UNIX timestamp for the time when the command was executed.
	// TODO(ehmaldonado): Consider converting to google.protobuf.Timestamp.
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The command that was executed. Must be one of the |knownCommands| defined
	// in metrics/constants.go.
	Command string `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	// The arguments passed to the command. All arguments must be |knownArguments|
	// as defined in metrics/constants.go.
	Arguments []string `protobuf:"bytes,4,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// The runtime of the command in seconds.
	// TODO(ehmaldonado): Consider converting to google.protobuf.Duration.
	ExecutionTime float64 `protobuf:"fixed64,5,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	// The exit code of the command.
	ExitCode int64 `protobuf:"varint,6,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	// Information on the sub-commands executed by this command.
	SubCommands []*SubCommand `protobuf:"bytes,7,rep,name=sub_commands,json=subCommands,proto3" json:"sub_commands,omitempty"`
	// Information on the HTTP requests made by this command.
	HttpRequests []*HttpRequest `protobuf:"bytes,8,rep,name=http_requests,json=httpRequests,proto3" json:"http_requests,omitempty"`
	// The URLs of the current project(s).
	// e.g. The project to which git-cl uploads a change; the projects gclient is
	// configured to manage; etc.
	// Must be one of the |knownProjectURLs| as defined in metrics/constants.go.
	ProjectUrls []string `protobuf:"bytes,9,rep,name=project_urls,json=projectUrls,proto3" json:"project_urls,omitempty"`
	// A UNIX timestamp for the time depot_tools was last modified.
	// TODO(ehmaldonado): Consider converting to google.protobuf.Timestamp.
	DepotToolsAge float64 `protobuf:"fixed64,10,opt,name=depot_tools_age,json=depotToolsAge,proto3" json:"depot_tools_age,omitempty"`
	// The arch the command was executed on. Must be one of the |knownHostArchs|
	// as defined in metrics/constants.go.
	// e.g. x86, arm
	HostArch string `protobuf:"bytes,11,opt,name=host_arch,json=hostArch,proto3" json:"host_arch,omitempty"`
	// The OS the command was executed on. Must be one of the |knownOSs| as
	// defined in metrics/constants.go.
	HostOs string `protobuf:"bytes,12,opt,name=host_os,json=hostOs,proto3" json:"host_os,omitempty"`
	// The python version the command was executed with. Must match the
	// |pythonVersionRegex| defined in metrics/constants.go.
	PythonVersion string `protobuf:"bytes,13,opt,name=python_version,json=pythonVersion,proto3" json:"python_version,omitempty"`
	// The git version the command used. Must match the |gitVersionRegex| defined
	// in metrics/constants.go.
	GitVersion string `protobuf:"bytes,14,opt,name=git_version,json=gitVersion,proto3" json:"git_version,omitempty"`
	// Metrics specific to commands run on bots.
	BotMetrics *BotMetrics `protobuf:"bytes,15,opt,name=bot_metrics,json=botMetrics,proto3" json:"bot_metrics,omitempty"`
	// The dependencies listed in DEPS that were synced by this command.
	GitDeps []*GitDependency `protobuf:"bytes,16,rep,name=git_deps,json=gitDeps,proto3" json:"git_deps,omitempty"`
	// The hooks listed in DEPS that were executed by this command.
	Hooks []*Hook `protobuf:"bytes,17,rep,name=hooks,proto3" json:"hooks,omitempty"`
	// Deprecated.
	//
	// Deprecated: Do not use.
	EnvVariables []string `protobuf:"bytes,18,rep,name=env_variables,json=envVariables,proto3" json:"env_variables,omitempty"`
	// The environment variables that depot_tools cares about.
	EnvVars []*EnvVar `protobuf:"bytes,19,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP(), []int{6}
}

func (x *Metrics) GetMetricsVersion() int64 {
	if x != nil {
		return x.MetricsVersion
	}
	return 0
}

func (x *Metrics) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Metrics) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Metrics) GetArguments() []string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *Metrics) GetExecutionTime() float64 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *Metrics) GetExitCode() int64 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *Metrics) GetSubCommands() []*SubCommand {
	if x != nil {
		return x.SubCommands
	}
	return nil
}

func (x *Metrics) GetHttpRequests() []*HttpRequest {
	if x != nil {
		return x.HttpRequests
	}
	return nil
}

func (x *Metrics) GetProjectUrls() []string {
	if x != nil {
		return x.ProjectUrls
	}
	return nil
}

func (x *Metrics) GetDepotToolsAge() float64 {
	if x != nil {
		return x.DepotToolsAge
	}
	return 0
}

func (x *Metrics) GetHostArch() string {
	if x != nil {
		return x.HostArch
	}
	return ""
}

func (x *Metrics) GetHostOs() string {
	if x != nil {
		return x.HostOs
	}
	return ""
}

func (x *Metrics) GetPythonVersion() string {
	if x != nil {
		return x.PythonVersion
	}
	return ""
}

func (x *Metrics) GetGitVersion() string {
	if x != nil {
		return x.GitVersion
	}
	return ""
}

func (x *Metrics) GetBotMetrics() *BotMetrics {
	if x != nil {
		return x.BotMetrics
	}
	return nil
}

func (x *Metrics) GetGitDeps() []*GitDependency {
	if x != nil {
		return x.GitDeps
	}
	return nil
}

func (x *Metrics) GetHooks() []*Hook {
	if x != nil {
		return x.Hooks
	}
	return nil
}

// Deprecated: Do not use.
func (x *Metrics) GetEnvVariables() []string {
	if x != nil {
		return x.EnvVariables
	}
	return nil
}

func (x *Metrics) GetEnvVars() []*EnvVar {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

var File_infra_appengine_depot_tools_metrics_schema_schema_proto protoreflect.FileDescriptor

var file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDesc = []byte{
	0x0a, 0x37, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x1a, 0x3b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8,
	0x01, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x53, 0x75,
	0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x0d, 0x47, 0x69, 0x74, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x33, 0x0a, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x04, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x77, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x77, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x32, 0x0a,
	0x06, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x5c, 0x0a, 0x0a, 0x42, 0x6f, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x49, 0x44, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x22,
	0xe5, 0x05, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x35,
	0x0a, 0x0c, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x75,
	0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x72,
	0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x65, 0x70,
	0x6f, 0x74, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x41, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x41, 0x72, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f,
	0x6f, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x4f, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x69, 0x74, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x69,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0b, 0x62, 0x6f, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x42, 0x6f, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x0a, 0x62, 0x6f, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x30, 0x0a,
	0x08, 0x67, 0x69, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x47, 0x69, 0x74, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x07, 0x67, 0x69, 0x74, 0x44, 0x65, 0x70, 0x73, 0x12,
	0x22, 0x0a, 0x05, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x68, 0x6f,
	0x6f, 0x6b, 0x73, 0x12, 0x27, 0x0a, 0x0d, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0c,
	0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x08,
	0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x52, 0x07,
	0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x2a, 0x66, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x59, 0x4e, 0x43,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x5f, 0x4f, 0x50, 0x10, 0x03, 0x42,
	0x33, 0x5a, 0x31, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3b, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescOnce sync.Once
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescData = file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDesc
)

func file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescGZIP() []byte {
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescOnce.Do(func() {
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescData)
	})
	return file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDescData
}

var file_infra_appengine_depot_tools_metrics_schema_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_infra_appengine_depot_tools_metrics_schema_schema_proto_goTypes = []interface{}{
	(SyncStatus)(0),         // 0: schema.SyncStatus
	(*HttpRequest)(nil),     // 1: schema.HttpRequest
	(*SubCommand)(nil),      // 2: schema.SubCommand
	(*GitDependency)(nil),   // 3: schema.GitDependency
	(*Hook)(nil),            // 4: schema.Hook
	(*EnvVar)(nil),          // 5: schema.EnvVar
	(*BotMetrics)(nil),      // 6: schema.BotMetrics
	(*Metrics)(nil),         // 7: schema.Metrics
	(*proto.BuilderID)(nil), // 8: buildbucket.v2.BuilderID
}
var file_infra_appengine_depot_tools_metrics_schema_schema_proto_depIdxs = []int32{
	0, // 0: schema.GitDependency.sync_status:type_name -> schema.SyncStatus
	8, // 1: schema.BotMetrics.builder:type_name -> buildbucket.v2.BuilderID
	2, // 2: schema.Metrics.sub_commands:type_name -> schema.SubCommand
	1, // 3: schema.Metrics.http_requests:type_name -> schema.HttpRequest
	6, // 4: schema.Metrics.bot_metrics:type_name -> schema.BotMetrics
	3, // 5: schema.Metrics.git_deps:type_name -> schema.GitDependency
	4, // 6: schema.Metrics.hooks:type_name -> schema.Hook
	5, // 7: schema.Metrics.env_vars:type_name -> schema.EnvVar
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_infra_appengine_depot_tools_metrics_schema_schema_proto_init() }
func file_infra_appengine_depot_tools_metrics_schema_schema_proto_init() {
	if File_infra_appengine_depot_tools_metrics_schema_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpRequest); i {
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
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubCommand); i {
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
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitDependency); i {
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
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hook); i {
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
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvVar); i {
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
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotMetrics); i {
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
		file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
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
			RawDescriptor: file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_depot_tools_metrics_schema_schema_proto_goTypes,
		DependencyIndexes: file_infra_appengine_depot_tools_metrics_schema_schema_proto_depIdxs,
		EnumInfos:         file_infra_appengine_depot_tools_metrics_schema_schema_proto_enumTypes,
		MessageInfos:      file_infra_appengine_depot_tools_metrics_schema_schema_proto_msgTypes,
	}.Build()
	File_infra_appengine_depot_tools_metrics_schema_schema_proto = out.File
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_rawDesc = nil
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_goTypes = nil
	file_infra_appengine_depot_tools_metrics_schema_schema_proto_depIdxs = nil
}
