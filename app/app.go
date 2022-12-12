package app

import (
	"github.com/cesarvspr/geot/app/stream"
	"github.com/nats-io/nats.go"
)

// Container for exporting app
type Container struct {
	Streams stream.App
}

// Instance creation struct
type Options struct {
	QueueMessager *nats.Conn
}

// Creates new instance of services
func New(opts Options) *Container {
	container := &Container{
		Streams: stream.NewApp(),
	}
	return container
}
