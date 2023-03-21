package routes

import (
	"mux-mysql-api/controllers"

	"github.com/gorilla/mux"
)

func BookRoute(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook()).Methods("POST")
}
