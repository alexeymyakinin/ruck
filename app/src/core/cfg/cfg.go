package cfg

import (
	"github.com/caarlos0/env/v6"
)

var Config *config

func init() {
	cfg := new(config)
	err := env.Parse(cfg)
	if err != nil {
		panic(err)
	}
	Config = cfg
}

type config struct {
	AppHost         string   `env:"APP_HOST" envDefault:"0.0.0.0"`
	AppPort         string   `env:"APP_PORT" envDefault:"8000"`
	AppCORS         []string `env:"APP_CORS" envDefault:"0.0.0.0"`
	AppMode         string   `env:"APP_MODE" envDefault:"release"`
	DBConnectionURL string   `env:"DB_URL"   envDefault:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
}
