package rpc_server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-go/test/framework/rpc_server/rpc_server"
)

type StreamServer interface {
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
}

type StreamHandler func(context.Context, StreamServer) (interface{}, error)

type Stream struct {
	Handler         StreamHandler
	StreamingClient bool
	StreamingServer bool
}

type ServiceDesc struct {
	Methods map[string]Method
	Streams map[string]Stream
}

type serviceSet struct {
	services map[string]*ServiceDesc
}

func (s *serviceSet) register(name string, desc *ServiceDesc) {
	if _, ok := s.services[name]; ok {
		panic(fmt.Errorf("duplicate service %v registered", name))
	}
	s.services[name] = desc
}

func (s *serviceSet) handler(ctx context.Context, req rpc_server.Request, response func(*status.Status, []byte, bool, bool) error) (*StreamHandler, error) {
	_, ok := s.services[req.Service]
	if !ok {
		// 使用code来对报错进行快速分类
		return nil, status.Errorf(codes.Unimplemented, "service %v", req.Service)
	}
	return nil, nil
}
