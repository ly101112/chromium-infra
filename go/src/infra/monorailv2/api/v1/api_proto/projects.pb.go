// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/api_proto/projects.proto

package monorail_v1

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for ListIssueTemplates
// Next available tag: 4
type ListIssueTemplatesRequest struct {
	// The name of the project these templates belong to.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return. The service may return fewer than
	// this value.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListIssueTemplatesRequest` call.
	// Provide this to retrieve the subsequent page.
	// When paginating, all other parameters provided to
	// `ListIssueTemplatesRequest` must match the call that provided the token.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListIssueTemplatesRequest) Reset()         { *m = ListIssueTemplatesRequest{} }
func (m *ListIssueTemplatesRequest) String() string { return proto.CompactTextString(m) }
func (*ListIssueTemplatesRequest) ProtoMessage()    {}
func (*ListIssueTemplatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c420c733c90f8ec5, []int{0}
}

func (m *ListIssueTemplatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIssueTemplatesRequest.Unmarshal(m, b)
}
func (m *ListIssueTemplatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIssueTemplatesRequest.Marshal(b, m, deterministic)
}
func (m *ListIssueTemplatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIssueTemplatesRequest.Merge(m, src)
}
func (m *ListIssueTemplatesRequest) XXX_Size() int {
	return xxx_messageInfo_ListIssueTemplatesRequest.Size(m)
}
func (m *ListIssueTemplatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIssueTemplatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListIssueTemplatesRequest proto.InternalMessageInfo

func (m *ListIssueTemplatesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListIssueTemplatesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListIssueTemplatesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response message for ListIssueTemplates
// Next available tag: 3
type ListIssueTemplatesResponse struct {
	Templates []*IssueTemplate `protobuf:"bytes,1,rep,name=templates,proto3" json:"templates,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListIssueTemplatesResponse) Reset()         { *m = ListIssueTemplatesResponse{} }
func (m *ListIssueTemplatesResponse) String() string { return proto.CompactTextString(m) }
func (*ListIssueTemplatesResponse) ProtoMessage()    {}
func (*ListIssueTemplatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c420c733c90f8ec5, []int{1}
}

func (m *ListIssueTemplatesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIssueTemplatesResponse.Unmarshal(m, b)
}
func (m *ListIssueTemplatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIssueTemplatesResponse.Marshal(b, m, deterministic)
}
func (m *ListIssueTemplatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIssueTemplatesResponse.Merge(m, src)
}
func (m *ListIssueTemplatesResponse) XXX_Size() int {
	return xxx_messageInfo_ListIssueTemplatesResponse.Size(m)
}
func (m *ListIssueTemplatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIssueTemplatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListIssueTemplatesResponse proto.InternalMessageInfo

func (m *ListIssueTemplatesResponse) GetTemplates() []*IssueTemplate {
	if m != nil {
		return m.Templates
	}
	return nil
}

func (m *ListIssueTemplatesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*ListIssueTemplatesRequest)(nil), "monorail.v1.ListIssueTemplatesRequest")
	proto.RegisterType((*ListIssueTemplatesResponse)(nil), "monorail.v1.ListIssueTemplatesResponse")
}

func init() {
	proto.RegisterFile("api/v1/api_proto/projects.proto", fileDescriptor_c420c733c90f8ec5)
}

var fileDescriptor_c420c733c90f8ec5 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x6a, 0xdb, 0x40,
	0x14, 0xc6, 0x19, 0x99, 0x1a, 0x7b, 0x4c, 0x29, 0x0c, 0x94, 0xaa, 0x2a, 0xa6, 0x46, 0x0b, 0xd7,
	0x2d, 0xad, 0x06, 0xbb, 0x14, 0x4c, 0x77, 0xee, 0xae, 0xd0, 0x85, 0x51, 0xbd, 0x17, 0x23, 0xf5,
	0x55, 0x9d, 0x56, 0x9e, 0x99, 0xce, 0x8c, 0x4c, 0xf0, 0x22, 0x8b, 0x1c, 0x20, 0x9b, 0x1c, 0x20,
	0x17, 0xc8, 0x6d, 0x72, 0x81, 0x2c, 0x72, 0x8a, 0xac, 0x82, 0xfe, 0x11, 0x19, 0x47, 0x64, 0x39,
	0x9f, 0x7e, 0xdf, 0xfb, 0xbe, 0xf7, 0x84, 0xdf, 0x32, 0xc5, 0xe9, 0x6e, 0x4e, 0x99, 0xe2, 0x91,
	0xd2, 0xd2, 0x4a, 0xaa, 0xb4, 0xfc, 0x0b, 0x89, 0x35, 0x41, 0xf9, 0x24, 0xa3, 0xad, 0x14, 0x52,
	0x33, 0x9e, 0x05, 0xbb, 0xb9, 0xf7, 0x31, 0x95, 0x32, 0xcd, 0xa0, 0x26, 0xab, 0x47, 0x61, 0xa5,
	0xbf, 0x39, 0x64, 0xbf, 0xa2, 0x18, 0xfe, 0xb0, 0x1d, 0x97, 0xba, 0xb2, 0x7a, 0xd3, 0x2e, 0x5a,
	0x83, 0x91, 0xb9, 0x4e, 0xa0, 0xe6, 0xde, 0x77, 0x71, 0x4c, 0x08, 0x69, 0x99, 0xe5, 0x52, 0x98,
	0x66, 0x64, 0x57, 0xdd, 0x48, 0xc6, 0xad, 0xd6, 0xfe, 0x39, 0xc2, 0xaf, 0x7f, 0x70, 0x63, 0xbf,
	0x1b, 0x93, 0xc3, 0x06, 0xb6, 0x2a, 0x63, 0x16, 0x4c, 0x08, 0xff, 0x73, 0x30, 0x96, 0x7c, 0xc1,
	0x7d, 0xc5, 0x34, 0x08, 0xeb, 0xa2, 0x09, 0x9a, 0x0d, 0xbf, 0x8d, 0x6f, 0x56, 0xce, 0xdd, 0xea,
	0x15, 0x7e, 0xc9, 0x14, 0x0f, 0x12, 0x1d, 0xe7, 0x69, 0x90, 0xc8, 0x2d, 0x5d, 0x57, 0xa3, 0xc3,
	0x1a, 0x26, 0x6f, 0xf0, 0x50, 0xb1, 0x14, 0x22, 0xc3, 0xf7, 0xe0, 0x3a, 0x13, 0x34, 0x7b, 0x16,
	0x0e, 0x0a, 0xe1, 0x27, 0xdf, 0x03, 0x19, 0x63, 0x5c, 0x7e, 0xb4, 0xf2, 0x1f, 0x08, 0xb7, 0x57,
	0xcc, 0x0d, 0x4b, 0x7c, 0x53, 0x08, 0xfe, 0x29, 0xf6, 0x1e, 0xeb, 0x63, 0x94, 0x14, 0x06, 0xc8,
	0x12, 0x0f, 0x6d, 0x23, 0xba, 0x68, 0xd2, 0x9b, 0x8d, 0x16, 0x5e, 0xd0, 0x3a, 0x7c, 0x70, 0xe0,
	0x0b, 0x1f, 0x60, 0x32, 0xc5, 0x2f, 0x04, 0x9c, 0xd8, 0xa8, 0x95, 0xed, 0x94, 0xd9, 0xcf, 0x0b,
	0x79, 0xdd, 0xe4, 0x2f, 0xae, 0x10, 0x1e, 0xd4, 0xfb, 0x18, 0x72, 0x89, 0x30, 0x39, 0x6e, 0x43,
	0xa6, 0x07, 0x91, 0x9d, 0xe7, 0xf3, 0xde, 0x3d, 0xc9, 0x55, 0x6b, 0xf9, 0xcb, 0xb3, 0xeb, 0xdb,
	0x0b, 0x67, 0xe1, 0x7f, 0xa2, 0x4a, 0xab, 0x84, 0xb6, 0x5d, 0x4d, 0x17, 0x7a, 0x6c, 0xff, 0x8a,
	0x3e, 0xc4, 0xfd, 0xf2, 0x37, 0x7e, 0xbe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x60, 0xb4, 0x3f,
	0x9f, 0x02, 0x00, 0x00,
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
	// Returns all templates for specified project.
	//
	// Raises:
	//   NOT_FOUND if the requested parent project is not found.
	//   INVALID_ARGUMENT if the given `parent` is not valid.
	ListIssueTemplates(ctx context.Context, in *ListIssueTemplatesRequest, opts ...grpc.CallOption) (*ListIssueTemplatesResponse, error)
}
type projectsPRPCClient struct {
	client *prpc.Client
}

func NewProjectsPRPCClient(client *prpc.Client) ProjectsClient {
	return &projectsPRPCClient{client}
}

func (c *projectsPRPCClient) ListIssueTemplates(ctx context.Context, in *ListIssueTemplatesRequest, opts ...grpc.CallOption) (*ListIssueTemplatesResponse, error) {
	out := new(ListIssueTemplatesResponse)
	err := c.client.Call(ctx, "monorail.v1.Projects", "ListIssueTemplates", in, out, opts...)
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

func (c *projectsClient) ListIssueTemplates(ctx context.Context, in *ListIssueTemplatesRequest, opts ...grpc.CallOption) (*ListIssueTemplatesResponse, error) {
	out := new(ListIssueTemplatesResponse)
	err := c.cc.Invoke(ctx, "/monorail.v1.Projects/ListIssueTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectsServer is the server API for Projects service.
type ProjectsServer interface {
	// Returns all templates for specified project.
	//
	// Raises:
	//   NOT_FOUND if the requested parent project is not found.
	//   INVALID_ARGUMENT if the given `parent` is not valid.
	ListIssueTemplates(context.Context, *ListIssueTemplatesRequest) (*ListIssueTemplatesResponse, error)
}

// UnimplementedProjectsServer can be embedded to have forward compatible implementations.
type UnimplementedProjectsServer struct {
}

func (*UnimplementedProjectsServer) ListIssueTemplates(ctx context.Context, req *ListIssueTemplatesRequest) (*ListIssueTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIssueTemplates not implemented")
}

func RegisterProjectsServer(s prpc.Registrar, srv ProjectsServer) {
	s.RegisterService(&_Projects_serviceDesc, srv)
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
		FullMethod: "/monorail.v1.Projects/ListIssueTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).ListIssueTemplates(ctx, req.(*ListIssueTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Projects_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monorail.v1.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListIssueTemplates",
			Handler:    _Projects_ListIssueTemplates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api_proto/projects.proto",
}
