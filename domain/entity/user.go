package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID           int
	Email        string
	PasswordHash string
}
