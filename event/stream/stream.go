package stream

import (
	"context"

	"github.com/cesarvspr/geot/app"
	"github.com/nats-io/nats.go"
)

const consumerGroupName = "realtime"

type event struct {
	apps *app.Container
	qm   *nats.Conn
	ctx  context.Context
}

func Register(apps *app.Container, queueMessager *nats.Conn) {
	e := &event{
		apps: apps,
		qm:   queueMessager,
		ctx:  context.Background(),
	}
	e.qm.Subscribe(consumerGroupName, e.sendToProcessQueue)
}

func (e *event) sendToProcessQueue(m *nats.Msg) {
	e.apps.Streams.Queue(string(m.Data))
}
