package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

//Config for loading toml documentation
type Config struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBname   string `toml:"dbname"`
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)

	}
}

var config Config

func init() {
	config.Read()
}

func GetConfig() *Config {
	return &config
}
