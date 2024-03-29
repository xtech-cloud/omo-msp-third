// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/third/holiday.proto

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

// Client API for HolidayService service

type HolidayService interface {
	AddOne(ctx context.Context, in *ReqHolidayAdd, opts ...client.CallOption) (*ReplyHolidayInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyHolidayInfo, error)
	UpdateBase(ctx context.Context, in *ReqHolidayUpdate, opts ...client.CallOption) (*ReplyHolidayInfo, error)
	GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyHolidayList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
}

type holidayService struct {
	c    client.Client
	name string
}

func NewHolidayService(name string, c client.Client) HolidayService {
	return &holidayService{
		c:    c,
		name: name,
	}
}

func (c *holidayService) AddOne(ctx context.Context, in *ReqHolidayAdd, opts ...client.CallOption) (*ReplyHolidayInfo, error) {
	req := c.c.NewRequest(c.name, "HolidayService.AddOne", in)
	out := new(ReplyHolidayInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidayService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyHolidayInfo, error) {
	req := c.c.NewRequest(c.name, "HolidayService.GetOne", in)
	out := new(ReplyHolidayInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidayService) UpdateBase(ctx context.Context, in *ReqHolidayUpdate, opts ...client.CallOption) (*ReplyHolidayInfo, error) {
	req := c.c.NewRequest(c.name, "HolidayService.UpdateBase", in)
	out := new(ReplyHolidayInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidayService) GetByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyHolidayList, error) {
	req := c.c.NewRequest(c.name, "HolidayService.GetByFilter", in)
	out := new(ReplyHolidayList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidayService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "HolidayService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidayService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "HolidayService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *holidayService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "HolidayService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HolidayService service

type HolidayServiceHandler interface {
	AddOne(context.Context, *ReqHolidayAdd, *ReplyHolidayInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyHolidayInfo) error
	UpdateBase(context.Context, *ReqHolidayUpdate, *ReplyHolidayInfo) error
	GetByFilter(context.Context, *RequestFilter, *ReplyHolidayList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
}

func RegisterHolidayServiceHandler(s server.Server, hdlr HolidayServiceHandler, opts ...server.HandlerOption) error {
	type holidayService interface {
		AddOne(ctx context.Context, in *ReqHolidayAdd, out *ReplyHolidayInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyHolidayInfo) error
		UpdateBase(ctx context.Context, in *ReqHolidayUpdate, out *ReplyHolidayInfo) error
		GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyHolidayList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
	}
	type HolidayService struct {
		holidayService
	}
	h := &holidayServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&HolidayService{h}, opts...))
}

type holidayServiceHandler struct {
	HolidayServiceHandler
}

func (h *holidayServiceHandler) AddOne(ctx context.Context, in *ReqHolidayAdd, out *ReplyHolidayInfo) error {
	return h.HolidayServiceHandler.AddOne(ctx, in, out)
}

func (h *holidayServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyHolidayInfo) error {
	return h.HolidayServiceHandler.GetOne(ctx, in, out)
}

func (h *holidayServiceHandler) UpdateBase(ctx context.Context, in *ReqHolidayUpdate, out *ReplyHolidayInfo) error {
	return h.HolidayServiceHandler.UpdateBase(ctx, in, out)
}

func (h *holidayServiceHandler) GetByFilter(ctx context.Context, in *RequestFilter, out *ReplyHolidayList) error {
	return h.HolidayServiceHandler.GetByFilter(ctx, in, out)
}

func (h *holidayServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.HolidayServiceHandler.GetStatistic(ctx, in, out)
}

func (h *holidayServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.HolidayServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *holidayServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.HolidayServiceHandler.RemoveOne(ctx, in, out)
}
