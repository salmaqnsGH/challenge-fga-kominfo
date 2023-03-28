package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	PORT        = ":4000"
	DB_HOST     = "localhost"
	DB_USER     = "root"
	DB_PASSWORD = "secret"
	DB_PORT     = 5432
	DB_NAME     = "book"
)

func main() {
	// router.StartServer().Run(PORT)

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", DB_USER, DB_NAME, DB_PASSWORD, DB_HOST, DB_PORT)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}
