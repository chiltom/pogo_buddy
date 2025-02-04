package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chiltom/pogo_buddy/db"
	"github.com/chiltom/pogo_buddy/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables
	env := "development"
	if err := godotenv.Load(".env." + env); err != nil {
		log.Fatal("Error loading .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")

	// Initialize db connection
	dbConn := db.ConnectDB()
	defer db.CloseDB()

	// Check the connection
	if err := dbConn.Ping(context.Background()); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	fmt.Println("Application started...")

	// db.RunMigrations()

	// Create a router
	router := mux.NewRouter()

	// Register routes
	routes.AuthRoutes(router)

	log.Println("Server is running on port" + serverPort)
	http.ListenAndServe(serverPort, router)
}
