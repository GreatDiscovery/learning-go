package log

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

var (
	// G is an alias for GetLogger.
	//
	// We may want to define this locally to a package to get package tagged log
	// messages.
	G = GetLogger

	// L is an alias for the standard logger.
	L = logrus.NewEntry(logrus.StandardLogger())
)

type (
	loggerKey struct{}

	// Fields type to pass to `WithFields`, alias from `logrus`.
	Fields = logrus.Fields

	// Level is a logging level
	Level = logrus.Level
)

const (
	// RFC3339NanoFixed is time.RFC3339Nano with nanoseconds padded using zeros to
	// ensure the formatted time is always the same number of characters.
	RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"

	// TextFormat represents the text logging format
	TextFormat = "text"

	// JSONFormat represents the JSON logging format
	JSONFormat = "json"

	// TraceLevel level.
	TraceLevel = logrus.TraceLevel

	// DebugLevel level.
	DebugLevel = logrus.DebugLevel

	// InfoLevel level.
	InfoLevel = logrus.InfoLevel
)

// SetLevel sets log level globally.
func SetLevel(level string) error {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}

	logrus.SetLevel(lvl)
	return nil
}

// GetLevel returns the current log level.
func GetLevel() Level {
	return logrus.GetLevel()
}

// SetFormat sets log output format
func SetFormat(format string) error {
	switch format {
	case TextFormat:
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: RFC3339NanoFixed,
			FullTimestamp:   true,
		})
	case JSONFormat:
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: RFC3339NanoFixed,
		})
	default:
		return fmt.Errorf("unknown log format: %s", format)
	}

	return nil
}

// WithLogger returns a new context with the provided logger. Use in
// combination with logger.WithField(s) for great effect.
func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	e := logger.WithContext(ctx)
	return context.WithValue(ctx, loggerKey{}, e)
}

// GetLogger retrieves the current logger from the context. If no logger is
// available, the default logger is returned.
func GetLogger(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})

	if logger == nil {
		return L.WithContext(ctx)
	}

	return logger.(*logrus.Entry)
}
