package repositories

import (
	"belajar-gin/models"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(book models.Book) models.Book
	// GetBooks() []models.Book
	// DeleteBook(id int) int
	// UpdateBook(book models.Book) (int, models.Book)
	GetBookByID(id int) (models.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(book models.Book) models.Book {
	err := r.db.Create(&book).Error

	if err != nil {
		panic(err)
	}
	fmt.Println(book)
	return book
}

func (r *repository) GetBookByID(id int) (models.Book, error) {
	var book models.Book
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

// func (r *repository) GetBooks() []models.Book {
// 	books := []models.Book{}

// 	query := "SELECT * FROM items"

// 	rows, err := r.db.Query(query)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		book := models.Book{}

// 		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
// 		if err != nil {
// 			panic(err)
// 		}

// 		books = append(books, book)
// 	}
// 	return books
// }

// func (r *repository) DeleteBook(id int) int {
// 	query := `
// 		DELETE FROM items
// 		WHERE id=$1
// 	`

// 	res, err := r.db.Exec(query, id)
// 	if err != nil {
// 		panic(err)
// 	}

// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return int(count)
// }

// func (r *repository) UpdateBook(book models.Book) (int, models.Book) {
// 	query := `
// 		UPDATE items
// 		SET title=$2, author=$3, description=$4
// 		WHERE id=$1
// 	`

// 	res, err := r.db.Exec(query, book.ID, book.Title, book.Author, book.Desc)
// 	if err != nil {
// 		panic(err)
// 	}

// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	return int(count), book
// }
