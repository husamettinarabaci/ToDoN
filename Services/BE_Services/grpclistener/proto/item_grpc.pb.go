// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: item.proto

package __

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

// SvcItemClient is the client API for SvcItem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SvcItemClient interface {
	RpcItem(ctx context.Context, in *PbItem, opts ...grpc.CallOption) (*PbResp, error)
	RpcItems(ctx context.Context, in *PbResp, opts ...grpc.CallOption) (*PbItems, error)
}

type svcItemClient struct {
	cc grpc.ClientConnInterface
}

func NewSvcItemClient(cc grpc.ClientConnInterface) SvcItemClient {
	return &svcItemClient{cc}
}

func (c *svcItemClient) RpcItem(ctx context.Context, in *PbItem, opts ...grpc.CallOption) (*PbResp, error) {
	out := new(PbResp)
	err := c.cc.Invoke(ctx, "/PbItem.SvcItem/RpcItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcItemClient) RpcItems(ctx context.Context, in *PbResp, opts ...grpc.CallOption) (*PbItems, error) {
	out := new(PbItems)
	err := c.cc.Invoke(ctx, "/PbItem.SvcItem/RpcItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SvcItemServer is the server API for SvcItem service.
// All implementations must embed UnimplementedSvcItemServer
// for forward compatibility
type SvcItemServer interface {
	RpcItem(context.Context, *PbItem) (*PbResp, error)
	RpcItems(context.Context, *PbResp) (*PbItems, error)
	mustEmbedUnimplementedSvcItemServer()
}

// UnimplementedSvcItemServer must be embedded to have forward compatible implementations.
type UnimplementedSvcItemServer struct {
}

func (UnimplementedSvcItemServer) RpcItem(context.Context, *PbItem) (*PbResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcItem not implemented")
}
func (UnimplementedSvcItemServer) RpcItems(context.Context, *PbResp) (*PbItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcItems not implemented")
}
func (UnimplementedSvcItemServer) mustEmbedUnimplementedSvcItemServer() {}

// UnsafeSvcItemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SvcItemServer will
// result in compilation errors.
type UnsafeSvcItemServer interface {
	mustEmbedUnimplementedSvcItemServer()
}

func RegisterSvcItemServer(s grpc.ServiceRegistrar, srv SvcItemServer) {
	s.RegisterService(&SvcItem_ServiceDesc, srv)
}

func _SvcItem_RpcItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PbItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcItemServer).RpcItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PbItem.SvcItem/RpcItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcItemServer).RpcItem(ctx, req.(*PbItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _SvcItem_RpcItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PbResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SvcItemServer).RpcItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PbItem.SvcItem/RpcItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SvcItemServer).RpcItems(ctx, req.(*PbResp))
	}
	return interceptor(ctx, in, info, handler)
}

// SvcItem_ServiceDesc is the grpc.ServiceDesc for SvcItem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SvcItem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PbItem.SvcItem",
	HandlerType: (*SvcItemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcItem",
			Handler:    _SvcItem_RpcItem_Handler,
		},
		{
			MethodName: "RpcItems",
			Handler:    _SvcItem_RpcItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "item.proto",
}
