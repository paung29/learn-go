package controller

import (
	"ep03-learn-gorm/model"
	"ep03-learn-gorm/service"
	"net/http"

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