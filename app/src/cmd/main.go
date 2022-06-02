package main

import (
	"app/src/core/env"
	"fmt"
	"log"

	"app/src/api/http/app"
)

func init() {

}

func main() {
	srv := app.NewApplication()
	err := srv.Start(fmt.Sprintf("%s:%s", env.AppHost, env.AppPort))
	if err != nil {
		log.Fatal(err)
	}
}
