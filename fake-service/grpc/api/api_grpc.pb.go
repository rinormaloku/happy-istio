// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// FakeServiceClient is the client API for FakeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FakeServiceClient interface {
	Handle(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type fakeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFakeServiceClient(cc grpc.ClientConnInterface) FakeServiceClient {
	return &fakeServiceClient{cc}
}

func (c *fakeServiceClient) Handle(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/FakeService/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FakeServiceServer is the server API for FakeService service.
// All implementations must embed UnimplementedFakeServiceServer
// for forward compatibility
type FakeServiceServer interface {
	Handle(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedFakeServiceServer()
}

// UnimplementedFakeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFakeServiceServer struct {
}

func (UnimplementedFakeServiceServer) Handle(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handle not implemented")
}
func (UnimplementedFakeServiceServer) mustEmbedUnimplementedFakeServiceServer() {}

// UnsafeFakeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FakeServiceServer will
// result in compilation errors.
type UnsafeFakeServiceServer interface {
	mustEmbedUnimplementedFakeServiceServer()
}

func RegisterFakeServiceServer(s grpc.ServiceRegistrar, srv FakeServiceServer) {
	s.RegisterService(&FakeService_ServiceDesc, srv)
}

func _FakeService_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakeServiceServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FakeService/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakeServiceServer).Handle(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// FakeService_ServiceDesc is the grpc.ServiceDesc for FakeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FakeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FakeService",
	HandlerType: (*FakeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _FakeService_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
