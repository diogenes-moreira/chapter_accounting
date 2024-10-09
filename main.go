package main

import (
	"argentina-tresury/controllers"
	"argentina-tresury/model"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	initEnvironment()
	InitDb()
	controllers.InitServer()
}

func InitDb() {
	model.AutoMigrate()
}

func initEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
