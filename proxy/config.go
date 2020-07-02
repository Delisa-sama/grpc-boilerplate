package main

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config structure
type Config struct {
	Entrypoint EntrypointConfig
	Endpoint   EndpointConfig
}

type EntrypointConfig struct {
	Host string `env:"ENTRYPOINT_HOST" env-description:"Address or addr-pattern what will listen to the proxy" env-default:""`
	Port string `env:"ENTRYPOINT_PORT" env-description:"Application port" env-default:"8080"`
}

type EndpointConfig struct {
	Host string `env:"ENDPOINT_HOST" env-description:"Address of gRPC server (endpoint)" env-required:"true"`
	Port string `env:"ENDPOINT_PORT" env-description:"Application port" env-required:"true"`
}

// GetConfig initial configuration from env
func GetConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		envHelp, _ := cleanenv.GetDescription(&cfg, nil)
		fmt.Println(envHelp)
		return nil, err
	}

	return &cfg, nil
}
