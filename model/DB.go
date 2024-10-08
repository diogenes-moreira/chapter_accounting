package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func getPostgresConnection() (*gorm.DB, error) {
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

func AutoMigrate() {
	var err error
	DB, err = getPostgresConnection()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// AutoMigrate your models
	err = DB.AutoMigrate(
		&Movement{},
		&Brother{},
		&RollingBalance{},
		&Chapter{},
		&Affiliation{},
		&Installment{},
		&Period{},
		&ChargeType{},
		&Deposit{},
		&MovementType{})
	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}

	fmt.Println("Successfully connected to the database and migrated models")
}
