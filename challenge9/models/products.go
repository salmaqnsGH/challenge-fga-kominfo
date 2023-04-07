package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GORMModel
	UserID      uint   `json:"user_id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	User        *User
}

func (u *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}
	return
}

func (u *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}
	return
}
