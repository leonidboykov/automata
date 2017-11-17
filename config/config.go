package config

import (
	"log"

	"github.com/caarlos0/env"
)

// Config ...
type Config struct {
	BotToken string `env:"AUTOMATA_TOKEN"`
}

// Parse ...
func Parse() Config {
	var conf Config
	if err := env.Parse(&conf); err != nil {
		log.Println(err)
	}
	return conf
}
