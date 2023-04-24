// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: satlabrpc.proto

package satlabrpcserver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SatlabRpcService_ListBuildTargets_FullMethodName     = "/satlabrpcserver.SatlabRpcService/list_build_targets"
	SatlabRpcService_ListMilestones_FullMethodName       = "/satlabrpcserver.SatlabRpcService/list_milestones"
	SatlabRpcService_ListAccessibleModels_FullMethodName = "/satlabrpcserver.SatlabRpcService/list_accessible_models"
	SatlabRpcService_ListBuildVersions_FullMethodName    = "/satlabrpcserver.SatlabRpcService/list_build_versions"
	SatlabRpcService_StageBuild_FullMethodName           = "/satlabrpcserver.SatlabRpcService/stage_build"
)

// SatlabRpcServiceClient is the client API for SatlabRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SatlabRpcServiceClient interface {
	ListBuildTargets(ctx context.Context, in *ListBuildTargetsRequest, opts ...grpc.CallOption) (*ListBuildTargetsResponse, error)
	ListMilestones(ctx context.Context, in *ListMilestonesRequest, opts ...grpc.CallOption) (*ListMilestonesResponse, error)
	ListAccessibleModels(ctx context.Context, in *ListAccessibleModelsRequest, opts ...grpc.CallOption) (*ListAccessibleModelsResponse, error)
	ListBuildVersions(ctx context.Context, in *ListBuildVersionsRequest, opts ...grpc.CallOption) (*ListBuildVersionsResponse, error)
	StageBuild(ctx context.Context, in *StageBuildRequest, opts ...grpc.CallOption) (*StageBuildResponse, error)
}

type satlabRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSatlabRpcServiceClient(cc grpc.ClientConnInterface) SatlabRpcServiceClient {
	return &satlabRpcServiceClient{cc}
}

