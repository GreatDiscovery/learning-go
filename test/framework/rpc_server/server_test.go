package rpc_server

import (
	"context"
	"errors"
	"net"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type socket string

func (s socket) path() string {
	path := strings.TrimPrefix(string(s), "unix://")
	return path
}

func TestServer(t *testing.T) {
	ctx := context.Background()
	server, err := NewServer()
	errChan := make(chan error, 1)
	if err != nil {
		t.Fatalf("failed to new ttrpc server: %v", err)
	}
	addr, listener, err := NewListener()
	defer listener.Close()
	if err != nil {
		t.Fatalf("failed to new ttrpc server: %v", err)
	}
	go func() {
		server.Server(ctx, listener)
	}()

	go func() {
		to := time.After(10 * time.Second)

		for {
			select {
			case <-to:
				errChan <- errors.New("timeout")
				return
			default:
			}

			conn, err := AnonDialer(addr, 1*time.Hour)
			if err != nil {
				errChan <- err
				return
			}
			conn.Close()
			errChan <- nil
			return
		}
	}()

	// it should be successful
	if err := <-errChan; err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
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

// AnonDialer returns a dialer for a socket
func AnonDialer(address string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout("unix", socket(address).path(), timeout)
}
