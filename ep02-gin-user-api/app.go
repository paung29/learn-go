package main

import (
	"ep02-gin-user-api/controller"
	"ep02-gin-user-api/database"
	"ep02-gin-user-api/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDatabase()

	migrationError := database.DB.AutoMigrate(&model.User{})

	if migrationError != nil {
		log.Fatalf("auto migrate failed : %v", migrationError)
	}

	router := gin.Default()

	router.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok" : true})
	})

	router.POST("/users", controller.CreateUserHandler)

	runServerError := router.Run(":8080")
	if runServerError != nil {
		log.Fatalf("server failed: %v", runServerError)
	}

}