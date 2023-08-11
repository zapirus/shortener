package config

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	HTTPAddr string `toml:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{
		HTTPAddr: ":8080",
	}
}

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-paths", "config/config.toml", "")
}

func (c *Config) ConfigPars(config *Config) error {
	flag.Parse()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
