package controller

import (
	"ep02-gin-user-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	FullName string		`json:"full_name" binding:"required"`
	Email	 string		`json:"email" binding:"required,email"`
}

func CreateUserHandler(context *gin.Context) {
	var createUserRequest CreateUserRequest
	bindError := context.ShouldBindJSON(&createUserRequest)

	if bindError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error" : bindError.Error(),
		})
		return
	}

	createdUser, serviceError := service.CreateUser(
		createUserRequest.FullName,
		createUserRequest.Email)

	if(serviceError != nil) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error" : serviceError.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, createdUser)
}

func ListUserHandler(context *gin.Context){
	userlist, err := service.ListAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error" : "failed to retrived users"})
	}
	context.JSON(http.StatusOK, userlist)
}

func GetUserByIDHandler(context *gin.Context) {

}