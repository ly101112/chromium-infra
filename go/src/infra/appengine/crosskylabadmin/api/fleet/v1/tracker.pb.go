// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/crosskylabadmin/api/fleet/v1/tracker.proto

/*
Package fleet is a generated protocol buffer package.

It is generated from these files:
	infra/appengine/crosskylabadmin/api/fleet/v1/tracker.proto

It has these top-level messages:
	RefreshBotsRequest
	RefreshBotsResponse
	SummarizeBotsRequest
	BotSelector
	SummarizeBotsResponse
	BotSummary
*/
package fleet

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

// RefreshBotsRequest can be used to restrict the Swarming bots to refresh via
// the Tracker.RefreshBots rpc.
type RefreshBotsRequest struct {
	// selectors whitelists the bots to refresh. This includes new bots
	// discovered from Swarming matching the selectors.
	// Bots selected via repeated selectors are unioned together.
	//
	// If no selectors are provided, all bots are selected.
	Selectors []*BotSelector `protobuf:"bytes,2,rep,name=selectors" json:"selectors,omitempty"`
}

func (m *RefreshBotsRequest) Reset()                    { *m = RefreshBotsRequest{} }
func (m *RefreshBotsRequest) String() string            { return proto.CompactTextString(m) }
func (*RefreshBotsRequest) ProtoMessage()               {}
func (*RefreshBotsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RefreshBotsRequest) GetSelectors() []*BotSelector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

// RefreshBotsResponse contains information about the Swarming bots actually
// refreshed in response to a Tracker.RefreshBots rpc.
type RefreshBotsResponse struct {
	// dut_ids lists the dut_id of of the bots refreshed.
	DutIds []string `protobuf:"bytes,1,rep,name=dut_ids,json=dutIds" json:"dut_ids,omitempty"`
}

func (m *RefreshBotsResponse) Reset()                    { *m = RefreshBotsResponse{} }
func (m *RefreshBotsResponse) String() string            { return proto.CompactTextString(m) }
func (*RefreshBotsResponse) ProtoMessage()               {}
func (*RefreshBotsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RefreshBotsResponse) GetDutIds() []string {
	if m != nil {
		return m.DutIds
	}
	return nil
}

// SummarizeBotsRequest can be used to restrict the Swarming bots to summarize
// via the Tracker.SummarizeBots rpc.
type SummarizeBotsRequest struct {
	// selectors whitelists the bots to refresh, from the already known bots to
	// Tracker. Bots selected via repeated selectors are unioned together.
	//
	// If no selectors are provided, all bots are selected.
	Selectors []*BotSelector `protobuf:"bytes,1,rep,name=selectors" json:"selectors,omitempty"`
}

func (m *SummarizeBotsRequest) Reset()                    { *m = SummarizeBotsRequest{} }
func (m *SummarizeBotsRequest) String() string            { return proto.CompactTextString(m) }
func (*SummarizeBotsRequest) ProtoMessage()               {}
func (*SummarizeBotsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SummarizeBotsRequest) GetSelectors() []*BotSelector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

// BotSelector is used to whitelist Swarming bots to select for various Tracker
// rpcs.
type BotSelector struct {
	// dut_ids selects bots by the dut_id dimension.
	DutId string `protobuf:"bytes,1,opt,name=dut_id,json=dutId" json:"dut_id,omitempty"`
}

func (m *BotSelector) Reset()                    { *m = BotSelector{} }
func (m *BotSelector) String() string            { return proto.CompactTextString(m) }
func (*BotSelector) ProtoMessage()               {}
func (*BotSelector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BotSelector) GetDutId() string {
	if m != nil {
		return m.DutId
	}
	return ""
}

// SummarizeBotsResponse contains summary information about Swarming bots
// returned by the Tracker.SummarizeBots rpc.
type SummarizeBotsResponse struct {
	Bots []*BotSummary `protobuf:"bytes,1,rep,name=bots" json:"bots,omitempty"`
}

func (m *SummarizeBotsResponse) Reset()                    { *m = SummarizeBotsResponse{} }
func (m *SummarizeBotsResponse) String() string            { return proto.CompactTextString(m) }
func (*SummarizeBotsResponse) ProtoMessage()               {}
func (*SummarizeBotsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SummarizeBotsResponse) GetBots() []*BotSummary {
	if m != nil {
		return m.Bots
	}
	return nil
}

// BotSummary contains the summary information tracked by Tracker for a single
// Skylab Swarming bot.
type BotSummary struct {
	// dut_id contains the dut_id dimension for the bot.
	DutId string `protobuf:"bytes,1,opt,name=dut_id,json=dutId" json:"dut_id,omitempty"`
}

