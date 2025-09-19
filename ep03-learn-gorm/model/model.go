package model


type User struct {
	ID 			uint	`json:"id" gorm:"primaryKey"`
	FullName 	string	`json:"full_name"`
	Email 		string	`json:"email" gorm:"uniqueIndex"`
	Posts 		[]Post	`json:"post" gorm:"foreignKey:UserID"`
}

type Post struct {
	ID 			uint	`json:"id" gorm:"primaryKey"`
	Title 		string	`json:"title"`
	UserID		uint	`json:"user_id"`
}