package server

import (
	"github.com/cesarvspr/geot/app"
	"github.com/cesarvspr/geot/config"
	_ "github.com/cesarvspr/geot/config"
	"github.com/cesarvspr/geot/event"
	"github.com/nats-io/nats.go"
)

// Server is a interface to define contract to server up
type Server interface {
	Start()
	Stop()
	ReloadConnections()
	Execute()
}

type server struct {
	queueMessager *nats.Conn
	app           *app.Container
}

// New is instance the server
func New() Server {
	return &server{}
}

func (s *server) Start() {

	// ---- start NATS connection ----
	nc, err := nats.Connect(config.ConfigGlobal.Nats.URL)
	if err != nil {
		panic(err)
	}
	s.queueMessager = nc

	// ---- setup App ----
	s.app = app.New(app.Options{
		QueueMessager: s.queueMessager,
	})

	// ---- setup Event ----
	event.Register(event.Options{
		Apps:          s.app,
		QueueMessager: s.queueMessager,
	})

}

func (s *server) Stop() {
	s.queueMessager = nil
}

// ReloadConnections all connections like DB, Nats, ...
func (s *server) ReloadConnections() {
	s.Stop()
	s.Start()
}

// Process that handle with all lists
func (s *server) Execute() {
	s.app.Streams.Cron()

	go s.app.Streams.Process()
	s.app.Streams.Cron()
}
