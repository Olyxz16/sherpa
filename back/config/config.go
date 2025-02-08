package config

import (
	"github.com/caarlos0/env"
)

type ServerConfig struct {
    Host    string  `env:"HOST" default:"0.0.0.0"`
    Port    int     `env:"PORT" default:"8080"`
}

type DatabaseConfig struct {
    DBName  string  `env:"DB_DATABASE" default:"sherpa"`
    Host    string  `env:"DB_HOST" default:"localhost"`
    Port    int     `env:"DB_PORT" default:"5432"`
    User    string  `env:"DB_USERNAME"`
    Pass    string  `env:"DB_PASSWORD"`
}

func NewServerConfig() (*ServerConfig, error) {
    cfg := &ServerConfig{}
    if err := env.Parse(cfg) ; err != nil {
        return nil, err
    }
    return cfg, nil
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
    cfg := &DatabaseConfig{}
    if err := env.Parse(cfg) ; err != nil {
        return nil, err
    }
    return cfg, nil
}
