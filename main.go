package main

import (
	"log"
	"mux-mysql-api/configs"
	"mux-mysql-api/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	configs.ConnectDB()

	routes.BookRoute(router)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
