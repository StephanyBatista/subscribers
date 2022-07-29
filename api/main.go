package main

import (
	"log"
	"subscribers/helpers"
	"subscribers/web/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvs()
	di := helpers.NewDI()
	r := gin.Default()

	r.GET("/healthcheck", di.HealthCheckHandler.Get)
	r.POST("/token", di.TokenHandler.Post)
	r.POST("/users", di.UserHandler.Post)
	secured := r.Group("").Use(middlewares.Auth())
	{
		secured.GET("/users/info", di.UserHandler.GetInfo)
	}

	r.Run(":6004")
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
