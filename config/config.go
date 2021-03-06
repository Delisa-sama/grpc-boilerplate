package config

import (
	"fmt"

	"github.com/Delisa-sama/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config structure
type Config struct {
	DB          DBConfig
	LogPath     string          `env:"LOG_PATH" env-description:"Path to log file. If putted filename only, log file will be created in the same directory as the binary" env-default:""`
	LogLevel    logger.LogLevel `env:"LOG_LEVEL" env-description:"Logging level [FATAL, ERROR, WARN, DEBUG, INFO, STASH]" env-default:"WARN"`
	Name        string          `env:"APP_NAME" env-description:"Name of service" env-default:"grpc-boilerplate"`
	Version     string          `env:"APP_VERSION" env-description:"Version of service" env-default:"1.0.0"`
	Environment string          `env:"APP_ENV" env-description:"Environment of service" env-required:"true"`
	Host        string          `env:"APP_HOST" env-description:"Address or addr-pattern what will listen to the application" env-default:""`
	Port        string          `env:"APP_PORT" env-description:"Application port" env-default:"8080"`
}

// DBConfig contain main fields for connect to DB
type DBConfig struct {
	Dialect  string `env:"DB_DIALECT" env-description:"The SQL dialect of your DBMS" env-default:"postgres"`
	Host     string `env:"DB_HOST" env-description:"DB address" env-required:"true"`
	Port     string `env:"DB_PORT" env-description:"DB port" env-required:"true"`
	Username string `env:"DB_USER" env-description:"DBMS name of the role" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-description:"DBMS auth password" env-required:"true"`
	Name     string `env:"DB_NAME" env-description:"Name of the database" env-required:"true"`
	Charset  string `env:"DB_CHARSET" env-description:""`
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
