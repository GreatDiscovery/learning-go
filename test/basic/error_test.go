package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"k8s.io/klog/v2"
	"net/http"
	"runtime"
	"testing"
)

var (
	// ReallyCrash controls the behavior of HandleCrash and defaults to
	// true. It's exposed so components can optionally set to false
	// to restore prior behavior. This flag is mostly used for tests to validate
	// crash conditions.
	ReallyCrash = true
)

var PanicHandlers = []func(interface{}){logPanic}

func HandleCrash(additionalHandlers ...func(interface{})) {
	if r := recover(); r != nil {
		for _, fn := range PanicHandlers {
			fn(r)
		}
		for _, fn := range additionalHandlers {
			fn(r)
		}
		if ReallyCrash {
			panic(r)
		}
	}
}

// logPanic logs the caller tree when a panic occurs (except in the special case of http.ErrAbortHandler).
func logPanic(r interface{}) {
	if r == http.ErrAbortHandler {
		// honor the http.ErrAbortHandler sentinel panic value:
		//   ErrAbortHandler is a sentinel panic value to abort a handler.
		//   While any panic from ServeHTTP aborts the response to the client,
		//   panicking with ErrAbortHandler also suppresses logging of a stack trace to the server's error log.
		return
	}

	// Same as stdlib http server code. Manually allocate stack trace buffer size
	// to prevent excessively large logs
	const size = 64 << 10
	stacktrace := make([]byte, size)
	stacktrace = stacktrace[:runtime.Stack(stacktrace, false)]
	if _, ok := r.(string); ok {
		klog.Errorf("Observed a panic: %s\n%s", r, stacktrace)
	} else {
		klog.Errorf("Observed a panic: %#v (%v)\n%s", r, r, stacktrace)
	}
}

func TestPanic(t *testing.T) {
	HandleCrash()
	panic(errors.New("manual error"))
}

func TestPrintErrorStack(t *testing.T) {
	err := errors.New("hello world")
	// 这里err没有携带错误堆栈
	log.Error("print stack,", err)
	fmt.Printf("err=%v", err)
	panic(err)
}
