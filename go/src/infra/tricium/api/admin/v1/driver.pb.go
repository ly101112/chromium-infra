// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/tricium/api/admin/v1/driver.proto

package admin

import prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TriggerRequest contains the details needed to launch a swarming task for a
// Tricium worker.
type TriggerRequest struct {
	RunId                int64    `protobuf:"varint,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	IsolatedInputHash    string   `protobuf:"bytes,2,opt,name=isolated_input_hash,json=isolatedInputHash,proto3" json:"isolated_input_hash,omitempty"`
	Worker               string   `protobuf:"bytes,3,opt,name=worker,proto3" json:"worker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TriggerRequest) Reset()         { *m = TriggerRequest{} }
func (m *TriggerRequest) String() string { return proto.CompactTextString(m) }
func (*TriggerRequest) ProtoMessage()    {}
func (*TriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_8471842ea4be6008, []int{0}
}
func (m *TriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerRequest.Unmarshal(m, b)
}
func (m *TriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerRequest.Marshal(b, m, deterministic)
}
func (dst *TriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerRequest.Merge(dst, src)
}
func (m *TriggerRequest) XXX_Size() int {
	return xxx_messageInfo_TriggerRequest.Size(m)
}
func (m *TriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerRequest proto.InternalMessageInfo

func (m *TriggerRequest) GetRunId() int64 {
	if m != nil {
		return m.RunId
	}
	return 0
}

func (m *TriggerRequest) GetIsolatedInputHash() string {
	if m != nil {
		return m.IsolatedInputHash
	}
	return ""
}

func (m *TriggerRequest) GetWorker() string {
	if m != nil {
		return m.Worker
	}
	return ""
}

type TriggerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TriggerResponse) Reset()         { *m = TriggerResponse{} }
func (m *TriggerResponse) String() string { return proto.CompactTextString(m) }
func (*TriggerResponse) ProtoMessage()    {}
func (*TriggerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_8471842ea4be6008, []int{1}
}
func (m *TriggerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerResponse.Unmarshal(m, b)
}
func (m *TriggerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerResponse.Marshal(b, m, deterministic)
}
func (dst *TriggerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerResponse.Merge(dst, src)
}
func (m *TriggerResponse) XXX_Size() int {
	return xxx_messageInfo_TriggerResponse.Size(m)
}
func (m *TriggerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerResponse proto.InternalMessageInfo

// CollectRequest contains the details needed to collect results from a swarming task
// running a Tricium worker and to launch succeeding Tricium workers.
type CollectRequest struct {
	RunId int64 `protobuf:"varint,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	// The isolated input of this worker is imported in the input hash of any successor
	// workers of this worker. Passing it along here to make sure it is available.
	IsolatedInputHash string `protobuf:"bytes,2,opt,name=isolated_input_hash,json=isolatedInputHash,proto3" json:"isolated_input_hash,omitempty"`
	// Points out which worker to collect results for. This worker name is used to
	// mangle the ID of the swarming task running the worker.
	Worker string `protobuf:"bytes,3,opt,name=worker,proto3" json:"worker,omitempty"`
	// Only one of task_id and build_id should be populated, depending on which
	// type of task is being collected (buildbucket or swarming.)
	// The Swarming task ID.
	//
	// Used to collect results from the completed swarming worker task.
	TaskId string `protobuf:"bytes,4,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// The Buildbucket build ID.
	//
	// Used to collect results from the completed buildbucket worker task.
	BuildId              int64    `protobuf:"varint,5,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectRequest) Reset()         { *m = CollectRequest{} }
func (m *CollectRequest) String() string { return proto.CompactTextString(m) }
func (*CollectRequest) ProtoMessage()    {}
func (*CollectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_8471842ea4be6008, []int{2}
}
func (m *CollectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectRequest.Unmarshal(m, b)
}
func (m *CollectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectRequest.Marshal(b, m, deterministic)
}
func (dst *CollectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectRequest.Merge(dst, src)
}
func (m *CollectRequest) XXX_Size() int {
	return xxx_messageInfo_CollectRequest.Size(m)
}
func (m *CollectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectRequest proto.InternalMessageInfo

func (m *CollectRequest) GetRunId() int64 {
	if m != nil {
		return m.RunId
	}
	return 0
}

func (m *CollectRequest) GetIsolatedInputHash() string {
	if m != nil {
		return m.IsolatedInputHash
	}
	return ""
}

func (m *CollectRequest) GetWorker() string {
	if m != nil {
		return m.Worker
	}
	return ""
}

func (m *CollectRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *CollectRequest) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

type CollectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectResponse) Reset()         { *m = CollectResponse{} }
func (m *CollectResponse) String() string { return proto.CompactTextString(m) }
func (*CollectResponse) ProtoMessage()    {}
func (*CollectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_8471842ea4be6008, []int{3}
}
func (m *CollectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectResponse.Unmarshal(m, b)
}
func (m *CollectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectResponse.Marshal(b, m, deterministic)
}
func (dst *CollectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectResponse.Merge(dst, src)
}
func (m *CollectResponse) XXX_Size() int {
	return xxx_messageInfo_CollectResponse.Size(m)
}
func (m *CollectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CollectResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TriggerRequest)(nil), "admin.TriggerRequest")
	proto.RegisterType((*TriggerResponse)(nil), "admin.TriggerResponse")
	proto.RegisterType((*CollectRequest)(nil), "admin.CollectRequest")
	proto.RegisterType((*CollectResponse)(nil), "admin.CollectResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DriverClient is the client API for Driver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriverClient interface {
	// Trigger triggers a swarming task for a Tricium worker.
	Trigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error)
	// Collect collects results from a swarming task running a Tricium worker.
	Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*CollectResponse, error)
}
type driverPRPCClient struct {
	client *prpc.Client
}

func NewDriverPRPCClient(client *prpc.Client) DriverClient {
	return &driverPRPCClient{client}
}

func (c *driverPRPCClient) Trigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error) {
	out := new(TriggerResponse)
	err := c.client.Call(ctx, "admin.Driver", "Trigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverPRPCClient) Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*CollectResponse, error) {
	out := new(CollectResponse)
	err := c.client.Call(ctx, "admin.Driver", "Collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type driverClient struct {
	cc *grpc.ClientConn
}

func NewDriverClient(cc *grpc.ClientConn) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) Trigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error) {
	out := new(TriggerResponse)
	err := c.cc.Invoke(ctx, "/admin.Driver/Trigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*CollectResponse, error) {
	out := new(CollectResponse)
	err := c.cc.Invoke(ctx, "/admin.Driver/Collect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServer is the server API for Driver service.
type DriverServer interface {
	// Trigger triggers a swarming task for a Tricium worker.
	Trigger(context.Context, *TriggerRequest) (*TriggerResponse, error)
	// Collect collects results from a swarming task running a Tricium worker.
	Collect(context.Context, *CollectRequest) (*CollectResponse, error)
}

func RegisterDriverServer(s prpc.Registrar, srv DriverServer) {
	s.RegisterService(&_Driver_serviceDesc, srv)
}

func _Driver_Trigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Trigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Driver/Trigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Trigger(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Driver/Collect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).Collect(ctx, req.(*CollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Driver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Trigger",
			Handler:    _Driver_Trigger_Handler,
		},
		{
			MethodName: "Collect",
			Handler:    _Driver_Collect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/tricium/api/admin/v1/driver.proto",
}

func init() {
	proto.RegisterFile("infra/tricium/api/admin/v1/driver.proto", fileDescriptor_driver_8471842ea4be6008)
}

var fileDescriptor_driver_8471842ea4be6008 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x51, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x55, 0x28, 0x4d, 0xe1, 0x86, 0xa2, 0x1a, 0xb5, 0x04, 0xa6, 0x2a, 0x0b, 0x9d, 0x12, 0x01,
	0x0b, 0x3b, 0x0c, 0x64, 0x8d, 0xd8, 0x23, 0xb7, 0x36, 0xcd, 0xa9, 0xa9, 0x1d, 0xce, 0x76, 0x59,
	0xf8, 0x37, 0xfc, 0x51, 0x64, 0x27, 0x44, 0x6a, 0x7f, 0x00, 0xe3, 0xbb, 0x97, 0xf7, 0x91, 0x67,
	0xb8, 0x47, 0xf5, 0x41, 0x3c, 0xb7, 0x84, 0x1b, 0x74, 0xfb, 0x9c, 0xb7, 0x98, 0x73, 0xb1, 0x47,
	0x95, 0x1f, 0x1e, 0x72, 0x41, 0x78, 0x90, 0x94, 0xb5, 0xa4, 0xad, 0x66, 0xe3, 0x70, 0x4e, 0x35,
	0x4c, 0xdf, 0x09, 0xb7, 0x5b, 0x49, 0xa5, 0xfc, 0x74, 0xd2, 0x58, 0x36, 0x87, 0x98, 0x9c, 0xaa,
	0x50, 0x24, 0xd1, 0x32, 0x5a, 0x8d, 0xca, 0x31, 0x39, 0x55, 0x08, 0x96, 0xc1, 0x35, 0x1a, 0xdd,
	0x70, 0x2b, 0x45, 0x85, 0xaa, 0x75, 0xb6, 0xaa, 0xb9, 0xa9, 0x93, 0xb3, 0x65, 0xb4, 0xba, 0x2c,
	0x67, 0x7f, 0x54, 0xe1, 0x99, 0x37, 0x6e, 0x6a, 0xb6, 0x80, 0xf8, 0x4b, 0xd3, 0x4e, 0x52, 0x32,
	0x0a, 0x9f, 0xf4, 0x28, 0x9d, 0xc1, 0xd5, 0x10, 0x68, 0x5a, 0xad, 0x8c, 0x4c, 0x7f, 0x22, 0x98,
	0xbe, 0xe8, 0xa6, 0x91, 0x1b, 0xfb, 0x3f, 0x25, 0xd8, 0x0d, 0x4c, 0x2c, 0x37, 0x3b, 0xef, 0x7f,
	0xde, 0x11, 0x1e, 0x16, 0x82, 0xdd, 0xc2, 0xc5, 0xda, 0x61, 0x23, 0x3c, 0x33, 0x0e, 0xc9, 0x93,
	0x80, 0x0b, 0xe1, 0x8b, 0x0f, 0x25, 0xbb, 0xe2, 0x8f, 0xdf, 0x10, 0xbf, 0x86, 0x4d, 0xd9, 0x33,
	0x4c, 0xfa, 0xbf, 0x62, 0xf3, 0x2c, 0x2c, 0x9b, 0x1d, 0xcf, 0x7a, 0xb7, 0x38, 0x3d, 0x77, 0x1e,
	0x5e, 0xd9, 0xdb, 0x0e, 0xca, 0xe3, 0x2d, 0x06, 0xe5, 0x49, 0xfa, 0x3a, 0x0e, 0x0f, 0xf9, 0xf4,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x06, 0x6f, 0xea, 0x76, 0xf3, 0x01, 0x00, 0x00,
}
