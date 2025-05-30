// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: works.proto

package works

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Works_CreatedWorks_FullMethodName = "/stream.Works/CreatedWorks"
)

// WorksClient is the client API for Works service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorksClient interface {
	CreatedWorks(ctx context.Context, in *CreatedWorksReq, opts ...grpc.CallOption) (*CreatedWorksResp, error)
}

type worksClient struct {
	cc grpc.ClientConnInterface
}

func NewWorksClient(cc grpc.ClientConnInterface) WorksClient {
	return &worksClient{cc}
}

func (c *worksClient) CreatedWorks(ctx context.Context, in *CreatedWorksReq, opts ...grpc.CallOption) (*CreatedWorksResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatedWorksResp)
	err := c.cc.Invoke(ctx, Works_CreatedWorks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorksServer is the server API for Works service.
// All implementations must embed UnimplementedWorksServer
// for forward compatibility
type WorksServer interface {
	CreatedWorks(context.Context, *CreatedWorksReq) (*CreatedWorksResp, error)
	mustEmbedUnimplementedWorksServer()
}

// UnimplementedWorksServer must be embedded to have forward compatible implementations.
type UnimplementedWorksServer struct {
}

func (UnimplementedWorksServer) CreatedWorks(context.Context, *CreatedWorksReq) (*CreatedWorksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedWorks not implemented")
}
func (UnimplementedWorksServer) mustEmbedUnimplementedWorksServer() {}

// UnsafeWorksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorksServer will
// result in compilation errors.
type UnsafeWorksServer interface {
	mustEmbedUnimplementedWorksServer()
}

func RegisterWorksServer(s grpc.ServiceRegistrar, srv WorksServer) {
	s.RegisterService(&Works_ServiceDesc, srv)
}

func _Works_CreatedWorks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatedWorksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorksServer).CreatedWorks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Works_CreatedWorks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorksServer).CreatedWorks(ctx, req.(*CreatedWorksReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Works_ServiceDesc is the grpc.ServiceDesc for Works service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Works_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.Works",
	HandlerType: (*WorksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatedWorks",
			Handler:    _Works_CreatedWorks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "works.proto",
}
