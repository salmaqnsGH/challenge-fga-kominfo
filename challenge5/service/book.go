package service

import (
	"belajar-gin/models"
	"belajar-gin/repositories"
	"fmt"
)

type Service interface {
	CreateBook(input models.BookInput) models.Book
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
