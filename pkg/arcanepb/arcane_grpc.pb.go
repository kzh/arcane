// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package arcanepb

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

// ArcaneClient is the client API for Arcane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArcaneClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	KVGet(ctx context.Context, in *KVGetRequest, opts ...grpc.CallOption) (*KVGetResponse, error)
	KVPut(ctx context.Context, in *KVPutRequest, opts ...grpc.CallOption) (*KVPutResponse, error)
}

type arcaneClient struct {
	cc grpc.ClientConnInterface
}

func NewArcaneClient(cc grpc.ClientConnInterface) ArcaneClient {
	return &arcaneClient{cc}
}

func (c *arcaneClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/arcane.Arcane/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arcaneClient) KVGet(ctx context.Context, in *KVGetRequest, opts ...grpc.CallOption) (*KVGetResponse, error) {
	out := new(KVGetResponse)
	err := c.cc.Invoke(ctx, "/arcane.Arcane/KVGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arcaneClient) KVPut(ctx context.Context, in *KVPutRequest, opts ...grpc.CallOption) (*KVPutResponse, error) {
	out := new(KVPutResponse)
	err := c.cc.Invoke(ctx, "/arcane.Arcane/KVPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArcaneServer is the server API for Arcane service.
// All implementations must embed UnimplementedArcaneServer
// for forward compatibility
type ArcaneServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	KVGet(context.Context, *KVGetRequest) (*KVGetResponse, error)
	KVPut(context.Context, *KVPutRequest) (*KVPutResponse, error)
	mustEmbedUnimplementedArcaneServer()
}

// UnimplementedArcaneServer must be embedded to have forward compatible implementations.
type UnimplementedArcaneServer struct {
}

func (UnimplementedArcaneServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedArcaneServer) KVGet(context.Context, *KVGetRequest) (*KVGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KVGet not implemented")
}
func (UnimplementedArcaneServer) KVPut(context.Context, *KVPutRequest) (*KVPutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KVPut not implemented")
}
func (UnimplementedArcaneServer) mustEmbedUnimplementedArcaneServer() {}

// UnsafeArcaneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArcaneServer will
// result in compilation errors.
type UnsafeArcaneServer interface {
	mustEmbedUnimplementedArcaneServer()
}

func RegisterArcaneServer(s grpc.ServiceRegistrar, srv ArcaneServer) {
	s.RegisterService(&Arcane_ServiceDesc, srv)
}

func _Arcane_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArcaneServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arcane.Arcane/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArcaneServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arcane_KVGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArcaneServer).KVGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arcane.Arcane/KVGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArcaneServer).KVGet(ctx, req.(*KVGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arcane_KVPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArcaneServer).KVPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arcane.Arcane/KVPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArcaneServer).KVPut(ctx, req.(*KVPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Arcane_ServiceDesc is the grpc.ServiceDesc for Arcane service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Arcane_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arcane.Arcane",
	HandlerType: (*ArcaneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Arcane_Status_Handler,
		},
		{
			MethodName: "KVGet",
			Handler:    _Arcane_KVGet_Handler,
		},
		{
			MethodName: "KVPut",
			Handler:    _Arcane_KVPut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/arcanepb/arcane.proto",
}
