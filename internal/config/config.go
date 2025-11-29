package config

import (
	"os"
	"sync"
)

type Config struct {
	Port        string
	DatabaseURL string
}

var (
	once            sync.Once
	configSingleton *Config
)

func GetConfig() Config {
	once.Do(func() {
		configSingleton = &Config{
			Port: "8080",
		}

		if envPort := os.Getenv("PORT"); envPort != "" {
			configSingleton.Port = envPort
		}

		if envDBURL := os.Getenv("DB_URL"); envDBURL == "" {
			panic("No url found")
		} else {
			configSingleton.DatabaseURL = envDBURL
		}
	})
	return *configSingleton
}
