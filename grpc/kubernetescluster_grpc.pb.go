// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: kubernetescluster.proto

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

// KubernetesClusterServiceClient is the client API for KubernetesClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubernetesClusterServiceClient interface {
	CreateKubernetesCluster(ctx context.Context, in *KubernetesCluster, opts ...grpc.CallOption) (*KubernetesCluster, error)
	GetKubernetesCluster(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*KubernetesCluster, error)
	ListKubernetesCluster(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*KubernetesClusterList, error)
	UpdateKubernetesCluster(ctx context.Context, in *KubernetesCluster, opts ...grpc.CallOption) (*KubernetesCluster, error)
	DeleteKubernetesCluster(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*KubernetesCluster, error)
}

type kubernetesClusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKubernetesClusterServiceClient(cc grpc.ClientConnInterface) KubernetesClusterServiceClient {
	return &kubernetesClusterServiceClient{cc}
}

func (c *kubernetesClusterServiceClient) CreateKubernetesCluster(ctx context.Context, in *KubernetesCluster, opts ...grpc.CallOption) (*KubernetesCluster, error) {
	out := new(KubernetesCluster)
	err := c.cc.Invoke(ctx, "/kubernetescluster.v1alpha1.KubernetesClusterService/CreateKubernetesCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterServiceClient) GetKubernetesCluster(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*KubernetesCluster, error) {
	out := new(KubernetesCluster)
	err := c.cc.Invoke(ctx, "/kubernetescluster.v1alpha1.KubernetesClusterService/GetKubernetesCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterServiceClient) ListKubernetesCluster(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*KubernetesClusterList, error) {
	out := new(KubernetesClusterList)
	err := c.cc.Invoke(ctx, "/kubernetescluster.v1alpha1.KubernetesClusterService/ListKubernetesCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterServiceClient) UpdateKubernetesCluster(ctx context.Context, in *KubernetesCluster, opts ...grpc.CallOption) (*KubernetesCluster, error) {
	out := new(KubernetesCluster)
	err := c.cc.Invoke(ctx, "/kubernetescluster.v1alpha1.KubernetesClusterService/UpdateKubernetesCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClusterServiceClient) DeleteKubernetesCluster(ctx context.Context, in *FilterOptions, opts ...grpc.CallOption) (*KubernetesCluster, error) {
	out := new(KubernetesCluster)
	err := c.cc.Invoke(ctx, "/kubernetescluster.v1alpha1.KubernetesClusterService/DeleteKubernetesCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesClusterServiceServer is the server API for KubernetesClusterService service.
// All implementations must embed UnimplementedKubernetesClusterServiceServer
// for forward compatibility
type KubernetesClusterServiceServer interface {
	CreateKubernetesCluster(context.Context, *KubernetesCluster) (*KubernetesCluster, error)
	GetKubernetesCluster(context.Context, *FilterOptions) (*KubernetesCluster, error)
	ListKubernetesCluster(context.Context, *FilterOptions) (*KubernetesClusterList, error)
	UpdateKubernetesCluster(context.Context, *KubernetesCluster) (*KubernetesCluster, error)
	DeleteKubernetesCluster(context.Context, *FilterOptions) (*KubernetesCluster, error)
	mustEmbedUnimplementedKubernetesClusterServiceServer()
}

// UnimplementedKubernetesClusterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKubernetesClusterServiceServer struct {
}

func (UnimplementedKubernetesClusterServiceServer) CreateKubernetesCluster(context.Context, *KubernetesCluster) (*KubernetesCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKubernetesCluster not implemented")
}
func (UnimplementedKubernetesClusterServiceServer) GetKubernetesCluster(context.Context, *FilterOptions) (*KubernetesCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKubernetesCluster not implemented")
}
func (UnimplementedKubernetesClusterServiceServer) ListKubernetesCluster(context.Context, *FilterOptions) (*KubernetesClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKubernetesCluster not implemented")
}
func (UnimplementedKubernetesClusterServiceServer) UpdateKubernetesCluster(context.Context, *KubernetesCluster) (*KubernetesCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKubernetesCluster not implemented")
}
func (UnimplementedKubernetesClusterServiceServer) DeleteKubernetesCluster(context.Context, *FilterOptions) (*KubernetesCluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKubernetesCluster not implemented")
}
func (UnimplementedKubernetesClusterServiceServer) mustEmbedUnimplementedKubernetesClusterServiceServer() {
}

// UnsafeKubernetesClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubernetesClusterServiceServer will
// result in compilation errors.
type UnsafeKubernetesClusterServiceServer interface {
	mustEmbedUnimplementedKubernetesClusterServiceServer()
}

func RegisterKubernetesClusterServiceServer(s grpc.ServiceRegistrar, srv KubernetesClusterServiceServer) {
	s.RegisterService(&KubernetesClusterService_ServiceDesc, srv)
}

func _KubernetesClusterService_CreateKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubernetesCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterServiceServer).CreateKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetescluster.v1alpha1.KubernetesClusterService/CreateKubernetesCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterServiceServer).CreateKubernetesCluster(ctx, req.(*KubernetesCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterService_GetKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterServiceServer).GetKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetescluster.v1alpha1.KubernetesClusterService/GetKubernetesCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterServiceServer).GetKubernetesCluster(ctx, req.(*FilterOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterService_ListKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterServiceServer).ListKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetescluster.v1alpha1.KubernetesClusterService/ListKubernetesCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterServiceServer).ListKubernetesCluster(ctx, req.(*FilterOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterService_UpdateKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubernetesCluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterServiceServer).UpdateKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetescluster.v1alpha1.KubernetesClusterService/UpdateKubernetesCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterServiceServer).UpdateKubernetesCluster(ctx, req.(*KubernetesCluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubernetesClusterService_DeleteKubernetesCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesClusterServiceServer).DeleteKubernetesCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubernetescluster.v1alpha1.KubernetesClusterService/DeleteKubernetesCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesClusterServiceServer).DeleteKubernetesCluster(ctx, req.(*FilterOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// KubernetesClusterService_ServiceDesc is the grpc.ServiceDesc for KubernetesClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubernetesClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubernetescluster.v1alpha1.KubernetesClusterService",
	HandlerType: (*KubernetesClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKubernetesCluster",
			Handler:    _KubernetesClusterService_CreateKubernetesCluster_Handler,
		},
		{
			MethodName: "GetKubernetesCluster",
			Handler:    _KubernetesClusterService_GetKubernetesCluster_Handler,
		},
		{
			MethodName: "ListKubernetesCluster",
			Handler:    _KubernetesClusterService_ListKubernetesCluster_Handler,
		},
		{
			MethodName: "UpdateKubernetesCluster",
			Handler:    _KubernetesClusterService_UpdateKubernetesCluster_Handler,
		},
		{
			MethodName: "DeleteKubernetesCluster",
			Handler:    _KubernetesClusterService_DeleteKubernetesCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubernetescluster.proto",
}