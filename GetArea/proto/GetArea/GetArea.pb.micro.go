// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GetArea/GetArea.proto

package go_micro_service_GetArea

import (
	fmt "fmt"
	"github.com/astaxie/beego"
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

// Client API for GetArea service

type GetAreaService interface {
	GetArea(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetArea_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (GetArea_PingPongService, error)
}

type getAreaService struct {
	c    client.Client
	name string
}

func NewGetAreaService(name string, c client.Client) GetAreaService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.service.GetArea"
	}
	return &getAreaService{
		c:    c,
		name: name,
	}
}

func (c *getAreaService) GetArea(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	beego.Info("paodaozhelilaile ")
	req := c.c.NewRequest(c.name, "GetArea.GetArea", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getAreaService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (GetArea_StreamService, error) {
	req := c.c.NewRequest(c.name, "GetArea.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &getAreaServiceStream{stream}, nil
}

type GetArea_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type getAreaServiceStream struct {
	stream client.Stream
}

func (x *getAreaServiceStream) Close() error {
	return x.stream.Close()
}

func (x *getAreaServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getAreaServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getAreaServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *getAreaService) PingPong(ctx context.Context, opts ...client.CallOption) (GetArea_PingPongService, error) {
	req := c.c.NewRequest(c.name, "GetArea.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &getAreaServicePingPong{stream}, nil
}

type GetArea_PingPongService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type getAreaServicePingPong struct {
	stream client.Stream
}

func (x *getAreaServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *getAreaServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getAreaServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getAreaServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *getAreaServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GetArea service

type GetAreaHandler interface {
	GetArea(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, GetArea_StreamStream) error
	PingPong(context.Context, GetArea_PingPongStream) error
}

func RegisterGetAreaHandler(s server.Server, hdlr GetAreaHandler, opts ...server.HandlerOption) error {
	fmt.Println("register my func")
	type getArea interface {
		GetArea(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type GetArea struct {
		getArea
	}
	h := &getAreaHandler{hdlr}
	return s.Handle(s.NewHandler(&GetArea{h}, opts...))
}

type getAreaHandler struct {
	GetAreaHandler
}

func (h *getAreaHandler) GetArea(ctx context.Context, in *Request, out *Response) error {
	return h.GetAreaHandler.GetArea(ctx, in, out)
}

func (h *getAreaHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GetAreaHandler.Stream(ctx, m, &getAreaStreamStream{stream})
}

type GetArea_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type getAreaStreamStream struct {
	stream server.Stream
}

func (x *getAreaStreamStream) Close() error {
	return x.stream.Close()
}

func (x *getAreaStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getAreaStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getAreaStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *getAreaHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GetAreaHandler.PingPong(ctx, &getAreaPingPongStream{stream})
}

type GetArea_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type getAreaPingPongStream struct {
	stream server.Stream
}

func (x *getAreaPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *getAreaPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *getAreaPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *getAreaPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *getAreaPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}