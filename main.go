package main

import (
	"argentina-tresury/db"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	initEnvironment()
	InitDb()
}

func InitDb() {
	db.Automigrate()
}

func initEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
