// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: opspillar/v1/clusters.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Clusters_CreateClusters_FullMethodName = "/api.opspillar.v1.Clusters/CreateClusters"
	Clusters_UpdateClusters_FullMethodName = "/api.opspillar.v1.Clusters/UpdateClusters"
	Clusters_DeleteClusters_FullMethodName = "/api.opspillar.v1.Clusters/DeleteClusters"
	Clusters_GetClusters_FullMethodName    = "/api.opspillar.v1.Clusters/GetClusters"
	Clusters_ListClusters_FullMethodName   = "/api.opspillar.v1.Clusters/ListClusters"
)

// ClustersClient is the client API for Clusters service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClustersClient interface {
	CreateClusters(ctx context.Context, in *CreateClustersRequest, opts ...grpc.CallOption) (*CreateClustersReply, error)
	UpdateClusters(ctx context.Context, in *UpdateClustersRequest, opts ...grpc.CallOption) (*UpdateClustersReply, error)
	DeleteClusters(ctx context.Context, in *DeleteClustersRequest, opts ...grpc.CallOption) (*DeleteClustersReply, error)
	GetClusters(ctx context.Context, in *GetClustersRequest, opts ...grpc.CallOption) (*GetClustersReply, error)
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersReply, error)
}

type clustersClient struct {
	cc grpc.ClientConnInterface
}

func NewClustersClient(cc grpc.ClientConnInterface) ClustersClient {
	return &clustersClient{cc}
}

func (c *clustersClient) CreateClusters(ctx context.Context, in *CreateClustersRequest, opts ...grpc.CallOption) (*CreateClustersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateClustersReply)
	err := c.cc.Invoke(ctx, Clusters_CreateClusters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) UpdateClusters(ctx context.Context, in *UpdateClustersRequest, opts ...grpc.CallOption) (*UpdateClustersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateClustersReply)
	err := c.cc.Invoke(ctx, Clusters_UpdateClusters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) DeleteClusters(ctx context.Context, in *DeleteClustersRequest, opts ...grpc.CallOption) (*DeleteClustersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteClustersReply)
	err := c.cc.Invoke(ctx, Clusters_DeleteClusters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) GetClusters(ctx context.Context, in *GetClustersRequest, opts ...grpc.CallOption) (*GetClustersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetClustersReply)
	err := c.cc.Invoke(ctx, Clusters_GetClusters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListClustersReply)
	err := c.cc.Invoke(ctx, Clusters_ListClusters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClustersServer is the server API for Clusters service.
// All implementations must embed UnimplementedClustersServer
// for forward compatibility.
type ClustersServer interface {
	CreateClusters(context.Context, *CreateClustersRequest) (*CreateClustersReply, error)
	UpdateClusters(context.Context, *UpdateClustersRequest) (*UpdateClustersReply, error)
	DeleteClusters(context.Context, *DeleteClustersRequest) (*DeleteClustersReply, error)
	GetClusters(context.Context, *GetClustersRequest) (*GetClustersReply, error)
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersReply, error)
	mustEmbedUnimplementedClustersServer()
}

// UnimplementedClustersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClustersServer struct{}

func (UnimplementedClustersServer) CreateClusters(context.Context, *CreateClustersRequest) (*CreateClustersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClusters not implemented")
}
func (UnimplementedClustersServer) UpdateClusters(context.Context, *UpdateClustersRequest) (*UpdateClustersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClusters not implemented")
}
func (UnimplementedClustersServer) DeleteClusters(context.Context, *DeleteClustersRequest) (*DeleteClustersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClusters not implemented")
}
func (UnimplementedClustersServer) GetClusters(context.Context, *GetClustersRequest) (*GetClustersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusters not implemented")
}
func (UnimplementedClustersServer) ListClusters(context.Context, *ListClustersRequest) (*ListClustersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (UnimplementedClustersServer) mustEmbedUnimplementedClustersServer() {}
func (UnimplementedClustersServer) testEmbeddedByValue()                  {}

// UnsafeClustersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClustersServer will
// result in compilation errors.
type UnsafeClustersServer interface {
	mustEmbedUnimplementedClustersServer()
}

func RegisterClustersServer(s grpc.ServiceRegistrar, srv ClustersServer) {
	// If the following call pancis, it indicates UnimplementedClustersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Clusters_ServiceDesc, srv)
}

func _Clusters_CreateClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).CreateClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Clusters_CreateClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).CreateClusters(ctx, req.(*CreateClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_UpdateClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).UpdateClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Clusters_UpdateClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).UpdateClusters(ctx, req.(*UpdateClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_DeleteClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).DeleteClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Clusters_DeleteClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).DeleteClusters(ctx, req.(*DeleteClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_GetClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).GetClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Clusters_GetClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).GetClusters(ctx, req.(*GetClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Clusters_ListClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Clusters_ServiceDesc is the grpc.ServiceDesc for Clusters service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Clusters_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.opspillar.v1.Clusters",
	HandlerType: (*ClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClusters",
			Handler:    _Clusters_CreateClusters_Handler,
		},
		{
			MethodName: "UpdateClusters",
			Handler:    _Clusters_UpdateClusters_Handler,
		},
		{
			MethodName: "DeleteClusters",
			Handler:    _Clusters_DeleteClusters_Handler,
		},
		{
			MethodName: "GetClusters",
			Handler:    _Clusters_GetClusters_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _Clusters_ListClusters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opspillar/v1/clusters.proto",
}
