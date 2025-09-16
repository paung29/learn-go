package service

import (
	"ep02-gin-user-api/model"
	"ep02-gin-user-api/database"
)

func CreateUser(fullName string, email string) (model.User, error) {

	newUser := model.User{
		FullName: fullName,
		Email: email,
	}
	result := database.DB.Create(&newUser)
	err := result.Error

	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
}

func ListAllUsers() ([]model.User, error) {

	var userList []model.User
	result :=  database.DB.Order("id ASC").Find(&userList)
	err := result.Error

	if err != nil {
		return nil, err
	}
	return userList, nil
}