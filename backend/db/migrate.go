package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations applies database migrations using the open connection pool
// from the db package. Hook this up to admin backend handlers later to handle
// migrations.
func RunMigrations() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
		dbURL = "postgres://user:oassword@localhost:5432/dbname?sslmode=disable"
	}

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Could not acquire database connection: %v", err)
	}
	defer db.Close()

	// Get the raw pgx.Conn from the pool connection
	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		log.Fatalf("Could not create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("Could not create migration instance: %v", err)
	}

	fmt.Println("Applying migrations...")
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migrations applied successfully")
}
