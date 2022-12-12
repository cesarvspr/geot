package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cesarvspr/geot/app/stream"
	_ "github.com/cesarvspr/geot/config"
	"github.com/cesarvspr/geot/logger"
	"github.com/cesarvspr/geot/server"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

// Create channel to listen for signals.
var c chan (os.Signal) = make(chan os.Signal)
var exit chan bool = make(chan bool)

func main() {
	s := server.New()
	s.Start()
	defer s.Stop()
	s.Execute()

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGHUP)

	go func() {
		for {
			// Of course, this nats publishing is just for development purposes. 
			nc, err := nats.Connect("localhost:4222")
			if err != nil {
				fmt.Println(err)
			}

			go func() {
				for {
					time.Sleep(time.Second / 2)
					nc.Publish("realtime", []byte(generateJWT()))
				}
			}()

			signal := <-c
			switch signal {
			case syscall.SIGINT:
				exit <- true
				os.Exit(0)
			case syscall.SIGTERM:
				s.Stop()
				exit <- true
				os.Exit(143)
			case syscall.SIGHUP:
				s.ReloadConnections()
				exit <- false
			}
		}
	}()

	logger := logger.FromContext(context.Background())
	logger.Info("Service started.")
	<-exit
}

func generateJWT() string {
	// Create token
	id := uuid.New()
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer:    "test",
		Id:        id.String(),
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	jwt, err := token.SignedString(stream.SecretKey)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return jwt
}
