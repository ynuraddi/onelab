package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port string `yaml:"port" env:"PORT" env-default:"8080"`
}

var (
	config Config
	once   sync.Once
)

func GetConfing() *Config {
	once.Do(func() {
		cleanenv.ReadConfig(".env", &config)
		log.Println("logs readed")
	})

	return &config
}
