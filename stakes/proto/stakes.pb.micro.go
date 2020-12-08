// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/stakes.proto

package stakes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Api Endpoints for Stakes service

func NewStakesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Stakes service

type StakesService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Stakes_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Stakes_PingPongService, error)
}

type stakesService struct {
	c    client.Client
	name string
}

func NewStakesService(name string, c client.Client) StakesService {
	return &stakesService{
		c:    c,
		name: name,
	}
}

func (c *stakesService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Stakes.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakesService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Stakes_StreamService, error) {
	req := c.c.NewRequest(c.name, "Stakes.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &stakesServiceStream{stream}, nil
}

type Stakes_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type stakesServiceStream struct {
	stream client.Stream
}

func (x *stakesServiceStream) Close() error {
	return x.stream.Close()
}

func (x *stakesServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *stakesServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *stakesServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *stakesServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stakesService) PingPong(ctx context.Context, opts ...client.CallOption) (Stakes_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Stakes.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &stakesServicePingPong{stream}, nil
}

type Stakes_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type stakesServicePingPong struct {
	stream client.Stream
}

func (x *stakesServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *stakesServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *stakesServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *stakesServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *stakesServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *stakesServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Stakes service

type StakesHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Stakes_StreamStream) error
	PingPong(context.Context, Stakes_PingPongStream) error
}

func RegisterStakesHandler(s server.Server, hdlr StakesHandler, opts ...server.HandlerOption) error {
	type stakes interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Stakes struct {
		stakes
	}
	h := &stakesHandler{hdlr}
	return s.Handle(s.NewHandler(&Stakes{h}, opts...))
}

type stakesHandler struct {
	StakesHandler
}

func (h *stakesHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.StakesHandler.Call(ctx, in, out)
}

func (h *stakesHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.StakesHandler.Stream(ctx, m, &stakesStreamStream{stream})
}

type Stakes_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type stakesStreamStream struct {
	stream server.Stream
}

func (x *stakesStreamStream) Close() error {
	return x.stream.Close()
}

func (x *stakesStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *stakesStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *stakesStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *stakesStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *stakesHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.StakesHandler.PingPong(ctx, &stakesPingPongStream{stream})
}

type Stakes_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type stakesPingPongStream struct {
	stream server.Stream
}

func (x *stakesPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *stakesPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *stakesPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *stakesPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *stakesPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *stakesPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}