// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v3.12.4
// source: opspillar/v1/products.proto

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

const OperationProductsCreateProducts = "/api.opspillar.v1.Products/CreateProducts"
const OperationProductsDeleteProducts = "/api.opspillar.v1.Products/DeleteProducts"
const OperationProductsGetProducts = "/api.opspillar.v1.Products/GetProducts"
const OperationProductsListProducts = "/api.opspillar.v1.Products/ListProducts"
const OperationProductsUpdateProducts = "/api.opspillar.v1.Products/UpdateProducts"

type ProductsHTTPServer interface {
	CreateProducts(context.Context, *CreateProductsRequest) (*CreateProductsReply, error)
	DeleteProducts(context.Context, *DeleteProductsRequest) (*DeleteProductsReply, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsReply, error)
	ListProducts(context.Context, *ListProductsRequest) (*ListProductsReply, error)
	UpdateProducts(context.Context, *UpdateProductsRequest) (*UpdateProductsReply, error)
}

func RegisterProductsHTTPServer(s *http.Server, srv ProductsHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/products/create", _Products_CreateProducts0_HTTP_Handler(srv))
	r.POST("/api/v1/products/update", _Products_UpdateProducts0_HTTP_Handler(srv))
	r.POST("/api/v1/products/delete", _Products_DeleteProducts0_HTTP_Handler(srv))
	r.GET("/api/v1/products/{id}", _Products_GetProducts0_HTTP_Handler(srv))
	r.POST("/api/v1/products", _Products_ListProducts0_HTTP_Handler(srv))
}

func _Products_CreateProducts0_HTTP_Handler(srv ProductsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateProductsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductsCreateProducts)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateProducts(ctx, req.(*CreateProductsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateProductsReply)
		return ctx.Result(200, reply)
	}
}

func _Products_UpdateProducts0_HTTP_Handler(srv ProductsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateProductsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductsUpdateProducts)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProducts(ctx, req.(*UpdateProductsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateProductsReply)
		return ctx.Result(200, reply)
	}
}

func _Products_DeleteProducts0_HTTP_Handler(srv ProductsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteProductsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductsDeleteProducts)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProducts(ctx, req.(*DeleteProductsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteProductsReply)
		return ctx.Result(200, reply)
	}
}

func _Products_GetProducts0_HTTP_Handler(srv ProductsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProductsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductsGetProducts)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProducts(ctx, req.(*GetProductsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProductsReply)
		return ctx.Result(200, reply)
	}
}

func _Products_ListProducts0_HTTP_Handler(srv ProductsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListProductsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProductsListProducts)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListProducts(ctx, req.(*ListProductsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListProductsReply)
		return ctx.Result(200, reply)
	}
}

type ProductsHTTPClient interface {
	CreateProducts(ctx context.Context, req *CreateProductsRequest, opts ...http.CallOption) (rsp *CreateProductsReply, err error)
	DeleteProducts(ctx context.Context, req *DeleteProductsRequest, opts ...http.CallOption) (rsp *DeleteProductsReply, err error)
	GetProducts(ctx context.Context, req *GetProductsRequest, opts ...http.CallOption) (rsp *GetProductsReply, err error)
	ListProducts(ctx context.Context, req *ListProductsRequest, opts ...http.CallOption) (rsp *ListProductsReply, err error)
	UpdateProducts(ctx context.Context, req *UpdateProductsRequest, opts ...http.CallOption) (rsp *UpdateProductsReply, err error)
}

type ProductsHTTPClientImpl struct {
	cc *http.Client
}

func NewProductsHTTPClient(client *http.Client) ProductsHTTPClient {
	return &ProductsHTTPClientImpl{client}
}

func (c *ProductsHTTPClientImpl) CreateProducts(ctx context.Context, in *CreateProductsRequest, opts ...http.CallOption) (*CreateProductsReply, error) {
	var out CreateProductsReply
	pattern := "/api/v1/products/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProductsCreateProducts))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductsHTTPClientImpl) DeleteProducts(ctx context.Context, in *DeleteProductsRequest, opts ...http.CallOption) (*DeleteProductsReply, error) {
	var out DeleteProductsReply
	pattern := "/api/v1/products/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProductsDeleteProducts))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductsHTTPClientImpl) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...http.CallOption) (*GetProductsReply, error) {
	var out GetProductsReply
	pattern := "/api/v1/products/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProductsGetProducts))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductsHTTPClientImpl) ListProducts(ctx context.Context, in *ListProductsRequest, opts ...http.CallOption) (*ListProductsReply, error) {
	var out ListProductsReply
	pattern := "/api/v1/products"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProductsListProducts))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ProductsHTTPClientImpl) UpdateProducts(ctx context.Context, in *UpdateProductsRequest, opts ...http.CallOption) (*UpdateProductsReply, error) {
	var out UpdateProductsReply
	pattern := "/api/v1/products/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProductsUpdateProducts))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
