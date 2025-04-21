package config

import (
	"github.com/caarlos0/env"
)

type ServerConfig struct {
    Host    			string  `env:"HOST"  envDefault:"0.0.0.0"`
    Port    			int     `env:"PORT"  envDefault:"8080"`
    Debug   			bool    `env:"DEBUG" envDefault:"false"`

	JwtKey				string  `env:"JWT_KEY,required"`
	GithubClientId		string	`env:"GITHUB_ID,required"`
	GithubClientSecret	string	`env:"GITHUB_SECRET,required"`
}

type DatabaseConfig struct {
    DBName  string  `env:"DB_DATABASE"  envDefault:"sherpa"`
    Host    string  `env:"DB_HOST"      envDefault:"localhost"`
    Port    int     `env:"DB_PORT"      envDefault:"5432"`
    User    string  `env:"DB_USERNAME,required"`
    Pass    string  `env:"DB_PASSWORD,required"`
}

func NewServerConfig() ServerConfig {
    cfg := ServerConfig{}
    if err := env.Parse(&cfg) ; err != nil {
		panic(err.Error());
    }
    return cfg
}

func NewDatabaseConfig() DatabaseConfig {
    cfg := DatabaseConfig{}
    if err := env.Parse(&cfg) ; err != nil {
        panic(err.Error());
    }
    return cfg
}
