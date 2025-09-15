package model

import "time"

type User struct {
	Id 			int 				`json:"id" gorm:"primaryKey"` 
	FullName 	string				`json:"full_name"`
	Email 		string				`json:"email" gorm:"uniqueIndex"`
	CreatedAt 	time.Time			`json:"created_at"`
	UpdatedAt	time.Time			`json:"updated_at"`
}