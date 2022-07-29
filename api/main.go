package main

import (
	"log"
	"subscribers/helpers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvs()
	di := helpers.NewDI()
	r := gin.Default()

	r.GET("/healthcheck", di.HealthCheckHandler.Get)
	r.POST("/users", di.UserHandler.Post)

	r.Run(":6004")
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
