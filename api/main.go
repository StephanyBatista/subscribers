package main

import (
	"log"
	"subscribers/infra/database"
	"subscribers/web/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvs()

	db := database.CreateConnection()

	r := gin.Default()
	healthcheck := handlers.NewHealthCheck(db)
	r.GET("/healthcheck", healthcheck.Get)

	r.Run(":6004")
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
