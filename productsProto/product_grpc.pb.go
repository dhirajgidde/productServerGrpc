// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: product.proto

package productsProto

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

// ProductMessageReceiverClient is the client API for ProductMessageReceiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductMessageReceiverClient interface {
	EnableProducts(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProdResponse, error)
}

type productMessageReceiverClient struct {
	cc grpc.ClientConnInterface
}

func NewProductMessageReceiverClient(cc grpc.ClientConnInterface) ProductMessageReceiverClient {
	return &productMessageReceiverClient{cc}
}

func (c *productMessageReceiverClient) EnableProducts(ctx context.Context, in *Product, opts ...grpc.CallOption) (*ProdResponse, error) {
	out := new(ProdResponse)
	err := c.cc.Invoke(ctx, "/messager.ProductMessageReceiver/EnableProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductMessageReceiverServer is the server API for ProductMessageReceiver service.
// All implementations must embed UnimplementedProductMessageReceiverServer
// for forward compatibility
type ProductMessageReceiverServer interface {
	EnableProducts(context.Context, *Product) (*ProdResponse, error)
	mustEmbedUnimplementedProductMessageReceiverServer()
}

// UnimplementedProductMessageReceiverServer must be embedded to have forward compatible implementations.
type UnimplementedProductMessageReceiverServer struct {
}

func (UnimplementedProductMessageReceiverServer) EnableProducts(context.Context, *Product) (*ProdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableProducts not implemented")
}
func (UnimplementedProductMessageReceiverServer) mustEmbedUnimplementedProductMessageReceiverServer() {
}

// UnsafeProductMessageReceiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductMessageReceiverServer will
// result in compilation errors.
type UnsafeProductMessageReceiverServer interface {
	mustEmbedUnimplementedProductMessageReceiverServer()
}

func RegisterProductMessageReceiverServer(s grpc.ServiceRegistrar, srv ProductMessageReceiverServer) {
	s.RegisterService(&ProductMessageReceiver_ServiceDesc, srv)
}

func _ProductMessageReceiver_EnableProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductMessageReceiverServer).EnableProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messager.ProductMessageReceiver/EnableProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductMessageReceiverServer).EnableProducts(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductMessageReceiver_ServiceDesc is the grpc.ServiceDesc for ProductMessageReceiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductMessageReceiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messager.ProductMessageReceiver",
	HandlerType: (*ProductMessageReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnableProducts",
			Handler:    _ProductMessageReceiver_EnableProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