func (c *satlabRpcServiceClient) ListBuildTargets(ctx context.Context, in *ListBuildTargetsRequest, opts ...grpc.CallOption) (*ListBuildTargetsResponse, error) {
	out := new(ListBuildTargetsResponse)
	err := c.cc.Invoke(ctx, SatlabRpcService_ListBuildTargets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *satlabRpcServiceClient) ListMilestones(ctx context.Context, in *ListMilestonesRequest, opts ...grpc.CallOption) (*ListMilestonesResponse, error) {
	out := new(ListMilestonesResponse)
	err := c.cc.Invoke(ctx, SatlabRpcService_ListMilestones_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *satlabRpcServiceClient) ListAccessibleModels(ctx context.Context, in *ListAccessibleModelsRequest, opts ...grpc.CallOption) (*ListAccessibleModelsResponse, error) {
	out := new(ListAccessibleModelsResponse)
	err := c.cc.Invoke(ctx, SatlabRpcService_ListAccessibleModels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *satlabRpcServiceClient) ListBuildVersions(ctx context.Context, in *ListBuildVersionsRequest, opts ...grpc.CallOption) (*ListBuildVersionsResponse, error) {
	out := new(ListBuildVersionsResponse)
	err := c.cc.Invoke(ctx, SatlabRpcService_ListBuildVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *satlabRpcServiceClient) StageBuild(ctx context.Context, in *StageBuildRequest, opts ...grpc.CallOption) (*StageBuildResponse, error) {
	out := new(StageBuildResponse)
	err := c.cc.Invoke(ctx, SatlabRpcService_StageBuild_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SatlabRpcServiceServer is the server API for SatlabRpcService service.
// All implementations must embed UnimplementedSatlabRpcServiceServer
// for forward compatibility
type SatlabRpcServiceServer interface {
	ListBuildTargets(context.Context, *ListBuildTargetsRequest) (*ListBuildTargetsResponse, error)
	ListMilestones(context.Context, *ListMilestonesRequest) (*ListMilestonesResponse, error)
	ListAccessibleModels(context.Context, *ListAccessibleModelsRequest) (*ListAccessibleModelsResponse, error)
	ListBuildVersions(context.Context, *ListBuildVersionsRequest) (*ListBuildVersionsResponse, error)
	StageBuild(context.Context, *StageBuildRequest) (*StageBuildResponse, error)
	mustEmbedUnimplementedSatlabRpcServiceServer()
}

// UnimplementedSatlabRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSatlabRpcServiceServer struct {
}

func (UnimplementedSatlabRpcServiceServer) ListBuildTargets(context.Context, *ListBuildTargetsRequest) (*ListBuildTargetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuildTargets not implemented")
}
func (UnimplementedSatlabRpcServiceServer) ListMilestones(context.Context, *ListMilestonesRequest) (*ListMilestonesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMilestones not implemented")
}
func (UnimplementedSatlabRpcServiceServer) ListAccessibleModels(context.Context, *ListAccessibleModelsRequest) (*ListAccessibleModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessibleModels not implemented")
}
func (UnimplementedSatlabRpcServiceServer) ListBuildVersions(context.Context, *ListBuildVersionsRequest) (*ListBuildVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuildVersions not implemented")
}
func (UnimplementedSatlabRpcServiceServer) StageBuild(context.Context, *StageBuildRequest) (*StageBuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StageBuild not implemented")
}
func (UnimplementedSatlabRpcServiceServer) mustEmbedUnimplementedSatlabRpcServiceServer() {}

// UnsafeSatlabRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SatlabRpcServiceServer will
// result in compilation errors.
type UnsafeSatlabRpcServiceServer interface {
	mustEmbedUnimplementedSatlabRpcServiceServer()
}

func RegisterSatlabRpcServiceServer(s grpc.ServiceRegistrar, srv SatlabRpcServiceServer) {
	s.RegisterService(&SatlabRpcService_ServiceDesc, srv)
}

func _SatlabRpcService_ListBuildTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildTargetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SatlabRpcServiceServer).ListBuildTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SatlabRpcService_ListBuildTargets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SatlabRpcServiceServer).ListBuildTargets(ctx, req.(*ListBuildTargetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SatlabRpcService_ListMilestones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMilestonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SatlabRpcServiceServer).ListMilestones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SatlabRpcService_ListMilestones_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SatlabRpcServiceServer).ListMilestones(ctx, req.(*ListMilestonesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SatlabRpcService_ListAccessibleModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccessibleModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SatlabRpcServiceServer).ListAccessibleModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SatlabRpcService_ListAccessibleModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SatlabRpcServiceServer).ListAccessibleModels(ctx, req.(*ListAccessibleModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SatlabRpcService_ListBuildVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBuildVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SatlabRpcServiceServer).ListBuildVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SatlabRpcService_ListBuildVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SatlabRpcServiceServer).ListBuildVersions(ctx, req.(*ListBuildVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SatlabRpcService_StageBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StageBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SatlabRpcServiceServer).StageBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SatlabRpcService_StageBuild_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SatlabRpcServiceServer).StageBuild(ctx, req.(*StageBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SatlabRpcService_ServiceDesc is the grpc.ServiceDesc for SatlabRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SatlabRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "satlabrpcserver.SatlabRpcService",
	HandlerType: (*SatlabRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "list_build_targets",
			Handler:    _SatlabRpcService_ListBuildTargets_Handler,
		},
		{
			MethodName: "list_milestones",
			Handler:    _SatlabRpcService_ListMilestones_Handler,
		},
		{
			MethodName: "list_accessible_models",
			Handler:    _SatlabRpcService_ListAccessibleModels_Handler,
		},
		{
			MethodName: "list_build_versions",
			Handler:    _SatlabRpcService_ListBuildVersions_Handler,
		},
		{
			MethodName: "stage_build",
			Handler:    _SatlabRpcService_StageBuild_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "satlabrpc.proto",
}
