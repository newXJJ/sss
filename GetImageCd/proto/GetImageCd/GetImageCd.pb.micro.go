// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetImageCd/GetImageCd.proto

package go_micro_service_GetImageCd

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

// Client API for GetImageCd service

type GetImageCdService interface {
	GetImageCode(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	//
	GetSmsCode(ctx context.Context, in *SmsRequest, opts ...client.CallOption) (*SmsResponse, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetImageCd_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (GetImageCd_PingPongService, error)
	//注册请求
	PostRet(ctx context.Context, in *PostRetRequest, opts ...client.CallOption) (*PostRetResponse, error)
}

type getImageCdService struct {
	c    client.Client
	name string
}

func NewGetImageCdService(name string, c client.Client) GetImageCdService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.service.GetImageCd"
	}
	return &getImageCdService{
		c:    c,
		name: name,
	}
}

func (c *getImageCdService) GetImageCode(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "GetImageCd.GetImageCode", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getImageCdService) GetSmsCode(ctx context.Context, in *SmsRequest, opts ...client.CallOption) (*SmsResponse, error) {
	req := c.c.NewRequest(c.name, "GetImageCd.GetSmsCode", in)
	out := new(SmsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getImageCdService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetImageCd_StreamService, error) {
	req := c.c.NewRequest(c.name, "GetImageCd.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &getImageCdServiceStream{stream}, nil
}

type GetImageCd_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type getImageCdServiceStream struct {
	stream client.Stream
}

func (x *getImageCdServiceStream) Close() error {
	return x.stream.Close()
}

func (x *getImageCdServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getImageCdServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getImageCdServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getImageCdService) PingPong(ctx context.Context, opts ...client.CallOption) (GetImageCd_PingPongService, error) {
	req := c.c.NewRequest(c.name, "GetImageCd.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &getImageCdServicePingPong{stream}, nil
}

type GetImageCd_PingPongService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type getImageCdServicePingPong struct {
	stream client.Stream
}

func (x *getImageCdServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *getImageCdServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getImageCdServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getImageCdServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *getImageCdServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getImageCdService) PostRet(ctx context.Context, in *PostRetRequest, opts ...client.CallOption) (*PostRetResponse, error) {
	req := c.c.NewRequest(c.name, "GetImageCd.PostRet", in)
	out := new(PostRetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetImageCd service

type GetImageCdHandler interface {
	GetImageCode(context.Context, *Request, *Response) error
	//
	GetSmsCode(context.Context, *SmsRequest, *SmsResponse) error
	Stream(context.Context, *StreamingRequest, GetImageCd_StreamStream) error
	PingPong(context.Context, GetImageCd_PingPongStream) error
	//注册请求
	PostRet(context.Context, *PostRetRequest, *PostRetResponse) error
}

func RegisterGetImageCdHandler(s server.Server, hdlr GetImageCdHandler, opts ...server.HandlerOption) error {
	type getImageCd interface {
		GetImageCode(ctx context.Context, in *Request, out *Response) error
		GetSmsCode(ctx context.Context, in *SmsRequest, out *SmsResponse) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
		PostRet(ctx context.Context, in *PostRetRequest, out *PostRetResponse) error
	}
	type GetImageCd struct {
		getImageCd
	}
	h := &getImageCdHandler{hdlr}
	return s.Handle(s.NewHandler(&GetImageCd{h}, opts...))
}

type getImageCdHandler struct {
	GetImageCdHandler
}

func (h *getImageCdHandler) GetImageCode(ctx context.Context, in *Request, out *Response) error {
	return h.GetImageCdHandler.GetImageCode(ctx, in, out)
}

func (h *getImageCdHandler) GetSmsCode(ctx context.Context, in *SmsRequest, out *SmsResponse) error {
	return h.GetImageCdHandler.GetSmsCode(ctx, in, out)
}

func (h *getImageCdHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GetImageCdHandler.Stream(ctx, m, &getImageCdStreamStream{stream})
}

type GetImageCd_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type getImageCdStreamStream struct {
	stream server.Stream
}

func (x *getImageCdStreamStream) Close() error {
	return x.stream.Close()
}

func (x *getImageCdStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getImageCdStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getImageCdStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *getImageCdHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GetImageCdHandler.PingPong(ctx, &getImageCdPingPongStream{stream})
}

type GetImageCd_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type getImageCdPingPongStream struct {
	stream server.Stream
}

func (x *getImageCdPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *getImageCdPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getImageCdPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getImageCdPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *getImageCdPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *getImageCdHandler) PostRet(ctx context.Context, in *PostRetRequest, out *PostRetResponse) error {
	return h.GetImageCdHandler.PostRet(ctx, in, out)
}
