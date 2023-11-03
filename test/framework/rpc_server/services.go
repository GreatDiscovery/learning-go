package rpc_server

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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
	services         map[string]*ServiceDesc
	unaryInterceptor UnaryServerInterceptor
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
	unmarshal := func(obj interface{}) error {
		return protoUnmarshal(data, obj)
	}
	resp, err := s.unaryInterceptor(ctx, unmarshal, info, method)
	if err == nil {
		if resp == nil {
			err = errors.New("ttrpc: marshal called with nil")
		} else {
			p, err = protoMarshal(resp)
		}
	}

	sts, ok := status.FromError(err)
	if !ok {
		sts = status.New(codes.Unknown, err.Error())
	}
	return p, sts
}

func protoUnmarshal(p []byte, obj interface{}) error {
	switch v := obj.(type) {
	case proto.Message:
		if err := proto.Unmarshal(p, v); err != nil {
			return status.Errorf(codes.Internal, "ttrpc: error unmarshalling payload: %v", err.Error())
		}
	default:
		return status.Errorf(codes.Internal, "ttrpc: error unsupported request type: %T", v)
	}
	return nil
}

func protoMarshal(obj interface{}) ([]byte, error) {
	if obj == nil {
		return nil, nil
	}
	switch v := obj.(type) {
	case proto.Message:
		marshal, err := proto.Marshal(v)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "ttrpc: error marshaling payload: %v", err.Error())
		}
		return marshal, nil
	default:
		return nil, status.Errorf(codes.Internal, "ttrpc: error unsupported response type: %T", v)
	}
}
