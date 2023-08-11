package main

import (
	"flag"
	"fmt"

	"gitlab.com/zapirus/shortener/config"
	"gitlab.com/zapirus/shortener/internal/pkg/app"
)

func main() {
	flag.Parse()
	conf := config.NewConfig()
	err := conf.ConfigPars(conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	server := app.New(conf)
	server.Run()

}
