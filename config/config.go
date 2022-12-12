package config

import (
	"fmt"
	"os"
)

// Nats is a struct to use in config
type Nats struct {
	URL string `mapstructure:"url"`
}

// Config is a struct to use in var ConfigGlobal
type Config struct {
	ENV   string `mapstructure:"env"`
	Nats  Nats   `mapstructure:"nats"`
}

// ConfigGlobal is you use in all app
var ConfigGlobal *Config = &Config{}

func init() {
	//load config
	//TODO: use viper here to load and decode json config
	temp := os.Getenv("NATS_URL")
	fmt.Println("NATS_URL: ", temp)
	ConfigGlobal.Nats.URL = os.Getenv("NATS_URL")
}
