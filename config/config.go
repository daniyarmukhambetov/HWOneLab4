package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	HTTP     HTTP
	Port     int    `env:"PORT" envDefault:"8080"`
	Database string `env:"DATABASE" envDefault:"easdasd"`
	JWTKey   string `env:"JWT_KEY" evnDefault:"supersecret"`
	PgURL    string `env:"PG_URL" envDefault:"user=admin password=password dbname=onelab2 sslmode=disable host=localhost port=5432" `
}

type HTTP struct {
	PORT   int
	URL    string
	UseSSL bool
}

func New() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
