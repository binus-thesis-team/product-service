// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pb/product_service/product_service.proto

package product_service

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

const (
	ProductService_FindAllProductsByIDs_FullMethodName = "/pb.product_service.ProductService/FindAllProductsByIDs"
	ProductService_FindByProductID_FullMethodName      = "/pb.product_service.ProductService/FindByProductID"
	ProductService_SearchAllProducts_FullMethodName    = "/pb.product_service.ProductService/SearchAllProducts"
	ProductService_UploadProducts_FullMethodName       = "/pb.product_service.ProductService/UploadProducts"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	FindAllProductsByIDs(ctx context.Context, in *FindByIDsRequest, opts ...grpc.CallOption) (*Products, error)
	FindByProductID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*Product, error)
	SearchAllProducts(ctx context.Context, in *ProductSearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	UploadProducts(ctx context.Context, in *UploadProductsRequest, opts ...grpc.CallOption) (*UploadProductsResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) FindAllProductsByIDs(ctx context.Context, in *FindByIDsRequest, opts ...grpc.CallOption) (*Products, error) {
	out := new(Products)
	err := c.cc.Invoke(ctx, ProductService_FindAllProductsByIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindByProductID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, ProductService_FindByProductID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SearchAllProducts(ctx context.Context, in *ProductSearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, ProductService_SearchAllProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UploadProducts(ctx context.Context, in *UploadProductsRequest, opts ...grpc.CallOption) (*UploadProductsResponse, error) {
	out := new(UploadProductsResponse)
	err := c.cc.Invoke(ctx, ProductService_UploadProducts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	FindAllProductsByIDs(context.Context, *FindByIDsRequest) (*Products, error)
	FindByProductID(context.Context, *FindByIDRequest) (*Product, error)
	SearchAllProducts(context.Context, *ProductSearchRequest) (*SearchResponse, error)
	UploadProducts(context.Context, *UploadProductsRequest) (*UploadProductsResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) FindAllProductsByIDs(context.Context, *FindByIDsRequest) (*Products, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllProductsByIDs not implemented")
}
func (UnimplementedProductServiceServer) FindByProductID(context.Context, *FindByIDRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByProductID not implemented")
}
func (UnimplementedProductServiceServer) SearchAllProducts(context.Context, *ProductSearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAllProducts not implemented")
}
func (UnimplementedProductServiceServer) UploadProducts(context.Context, *UploadProductsRequest) (*UploadProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadProducts not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_FindAllProductsByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindAllProductsByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindAllProductsByIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindAllProductsByIDs(ctx, req.(*FindByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindByProductID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindByProductID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindByProductID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindByProductID(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SearchAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SearchAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_SearchAllProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SearchAllProducts(ctx, req.(*ProductSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UploadProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UploadProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_UploadProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UploadProducts(ctx, req.(*UploadProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.product_service.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllProductsByIDs",
			Handler:    _ProductService_FindAllProductsByIDs_Handler,
		},
		{
			MethodName: "FindByProductID",
			Handler:    _ProductService_FindByProductID_Handler,
		},
		{
			MethodName: "SearchAllProducts",
			Handler:    _ProductService_SearchAllProducts_Handler,
		},
		{
			MethodName: "UploadProducts",
			Handler:    _ProductService_UploadProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/product_service/product_service.proto",
}
