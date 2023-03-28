package repositories

import (
	"belajar-gin/models"
	"database/sql"
)

type Repository interface {
	CreateBook(book models.Book) models.Book
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
