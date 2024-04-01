// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package m7client

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

// M7Client is the client API for M7 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type M7Client interface {
	// 用户登录
	Test(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestResp, error)
}

type m7Client struct {
	cc grpc.ClientConnInterface
}

func NewM7Client(cc grpc.ClientConnInterface) M7Client {
	return &m7Client{cc}
}

func (c *m7Client) Test(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestResp, error) {
	out := new(TestResp)
	err := c.cc.Invoke(ctx, "/m7client.M7/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// M7Server is the server API for M7 service.
// All implementations must embed UnimplementedM7Server
// for forward compatibility
type M7Server interface {
	// 用户登录
	Test(context.Context, *TestReq) (*TestResp, error)
	mustEmbedUnimplementedM7Server()
}

// UnimplementedM7Server must be embedded to have forward compatible implementations.
type UnimplementedM7Server struct {
}

func (UnimplementedM7Server) Test(context.Context, *TestReq) (*TestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedM7Server) mustEmbedUnimplementedM7Server() {}

// UnsafeM7Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to M7Server will
// result in compilation errors.
type UnsafeM7Server interface {
	mustEmbedUnimplementedM7Server()
}

func RegisterM7Server(s grpc.ServiceRegistrar, srv M7Server) {
	s.RegisterService(&M7_ServiceDesc, srv)
}

func _M7_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(M7Server).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m7client.M7/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(M7Server).Test(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

// M7_ServiceDesc is the grpc.ServiceDesc for M7 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var M7_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "m7client.M7",
	HandlerType: (*M7Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _M7_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "m7.proto",
}