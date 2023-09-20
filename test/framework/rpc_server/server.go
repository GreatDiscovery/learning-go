package rpc_server

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
)

type Server struct {
	config      *serverConfig
	mu          sync.Mutex
	listeners   map[net.Listener]struct{}
	connections map[*serverConn]struct{} // all connections to current state
	done        chan struct{}            // marks point at which we stop serving requests
}

type serverConn struct {
	server    *Server
	conn      net.Conn
	handshake interface{}
	state     atomic.Value

	shutdownOnce sync.Once     // 只关闭一次
	shutdown     chan struct{} // forced shutdown, used by close
}

func NewServer() (*Server, error) {
	config := &serverConfig{}
	return &Server{
		config:      config,
		mu:          sync.Mutex{},
		listeners:   make(map[net.Listener]struct{}),
		connections: make(map[*serverConn]struct{}),
		done:        make(chan struct{}),
	}, nil
}

func (s *Server) Server(ctx context.Context, listener net.Listener) {

}
