// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v3.12.4
// source: api/appix/v1/teams.proto

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

const OperationTeamsCreateTeams = "/api.appix.v1.Teams/CreateTeams"
const OperationTeamsDeleteTeams = "/api.appix.v1.Teams/DeleteTeams"
const OperationTeamsGetTeams = "/api.appix.v1.Teams/GetTeams"
const OperationTeamsListTeams = "/api.appix.v1.Teams/ListTeams"
const OperationTeamsUpdateTeams = "/api.appix.v1.Teams/UpdateTeams"

type TeamsHTTPServer interface {
	CreateTeams(context.Context, *CreateTeamsRequest) (*CreateTeamsReply, error)
	DeleteTeams(context.Context, *DeleteTeamsRequest) (*DeleteTeamsReply, error)
	GetTeams(context.Context, *GetTeamsRequest) (*GetTeamsReply, error)
	ListTeams(context.Context, *ListTeamsRequest) (*ListTeamsReply, error)
	UpdateTeams(context.Context, *UpdateTeamsRequest) (*UpdateTeamsReply, error)
}

func RegisterTeamsHTTPServer(s *http.Server, srv TeamsHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/teams/create", _Teams_CreateTeams0_HTTP_Handler(srv))
	r.POST("/api/v1/teams/update", _Teams_UpdateTeams0_HTTP_Handler(srv))
	r.POST("/api/v1/teams/delete", _Teams_DeleteTeams0_HTTP_Handler(srv))
	r.GET("/api/v1/teams/{id}", _Teams_GetTeams0_HTTP_Handler(srv))
	r.POST("/api/v1/teams/list", _Teams_ListTeams0_HTTP_Handler(srv))
}

func _Teams_CreateTeams0_HTTP_Handler(srv TeamsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTeamsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamsCreateTeams)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTeams(ctx, req.(*CreateTeamsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTeamsReply)
		return ctx.Result(200, reply)
	}
}

func _Teams_UpdateTeams0_HTTP_Handler(srv TeamsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTeamsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamsUpdateTeams)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTeams(ctx, req.(*UpdateTeamsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTeamsReply)
		return ctx.Result(200, reply)
	}
}

func _Teams_DeleteTeams0_HTTP_Handler(srv TeamsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTeamsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamsDeleteTeams)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTeams(ctx, req.(*DeleteTeamsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTeamsReply)
		return ctx.Result(200, reply)
	}
}

func _Teams_GetTeams0_HTTP_Handler(srv TeamsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTeamsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamsGetTeams)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTeams(ctx, req.(*GetTeamsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTeamsReply)
		return ctx.Result(200, reply)
	}
}

func _Teams_ListTeams0_HTTP_Handler(srv TeamsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTeamsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTeamsListTeams)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTeams(ctx, req.(*ListTeamsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTeamsReply)
		return ctx.Result(200, reply)
	}
}

type TeamsHTTPClient interface {
	CreateTeams(ctx context.Context, req *CreateTeamsRequest, opts ...http.CallOption) (rsp *CreateTeamsReply, err error)
	DeleteTeams(ctx context.Context, req *DeleteTeamsRequest, opts ...http.CallOption) (rsp *DeleteTeamsReply, err error)
	GetTeams(ctx context.Context, req *GetTeamsRequest, opts ...http.CallOption) (rsp *GetTeamsReply, err error)
	ListTeams(ctx context.Context, req *ListTeamsRequest, opts ...http.CallOption) (rsp *ListTeamsReply, err error)
	UpdateTeams(ctx context.Context, req *UpdateTeamsRequest, opts ...http.CallOption) (rsp *UpdateTeamsReply, err error)
}

type TeamsHTTPClientImpl struct {
	cc *http.Client
}

func NewTeamsHTTPClient(client *http.Client) TeamsHTTPClient {
	return &TeamsHTTPClientImpl{client}
}

func (c *TeamsHTTPClientImpl) CreateTeams(ctx context.Context, in *CreateTeamsRequest, opts ...http.CallOption) (*CreateTeamsReply, error) {
	var out CreateTeamsReply
	pattern := "/api/v1/teams/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamsCreateTeams))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamsHTTPClientImpl) DeleteTeams(ctx context.Context, in *DeleteTeamsRequest, opts ...http.CallOption) (*DeleteTeamsReply, error) {
	var out DeleteTeamsReply
	pattern := "/api/v1/teams/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamsDeleteTeams))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamsHTTPClientImpl) GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...http.CallOption) (*GetTeamsReply, error) {
	var out GetTeamsReply
	pattern := "/api/v1/teams/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTeamsGetTeams))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamsHTTPClientImpl) ListTeams(ctx context.Context, in *ListTeamsRequest, opts ...http.CallOption) (*ListTeamsReply, error) {
	var out ListTeamsReply
	pattern := "/api/v1/teams/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamsListTeams))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TeamsHTTPClientImpl) UpdateTeams(ctx context.Context, in *UpdateTeamsRequest, opts ...http.CallOption) (*UpdateTeamsReply, error) {
	var out UpdateTeamsReply
	pattern := "/api/v1/teams/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTeamsUpdateTeams))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
