package db

import (
	"fmt"
	"latihan-jwt/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "root"
	DB_PASSWORD = "secret"
	DB_PORT     = 5432
	DB_NAME     = "tests"
	DEBUG_MODE  = false // true/false
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", DB_USER, DB_NAME, DB_PASSWORD, DB_HOST, DB_PORT)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})

	fmt.Println("Successfully connected to database!")
}

func GetDB() *gorm.DB {
	if DEBUG_MODE {
		return db.Debug()
	}

	return db
}
