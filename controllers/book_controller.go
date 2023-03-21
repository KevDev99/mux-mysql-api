package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mux-mysql-api/configs"
	"mux-mysql-api/models"
	"mux-mysql-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var book models.Book
		defer cancel()

		// validate the request body
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&book); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		fmt.Println(book)

		newBook := models.Book{
			Title:        book.Title,
			Descr:        book.Descr,
			ThumbnailUrl: book.ThumbnailUrl,
		}

		err := configs.DB.Create(&newBook).Error
		if err != nil {
			log.Fatal(err)
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.BookResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": newBook}}
		json.NewEncoder(rw).Encode(response)
	}
}

func ListBooks() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

	}
}

func UpdateBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

	}
}

func DeleteBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

	}
}
