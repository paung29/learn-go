package service

import (
	"ep03-learn-gorm/database"
	"ep03-learn-gorm/model"
	"ep03-learn-gorm/utils"
	"errors"
)

func LoginUser(emial string, password string) (string, error){
	var user model.User
	if err := database.DB.Where("email = ?", emial).First(&user).Error; err != nil {
		return "", errors.New("invalid email or password")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, _ := utils.GenerateJWt(user.ID)
	return token, nil
}

func CreateUser(user model.User) (model.User, error){

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword

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

func DeleteUser(id uint) error {
	result := database.DB.Delete(&model.User{}, id)
	return result.Error
}