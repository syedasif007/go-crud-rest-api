package main

import (
	"log"
	"net/http"

	"rest-api-crud-2/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()
	// Route handles & endpoints
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/user/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/user", handlers.AddUser).Methods("POST")
	r.HandleFunc("/user/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", handlers.DeleteUser).Methods("DELETE")
	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
