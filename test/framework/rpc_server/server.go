package rpc_server

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"sync"
	"sync/atomic"
)

const (
	connStateActive = iota + 1 // outstanding requests
	connStateIdle              // no requests
	connStateClosed            // closed connection
)

var (
	ErrServerClosed = errors.New("server closed")
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

func (s *Server) Server(ctx context.Context, listener net.Listener) error {
	s.addListener(listener)
	defer s.closeListener(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-s.done:
				return ErrServerClosed
			default:
			}
		}
		sc, err := s.newConn(conn)
		go sc.run(ctx)
	}
}

func (s *Server) addListener(l net.Listener) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.listeners[l] = struct{}{}
}

func (s *Server) closeListener(l net.Listener) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	defer delete(s.listeners, l)
	err := l.Close()
	return err
}

func (s *Server) newConn(conn net.Conn) (*serverConn, error) {
	c := &serverConn{
		server:    s,
		conn:      conn,
		handshake: nil,
		state:     atomic.Value{},
		shutdown:  make(chan struct{}),
	}
	c.state.Store(connStateIdle)
	err := s.addConnection(c)
	if err != nil {
		c.close()
		return nil, err
	}
	return c, nil
}

func (s *Server) addConnection(conn *serverConn) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	select {
	case <-s.done:
		return ErrServerClosed
	default:
	}

	s.connections[conn] = struct{}{}
	return nil
}

func (s *Server) deleteConnection(conn *serverConn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.connections, conn)
}

func (c *serverConn) close() error {
	c.shutdownOnce.Do(
		func() {
			close(c.shutdown)
		})
	return nil
}

func (c *serverConn) run(sctx context.Context) error {
	type (
		response struct {
			id          uint32
			status      *status.Status
			data        []byte
			closeStream bool
			streaming   bool
		}
	)

	var (
		ch        = newChannel(c.conn)
		_, cancel = context.WithCancel(sctx)
		done      = make(chan struct{})
		responses = make(chan response)
		recvErr   = make(chan error, 1)
	)

	defer c.close()
	defer cancel()
	defer close(done)
	defer c.server.deleteConnection(c)

	sendStatus := func(id uint32, st *status.Status) bool {
		select {
		case responses <- response{
			id:          id,
			status:      st,
			closeStream: true,
		}:
			return true
		case <-c.shutdown:
			return false
		case <-done:
			return false
		}
	}

	go func(recvErr chan error) {
		defer close(recvErr)
		for {
			select {
			case <-c.shutdown:
				return
			case <-done:
				return
			default:
			}

			ch.recv()
			if !sendStatus(0, status.Newf(codes.InvalidArgument, "StreamID must be odd for client initiated streams")) {
				return
			}
		}
	}(recvErr)

	return nil
}
