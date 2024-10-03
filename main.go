package main

import (
	"argentina-tresury/controllers"
	"argentina-tresury/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	initEnvironment()
	InitDb()

	InitServer()

}

func InitServer() {
	r := mux.NewRouter()
	controllers.InitView(r)
	controllers.RegisterChapterRoutesOn(r)
	controllers.RegisterBrotherRoutesOn(r)
	controllers.RegisterIndex(r)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTENER"), r))
}

func InitDb() {
	db.AutoMigrate()
}

func initEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
