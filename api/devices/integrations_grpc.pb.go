// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/devices/integrations.proto

package devices

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IntegrationServiceClient is the client API for IntegrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntegrationServiceClient interface {
	ListIntegrations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListIntegrationsResponse, error)
}

type integrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIntegrationServiceClient(cc grpc.ClientConnInterface) IntegrationServiceClient {
	return &integrationServiceClient{cc}
}

func (c *integrationServiceClient) ListIntegrations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListIntegrationsResponse, error) {
	out := new(ListIntegrationsResponse)
	err := c.cc.Invoke(ctx, "/devices.IntegrationService/ListIntegrations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegrationServiceServer is the server API for IntegrationService service.
// All implementations must embed UnimplementedIntegrationServiceServer
// for forward compatibility
type IntegrationServiceServer interface {
	ListIntegrations(context.Context, *emptypb.Empty) (*ListIntegrationsResponse, error)
	mustEmbedUnimplementedIntegrationServiceServer()
}

// UnimplementedIntegrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIntegrationServiceServer struct {
}

func (UnimplementedIntegrationServiceServer) ListIntegrations(context.Context, *emptypb.Empty) (*ListIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIntegrations not implemented")
}
func (UnimplementedIntegrationServiceServer) mustEmbedUnimplementedIntegrationServiceServer() {}

// UnsafeIntegrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntegrationServiceServer will
// result in compilation errors.
type UnsafeIntegrationServiceServer interface {
	mustEmbedUnimplementedIntegrationServiceServer()
}

func RegisterIntegrationServiceServer(s grpc.ServiceRegistrar, srv IntegrationServiceServer) {
	s.RegisterService(&IntegrationService_ServiceDesc, srv)
}

func _IntegrationService_ListIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).ListIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devices.IntegrationService/ListIntegrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).ListIntegrations(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// IntegrationService_ServiceDesc is the grpc.ServiceDesc for IntegrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntegrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "devices.IntegrationService",
	HandlerType: (*IntegrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListIntegrations",
			Handler:    _IntegrationService_ListIntegrations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/devices/integrations.proto",
}
