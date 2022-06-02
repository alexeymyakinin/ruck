package main

import (
	"fmt"
	"github.com/alexeymyakinin/ruck/app/src/api/http/app"
	"github.com/alexeymyakinin/ruck/app/src/core/env"
	"log"
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
