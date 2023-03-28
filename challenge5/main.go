package main

import (
	router "belajar-gin/routers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

const (
	DB_HOST     = "localhost"
	DB_USER     = "root"
	DB_PASSWORD = "secret"
	DB_PORT     = 5432
	DB_NAME     = "books"
	PORT        = ":4000"
)

func main() {
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

	router.StartServer(db).Run(PORT)

	// UpdateBook()
	// DeleteBook()
	// GetBooks()
}

// func UpdateBook() {
// 	query := `
// 		UPDATE items
// 		SET title=$2, author=$3, description=$4
// 		WHERE id=$1
// 	`

// 	res, err := db.Exec(query, 1, "Rich Dad Poor Dad", "Robert Kiyosaki", "financial book")
// 	if err != nil {
// 		panic(err)
// 	}

// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	if count == 0 {
// 		fmt.Println("Data Not Found", count)
// 		return
// 	}

// 	fmt.Println("Successfully updated book", count)
// }
