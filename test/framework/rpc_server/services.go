package rpc_server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learning-go/test/framework/rpc_server/rpc_server"
	"path"
	"time"
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

func (s *serviceSet) handler(ctx context.Context, req *rpc_server.Request, response func(*status.Status, []byte, bool, bool) error) (*StreamHandler, error) {
	srv, ok := s.services[req.Service]
	if !ok {
		// 使用code来对报错进行快速分类
		return nil, status.Errorf(codes.Unimplemented, "service %v", req.Service)
	}
	if method, ok := srv.Methods[req.Method]; ok {
		go func() {
			// 从request那里获取到了meta信息，超时信息
			ctx, cancel := getRequestContext(ctx, req)
			defer cancel()

			info := &UnaryServerInfo{fullPath(req.Service, req.Method)}
			p, st := s.unaryCall(ctx, method, info, req.Payload)
			response(st, p, false, true)
		}()
		return nil, nil
	}
	return nil, nil
}

type MD map[string][]string
type metadataKey struct{}

var noopFunc = func() {}

func (m MD) fromContext(request *rpc_server.Request) {
	for _, meta := range request.Metadata {
		m[meta.Key] = append(m[meta.Key], meta.Value)
	}
}

func getRequestContext(ctx context.Context, req *rpc_server.Request) (retCtx context.Context, cancel func()) {
	if len(req.Metadata) > 0 {
		md := MD{}
		md.fromContext(req)
		ctx = context.WithValue(ctx, metadataKey{}, md)
	}
	cancel = noopFunc
	if req.TimeoutNano == 0 {
		return ctx, cancel
	}
	ctx, cancel = context.WithTimeout(ctx, time.Duration(req.TimeoutNano))
	return ctx, cancel
}

func fullPath(service string, method string) string {
	return "/" + path.Join(service, method)
}

func (s *serviceSet) unaryCall(ctx context.Context, method Method, info *UnaryServerInfo, data []byte) (p []byte, st *status.Status) {
	return nil, nil
}
