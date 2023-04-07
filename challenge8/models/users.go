package models

import (
	"latihan-jwt/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GORMModel
	FirstName string    `gorm:"not null" json:"first_name" validate:"required-First name is required"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" validate:"required-Email is required"`
	Password  string    `gorm:"not null" json:"password" validate:"required-Password is required,minstringlength(6)-Password has to have minimum 6 charracter"`
	Products  []Product `json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	hashedPass, err := helpers.HashPassword(u.Password)
	if err != nil {
		return
	}

	u.Password = hashedPass
	return
}
