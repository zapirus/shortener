package main

import (
	"flag"
	"fmt"

	"github.com/BurntSushi/toml"

	"gitlab.com/zapirus/shortener/internal/handlers"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-paths", "config/config.toml", "")
}

func main() {
	flag.Parse()
	config := handlers.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	server := handlers.New(config)

	server.Run()

}
