package repositories

import (
	"belajar-gin/models"
	"database/sql"
)

type Repository interface {
	CreateBook(book models.Book) models.Book
	GetBooks() []models.Book
	DeleteBook(id int) int
	UpdateBook(book models.Book) (int, models.Book)
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

func (r *repository) DeleteBook(id int) int {
	query := `
		DELETE FROM items
		WHERE id=$1
	`

	res, err := r.db.Exec(query, id)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return int(count)
}

func (r *repository) UpdateBook(book models.Book) (int, models.Book) {
	query := `
		UPDATE items
		SET title=$2, author=$3, description=$4
		WHERE id=$1
	`

	res, err := r.db.Exec(query, book.ID, book.Title, book.Author, book.Desc)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	return int(count), book
}
