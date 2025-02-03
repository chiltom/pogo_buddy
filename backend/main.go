package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/chiltom/pogo_buddy/db"
)

func main() {
	// Initialize db connection
	dbConn := db.ConnectDB()
	defer db.CloseDB()

	// Check the connection
	if err := dbConn.Ping(context.Background()); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	fmt.Println("Application started...")

	// db.RunMigrations()

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("Shutting down application...")
}
