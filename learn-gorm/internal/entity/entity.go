package entity

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username string
	Email    string
	Password string
}

type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
