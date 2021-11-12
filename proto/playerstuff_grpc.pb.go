// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package webscraper

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

// WebscraperServiceClient is the client API for WebscraperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebscraperServiceClient interface {
	GetPlayers(ctx context.Context, in *PlayerRequest, opts ...grpc.CallOption) (*PlayerArray, error)
}

type webscraperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebscraperServiceClient(cc grpc.ClientConnInterface) WebscraperServiceClient {
	return &webscraperServiceClient{cc}
}

func (c *webscraperServiceClient) GetPlayers(ctx context.Context, in *PlayerRequest, opts ...grpc.CallOption) (*PlayerArray, error) {
	out := new(PlayerArray)
	err := c.cc.Invoke(ctx, "/webscraper.webscraperService/GetPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebscraperServiceServer is the server API for WebscraperService service.
// All implementations must embed UnimplementedWebscraperServiceServer
// for forward compatibility
type WebscraperServiceServer interface {
	GetPlayers(context.Context, *PlayerRequest) (*PlayerArray, error)
	mustEmbedUnimplementedWebscraperServiceServer()
}

// UnimplementedWebscraperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWebscraperServiceServer struct {
}

func (UnimplementedWebscraperServiceServer) GetPlayers(context.Context, *PlayerRequest) (*PlayerArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayers not implemented")
}
func (UnimplementedWebscraperServiceServer) mustEmbedUnimplementedWebscraperServiceServer() {}

// UnsafeWebscraperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebscraperServiceServer will
// result in compilation errors.
type UnsafeWebscraperServiceServer interface {
	mustEmbedUnimplementedWebscraperServiceServer()
}

func RegisterWebscraperServiceServer(s grpc.ServiceRegistrar, srv WebscraperServiceServer) {
	s.RegisterService(&WebscraperService_ServiceDesc, srv)
}

func _WebscraperService_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebscraperServiceServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webscraper.webscraperService/GetPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebscraperServiceServer).GetPlayers(ctx, req.(*PlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebscraperService_ServiceDesc is the grpc.ServiceDesc for WebscraperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebscraperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "webscraper.webscraperService",
	HandlerType: (*WebscraperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayers",
			Handler:    _WebscraperService_GetPlayers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/playerstuff.proto",
}
