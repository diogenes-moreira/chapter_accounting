package db

import (
	"argentina-tresury/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetPostgresConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Automigrate() {
	db, err := GetPostgresConnection()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// AutoMigrate your models
	err = db.AutoMigrate(&model.Brother{}, &model.RollingBalance{}, &model.Chapter{})
	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	fmt.Println("Successfully connected to the database and migrated models")
}