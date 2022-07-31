package main

import (
	"log"
	"subscribers/helpers"
	"subscribers/web/routers"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvs()
	di := helpers.NewDI()
	r := routers.CreateRouter(di)
	r.Run(":6004")
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
