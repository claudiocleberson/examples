// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: micro/examples/blog/tag/proto/tag/tag.proto

package go_micro_service_tag

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for TagService service

func NewTagServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TagService service

type TagService interface {
	// Increase count creates the tag or bumps the counter
	IncreaseCount(ctx context.Context, in *IncreaseCountRequest, opts ...client.CallOption) (*IncreaseCountResponse, error)
	// Decreases the counter
	DecreaseCount(ctx context.Context, in *DecreaseCountRequest, opts ...client.CallOption) (*DecreaseCountResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	// Change properties of a tag, currently only the title
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
}

type tagService struct {
	c    client.Client
	name string
}

func NewTagService(name string, c client.Client) TagService {
	return &tagService{
		c:    c,
		name: name,
	}
}

func (c *tagService) IncreaseCount(ctx context.Context, in *IncreaseCountRequest, opts ...client.CallOption) (*IncreaseCountResponse, error) {
	req := c.c.NewRequest(c.name, "TagService.IncreaseCount", in)
	out := new(IncreaseCountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagService) DecreaseCount(ctx context.Context, in *DecreaseCountRequest, opts ...client.CallOption) (*DecreaseCountResponse, error) {
	req := c.c.NewRequest(c.name, "TagService.DecreaseCount", in)
	out := new(DecreaseCountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "TagService.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "TagService.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TagService service

type TagServiceHandler interface {
	// Increase count creates the tag or bumps the counter
	IncreaseCount(context.Context, *IncreaseCountRequest, *IncreaseCountResponse) error
	// Decreases the counter
	DecreaseCount(context.Context, *DecreaseCountRequest, *DecreaseCountResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	// Change properties of a tag, currently only the title
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
}

func RegisterTagServiceHandler(s server.Server, hdlr TagServiceHandler, opts ...server.HandlerOption) error {
	type tagService interface {
		IncreaseCount(ctx context.Context, in *IncreaseCountRequest, out *IncreaseCountResponse) error
		DecreaseCount(ctx context.Context, in *DecreaseCountRequest, out *DecreaseCountResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
	}
	type TagService struct {
		tagService
	}
	h := &tagServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TagService{h}, opts...))
}

type tagServiceHandler struct {
	TagServiceHandler
}

func (h *tagServiceHandler) IncreaseCount(ctx context.Context, in *IncreaseCountRequest, out *IncreaseCountResponse) error {
	return h.TagServiceHandler.IncreaseCount(ctx, in, out)
}

func (h *tagServiceHandler) DecreaseCount(ctx context.Context, in *DecreaseCountRequest, out *DecreaseCountResponse) error {
	return h.TagServiceHandler.DecreaseCount(ctx, in, out)
}

func (h *tagServiceHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.TagServiceHandler.List(ctx, in, out)
}

func (h *tagServiceHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.TagServiceHandler.Update(ctx, in, out)
}
