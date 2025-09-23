package controller

import (
	"ep03-learn-gorm/controller/dto"
	"ep03-learn-gorm/model"
	"ep03-learn-gorm/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var form dto.LoginForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error : " : err.Error()})
		return
	}

	token, err := service.LoginUser(form.Email, form.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error : " : err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token :" : token})
}

func CreateUser(c *gin.Context) {
	var form dto.CreateUserForm
	if err := c.ShouldBindJSON(&form);  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error :" : err.Error()})
		return
	}

	user := model.User{
		Name : form.Name,
		Email: form.Email,
		Password: form.Password,
	}

	newUser, err := service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error : ": "Failed to create User"})
		return
	}
	c.JSON(http.StatusOK, newUser)

}

func GetUserById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error : " : "id must be a number"})
		return
	}

	user, err := service.FindUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error :" : "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetAllUser(c *gin.Context) {
	users, err := service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error :" : "failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func DeleteUserById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error : " : "id must be a number"})
		return
	}

	if err := service.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error : " : "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message : " : "user deleted successfully"})
}