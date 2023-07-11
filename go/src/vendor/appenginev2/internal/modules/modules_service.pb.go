// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: modules_service.proto

package modules

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

type ModulesServiceError_ErrorCode int32

const (
	ModulesServiceError_OK                ModulesServiceError_ErrorCode = 0
	ModulesServiceError_INVALID_MODULE    ModulesServiceError_ErrorCode = 1
	ModulesServiceError_INVALID_VERSION   ModulesServiceError_ErrorCode = 2
	ModulesServiceError_INVALID_INSTANCES ModulesServiceError_ErrorCode = 3
	ModulesServiceError_TRANSIENT_ERROR   ModulesServiceError_ErrorCode = 4
	ModulesServiceError_UNEXPECTED_STATE  ModulesServiceError_ErrorCode = 5
)

// Enum value maps for ModulesServiceError_ErrorCode.
var (
	ModulesServiceError_ErrorCode_name = map[int32]string{
		0: "OK",
		1: "INVALID_MODULE",
		2: "INVALID_VERSION",
		3: "INVALID_INSTANCES",
		4: "TRANSIENT_ERROR",
		5: "UNEXPECTED_STATE",
	}
	ModulesServiceError_ErrorCode_value = map[string]int32{
		"OK":                0,
		"INVALID_MODULE":    1,
		"INVALID_VERSION":   2,
		"INVALID_INSTANCES": 3,
		"TRANSIENT_ERROR":   4,
		"UNEXPECTED_STATE":  5,
	}
)

func (x ModulesServiceError_ErrorCode) Enum() *ModulesServiceError_ErrorCode {
	p := new(ModulesServiceError_ErrorCode)
	*p = x
	return p
}

func (x ModulesServiceError_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModulesServiceError_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_modules_service_proto_enumTypes[0].Descriptor()
}

func (ModulesServiceError_ErrorCode) Type() protoreflect.EnumType {
	return &file_modules_service_proto_enumTypes[0]
}

func (x ModulesServiceError_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ModulesServiceError_ErrorCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ModulesServiceError_ErrorCode(num)
	return nil
}

// Deprecated: Use ModulesServiceError_ErrorCode.Descriptor instead.
func (ModulesServiceError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{0, 0}
}

type ModulesServiceError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ModulesServiceError) Reset() {
	*x = ModulesServiceError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModulesServiceError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModulesServiceError) ProtoMessage() {}

func (x *ModulesServiceError) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModulesServiceError.ProtoReflect.Descriptor instead.
func (*ModulesServiceError) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{0}
}

type GetModulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetModulesRequest) Reset() {
	*x = GetModulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModulesRequest) ProtoMessage() {}

func (x *GetModulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModulesRequest.ProtoReflect.Descriptor instead.
func (*GetModulesRequest) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{1}
}

type GetModulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module []string `protobuf:"bytes,1,rep,name=module" json:"module,omitempty"`
}

func (x *GetModulesResponse) Reset() {
	*x = GetModulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModulesResponse) ProtoMessage() {}

func (x *GetModulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModulesResponse.ProtoReflect.Descriptor instead.
func (*GetModulesResponse) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetModulesResponse) GetModule() []string {
	if x != nil {
		return x.Module
	}
	return nil
}

type GetVersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
}

func (x *GetVersionsRequest) Reset() {
	*x = GetVersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionsRequest) ProtoMessage() {}

func (x *GetVersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionsRequest.ProtoReflect.Descriptor instead.
func (*GetVersionsRequest) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetVersionsRequest) GetModule() string {
	if x != nil && x.Module != nil {
		return *x.Module
	}
	return ""
}

type GetVersionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version []string `protobuf:"bytes,1,rep,name=version" json:"version,omitempty"`
}

func (x *GetVersionsResponse) Reset() {
	*x = GetVersionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionsResponse) ProtoMessage() {}

func (x *GetVersionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionsResponse.ProtoReflect.Descriptor instead.
func (*GetVersionsResponse) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetVersionsResponse) GetVersion() []string {
	if x != nil {
		return x.Version
	}
	return nil
}

type GetDefaultVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
}

func (x *GetDefaultVersionRequest) Reset() {
	*x = GetDefaultVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefaultVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefaultVersionRequest) ProtoMessage() {}

func (x *GetDefaultVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefaultVersionRequest.ProtoReflect.Descriptor instead.
func (*GetDefaultVersionRequest) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetDefaultVersionRequest) GetModule() string {
	if x != nil && x.Module != nil {
		return *x.Module
	}
	return ""
}

type GetDefaultVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *string `protobuf:"bytes,1,req,name=version" json:"version,omitempty"`
}

func (x *GetDefaultVersionResponse) Reset() {
	*x = GetDefaultVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDefaultVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDefaultVersionResponse) ProtoMessage() {}

