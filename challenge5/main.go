package main

import (
	"belajar-gin/models"
	router "belajar-gin/routers"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "root"
	DB_PASSWORD = "secret"
	DB_PORT     = 5432
	DB_NAME     = "tests"
	PORT        = ":8080"
)

func main() {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", DB_USER, DB_NAME, DB_PASSWORD, DB_HOST, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Book{})
	fmt.Println("Successfully connected to database!")

	router.StartServer(db).Run(PORT)
}