func (m *BotSummary) Reset()                    { *m = BotSummary{} }
func (m *BotSummary) String() string            { return proto.CompactTextString(m) }
func (*BotSummary) ProtoMessage()               {}
func (*BotSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BotSummary) GetDutId() string {
	if m != nil {
		return m.DutId
	}
	return ""
}

func init() {
	proto.RegisterType((*RefreshBotsRequest)(nil), "crosskylabadmin.fleet.RefreshBotsRequest")
	proto.RegisterType((*RefreshBotsResponse)(nil), "crosskylabadmin.fleet.RefreshBotsResponse")
	proto.RegisterType((*SummarizeBotsRequest)(nil), "crosskylabadmin.fleet.SummarizeBotsRequest")
	proto.RegisterType((*BotSelector)(nil), "crosskylabadmin.fleet.BotSelector")
	proto.RegisterType((*SummarizeBotsResponse)(nil), "crosskylabadmin.fleet.SummarizeBotsResponse")
	proto.RegisterType((*BotSummary)(nil), "crosskylabadmin.fleet.BotSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tracker service

type TrackerClient interface {
	// RefreshBots instructs the Tracker service to update Swarming bot
	// information from the Swarming server hosting ChromeOS Skylab bots.
	//
	// RefreshBots stops at the first error encountered and returns the error. A
	// failed RefreshBots call may have refreshed some of the bots requested.
	// It is safe to call RefreshBots to continue from a partially failed call.
	RefreshBots(ctx context.Context, in *RefreshBotsRequest, opts ...grpc.CallOption) (*RefreshBotsResponse, error)
	// SummarizeBots returns summary information about Swarming bots.
	// This includes ChromeOS Skylab specific dimensions/state information as
	// well as a summary of the recenty history of administrative tasks.
	//
	// SummarizeBots stops at the first error encountered and returns the error.
	SummarizeBots(ctx context.Context, in *SummarizeBotsRequest, opts ...grpc.CallOption) (*SummarizeBotsResponse, error)
}
type trackerPRPCClient struct {
	client *prpc.Client
}

func NewTrackerPRPCClient(client *prpc.Client) TrackerClient {
	return &trackerPRPCClient{client}
}

func (c *trackerPRPCClient) RefreshBots(ctx context.Context, in *RefreshBotsRequest, opts ...grpc.CallOption) (*RefreshBotsResponse, error) {
	out := new(RefreshBotsResponse)
	err := c.client.Call(ctx, "crosskylabadmin.fleet.Tracker", "RefreshBots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerPRPCClient) SummarizeBots(ctx context.Context, in *SummarizeBotsRequest, opts ...grpc.CallOption) (*SummarizeBotsResponse, error) {
	out := new(SummarizeBotsResponse)
	err := c.client.Call(ctx, "crosskylabadmin.fleet.Tracker", "SummarizeBots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type trackerClient struct {
	cc *grpc.ClientConn
}

func NewTrackerClient(cc *grpc.ClientConn) TrackerClient {
	return &trackerClient{cc}
}

func (c *trackerClient) RefreshBots(ctx context.Context, in *RefreshBotsRequest, opts ...grpc.CallOption) (*RefreshBotsResponse, error) {
	out := new(RefreshBotsResponse)
	err := grpc.Invoke(ctx, "/crosskylabadmin.fleet.Tracker/RefreshBots", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackerClient) SummarizeBots(ctx context.Context, in *SummarizeBotsRequest, opts ...grpc.CallOption) (*SummarizeBotsResponse, error) {
	out := new(SummarizeBotsResponse)
	err := grpc.Invoke(ctx, "/crosskylabadmin.fleet.Tracker/SummarizeBots", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tracker service

type TrackerServer interface {
	// RefreshBots instructs the Tracker service to update Swarming bot
	// information from the Swarming server hosting ChromeOS Skylab bots.
	//
	// RefreshBots stops at the first error encountered and returns the error. A
	// failed RefreshBots call may have refreshed some of the bots requested.
	// It is safe to call RefreshBots to continue from a partially failed call.
	RefreshBots(context.Context, *RefreshBotsRequest) (*RefreshBotsResponse, error)
	// SummarizeBots returns summary information about Swarming bots.
	// This includes ChromeOS Skylab specific dimensions/state information as
	// well as a summary of the recenty history of administrative tasks.
	//
	// SummarizeBots stops at the first error encountered and returns the error.
	SummarizeBots(context.Context, *SummarizeBotsRequest) (*SummarizeBotsResponse, error)
}

func RegisterTrackerServer(s prpc.Registrar, srv TrackerServer) {
	s.RegisterService(&_Tracker_serviceDesc, srv)
}

func _Tracker_RefreshBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshBotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).RefreshBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crosskylabadmin.fleet.Tracker/RefreshBots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).RefreshBots(ctx, req.(*RefreshBotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tracker_SummarizeBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummarizeBotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).SummarizeBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crosskylabadmin.fleet.Tracker/SummarizeBots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).SummarizeBots(ctx, req.(*SummarizeBotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crosskylabadmin.fleet.Tracker",
	HandlerType: (*TrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RefreshBots",
			Handler:    _Tracker_RefreshBots_Handler,
		},
		{
			MethodName: "SummarizeBots",
			Handler:    _Tracker_SummarizeBots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/appengine/crosskylabadmin/api/fleet/v1/tracker.proto",
}

func init() {
	proto.RegisterFile("infra/appengine/crosskylabadmin/api/fleet/v1/tracker.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3d, 0x4f, 0xf3, 0x30,
	0x10, 0xc7, 0x95, 0xa7, 0x4f, 0x5b, 0xf5, 0x2a, 0x16, 0x43, 0x45, 0xd5, 0xa9, 0x18, 0x86, 0xf2,
	0xa2, 0x58, 0x14, 0xb1, 0x30, 0xa1, 0x6e, 0x2c, 0x0c, 0x2e, 0x42, 0x88, 0x05, 0xb9, 0xf5, 0x05,
	0x4c, 0x93, 0x38, 0xd8, 0x0e, 0x52, 0xf9, 0xaa, 0x7c, 0x19, 0x24, 0x27, 0xa8, 0x2f, 0x34, 0xa8,
	0x62, 0xcd, 0xfd, 0xee, 0x7e, 0xff, 0xcb, 0x19, 0xae, 0x54, 0x1a, 0x19, 0xc1, 0x44, 0x96, 0x61,
	0xfa, 0xac, 0x52, 0x64, 0x53, 0xa3, 0xad, 0x9d, 0xcd, 0x63, 0x31, 0x11, 0x32, 0x51, 0x29, 0x13,
	0x99, 0x62, 0x51, 0x8c, 0xe8, 0xd8, 0xfb, 0x39, 0x73, 0x46, 0x4c, 0x67, 0x68, 0xc2, 0xcc, 0x68,
	0xa7, 0x49, 0x67, 0x8d, 0x0d, 0x3d, 0x47, 0xef, 0x81, 0x70, 0x8c, 0x0c, 0xda, 0x97, 0x91, 0x76,
	0x96, 0xe3, 0x5b, 0x8e, 0xd6, 0x91, 0x6b, 0x68, 0x59, 0x8c, 0x71, 0xea, 0xb4, 0xb1, 0xdd, 0x7f,
	0xfd, 0xda, 0xa0, 0x3d, 0xa4, 0xe1, 0xc6, 0x01, 0xe1, 0x48, 0xbb, 0x71, 0x89, 0xf2, 0x45, 0x13,
	0x0d, 0x61, 0x77, 0x65, 0xae, 0xcd, 0x74, 0x6a, 0x91, 0xec, 0x43, 0x53, 0xe6, 0xee, 0x49, 0x49,
	0xdb, 0x0d, 0xfa, 0xb5, 0x41, 0x8b, 0x37, 0x64, 0xee, 0x6e, 0xa4, 0xa5, 0x0f, 0xb0, 0x37, 0xce,
	0x93, 0x44, 0x18, 0xf5, 0x81, 0x95, 0x49, 0x82, 0xbf, 0x24, 0x39, 0x82, 0xf6, 0x52, 0x85, 0x74,
	0xa0, 0x51, 0x24, 0xe8, 0x06, 0xfd, 0x60, 0xd0, 0xe2, 0x75, 0x1f, 0x80, 0xde, 0x42, 0x67, 0xcd,
	0x5f, 0x26, 0xbe, 0x84, 0xff, 0x13, 0xed, 0xbe, 0xdd, 0x07, 0xbf, 0xb8, 0x7d, 0xfb, 0x9c, 0x7b,
	0x9c, 0x1e, 0x02, 0x2c, 0xbe, 0x55, 0x48, 0x87, 0x9f, 0x01, 0x34, 0xef, 0x8a, 0x2b, 0x11, 0x09,
	0xed, 0xa5, 0x1f, 0x46, 0x8e, 0x2b, 0x44, 0x3f, 0x8f, 0xd5, 0x3b, 0xd9, 0x06, 0x2d, 0xb7, 0x79,
	0x85, 0x9d, 0x95, 0x35, 0xc9, 0x69, 0x45, 0xf3, 0xa6, 0x63, 0xf4, 0xce, 0xb6, 0x83, 0x0b, 0xd7,
	0xa8, 0xf9, 0x58, 0xf7, 0xe5, 0x49, 0xc3, 0xbf, 0xc0, 0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbb, 0x5c, 0xe3, 0x52, 0xbf, 0x02, 0x00, 0x00,
}
