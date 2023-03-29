package models

import "time"

type Book struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	NameBook  string    `gorm:"type:varchar(255)" json:"name_book"`
	Author    string    `gorm:"type:varchar(255)" json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookInput struct {
	NameBook string `json:"name_book"`
	Author   string `json:"author"`
}
