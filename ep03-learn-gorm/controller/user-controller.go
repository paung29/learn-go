package controller

import (
	"ep03-learn-gorm/model"
	"ep03-learn-gorm/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user);  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error :" : err.Error()})
		return
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