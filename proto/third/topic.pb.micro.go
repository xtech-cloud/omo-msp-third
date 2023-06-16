// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/third/topic.proto

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

// Client API for TopicService service

type TopicService interface {
	AddOne(ctx context.Context, in *ReqTopicAdd, opts ...client.CallOption) (*ReplyTopicInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTopicInfo, error)
	GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyTopicList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	AddRecord(ctx context.Context, in *ReqTopicRecord, opts ...client.CallOption) (*ReplyInfo, error)
	GetRecords(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyTopicRecords, error)
}

type topicService struct {
	c    client.Client
	name string
}

func NewTopicService(name string, c client.Client) TopicService {
	return &topicService{
		c:    c,
		name: name,
	}
}

func (c *topicService) AddOne(ctx context.Context, in *ReqTopicAdd, opts ...client.CallOption) (*ReplyTopicInfo, error) {
	req := c.c.NewRequest(c.name, "TopicService.AddOne", in)
	out := new(ReplyTopicInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTopicInfo, error) {
	req := c.c.NewRequest(c.name, "TopicService.GetOne", in)
	out := new(ReplyTopicInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicService) GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyTopicList, error) {
	req := c.c.NewRequest(c.name, "TopicService.GetByFilter", in)
	out := new(ReplyTopicList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "TopicService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TopicService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TopicService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicService) AddRecord(ctx context.Context, in *ReqTopicRecord, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TopicService.AddRecord", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicService) GetRecords(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyTopicRecords, error) {
	req := c.c.NewRequest(c.name, "TopicService.GetRecords", in)
	out := new(ReplyTopicRecords)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TopicService service

type TopicServiceHandler interface {
	AddOne(context.Context, *ReqTopicAdd, *ReplyTopicInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyTopicInfo) error
	GetByFilter(context.Context, *RequestFilter, *ReplyTopicList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	AddRecord(context.Context, *ReqTopicRecord, *ReplyInfo) error
	GetRecords(context.Context, *RequestFilter, *ReplyTopicRecords) error
}

func RegisterTopicServiceHandler(s server.Server, hdlr TopicServiceHandler, opts ...server.HandlerOption) error {
	type topicService interface {
		AddOne(ctx context.Context, in *ReqTopicAdd, out *ReplyTopicInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyTopicInfo) error
		GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyTopicList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		AddRecord(ctx context.Context, in *ReqTopicRecord, out *ReplyInfo) error
		GetRecords(ctx context.Context, in *RequestFilter, out *ReplyTopicRecords) error
	}
	type TopicService struct {
		topicService
	}
	h := &topicServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TopicService{h}, opts...))
}

type topicServiceHandler struct {
	TopicServiceHandler
}

func (h *topicServiceHandler) AddOne(ctx context.Context, in *ReqTopicAdd, out *ReplyTopicInfo) error {
	return h.TopicServiceHandler.AddOne(ctx, in, out)
}

func (h *topicServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyTopicInfo) error {
	return h.TopicServiceHandler.GetOne(ctx, in, out)
}

func (h *topicServiceHandler) GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyTopicList) error {
	return h.TopicServiceHandler.GetByFilter(ctx, in, out)
}

func (h *topicServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.TopicServiceHandler.GetStatistic(ctx, in, out)
}

func (h *topicServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.TopicServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *topicServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.TopicServiceHandler.RemoveOne(ctx, in, out)
}

func (h *topicServiceHandler) AddRecord(ctx context.Context, in *ReqTopicRecord, out *ReplyInfo) error {
	return h.TopicServiceHandler.AddRecord(ctx, in, out)
}

func (h *topicServiceHandler) GetRecords(ctx context.Context, in *RequestFilter, out *ReplyTopicRecords) error {
	return h.TopicServiceHandler.GetRecords(ctx, in, out)
}