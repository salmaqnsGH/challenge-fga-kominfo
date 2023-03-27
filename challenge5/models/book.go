package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var Books = []Book{
	{ID: 1, Title: "Naruto", Author: "Masashi K", Desc: "a descrition of thebook"},
	{ID: 2, Title: "Rich Dad Poor Dad", Author: "Robert Kiyosaki", Desc: "a descrition of thebook"},
	{ID: 3, Title: "Harry Potter", Author: "J.K Rowling", Desc: "a descrition of thebook"},
}
