package book

import "belajar-gin/models"

func CreateBook(book models.Book) (data models.Book, err error) {
	book.ID = len(models.Books) + 1

	models.Books = append(models.Books, book)

	return book, nil
}
