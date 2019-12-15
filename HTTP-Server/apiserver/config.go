package apiserver

import "server/store"

type Config struct {
	Port string
	LogLevel string
	Store *store.Config
}

func NewConfig() *Config {
	return &Config{
		Port: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}