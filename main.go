package main

import (
	"argentina-tresury/controllers"
	"argentina-tresury/model"
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
	controllers.RegisterTreasuryRoutesOn(r)
	controllers.RegisterIndex(r)
	controllers.RegisterMovementTypeRoutesOn(r)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTENER"), r))
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
