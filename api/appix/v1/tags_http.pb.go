// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v3.12.4
// source: api/appix/v1/tags.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTagsCreateTags = "/api.appix.v1.Tags/CreateTags"
const OperationTagsDeleteTags = "/api.appix.v1.Tags/DeleteTags"
const OperationTagsGetTags = "/api.appix.v1.Tags/GetTags"
const OperationTagsListTags = "/api.appix.v1.Tags/ListTags"
const OperationTagsUpdateTags = "/api.appix.v1.Tags/UpdateTags"

type TagsHTTPServer interface {
	CreateTags(context.Context, *CreateTagsRequest) (*CreateTagsReply, error)
	DeleteTags(context.Context, *DeleteTagsRequest) (*DeleteTagsReply, error)
	GetTags(context.Context, *GetTagsRequest) (*GetTagsReply, error)
	ListTags(context.Context, *ListTagsRequest) (*ListTagsReply, error)
	UpdateTags(context.Context, *UpdateTagsRequest) (*UpdateTagsReply, error)
}

func RegisterTagsHTTPServer(s *http.Server, srv TagsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/tags/create", _Tags_CreateTags0_HTTP_Handler(srv))
	r.POST("/v1/tags/update", _Tags_UpdateTags0_HTTP_Handler(srv))
	r.POST("/v1/tags/delete", _Tags_DeleteTags0_HTTP_Handler(srv))
	r.GET("/v1/tags/{id}", _Tags_GetTags0_HTTP_Handler(srv))
	r.POST("/v1/tags/list", _Tags_ListTags0_HTTP_Handler(srv))
}

func _Tags_CreateTags0_HTTP_Handler(srv TagsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTagsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagsCreateTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTags(ctx, req.(*CreateTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTagsReply)
		return ctx.Result(200, reply)
	}
}

func _Tags_UpdateTags0_HTTP_Handler(srv TagsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTagsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagsUpdateTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTags(ctx, req.(*UpdateTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTagsReply)
		return ctx.Result(200, reply)
	}
}

func _Tags_DeleteTags0_HTTP_Handler(srv TagsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTagsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagsDeleteTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTags(ctx, req.(*DeleteTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTagsReply)
		return ctx.Result(200, reply)
	}
}

func _Tags_GetTags0_HTTP_Handler(srv TagsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTagsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagsGetTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTags(ctx, req.(*GetTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTagsReply)
		return ctx.Result(200, reply)
	}
}

func _Tags_ListTags0_HTTP_Handler(srv TagsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTagsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagsListTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTags(ctx, req.(*ListTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTagsReply)
		return ctx.Result(200, reply)
	}
}

type TagsHTTPClient interface {
	CreateTags(ctx context.Context, req *CreateTagsRequest, opts ...http.CallOption) (rsp *CreateTagsReply, err error)
	DeleteTags(ctx context.Context, req *DeleteTagsRequest, opts ...http.CallOption) (rsp *DeleteTagsReply, err error)
	GetTags(ctx context.Context, req *GetTagsRequest, opts ...http.CallOption) (rsp *GetTagsReply, err error)
	ListTags(ctx context.Context, req *ListTagsRequest, opts ...http.CallOption) (rsp *ListTagsReply, err error)
	UpdateTags(ctx context.Context, req *UpdateTagsRequest, opts ...http.CallOption) (rsp *UpdateTagsReply, err error)
}

type TagsHTTPClientImpl struct {
	cc *http.Client
}

func NewTagsHTTPClient(client *http.Client) TagsHTTPClient {
	return &TagsHTTPClientImpl{client}
}

func (c *TagsHTTPClientImpl) CreateTags(ctx context.Context, in *CreateTagsRequest, opts ...http.CallOption) (*CreateTagsReply, error) {
	var out CreateTagsReply
	pattern := "/v1/tags/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagsCreateTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TagsHTTPClientImpl) DeleteTags(ctx context.Context, in *DeleteTagsRequest, opts ...http.CallOption) (*DeleteTagsReply, error) {
	var out DeleteTagsReply
	pattern := "/v1/tags/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagsDeleteTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TagsHTTPClientImpl) GetTags(ctx context.Context, in *GetTagsRequest, opts ...http.CallOption) (*GetTagsReply, error) {
	var out GetTagsReply
	pattern := "/v1/tags/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTagsGetTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TagsHTTPClientImpl) ListTags(ctx context.Context, in *ListTagsRequest, opts ...http.CallOption) (*ListTagsReply, error) {
	var out ListTagsReply
	pattern := "/v1/tags/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagsListTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TagsHTTPClientImpl) UpdateTags(ctx context.Context, in *UpdateTagsRequest, opts ...http.CallOption) (*UpdateTagsReply, error) {
	var out UpdateTagsReply
	pattern := "/v1/tags/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagsUpdateTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
