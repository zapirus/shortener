package main

import (
	"log"

	"gitlab.com/zapirus/shortener/config"
	"gitlab.com/zapirus/shortener/internal/pkg/app"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := app.New(conf)
	server.Run()
}
