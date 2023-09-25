package rpc_server

import (
	"context"
	"net"
	"path/filepath"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	server, err := NewServer()
	if err != nil {
		t.Fatalf("failed to new ttrpc server: %v", err)
	}
	_, listener, err := NewListener()
	defer listener.Close()
	if err != nil {
		t.Fatalf("failed to new ttrpc server: %v", err)
	}
	go func() {
		server.Server(ctx, listener)
	}()
}

func NewListener() (string, net.Listener, error) {
	curDir, err2 := filepath.Abs(".")
	if err2 != nil {
		return "", nil, err2
	}
	addr := "unix://" + filepath.Join(curDir, "uds.socket")
	listener, err := net.Listen("unix", strings.TrimPrefix(addr, "unix://"))
	if err != nil {
		return addr, nil, err
	}
	return addr, listener, nil
}
