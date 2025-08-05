package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DatabaseURL string
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
	
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}
	
	return &Config{
		Port: port,
		DatabaseURL: databaseURL,
	}, nil
}

