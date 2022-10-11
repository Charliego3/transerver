// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: accounts/rsa.proto

package acctspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RsaServiceClient is the client API for RsaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RsaServiceClient interface {
	PublicKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.BytesValue, error)
}

type rsaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRsaServiceClient(cc grpc.ClientConnInterface) RsaServiceClient {
	return &rsaServiceClient{cc}
}

func (c *rsaServiceClient) PublicKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*wrapperspb.BytesValue, error) {
	out := new(wrapperspb.BytesValue)
	err := c.cc.Invoke(ctx, "/org.github.transerver.accounts.RsaService/PublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RsaServiceServer is the server API for RsaService service.
// All implementations must embed UnimplementedRsaServiceServer
// for forward compatibility
type RsaServiceServer interface {
	PublicKey(context.Context, *emptypb.Empty) (*wrapperspb.BytesValue, error)
	mustEmbedUnimplementedRsaServiceServer()
}

// UnimplementedRsaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRsaServiceServer struct {
}

func (UnimplementedRsaServiceServer) PublicKey(context.Context, *emptypb.Empty) (*wrapperspb.BytesValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicKey not implemented")
}
func (UnimplementedRsaServiceServer) mustEmbedUnimplementedRsaServiceServer() {}

// UnsafeRsaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RsaServiceServer will
// result in compilation errors.
type UnsafeRsaServiceServer interface {
	mustEmbedUnimplementedRsaServiceServer()
}

func RegisterRsaServiceServer(s grpc.ServiceRegistrar, srv RsaServiceServer) {
	s.RegisterService(&RsaService_ServiceDesc, srv)
}

func _RsaService_PublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RsaServiceServer).PublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.github.transerver.accounts.RsaService/PublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RsaServiceServer).PublicKey(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RsaService_ServiceDesc is the grpc.ServiceDesc for RsaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RsaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.github.transerver.accounts.RsaService",
	HandlerType: (*RsaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublicKey",
			Handler:    _RsaService_PublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts/rsa.proto",
}
