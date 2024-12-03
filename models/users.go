package models

import "gorm.io/gorm"

type UsersModel struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Phone    string `json:"phone" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
