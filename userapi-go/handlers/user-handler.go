package handlers 

import(
	"encoding/json"
	"net/http"
	"user.com/userapi/model"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []model.User{
		{Id: 1, Name: "Aike", Email: "aikepaung767@gmail.com"},
		{Id: 2, Name: "Deborah", Email: "deborah@gmail.com"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}