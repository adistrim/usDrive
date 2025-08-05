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
	R2AccessKeyID string
	R2SecretAccessKey string
	R2Bucket string
	R2Region string
	R2API string
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
	
	// Object Storage environment variables
	r2AccessKeyID := os.Getenv("R2_ACCESS_KEY_ID")
	r2SecretAccessKey := os.Getenv("R2_SECRET_ACCESS_KEY")
	r2Bucket := os.Getenv("R2_BUCKET")
	r2Region := os.Getenv("R2_REGION")
	r2API := os.Getenv("R2_API")
	if r2AccessKeyID == "" || r2SecretAccessKey == "" || r2Bucket == "" || r2Region == "" || r2API == "" {
		return nil, fmt.Errorf("R2 environment variables are not set")
	}
	
	return &Config{
		Port: port,
		DatabaseURL: databaseURL,
		R2AccessKeyID: r2AccessKeyID,
		R2SecretAccessKey: r2SecretAccessKey,
		R2Bucket: r2Bucket,
		R2Region: r2Region,
		R2API: r2API,
	}, nil
}

