package repositories

import (
	"belajar-gin/models"
	"database/sql"
)

type Repository interface {
	CreateBook(book models.Book) models.Book
	GetBooks() []models.Book
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBook(book models.Book) models.Book {
	query := `
		INSERT INTO items(title, author, description)
		VALUES ($1, $2, $3)
		RETURNING *
	`
	err := r.db.QueryRow(query, book.Title, book.Author, book.Desc).
		Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		panic(err)
	}

	return book
}

func (r *repository) GetBooks() []models.Book {
	books := []models.Book{}

	query := "SELECT * FROM items"

	rows, err := r.db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		book := models.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			panic(err)
		}

		books = append(books, book)
	}
	return books
}
