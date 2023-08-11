package config

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-paths", "config/config.toml", "")
}

type Config struct {
	HTTPAddr string `toml:"bind_addr"`
}

func NewConfig() (*Config, error) {
	config := &Config{
		HTTPAddr: ":8080",
	}
	flag.Parse()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return config, nil
}
