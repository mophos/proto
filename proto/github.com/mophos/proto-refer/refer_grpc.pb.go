// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto_refer

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

// ReferServiceClient is the client API for ReferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReferServiceClient interface {
	ReferOut(ctx context.Context, in *ReferOutRequest, opts ...grpc.CallOption) (*ReferOutResponse, error)
}

type referServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReferServiceClient(cc grpc.ClientConnInterface) ReferServiceClient {
	return &referServiceClient{cc}
}

func (c *referServiceClient) ReferOut(ctx context.Context, in *ReferOutRequest, opts ...grpc.CallOption) (*ReferOutResponse, error) {
	out := new(ReferOutResponse)
	err := c.cc.Invoke(ctx, "/proto.ReferService/ReferOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReferServiceServer is the server API for ReferService service.
// All implementations must embed UnimplementedReferServiceServer
// for forward compatibility
type ReferServiceServer interface {
	ReferOut(context.Context, *ReferOutRequest) (*ReferOutResponse, error)
	mustEmbedUnimplementedReferServiceServer()
}

// UnimplementedReferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReferServiceServer struct {
}

func (UnimplementedReferServiceServer) ReferOut(context.Context, *ReferOutRequest) (*ReferOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReferOut not implemented")
}
func (UnimplementedReferServiceServer) mustEmbedUnimplementedReferServiceServer() {}

// UnsafeReferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReferServiceServer will
// result in compilation errors.
type UnsafeReferServiceServer interface {
	mustEmbedUnimplementedReferServiceServer()
}

func RegisterReferServiceServer(s grpc.ServiceRegistrar, srv ReferServiceServer) {
	s.RegisterService(&ReferService_ServiceDesc, srv)
}

func _ReferService_ReferOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReferOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferServiceServer).ReferOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ReferService/ReferOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferServiceServer).ReferOut(ctx, req.(*ReferOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReferService_ServiceDesc is the grpc.ServiceDesc for ReferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ReferService",
	HandlerType: (*ReferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReferOut",
			Handler:    _ReferService_ReferOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "refer.proto",
}