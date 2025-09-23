package dto

type LoginForm struct {
	Email 		string	`binding:"required"`
	Password	string	`binding:"required"`
}

type CreateUserForm struct {
	Name string
	Email string 
	Password string
}