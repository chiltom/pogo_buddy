package main

import (
  "log"
  "net/http"
  "os"

  "github.com/chiltom/pogo_buddy/internal/db"
  "github.com/chiltom/pogo_buddy/internal/handlers"
)

func main() {
  cfg := db.DbConfig{
    Host:     os.Getenv("DB_HOST"),
    Port:     os.Getenv("DB_PORT"),
    User:     os.Getenv("DB_USER"),
    Password: os.Getenv("DB_PASSWORD"),
    DBName:   os.Getenv("DB_NAME"),
    SSLMode: "disable",
  }

  dbConn, err := db.New(cfg)
}
