package models

import "time"

// type Book struct {
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// 	Desc   string `json:"desc"`
// }

type Book struct {
	ID        int    `gorm:"primaryKey"`
	NameBook  string `gorm:"type:varchar(255)"`
	Author    string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookInput struct {
	NameBook string `json:"name_book"`
	Author   string `json:"author"`
}
