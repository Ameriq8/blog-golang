package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Phone    string `json:"phone" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
