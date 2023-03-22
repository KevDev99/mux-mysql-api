package routes

import (
	"mux-mysql-api/controllers"

	"github.com/gorilla/mux"
)

func BookRoute(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook()).Methods("POST")
	router.HandleFunc("/books", controllers.ListBooks()).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.ListBook()).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.DeleteBook()).Methods("DELETE")
	router.HandleFunc("/book/{id}", controllers.UpdateBook()).Methods("PATCH")

}
