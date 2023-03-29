package service

import (
	"belajar-gin/models"
	"belajar-gin/repositories"
	"fmt"
	"time"
)

type Service interface {
	CreateBook(input models.BookInput) models.Book
	GetBookyID(id int) (models.Book, error)
	UpdateBook(id int, inputData models.BookInput) (models.Book, error)
}

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) CreateBook(input models.BookInput) models.Book {
	book := models.Book{}
	book.NameBook = input.NameBook
	book.Author = input.Author

	newBook := s.repository.Create(book)
	fmt.Println("newBook", newBook)

	return newBook
}

func (s *service) GetBookyID(id int) (models.Book, error) {
	book, err := s.repository.GetBookByID(id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *service) UpdateBook(id int, inputData models.BookInput) (models.Book, error) {
	book, err := s.repository.GetBookByID(id)
	if err != nil {
		return book, err
	}

	book.NameBook = inputData.NameBook
	book.Author = inputData.Author
	book.UpdatedAt = time.Now()
	book.CreatedAt = time.Now()

	updatedBook, err := s.repository.UpdateBook(book)
	if err != nil {
		return updatedBook, err
	}

	return updatedBook, nil
}
