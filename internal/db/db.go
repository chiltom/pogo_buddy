package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/chiltom/pogo_buddy/internal/utils"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// getDBConnString constructs the Postgres connection string from env variables
func getDBConnString() string {
	utils.LoadEnv("../../.env")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}

// Connect to the Postgres database
func Connect() (*sql.DB, error) {
	var err error
	once.Do(func() {
		connStr := getDBConnString()
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Failed to open DB connection: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to ping DB: %v", err)
		}

		log.Println("Connected to DB")
	})

	return db, err
}

// Close the database connection
func Close() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Fatalf("Failed to close DB: %v", err)
		} else {
			log.Println("Closed DB connection")
		}
	}
}
