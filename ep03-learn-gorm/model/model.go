package model

import "gorm.io/gorm"


type User struct {
	gorm.Model	
	Name 		string	
	Email 		string	
	Posts 		[]Post	
	Comments	[]Comment
}

type Category struct {
	gorm.Model
	Name string
	Posts []Post
}

type Post struct {
	gorm.Model	
	Title 		string	
	UserID		uint	
	User		User
	CategoryID 	uint
	Category 	Category
	Comments	[]Comment
}

type Comment struct {
	gorm.Model
	Content 	string
	UserID		uint
	User		User
	PostID		uint
	Post		Post
}