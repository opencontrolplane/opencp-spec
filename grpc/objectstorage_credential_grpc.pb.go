// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: objectstorage_credential.proto

package grpc

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

// ObjectStorageCredentialServiceClient is the client API for ObjectStorageCredentialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectStorageCredentialServiceClient interface {
	CreateObjectStorageCredential(ctx context.Context, in *ObjectStorageCredential, opts ...grpc.CallOption) (*ObjectStorageCredential, error)
	// rpc UpdateObjectStorageCredential(ObjectStorageCredential) returns (ObjectStorageCredential) {}
	DeleteObjectStorageCredential(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*ObjectStorageCredential, error)
	GetObjectStorageCredential(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*ObjectStorageCredential, error)
	ListObjectStorageCredential(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*ObjectStorageCredentialList, error)
}

type objectStorageCredentialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectStorageCredentialServiceClient(cc grpc.ClientConnInterface) ObjectStorageCredentialServiceClient {
	return &objectStorageCredentialServiceClient{cc}
}

func (c *objectStorageCredentialServiceClient) CreateObjectStorageCredential(ctx context.Context, in *ObjectStorageCredential, opts ...grpc.CallOption) (*ObjectStorageCredential, error) {
	out := new(ObjectStorageCredential)
	err := c.cc.Invoke(ctx, "/objectstorageCredential.v1alpha1.ObjectStorageCredentialService/CreateObjectStorageCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStorageCredentialServiceClient) DeleteObjectStorageCredential(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*ObjectStorageCredential, error) {
	out := new(ObjectStorageCredential)
	err := c.cc.Invoke(ctx, "/objectstorageCredential.v1alpha1.ObjectStorageCredentialService/DeleteObjectStorageCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStorageCredentialServiceClient) GetObjectStorageCredential(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*ObjectStorageCredential, error) {
	out := new(ObjectStorageCredential)
	err := c.cc.Invoke(ctx, "/objectstorageCredential.v1alpha1.ObjectStorageCredentialService/GetObjectStorageCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStorageCredentialServiceClient) ListObjectStorageCredential(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*ObjectStorageCredentialList, error) {
	out := new(ObjectStorageCredentialList)
	err := c.cc.Invoke(ctx, "/objectstorageCredential.v1alpha1.ObjectStorageCredentialService/ListObjectStorageCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectStorageCredentialServiceServer is the server API for ObjectStorageCredentialService service.
// All implementations must embed UnimplementedObjectStorageCredentialServiceServer
// for forward compatibility
type ObjectStorageCredentialServiceServer interface {
	CreateObjectStorageCredential(context.Context, *ObjectStorageCredential) (*ObjectStorageCredential, error)
	// rpc UpdateObjectStorageCredential(ObjectStorageCredential) returns (ObjectStorageCredential) {}
	DeleteObjectStorageCredential(context.Context, *FilterOptions) (*ObjectStorageCredential, error)
	GetObjectStorageCredential(context.Context, *FilterOptions) (*ObjectStorageCredential, error)
	ListObjectStorageCredential(context.Context, *FilterOptions) (*ObjectStorageCredentialList, error)
	mustEmbedUnimplementedObjectStorageCredentialServiceServer()
}

// UnimplementedObjectStorageCredentialServiceServer must be embedded to have forward compatible implementations.
type UnimplementedObjectStorageCredentialServiceServer struct {
}

func (UnimplementedObjectStorageCredentialServiceServer) CreateObjectStorageCredential(context.Context, *ObjectStorageCredential) (*ObjectStorageCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectStorageCredential not implemented")
}
func (UnimplementedObjectStorageCredentialServiceServer) DeleteObjectStorageCredential(context.Context, *FilterOptions) (*ObjectStorageCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjectStorageCredential not implemented")
}
func (UnimplementedObjectStorageCredentialServiceServer) GetObjectStorageCredential(context.Context, *FilterOptions) (*ObjectStorageCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectStorageCredential not implemented")
}
func (UnimplementedObjectStorageCredentialServiceServer) ListObjectStorageCredential(context.Context, *FilterOptions) (*ObjectStorageCredentialList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjectStorageCredential not implemented")
}
func (UnimplementedObjectStorageCredentialServiceServer) mustEmbedUnimplementedObjectStorageCredentialServiceServer() {
}

// UnsafeObjectStorageCredentialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectStorageCredentialServiceServer will
// result in compilation errors.
type UnsafeObjectStorageCredentialServiceServer interface {
	mustEmbedUnimplementedObjectStorageCredentialServiceServer()
}

func RegisterObjectStorageCredentialServiceServer(s grpc.ServiceRegistrar, srv ObjectStorageCredentialServiceServer) {
	s.RegisterService(&ObjectStorageCredentialService_ServiceDesc, srv)
}

func _ObjectStorageCredentialService_CreateObjectStorageCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectStorageCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStorageCredentialServiceServer).CreateObjectStorageCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/objectstorageCredential.v1alpha1.ObjectStorageCredentialService/CreateObjectStorageCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStorageCredentialServiceServer).CreateObjectStorageCredential(ctx, req.(*ObjectStorageCredential))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStorageCredentialService_DeleteObjectStorageCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStorageCredentialServiceServer).DeleteObjectStorageCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/objectstorageCredential.v1alpha1.ObjectStorageCredentialService/DeleteObjectStorageCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStorageCredentialServiceServer).DeleteObjectStorageCredential(ctx, req.(*FilterOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStorageCredentialService_GetObjectStorageCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStorageCredentialServiceServer).GetObjectStorageCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/objectstorageCredential.v1alpha1.ObjectStorageCredentialService/GetObjectStorageCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStorageCredentialServiceServer).GetObjectStorageCredential(ctx, req.(*FilterOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStorageCredentialService_ListObjectStorageCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStorageCredentialServiceServer).ListObjectStorageCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/objectstorageCredential.v1alpha1.ObjectStorageCredentialService/ListObjectStorageCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStorageCredentialServiceServer).ListObjectStorageCredential(ctx, req.(*FilterOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectStorageCredentialService_ServiceDesc is the grpc.ServiceDesc for ObjectStorageCredentialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectStorageCredentialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "objectstorageCredential.v1alpha1.ObjectStorageCredentialService",
	HandlerType: (*ObjectStorageCredentialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObjectStorageCredential",
			Handler:    _ObjectStorageCredentialService_CreateObjectStorageCredential_Handler,
		},
		{
			MethodName: "DeleteObjectStorageCredential",
			Handler:    _ObjectStorageCredentialService_DeleteObjectStorageCredential_Handler,
		},
		{
			MethodName: "GetObjectStorageCredential",
			Handler:    _ObjectStorageCredentialService_GetObjectStorageCredential_Handler,
		},
		{
			MethodName: "ListObjectStorageCredential",
			Handler:    _ObjectStorageCredentialService_ListObjectStorageCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "objectstorage_credential.proto",
}
