package main

import (
	router "belajar-gin/routers"
)

const PORT = ":4000"

func main() {
	router.StartServer().Run(PORT)
}

// func main() {
// 	http.HandleFunc("/books", handleBooks)
// 	http.HandleFunc("/books/", handleBook)

// 	http.ListenAndServe(PORT, nil)
// }

// func handleBooks(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		getBooks(w, r)
// 	case http.MethodPost:
// 		addBook(w, r)
// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func handleBook(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		getBook(w, r)
// 	case http.MethodPut:
// 		updateBook(w, r)
// 	case http.MethodDelete:
// 		deleteBook(w, r)
// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-type", "application/json")

// 	json.NewEncoder(w).Encode(books)
// }

// func addBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-type", "application/json")

// 	body, _ := ioutil.ReadAll(r.Body)

// 	var book Book

// 	err := json.Unmarshal(body, &book)
// 	if err != nil {
// 		http.Error(w, "Error parsing request body", http.StatusBadRequest)
// 		return
// 	}

// 	book.ID = len(books) + 1

// 	newBook := Book{
// 		ID:     book.ID,
// 		Title:  book.Title,
// 		Author: book.Author,
// 		Desc:   book.Desc,
// 	}
// 	books = append(books, newBook)

// 	json.NewEncoder(w).Encode("Created")
// }

// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-type", "application/json")

// 	params := strings.TrimPrefix(r.URL.Path, "/books/")
// 	bookID, _ := strconv.Atoi(params)

// 	var isBookExist bool
// 	for _, book := range books {
// 		if book.ID == bookID {
// 			isBookExist = true
// 		}
// 	}
// 	if !isBookExist {
// 		http.Error(w, "Book not found", http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(books[bookID-1])
// }

// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-type", "application/json")

// 	params := strings.TrimPrefix(r.URL.Path, "/books/")
// 	bookID, _ := strconv.Atoi(params)

// 	var isBookExist bool
// 	for _, book := range books {
// 		if book.ID == bookID {
// 			isBookExist = true
// 		}
// 	}
// 	if !isBookExist {
// 		http.Error(w, "Book not found", http.StatusBadRequest)
// 		return
// 	}

// 	body, _ := ioutil.ReadAll(r.Body)

// 	var book Book

// 	err := json.Unmarshal(body, &book)
// 	if err != nil {
// 		http.Error(w, "Error parsing request body", http.StatusBadRequest)
// 		return
// 	}

// 	books[bookID-1].ID = bookID
// 	books[bookID-1].Title = book.Title
// 	books[bookID-1].Author = book.Author
// 	books[bookID-1].Desc = book.Desc

// 	json.NewEncoder(w).Encode("Updated")
// }

// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-type", "application/json")

// 	params := strings.TrimPrefix(r.URL.Path, "/books/")
// 	bookID, _ := strconv.Atoi(params)

// 	var isBookExist bool
// 	for _, book := range books {
// 		if book.ID == bookID {
// 			isBookExist = true
// 		}
// 	}
// 	if !isBookExist {
// 		http.Error(w, "Book not found", http.StatusBadRequest)
// 		return
// 	}

// 	for i, book := range books {
// 		if book.ID == bookID {
// 			books = append(books[:i], books[i+1:]...)
// 		}
// 	}
// 	json.NewEncoder(w).Encode("Deleted")

// }
