package event

import (
	"context"

	"github.com/cesarvspr/geot/app"
	"github.com/cesarvspr/geot/event/stream"
	"github.com/cesarvspr/geot/logger"
	"github.com/nats-io/nats.go"
)

// Here we can set register options
type Options struct {
	Apps          *app.Container
	QueueMessager *nats.Conn
}

// Register handler instance
func Register(opts Options) {
	stream.Register(opts.Apps, opts.QueueMessager)
	logger := logger.FromContext(context.Background())
	logger.Info("Registered EVENT")
}
