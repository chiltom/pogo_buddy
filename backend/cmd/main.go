package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	env := "development"

	if err := godotenv.Load(".env." + env); err != nil {
		log.Fatal("Error loading .env file")
	}

	db_port := os.Getenv("DB_PORT")

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bing",
		})
	})

	router.Run(":" + db_port)
}
