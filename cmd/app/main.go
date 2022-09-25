package main

import (
	"log"
	"os"

	"github.com/Rubemlrm/go-gin-sanbox/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	log.Println(os.Getenv("GIN_MODE"))

	controller.testing()
	return router
}
