package service

import (
	"github.com/diyor200/learn-gorm/internal/entity"
	"gorm.io/gorm"
	"log"
)

func Create(db *gorm.DB, user entity.UserInput) error {
	err := db.Table("users").Create(&user)
	return err.Error
}

func Read(db *gorm.DB) ([]entity.User, error) {
	var users []entity.User
	res := db.Table("users").Scan(&users)
	log.Println(res)
	return users, res.Error
}

//func ReadOne(db *gorm.DB, userID int) (entity.User, error) {
//	var user entity.User
//	err := db.Table("users").First(&user, userID)
//	return user, err.Error
//}

func Delete(db *gorm.DB, id int) error {
	log.Println("id=====", id)
	err := db.Table("users").Where("id=?", id).Delete(&entity.User{})
	return err.Error
}

func Update(db *gorm.DB, input entity.UserInput) (entity.User, error) {
	var user entity.User
	res := db.Table("users").
		Where("email=?", input.Email).
		Updates(entity.User{Username: input.Username, Password: input.Password}).Scan(&user)
	return user, res.Error
}
