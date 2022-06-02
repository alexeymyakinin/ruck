package env

import (
	"github.com/gobuffalo/envy"
	"github.com/labstack/gommon/log"
	"os"
	"path"
)

func init() {
	wd, _ := os.Getwd()
	if err := envy.Load(path.Join(path.Dir(wd), ".env")); err != nil {
		log.Error(err)
	}
}

var (
	DBConnectionURL = envy.Get("RUCK_DB_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	AppHost         = envy.Get("RUCK_APP_HOST", "0.0.0.0")
	AppPort         = envy.Get("RUCK_APP_PORT", "8000")
)
