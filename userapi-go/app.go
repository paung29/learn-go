package main

import (
	"log"
	"net/http"
	"user.com/userapi/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handlers.GetUsers)

	log.Println("Server running")
	log.Fatal(http.ListenAndServe(":8080", mux))
}