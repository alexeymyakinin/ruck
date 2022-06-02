package env

import (
	"github.com/gobuffalo/envy"
	"github.com/labstack/gommon/log"
)

func init() {
	if err := envy.Load(); err != nil {
		log.Error(err)
	}
}

var (
	DBConnectionURL = envy.Get("RUCK_DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	AppHost         = envy.Get("RUCK_APP_HOST", "0.0.0.0")
	AppPort         = envy.Get("RUCK_APP_PORT", "8000")
	JWTSecret       = envy.Get("RUCK_JWT_SECRET", "secret")
)
