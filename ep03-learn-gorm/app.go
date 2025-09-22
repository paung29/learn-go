package main

import (
	"ep03-learn-gorm/controller"
	"ep03-learn-gorm/database"
	"ep03-learn-gorm/model"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDatabase()

	migrationError := database.DB.AutoMigrate(&model.User{}, &model.Category{}, &model.Post{},  &model.Comment{})

	if migrationError != nil {
		log.Fatalf("auto migration failed : %v", migrationError)
	}
	
	route := gin.Default()
	route.POST("/user", controller.CreateUser)
	route.GET("/user/:id", controller.GetUserById)
	route.GET("/users", controller.GetAllUser)
	route.DELETE("/user/:id", controller.DeleteUserById)

	route.Run(":8080")

}