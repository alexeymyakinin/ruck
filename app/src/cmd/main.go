package main

import (
	"fmt"
	"log"

	"app/src/api/http/app"
	"app/src/core/cfg"
)

func init() {

}

func main() {
	srv := app.NewApplication()
	err := srv.Start(fmt.Sprintf("%s:%s", cfg.Config.AppHost, cfg.Config.AppPort))
	if err != nil {
		log.Fatal(err)
	}
}
