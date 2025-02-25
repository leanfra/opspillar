// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v3.12.4
// source: opspillar/v1/datacenters.proto

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

const OperationDatacentersCreateDatacenters = "/api.opspillar.v1.Datacenters/CreateDatacenters"
const OperationDatacentersDeleteDatacenters = "/api.opspillar.v1.Datacenters/DeleteDatacenters"
const OperationDatacentersGetDatacenters = "/api.opspillar.v1.Datacenters/GetDatacenters"
const OperationDatacentersListDatacenters = "/api.opspillar.v1.Datacenters/ListDatacenters"
const OperationDatacentersUpdateDatacenters = "/api.opspillar.v1.Datacenters/UpdateDatacenters"

type DatacentersHTTPServer interface {
	CreateDatacenters(context.Context, *CreateDatacentersRequest) (*CreateDatacentersReply, error)
	DeleteDatacenters(context.Context, *DeleteDatacentersRequest) (*DeleteDatacentersReply, error)
	GetDatacenters(context.Context, *GetDatacentersRequest) (*GetDatacentersReply, error)
	ListDatacenters(context.Context, *ListDatacentersRequest) (*ListDatacentersReply, error)
	UpdateDatacenters(context.Context, *UpdateDatacentersRequest) (*UpdateDatacentersReply, error)
}

func RegisterDatacentersHTTPServer(s *http.Server, srv DatacentersHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/datacenters/create", _Datacenters_CreateDatacenters0_HTTP_Handler(srv))
	r.POST("/api/v1/datacenters/update", _Datacenters_UpdateDatacenters0_HTTP_Handler(srv))
	r.POST("/api/v1/datacenters/delete", _Datacenters_DeleteDatacenters0_HTTP_Handler(srv))
	r.GET("/api/v1/datacenters/{id}", _Datacenters_GetDatacenters0_HTTP_Handler(srv))
	r.POST("/api/v1/datacenters/list", _Datacenters_ListDatacenters0_HTTP_Handler(srv))
}

func _Datacenters_CreateDatacenters0_HTTP_Handler(srv DatacentersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDatacentersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDatacentersCreateDatacenters)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDatacenters(ctx, req.(*CreateDatacentersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDatacentersReply)
		return ctx.Result(200, reply)
	}
}

func _Datacenters_UpdateDatacenters0_HTTP_Handler(srv DatacentersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDatacentersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDatacentersUpdateDatacenters)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDatacenters(ctx, req.(*UpdateDatacentersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDatacentersReply)
		return ctx.Result(200, reply)
	}
}

func _Datacenters_DeleteDatacenters0_HTTP_Handler(srv DatacentersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDatacentersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDatacentersDeleteDatacenters)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDatacenters(ctx, req.(*DeleteDatacentersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDatacentersReply)
		return ctx.Result(200, reply)
	}
}

func _Datacenters_GetDatacenters0_HTTP_Handler(srv DatacentersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDatacentersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDatacentersGetDatacenters)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDatacenters(ctx, req.(*GetDatacentersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDatacentersReply)
		return ctx.Result(200, reply)
	}
}

func _Datacenters_ListDatacenters0_HTTP_Handler(srv DatacentersHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDatacentersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDatacentersListDatacenters)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDatacenters(ctx, req.(*ListDatacentersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDatacentersReply)
		return ctx.Result(200, reply)
	}
}

type DatacentersHTTPClient interface {
	CreateDatacenters(ctx context.Context, req *CreateDatacentersRequest, opts ...http.CallOption) (rsp *CreateDatacentersReply, err error)
	DeleteDatacenters(ctx context.Context, req *DeleteDatacentersRequest, opts ...http.CallOption) (rsp *DeleteDatacentersReply, err error)
	GetDatacenters(ctx context.Context, req *GetDatacentersRequest, opts ...http.CallOption) (rsp *GetDatacentersReply, err error)
	ListDatacenters(ctx context.Context, req *ListDatacentersRequest, opts ...http.CallOption) (rsp *ListDatacentersReply, err error)
	UpdateDatacenters(ctx context.Context, req *UpdateDatacentersRequest, opts ...http.CallOption) (rsp *UpdateDatacentersReply, err error)
}

type DatacentersHTTPClientImpl struct {
	cc *http.Client
}

func NewDatacentersHTTPClient(client *http.Client) DatacentersHTTPClient {
	return &DatacentersHTTPClientImpl{client}
}

func (c *DatacentersHTTPClientImpl) CreateDatacenters(ctx context.Context, in *CreateDatacentersRequest, opts ...http.CallOption) (*CreateDatacentersReply, error) {
	var out CreateDatacentersReply
	pattern := "/api/v1/datacenters/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDatacentersCreateDatacenters))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DatacentersHTTPClientImpl) DeleteDatacenters(ctx context.Context, in *DeleteDatacentersRequest, opts ...http.CallOption) (*DeleteDatacentersReply, error) {
	var out DeleteDatacentersReply
	pattern := "/api/v1/datacenters/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDatacentersDeleteDatacenters))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DatacentersHTTPClientImpl) GetDatacenters(ctx context.Context, in *GetDatacentersRequest, opts ...http.CallOption) (*GetDatacentersReply, error) {
	var out GetDatacentersReply
	pattern := "/api/v1/datacenters/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDatacentersGetDatacenters))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DatacentersHTTPClientImpl) ListDatacenters(ctx context.Context, in *ListDatacentersRequest, opts ...http.CallOption) (*ListDatacentersReply, error) {
	var out ListDatacentersReply
	pattern := "/api/v1/datacenters/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDatacentersListDatacenters))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DatacentersHTTPClientImpl) UpdateDatacenters(ctx context.Context, in *UpdateDatacentersRequest, opts ...http.CallOption) (*UpdateDatacentersReply, error) {
	var out UpdateDatacentersReply
	pattern := "/api/v1/datacenters/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDatacentersUpdateDatacenters))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