func (x *GetDefaultVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDefaultVersionResponse.ProtoReflect.Descriptor instead.
func (*GetDefaultVersionResponse) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetDefaultVersionResponse) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

type GetNumInstancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module  *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version *string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (x *GetNumInstancesRequest) Reset() {
	*x = GetNumInstancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumInstancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumInstancesRequest) ProtoMessage() {}

func (x *GetNumInstancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumInstancesRequest.ProtoReflect.Descriptor instead.
func (*GetNumInstancesRequest) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetNumInstancesRequest) GetModule() string {
	if x != nil && x.Module != nil {
		return *x.Module
	}
	return ""
}

func (x *GetNumInstancesRequest) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

type GetNumInstancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instances *int64 `protobuf:"varint,1,req,name=instances" json:"instances,omitempty"`
}

func (x *GetNumInstancesResponse) Reset() {
	*x = GetNumInstancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumInstancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumInstancesResponse) ProtoMessage() {}

func (x *GetNumInstancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumInstancesResponse.ProtoReflect.Descriptor instead.
func (*GetNumInstancesResponse) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetNumInstancesResponse) GetInstances() int64 {
	if x != nil && x.Instances != nil {
		return *x.Instances
	}
	return 0
}

type SetNumInstancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module    *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version   *string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Instances *int64  `protobuf:"varint,3,req,name=instances" json:"instances,omitempty"`
}

func (x *SetNumInstancesRequest) Reset() {
	*x = SetNumInstancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNumInstancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNumInstancesRequest) ProtoMessage() {}

func (x *SetNumInstancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNumInstancesRequest.ProtoReflect.Descriptor instead.
func (*SetNumInstancesRequest) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{9}
}

func (x *SetNumInstancesRequest) GetModule() string {
	if x != nil && x.Module != nil {
		return *x.Module
	}
	return ""
}

func (x *SetNumInstancesRequest) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *SetNumInstancesRequest) GetInstances() int64 {
	if x != nil && x.Instances != nil {
		return *x.Instances
	}
	return 0
}

type SetNumInstancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetNumInstancesResponse) Reset() {
	*x = SetNumInstancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNumInstancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNumInstancesResponse) ProtoMessage() {}

func (x *SetNumInstancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNumInstancesResponse.ProtoReflect.Descriptor instead.
func (*SetNumInstancesResponse) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{10}
}

type StartModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module  *string `protobuf:"bytes,1,req,name=module" json:"module,omitempty"`
	Version *string `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
}

func (x *StartModuleRequest) Reset() {
	*x = StartModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartModuleRequest) ProtoMessage() {}

func (x *StartModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartModuleRequest.ProtoReflect.Descriptor instead.
func (*StartModuleRequest) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{11}
}

func (x *StartModuleRequest) GetModule() string {
	if x != nil && x.Module != nil {
		return *x.Module
	}
	return ""
}

func (x *StartModuleRequest) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

type StartModuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartModuleResponse) Reset() {
	*x = StartModuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartModuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartModuleResponse) ProtoMessage() {}

func (x *StartModuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartModuleResponse.ProtoReflect.Descriptor instead.
func (*StartModuleResponse) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{12}
}

type StopModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module  *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version *string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (x *StopModuleRequest) Reset() {
	*x = StopModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopModuleRequest) ProtoMessage() {}

func (x *StopModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopModuleRequest.ProtoReflect.Descriptor instead.
func (*StopModuleRequest) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{13}
}

func (x *StopModuleRequest) GetModule() string {
	if x != nil && x.Module != nil {
		return *x.Module
	}
	return ""
}

func (x *StopModuleRequest) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

type StopModuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopModuleResponse) Reset() {
	*x = StopModuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopModuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopModuleResponse) ProtoMessage() {}

func (x *StopModuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopModuleResponse.ProtoReflect.Descriptor instead.
func (*StopModuleResponse) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{14}
}

type GetHostnameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module   *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version  *string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Instance *string `protobuf:"bytes,3,opt,name=instance" json:"instance,omitempty"`
}

func (x *GetHostnameRequest) Reset() {
	*x = GetHostnameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostnameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostnameRequest) ProtoMessage() {}

func (x *GetHostnameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostnameRequest.ProtoReflect.Descriptor instead.
func (*GetHostnameRequest) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{15}
}

func (x *GetHostnameRequest) GetModule() string {
	if x != nil && x.Module != nil {
		return *x.Module
	}
	return ""
}

func (x *GetHostnameRequest) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *GetHostnameRequest) GetInstance() string {
	if x != nil && x.Instance != nil {
		return *x.Instance
	}
	return ""
}

type GetHostnameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname *string `protobuf:"bytes,1,req,name=hostname" json:"hostname,omitempty"`
}

func (x *GetHostnameResponse) Reset() {
	*x = GetHostnameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_service_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostnameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostnameResponse) ProtoMessage() {}

func (x *GetHostnameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_modules_service_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostnameResponse.ProtoReflect.Descriptor instead.
func (*GetHostnameResponse) Descriptor() ([]byte, []int) {
	return file_modules_service_proto_rawDescGZIP(), []int{16}
}

func (x *GetHostnameResponse) GetHostname() string {
	if x != nil && x.Hostname != nil {
		return *x.Hostname
	}
	return ""
}

var File_modules_service_proto protoreflect.FileDescriptor

var file_modules_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x32, 0x22, 0x95, 0x01, 0x0a, 0x13, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x7e, 0x0a,
	0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4d, 0x4f,
	0x44, 0x55, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x53,
	0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x45, 0x4e, 0x54, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x45, 0x58, 0x50,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x05, 0x22, 0x13, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x2f,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x32, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x22, 0x35, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22,
	0x68, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x03, 0x52, 0x09,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x65, 0x74,
	0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x15, 0x0a, 0x13,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x70, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x74,
	0x6f, 0x70, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x62, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73,
}

var (
	file_modules_service_proto_rawDescOnce sync.Once
	file_modules_service_proto_rawDescData = file_modules_service_proto_rawDesc
)

func file_modules_service_proto_rawDescGZIP() []byte {
	file_modules_service_proto_rawDescOnce.Do(func() {
		file_modules_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_service_proto_rawDescData)
	})
	return file_modules_service_proto_rawDescData
}

var file_modules_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_modules_service_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_modules_service_proto_goTypes = []interface{}{
	(ModulesServiceError_ErrorCode)(0), // 0: appengine.v2.ModulesServiceError.ErrorCode
	(*ModulesServiceError)(nil),        // 1: appengine.v2.ModulesServiceError
	(*GetModulesRequest)(nil),          // 2: appengine.v2.GetModulesRequest
	(*GetModulesResponse)(nil),         // 3: appengine.v2.GetModulesResponse
	(*GetVersionsRequest)(nil),         // 4: appengine.v2.GetVersionsRequest
	(*GetVersionsResponse)(nil),        // 5: appengine.v2.GetVersionsResponse
	(*GetDefaultVersionRequest)(nil),   // 6: appengine.v2.GetDefaultVersionRequest
	(*GetDefaultVersionResponse)(nil),  // 7: appengine.v2.GetDefaultVersionResponse
	(*GetNumInstancesRequest)(nil),     // 8: appengine.v2.GetNumInstancesRequest
	(*GetNumInstancesResponse)(nil),    // 9: appengine.v2.GetNumInstancesResponse
	(*SetNumInstancesRequest)(nil),     // 10: appengine.v2.SetNumInstancesRequest
	(*SetNumInstancesResponse)(nil),    // 11: appengine.v2.SetNumInstancesResponse
	(*StartModuleRequest)(nil),         // 12: appengine.v2.StartModuleRequest
	(*StartModuleResponse)(nil),        // 13: appengine.v2.StartModuleResponse
	(*StopModuleRequest)(nil),          // 14: appengine.v2.StopModuleRequest
	(*StopModuleResponse)(nil),         // 15: appengine.v2.StopModuleResponse
	(*GetHostnameRequest)(nil),         // 16: appengine.v2.GetHostnameRequest
	(*GetHostnameResponse)(nil),        // 17: appengine.v2.GetHostnameResponse
}
var file_modules_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_service_proto_init() }
func file_modules_service_proto_init() {
	if File_modules_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModulesServiceError); i {
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
		file_modules_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModulesRequest); i {
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
		file_modules_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModulesResponse); i {
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
		file_modules_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionsRequest); i {
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
		file_modules_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionsResponse); i {
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
		file_modules_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefaultVersionRequest); i {
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
		file_modules_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDefaultVersionResponse); i {
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
		file_modules_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumInstancesRequest); i {
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
		file_modules_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumInstancesResponse); i {
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
		file_modules_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNumInstancesRequest); i {
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
		file_modules_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNumInstancesResponse); i {
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
		file_modules_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartModuleRequest); i {
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
		file_modules_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartModuleResponse); i {
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
		file_modules_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopModuleRequest); i {
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
		file_modules_service_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopModuleResponse); i {
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
		file_modules_service_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostnameRequest); i {
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
		file_modules_service_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostnameResponse); i {
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
			RawDescriptor: file_modules_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_modules_service_proto_goTypes,
		DependencyIndexes: file_modules_service_proto_depIdxs,
		EnumInfos:         file_modules_service_proto_enumTypes,
		MessageInfos:      file_modules_service_proto_msgTypes,
	}.Build()
	File_modules_service_proto = out.File
	file_modules_service_proto_rawDesc = nil
	file_modules_service_proto_goTypes = nil
	file_modules_service_proto_depIdxs = nil
}
