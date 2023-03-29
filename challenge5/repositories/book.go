package repositories

import (
	"belajar-gin/models"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(book models.Book) models.Book
	GetBooks() ([]models.Book, error)
	DeleteBook(id int) error
	UpdateBook(book models.Book) (models.Book, error)
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

func (r *repository) UpdateBook(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) GetBooks() ([]models.Book, error) {
	var books []models.Book

	err := r.db.Find(&books).Error

	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *repository) DeleteBook(id int) error {
	var book models.Book
	if err := r.db.Where("id = ?", id).First(&book).Delete(&book).Error; err != nil {
		return err
	}

	return nil
}
