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
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

		var Books []models.Book
		var err error

		err = configs.DB.Find(&Books).Error

		if err != nil {
			fmt.Println("could not fetch all books")
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": Books}}
		json.NewEncoder(rw).Encode(response)

	}
}

func ListBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		book_id := mux.Vars(r)["id"]

		if book_id == "" {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "no id provided"}}
			json.NewEncoder(rw).Encode(response)
		}

		if _, err := strconv.Atoi(book_id); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "id is not an integer"}}
			json.NewEncoder(rw).Encode(response)
		}

		// fetch book
		var book models.Book

		err := configs.DB.Take(&book, "book_id = ?", book_id).Error

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err}}
			json.NewEncoder(rw).Encode(response)
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": book}}
		json.NewEncoder(rw).Encode(response)

	}
}

func contains(stringSlice []string, text string) bool {
	for _, a := range stringSlice {
		if a == text {
			return true
		}
	}
	return false
}

func UpdateBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// get url param id
		book_id := mux.Vars(r)["id"]

		// decode body
		var generic map[string]interface{}
		json.NewDecoder(r.Body).Decode(&generic)

		// get json fields
		jsonFields := new(models.Book).GetJsonFields()

		// loop through the body
		// if one key does not match -> return http.StatusBadRequest
		for k, _ := range generic {
			if contains(jsonFields, k) == false {
				rw.WriteHeader(http.StatusBadRequest)
				response := responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "a provided field is unknown."}}
				json.NewEncoder(rw).Encode(response)
				return
			}
		}

		if book_id == "" {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "no id provided"}}
			json.NewEncoder(rw).Encode(response)
		}

		var book models.Book

		if err := configs.DB.Model(&book).Where("book_id = ?", book_id).Updates(generic).Error; err != nil {
			fmt.Println(err)
		}

		configs.DB.Take(&book, "book_id = ?", book_id)

		rw.WriteHeader(http.StatusOK)
		response := responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"updated_book": book}}
		json.NewEncoder(rw).Encode(response)

	}
}

func DeleteBook() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		book_id := mux.Vars(r)["id"]

		if book_id == "" {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "no id provided"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		if _, err := strconv.Atoi(book_id); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "id is not an integer"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		var book models.Book

		err := configs.DB.Where("book_id = ?", book_id).Delete(&book).Error

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "book deleted"}}
		json.NewEncoder(rw).Encode(response)
	}
}
