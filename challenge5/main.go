package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const PORT = ":4000"

type Book struct {
	ID     int
	Title  string
	Author string
	Desc   string
}

var books = []Book{
	{ID: 1, Title: "Naruto", Author: "Masashi K", Desc: "a descrition of thebook"},
	{ID: 2, Title: "Rich Dad Poor Dad", Author: "Robert Kiyosaki", Desc: "a descrition of thebook"},
	{ID: 3, Title: "Harry Potter", Author: "J.K Rowling", Desc: "a descrition of thebook"},
}

func main() {
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/books/new", addBook)
	http.HandleFunc("/books/", getBook)
	http.HandleFunc("/books/update/", updateBook)
	http.HandleFunc("/books/delete/", deleteBook)

	http.ListenAndServe(PORT, nil)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(books)
		return
	}

	http.Error(w, "Method is not allowed", http.StatusBadRequest)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)

		var book Book

		err := json.Unmarshal(body, &book)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		book.ID = len(books) + 1

		newBook := Book{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Desc:   book.Desc,
		}
		books = append(books, newBook)

		json.NewEncoder(w).Encode("Created")
		return
	}

	http.Error(w, "Method is not allowed", http.StatusBadRequest)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	params := strings.TrimPrefix(r.URL.Path, "/books/")
	bookID, _ := strconv.Atoi(params)

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(books[bookID-1])
		return
	}

	http.Error(w, "Method is not allowed", http.StatusBadRequest)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	params := strings.TrimPrefix(r.URL.Path, "/books/update/")
	bookID, _ := strconv.Atoi(params)

	if r.Method == "PUT" {
		body, _ := ioutil.ReadAll(r.Body)

		var book Book

		err := json.Unmarshal(body, &book)
		if err != nil {
			http.Error(w, "Error parsing request body", http.StatusBadRequest)
			return
		}

		books[bookID-1].ID = bookID
		books[bookID-1].Title = book.Title
		books[bookID-1].Author = book.Author
		books[bookID-1].Desc = book.Desc

		json.NewEncoder(w).Encode("Updated")
		return
	}

	http.Error(w, "Method is not allowed", http.StatusBadRequest)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	params := strings.TrimPrefix(r.URL.Path, "/books/delete/")
	bookID, _ := strconv.Atoi(params)

	if r.Method == "DELETE" {
		for i, book := range books {
			if book.ID == bookID {
				books = append(books[:i], books[i+1:]...)
			}
		}
		json.NewEncoder(w).Encode("Deleted")
		return
	}

	http.Error(w, "Method is not allowed", http.StatusBadRequest)
}
