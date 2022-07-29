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
	healthcheck := handlers.NewHealthCheckHandler(db)
	r.GET("/healthcheck", healthcheck.Get)
	user := handlers.NewUserHandler(db)
	r.POST("/users", user.Post)

	r.Run(":6004")
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
