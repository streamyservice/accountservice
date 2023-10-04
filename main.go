package main

import (
	"accountservice/app/router"
	"accountservice/config"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to read environment variables %s", err)
	}
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	err := app.Run(":" + port)
	if err != nil {
		log.Fatalf("Unable to start application on the provided port %s", err)
	}
}
