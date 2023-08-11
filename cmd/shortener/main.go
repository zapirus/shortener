package main

import (
	"fmt"

	"gitlab.com/zapirus/shortener/config"
	"gitlab.com/zapirus/shortener/internal/pkg/app"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	server := app.New(conf)
	server.Run()
}
