package main

import (
	"fmt"
	"log"

	"app/src/core/cfg"
	"app/src/web/http/app"
)

func init() {

}

func main() {
	srv := app.NewApplication()
	err := srv.Run(fmt.Sprintf("%s:%s", cfg.Config.AppHost, cfg.Config.AppPort))
	if err != nil {
		log.Fatal(err)
	}
}
