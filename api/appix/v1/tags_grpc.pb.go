// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: api/appix/v1/tags.proto

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
	Tags_CreateTags_FullMethodName = "/api.appix.v1.Tags/CreateTags"
	Tags_UpdateTags_FullMethodName = "/api.appix.v1.Tags/UpdateTags"
	Tags_DeleteTags_FullMethodName = "/api.appix.v1.Tags/DeleteTags"
	Tags_GetTags_FullMethodName    = "/api.appix.v1.Tags/GetTags"
	Tags_ListTags_FullMethodName   = "/api.appix.v1.Tags/ListTags"
)

// TagsClient is the client API for Tags service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagsClient interface {
	CreateTags(ctx context.Context, in *CreateTagsRequest, opts ...grpc.CallOption) (*CreateTagsReply, error)
	UpdateTags(ctx context.Context, in *UpdateTagsRequest, opts ...grpc.CallOption) (*UpdateTagsReply, error)
	DeleteTags(ctx context.Context, in *DeleteTagsRequest, opts ...grpc.CallOption) (*DeleteTagsReply, error)
	GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*GetTagsReply, error)
	ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsReply, error)
}

type tagsClient struct {
	cc grpc.ClientConnInterface
}

func NewTagsClient(cc grpc.ClientConnInterface) TagsClient {
	return &tagsClient{cc}
}

func (c *tagsClient) CreateTags(ctx context.Context, in *CreateTagsRequest, opts ...grpc.CallOption) (*CreateTagsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTagsReply)
	err := c.cc.Invoke(ctx, Tags_CreateTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsClient) UpdateTags(ctx context.Context, in *UpdateTagsRequest, opts ...grpc.CallOption) (*UpdateTagsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTagsReply)
	err := c.cc.Invoke(ctx, Tags_UpdateTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsClient) DeleteTags(ctx context.Context, in *DeleteTagsRequest, opts ...grpc.CallOption) (*DeleteTagsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTagsReply)
	err := c.cc.Invoke(ctx, Tags_DeleteTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsClient) GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*GetTagsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTagsReply)
	err := c.cc.Invoke(ctx, Tags_GetTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsClient) ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTagsReply)
	err := c.cc.Invoke(ctx, Tags_ListTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagsServer is the server API for Tags service.
// All implementations must embed UnimplementedTagsServer
// for forward compatibility.
type TagsServer interface {
	CreateTags(context.Context, *CreateTagsRequest) (*CreateTagsReply, error)
	UpdateTags(context.Context, *UpdateTagsRequest) (*UpdateTagsReply, error)
	DeleteTags(context.Context, *DeleteTagsRequest) (*DeleteTagsReply, error)
	GetTags(context.Context, *GetTagsRequest) (*GetTagsReply, error)
	ListTags(context.Context, *ListTagsRequest) (*ListTagsReply, error)
	mustEmbedUnimplementedTagsServer()
}

// UnimplementedTagsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTagsServer struct{}

func (UnimplementedTagsServer) CreateTags(context.Context, *CreateTagsRequest) (*CreateTagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTags not implemented")
}
func (UnimplementedTagsServer) UpdateTags(context.Context, *UpdateTagsRequest) (*UpdateTagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTags not implemented")
}
func (UnimplementedTagsServer) DeleteTags(context.Context, *DeleteTagsRequest) (*DeleteTagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTags not implemented")
}
func (UnimplementedTagsServer) GetTags(context.Context, *GetTagsRequest) (*GetTagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (UnimplementedTagsServer) ListTags(context.Context, *ListTagsRequest) (*ListTagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTags not implemented")
}
func (UnimplementedTagsServer) mustEmbedUnimplementedTagsServer() {}
func (UnimplementedTagsServer) testEmbeddedByValue()              {}

// UnsafeTagsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagsServer will
// result in compilation errors.
type UnsafeTagsServer interface {
	mustEmbedUnimplementedTagsServer()
}

func RegisterTagsServer(s grpc.ServiceRegistrar, srv TagsServer) {
	// If the following call pancis, it indicates UnimplementedTagsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Tags_ServiceDesc, srv)
}

func _Tags_CreateTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServer).CreateTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tags_CreateTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServer).CreateTags(ctx, req.(*CreateTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tags_UpdateTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServer).UpdateTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tags_UpdateTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServer).UpdateTags(ctx, req.(*UpdateTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tags_DeleteTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServer).DeleteTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tags_DeleteTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServer).DeleteTags(ctx, req.(*DeleteTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tags_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tags_GetTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServer).GetTags(ctx, req.(*GetTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tags_ListTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServer).ListTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tags_ListTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServer).ListTags(ctx, req.(*ListTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tags_ServiceDesc is the grpc.ServiceDesc for Tags service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tags_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.appix.v1.Tags",
	HandlerType: (*TagsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTags",
			Handler:    _Tags_CreateTags_Handler,
		},
		{
			MethodName: "UpdateTags",
			Handler:    _Tags_UpdateTags_Handler,
		},
		{
			MethodName: "DeleteTags",
			Handler:    _Tags_DeleteTags_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _Tags_GetTags_Handler,
		},
		{
			MethodName: "ListTags",
			Handler:    _Tags_ListTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/appix/v1/tags.proto",
}
