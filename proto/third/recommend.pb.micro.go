// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/third/recommend.proto

package third

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for RecommendService service

type RecommendService interface {
	AddOne(ctx context.Context, in *ReqRecommendAdd, opts ...client.CallOption) (*ReplyRecommendInfo, error)
	UpdateOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRecommendInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRecommendInfo, error)
	GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyRecommendList, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyRecommendList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
}

type recommendService struct {
	c    client.Client
	name string
}

func NewRecommendService(name string, c client.Client) RecommendService {
	return &recommendService{
		c:    c,
		name: name,
	}
}

func (c *recommendService) AddOne(ctx context.Context, in *ReqRecommendAdd, opts ...client.CallOption) (*ReplyRecommendInfo, error) {
	req := c.c.NewRequest(c.name, "RecommendService.AddOne", in)
	out := new(ReplyRecommendInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendService) UpdateOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRecommendInfo, error) {
	req := c.c.NewRequest(c.name, "RecommendService.UpdateOne", in)
	out := new(ReplyRecommendInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRecommendInfo, error) {
	req := c.c.NewRequest(c.name, "RecommendService.GetOne", in)
	out := new(ReplyRecommendInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendService) GetList(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyRecommendList, error) {
	req := c.c.NewRequest(c.name, "RecommendService.GetList", in)
	out := new(ReplyRecommendList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "RecommendService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendService) GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyRecommendList, error) {
	req := c.c.NewRequest(c.name, "RecommendService.GetByFilter", in)
	out := new(ReplyRecommendList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "RecommendService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "RecommendService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RecommendService service

type RecommendServiceHandler interface {
	AddOne(context.Context, *ReqRecommendAdd, *ReplyRecommendInfo) error
	UpdateOne(context.Context, *RequestInfo, *ReplyRecommendInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyRecommendInfo) error
	GetList(context.Context, *RequestPage, *ReplyRecommendList) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetByFilter(context.Context, *RequestFilter, *ReplyRecommendList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
}

func RegisterRecommendServiceHandler(s server.Server, hdlr RecommendServiceHandler, opts ...server.HandlerOption) error {
	type recommendService interface {
		AddOne(ctx context.Context, in *ReqRecommendAdd, out *ReplyRecommendInfo) error
		UpdateOne(ctx context.Context, in *RequestInfo, out *ReplyRecommendInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyRecommendInfo) error
		GetList(ctx context.Context, in *RequestPage, out *ReplyRecommendList) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyRecommendList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
	}
	type RecommendService struct {
		recommendService
	}
	h := &recommendServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RecommendService{h}, opts...))
}

type recommendServiceHandler struct {
	RecommendServiceHandler
}

func (h *recommendServiceHandler) AddOne(ctx context.Context, in *ReqRecommendAdd, out *ReplyRecommendInfo) error {
	return h.RecommendServiceHandler.AddOne(ctx, in, out)
}

func (h *recommendServiceHandler) UpdateOne(ctx context.Context, in *RequestInfo, out *ReplyRecommendInfo) error {
	return h.RecommendServiceHandler.UpdateOne(ctx, in, out)
}

func (h *recommendServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyRecommendInfo) error {
	return h.RecommendServiceHandler.GetOne(ctx, in, out)
}

func (h *recommendServiceHandler) GetList(ctx context.Context, in *RequestPage, out *ReplyRecommendList) error {
	return h.RecommendServiceHandler.GetList(ctx, in, out)
}

func (h *recommendServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.RecommendServiceHandler.RemoveOne(ctx, in, out)
}

func (h *recommendServiceHandler) GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyRecommendList) error {
	return h.RecommendServiceHandler.GetByFilter(ctx, in, out)
}

func (h *recommendServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.RecommendServiceHandler.GetStatistic(ctx, in, out)
}

func (h *recommendServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.RecommendServiceHandler.UpdateByFilter(ctx, in, out)
}
