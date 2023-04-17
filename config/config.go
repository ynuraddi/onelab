package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP struct {
		Host string `env:"HOST" env-default:"localhost"`
		Port string `env:"PORT" env-default:"8080"`
	}

	JWTKey string `env:"JWTKey" env-default:"secret"`

	Database struct {
		Host   string `env:"DB_HOST" env-default:"localhost"`
		Port   string `env:"DB_PORT" env-default:"5432"`
		User   string `env:"DB_USER" env-default:"onelab"`
		Pass   string `env:"DB_PASSWORD" env-default:"onelab"`
		DBName string `env:"DB_DATABASE" env-default:"onelab"`
	}
}

var (
	config Config
	once   sync.Once
)

func GetConfing() *Config {
	once.Do(func() {
		if err := cleanenv.ReadConfig(".env", &config); err != nil {
			log.Fatalln(err)
		}
		log.Println("configs readed")
	})

	return &config
}
