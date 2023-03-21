package main

import (
	"log"
	"mux-mysql-api/configs"
	"mux-mysql-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	configs.ConnectDB()

	routes.BookRoute(router)

	log.Fatal(http.ListenAndServe(":6000", router))
}
