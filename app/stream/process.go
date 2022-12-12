package stream

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/cesarvspr/geot/logger"
	"github.com/golang-jwt/jwt"
)

// ----- lock goroutines counter -----
var counterLock sync.RWMutex

var SecretKey = []byte(os.Getenv("JWT_SECRET"))

// App interface
type App interface {
	Process()
	Cron()
	Queue(data string)
	Status()
}

// NewApp create a new APP instance, we can have multiple of this
func NewApp() App {
	return &appImpl{
		stream:  make(chan string),
		counter: 0,
	}
}

type appImpl struct {
	stream  chan string
	counter int
}

// Send message to process
func (s *appImpl) Queue(data string) {
	s.stream <- data
}

// Process incoming data concurrently
func (s *appImpl) Process() {
	for data := range s.stream {
		counterLock.Lock()
		s.counter++
		counterLock.Unlock()
		go WorkWithData(data)
	}
}

func (s *appImpl) Status() {
	logger := logger.FromContext(context.Background())
	counterLock.RLock()
	defer counterLock.RUnlock()
	logger.Info(fmt.Sprintf("[Cron][Goroutines counter -> %d were executed]\n", s.counter))
}

func WorkWithData(data string) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(data, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
