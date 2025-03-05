package db

import (
  "database/sql"
  "fmt"
  "log"

  _ "github.com/lib/pq"
)

type DbConfig struct {
  Host     string
  Port     string
  User     string
  Password string
  DBName   string
  SSLMode  string
}

type DB struct {
  *sql.DB
}

func New(cfg DbConfig) (*DB, error) {
  connStr := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
  )

  db, err := sql.Open("postgres", connStr)
  if err != nil {
    return nil, fmt.Errorf("failed to open db connection: %w", err)
  }

  if err := db.Ping(); err != nil {
    db.Close()
    return nil, fmt.Errorf("failed to ping db: %w", err)
  }

  log.Println("connected to db")
  return &DB{db}, nil
}

func (d *DB) Close() error {
  if d.DB != nil {
    err := d.DB.Close()
    if err != nil {
      log.Printf("failed to close db: %v", err)
      return fmt.Errorf("failed to close db: %w", err)
    }
    log.Println("close db connection")
  }
  return nil
}
