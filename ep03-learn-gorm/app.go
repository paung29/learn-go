package main

import (
	"ep03-learn-gorm/database"
	"ep03-learn-gorm/model"
	"log"
)

func main() {

	database.ConnectDatabase()

	migrationError := database.DB.AutoMigrate(&model.User{}, &model.Post{})

	if migrationError != nil {
		log.Fatalf("auto migration failed : %v", migrationError)
	}

	

}