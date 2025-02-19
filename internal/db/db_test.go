package db

import (
	"os"
	"testing"

	"github.com/chiltom/pogo_buddy/internal/utils"
)

// TestConnect checks if the database connection function works correctly
func TestConnect(t *testing.T) {
	utils.LoadEnv("../../.env")

	os.Getenv("DB_HOST")
	os.Getenv("DB_PORT")
	os.Getenv("DB_USER")
	os.Getenv("DB_PASSWORD")
	os.Getenv("DB_NAME")

	db, err := Connect()
	if err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	if db == nil {
		t.Fatal("DB is nil")
	}
}
