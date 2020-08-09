package logger

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// Logger ...
type Logger interface {
	Log(kind string, message string, err error)
}

type svcLogger struct {
	logger log.Logger
}

// Log ...
func (l *svcLogger) Log(kind string, message string, err error) {
	if err != nil {
		level.Error(l.logger) // TODO
	} else {
		level.Debug(l.logger)
	}
	_ = l.logger.Log("kind", kind, "message", message, "error", err)
}

// New ...
func New(logger log.Logger) Logger {
	return &svcLogger{
		logger: logger,
	}
}
