package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"learning-go/test/framework/log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestServerStart(t *testing.T) {
	l, err := serveListener("")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	go func() {
		defer l.Close()
		if err := serverStart(ctx, l); err != nil && !errors.Is(err, net.ErrClosed) {
			log.G(ctx).WithError(err).Fatal("containerd-shim: ttrpc rpc_server failure")
		}
	}()

	ctx, cancel := context.WithCancel(ctx)
	go handleExitSignals(ctx, log.G(ctx), cancel)
}

func serveListener(path string) (net.Listener, error) {
	var (
		l   net.Listener
		err error
	)
	if path == "" {
		l, err = net.FileListener(os.NewFile(3, "socket"))
		path = "[inherited from parent]"
	} else {
		if len(path) > 106 {
			return nil, fmt.Errorf("%q: unix socket path too long (> %d)", path, 106)
		}
		l, err = net.Listen("unix", path)
	}
	if err != nil {
		return nil, err
	}
	logrus.WithField("socket", path).Debug("serving api on socket")
	return l, nil
}

func serverStart(ctx context.Context, l net.Listener) error {
	for {

	}
}

func handleExitSignals(ctx context.Context, logger *logrus.Entry, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 32)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case s := <-ch:
			logger.WithField("signal", s).Debugf("Caught exit signal")
			cancel()
			return
		case <-ctx.Done():
			return
		}
	}
}
