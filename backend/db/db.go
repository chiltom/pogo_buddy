package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// Global DB pool instance
var (
	pool *pgxpool.Pool
	once sync.Once
)

// ConnectDB initializes and returns a database connection pool
func ConnectDB() *pgxpool.Pool {
	once.Do(func() {
		fmt.Println("Connecting to PostgreSQL...")

		env := "development"

		if err := godotenv.Load(".env." + env); err != nil {
			log.Fatal("Error loading .env file")
		}

		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			dbURL = "postgres://user:password@localhost:5432/dbname"
		}

		var err error
		pool, err = pgxpool.New(context.Background(), dbURL)
		if err != nil {
			log.Fatalf("Unable to create connection pool: %v", err)
		}

		fmt.Println("Connected to PostgreSQL successfully")
	})

	return pool
}

// GetDB returns the initialized database connection pool
func GetDB() *pgxpool.Pool {
	if pool == nil {
		log.Fatal("Database connection pool is not initialized. Call ConnectDB() first.")
	}
	return pool
}

// CloseDB closes the database connection pool
func CloseDB() {
	if pool != nil {
		pool.Close()
		fmt.Println("Database connection pool closed.")
	}
}
