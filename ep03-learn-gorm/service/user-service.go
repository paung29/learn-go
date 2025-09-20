package service

import (
	"ep03-learn-gorm/database"
	"ep03-learn-gorm/model"
)

func CreateUser(user model.User) (model.User, error){
	result := database.DB.Create(&user)
	return user, result.Error
}