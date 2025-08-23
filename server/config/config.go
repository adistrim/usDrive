package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	AllowedOrigins []string
	DatabaseURL string
	R2AccessKeyID string
	R2SecretAccessKey string
	R2Bucket string
	R2Region string
	R2API string
	GClientID string
	JWTSecret string
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
	allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
	var allowedOrigins []string
	if allowedOriginsEnv == "" {
			log.Println("Please make sure to put the allowed origins, going with default: localhost:3000")
			allowedOrigins = []string{"http://localhost:3000"}
	} else {
			allowedOrigins = strings.Split(allowedOriginsEnv, ",")
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
	
	gClientID := os.Getenv("GOOGLE_CLIENT_ID")
	if gClientID == "" {
		return nil, fmt.Errorf("Google Client ID missing from environment variables")
	}
	
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable not set")
	}
	
	return &Config{
		Port: port,
		AllowedOrigins: allowedOrigins,
		DatabaseURL: databaseURL,
		R2AccessKeyID: r2AccessKeyID,
		R2SecretAccessKey: r2SecretAccessKey,
		R2Bucket: r2Bucket,
		R2Region: r2Region,
		R2API: r2API,
		GClientID: gClientID,
		JWTSecret: jwtSecret,
	}, nil
}

