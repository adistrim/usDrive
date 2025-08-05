package db

import (
	"log"
	"usdrive/config"

	"github.com/jackc/pgx/v5/pgxpool"

	"context"
	"sync"
)

var dbInstance *pgxpool.Pool

var once sync.Once

func GetDBInstance() *pgxpool.Pool {
	
	once.Do(func(){
		connStr := config.ENV.DatabaseURL
		var err error
		
		config, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			log.Fatalf("Error: Unable to parse database URL: %v", err)
		}
		
		dbInstance, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("Error: Unable to connect to database: %v", err)
		}
		
		log.Println("Database connection pool created successfully")
	})
	
	return dbInstance
}
