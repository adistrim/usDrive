package config

import (
	"github.com/joho/godotenv"
	"os"
	"log"
)

type Config struct {
	Port string
}

var ENV *Config

func init() {
	var err error
	ENV, err = LoadConfig()
	if err != nil {
		log.Fatalf("Error: Failed to load configuration. %v", err)
	}
}

func LoadConfig() (*Config, error) {
	godotenv.Load()
	
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("No PORT environment variable found, using default port 8080")
		port = "8080"
	}
	
	return &Config{
		Port: port,
	}, nil
}

