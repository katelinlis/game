package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	UrlDB         string `env:"DB_URL"`
	Redis         string `env:"REDIS_URL"`
	RedisPassword string `env:"REDIS_PASSWORD"`
}

func InitConfig() Config {
	// parse

	var flagDontloaddotent bool

	if err := godotenv.Load(); err != nil {
		flagDontloaddotent = true
	}

	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		if flagDontloaddotent {
			log.Println("Error loading .env file")
			panic(err)
		}

	}

	// parse with generics
	cfg, err = env.ParseAs[Config]()
	if err != nil {
		panic(err)
	}

	return cfg
}
