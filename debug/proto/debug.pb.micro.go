// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/debug.proto

package debug

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "go-micro.dev/v5/client"
	server "go-micro.dev/v5/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Debug service

type DebugService interface {
	Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (Debug_LogService, error)
	Health(ctx context.Context, in *HealthRequest, opts ...client.CallOption) (*HealthResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error)
	Trace(ctx context.Context, in *TraceRequest, opts ...client.CallOption) (*TraceResponse, error)
	MessageBus(ctx context.Context, opts ...client.CallOption) (Debug_MessageBusService, error)
}

type debugService struct {
	c    client.Client
	name string
}

func NewDebugService(name string, c client.Client) DebugService {
	return &debugService{
		c:    c,
		name: name,
	}
}

func (c *debugService) Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (Debug_LogService, error) {
	req := c.c.NewRequest(c.name, "Debug.Log", &LogRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &debugServiceLog{stream}, nil
}

type Debug_LogService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*Record, error)
}

type debugServiceLog struct {
	stream client.Stream
}

func (x *debugServiceLog) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *debugServiceLog) Close() error {
	return x.stream.Close()
}

func (x *debugServiceLog) Context() context.Context {
	return x.stream.Context()
}

func (x *debugServiceLog) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugServiceLog) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *debugServiceLog) Recv() (*Record, error) {
	m := new(Record)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugService) Health(ctx context.Context, in *HealthRequest, opts ...client.CallOption) (*HealthResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Health", in)
	out := new(HealthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Stats", in)
	out := new(StatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) Trace(ctx context.Context, in *TraceRequest, opts ...client.CallOption) (*TraceResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Trace", in)
	out := new(TraceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) MessageBus(ctx context.Context, opts ...client.CallOption) (Debug_MessageBusService, error) {
	req := c.c.NewRequest(c.name, "Debug.MessageBus", &BusMsg{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &debugServiceMessageBus{stream}, nil
}

type Debug_MessageBusService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BusMsg) error
	Recv() (*BusMsg, error)
}

type debugServiceMessageBus struct {
	stream client.Stream
}

func (x *debugServiceMessageBus) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *debugServiceMessageBus) Close() error {
	return x.stream.Close()
}

func (x *debugServiceMessageBus) Context() context.Context {
	return x.stream.Context()
}

func (x *debugServiceMessageBus) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugServiceMessageBus) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *debugServiceMessageBus) Send(m *BusMsg) error {
	return x.stream.Send(m)
}

func (x *debugServiceMessageBus) Recv() (*BusMsg, error) {
	m := new(BusMsg)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Debug service

type DebugHandler interface {
	Log(context.Context, *LogRequest, Debug_LogStream) error
	Health(context.Context, *HealthRequest, *HealthResponse) error
	Stats(context.Context, *StatsRequest, *StatsResponse) error
	Trace(context.Context, *TraceRequest, *TraceResponse) error
	MessageBus(context.Context, Debug_MessageBusStream) error
}

func RegisterDebugHandler(s server.Server, hdlr DebugHandler, opts ...server.HandlerOption) error {
	type debug interface {
		Log(ctx context.Context, stream server.Stream) error
		Health(ctx context.Context, in *HealthRequest, out *HealthResponse) error
		Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error
		Trace(ctx context.Context, in *TraceRequest, out *TraceResponse) error
		MessageBus(ctx context.Context, stream server.Stream) error
	}
	type Debug struct {
		debug
	}
	h := &debugHandler{hdlr}
	return s.Handle(s.NewHandler(&Debug{h}, opts...))
}

type debugHandler struct {
	DebugHandler
}

func (h *debugHandler) Log(ctx context.Context, stream server.Stream) error {
	m := new(LogRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.DebugHandler.Log(ctx, m, &debugLogStream{stream})
}

type Debug_LogStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Record) error
}

type debugLogStream struct {
	stream server.Stream
}

func (x *debugLogStream) Close() error {
	return x.stream.Close()
}

func (x *debugLogStream) Context() context.Context {
	return x.stream.Context()
}

func (x *debugLogStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugLogStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *debugLogStream) Send(m *Record) error {
	return x.stream.Send(m)
}

func (h *debugHandler) Health(ctx context.Context, in *HealthRequest, out *HealthResponse) error {
	return h.DebugHandler.Health(ctx, in, out)
}

func (h *debugHandler) Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error {
	return h.DebugHandler.Stats(ctx, in, out)
}

func (h *debugHandler) Trace(ctx context.Context, in *TraceRequest, out *TraceResponse) error {
	return h.DebugHandler.Trace(ctx, in, out)
}

func (h *debugHandler) MessageBus(ctx context.Context, stream server.Stream) error {
	return h.DebugHandler.MessageBus(ctx, &debugMessageBusStream{stream})
}

type Debug_MessageBusStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BusMsg) error
	Recv() (*BusMsg, error)
}

type debugMessageBusStream struct {
	stream server.Stream
}

func (x *debugMessageBusStream) Close() error {
	return x.stream.Close()
}

func (x *debugMessageBusStream) Context() context.Context {
	return x.stream.Context()
}

func (x *debugMessageBusStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *debugMessageBusStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *debugMessageBusStream) Send(m *BusMsg) error {
	return x.stream.Send(m)
}

func (x *debugMessageBusStream) Recv() (*BusMsg, error) {
	m := new(BusMsg)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
