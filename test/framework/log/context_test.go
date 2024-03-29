package log

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintInfo(t *testing.T) {
	ctx := context.Background()
	ctx = WithLogger(ctx, G(ctx).WithField("test", "one"))
	G(ctx).Info("hello world")
}

func TestLoggerContext(t *testing.T) {
	ctx := context.Background()
	ctx = WithLogger(ctx, G(ctx).WithField("test", "one"))
	assert.Equal(t, GetLogger(ctx).Data["test"], "one")
	assert.Same(t, G(ctx), GetLogger(ctx)) // these should be the same.
}
