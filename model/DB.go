package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func getPostgresConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{FullSaveAssociations: true})
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
		&MovementType{},
		&User{})
	if err != nil {
		log.Fatalf("failed to auto migrate: %v", err)
	}
	// TODO: Remove this after testing
	InitPeriods()
	InitMovementTypes()
	chapter := InitChapter()
	if err := DB.Create(chapter).Error; err != nil {
		log.Printf("failed to create chapter: %v", err)
	}
	InitUsers(chapter)

}

func InitPeriods() {
	if DB.Find(&Period{}).RowsAffected > 0 {
		return
	}
	if err := DB.Create(&Period{
		Year:                  time.Now().Year(),
		TotalInstallments:     10,
		FirstMonthInstallment: 3,
		Current:               true,
	}).Error; err != nil {
		log.Printf("failed to create period: %v", err)
	}
}

func InitUsers(chapter *Chapter) {
	if DB.Find(&User{}).RowsAffected > 0 {
		return
	}
	hashedPassword := HashAndSalt([]byte("password"))
	if err := DB.Create(&User{
		UserName: "admin",
		Password: hashedPassword,
		Chapter:  chapter,
	}).Error; err != nil {
		log.Printf("failed to create user: %v", err)
	}

}

func InitChapter() *Chapter {
	chapter := &Chapter{}
	chapter.Init()
	chapter.Name = "Capitulo Argentino 1"
	return chapter
}

func HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
