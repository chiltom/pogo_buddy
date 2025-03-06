package main

import (
  "log"
  "net/http"
  "os"

  "github.com/chiltom/pogo_buddy/internal/db"
  "github.com/chiltom/pogo_buddy/internal/handlers"
  "github.com/chiltom/pogo_buddy/internal/utils"
)

func main() {
  utils.LoadEnv(".env")

  cfg := db.DbConfig{
    Host:     os.Getenv("DB_HOST"),
    Port:     os.Getenv("DB_PORT"),
    User:     os.Getenv("DB_USER"),
    Password: os.Getenv("DB_PASSWORD"),
    DBName:   os.Getenv("DB_NAME"),
    SSLMode: "disable",
  }

  dbConn, err := db.New(cfg)
  if err != nil {
    log.Fatalf("Failed to connect to database: %v", err)
  }
  defer dbConn.Close()

  h := handlers.New(dbConn)

  http.HandleFunc("/users/create", h.User.CreateUser)

  log.Println("Starting server on :8080")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatalf("Server failed: %v", err)
  }
}
