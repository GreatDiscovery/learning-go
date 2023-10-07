package rpc_server

import (
	"context"
	"net"
)

type serverConfig struct {
	handshaker  Handshaker
	interceptor UnaryServerInterceptor
}

type UnaryServerInterceptor func(context.Context, Unmarshaler, *UnaryServerInfo, Method) (interface{}, error)

// Unmarshaler contains the server request data and allows it to be unmarshaled
// into a concrete type
type Unmarshaler func(interface{}) error

// UnaryServerInfo provides information about the server request
type UnaryServerInfo struct {
	FullMethod string
}

type Method func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error)

type Handshaker interface {
	// Handshake should confirm or decorate a connection that may be incoming
	// to a server or outgoing from a client.
	//
	// If this returns without an error, the caller should use the connection
	// in place of the original connection.
	//
	// The second return value can contain credential specific data, such as
	// unix socket credentials or TLS information.
	//
	// While we currently only have implementations on the server-side, this
	// interface should be sufficient to implement similar handshakes on the
	// client-side.
	Handshake(ctx context.Context, conn net.Conn) (net.Conn, interface{}, error)
}
