package service

import (
	"ep03-learn-gorm/database"
	"ep03-learn-gorm/model"
)

func CreateUser(user model.User) (model.User, error){
	result := database.DB.Create(&user)
	return user, result.Error
}

func FindUserByID(id uint) (model.User, error) {
	var user model.User
	result := database.DB.First(&user, id)
	return user, result.Error
}

func GetAllUser() ([]model.User, error){
	var users []model.User
	result := database.DB.Find(&users)
	return users, result.Error
}