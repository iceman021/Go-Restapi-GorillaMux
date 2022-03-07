package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book


// Get All Books
func getBooks(w http.ResponseWriter, router *http.Request) {
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(books)
}

// Get A Single Book
func getBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router) // Get Params
	//Loop through Books and find with Id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Create a New Book
func createBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_=json.NewDecoder(router.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // MOCK ID - NOT SAFE IN PRODUCTION
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router) // Get Params
	//Loop through Books and find with Id
	for index, item := range books {
		if item.ID == params["id"] { 
		books = append(books[:index], books[index+1:]...)
		var book Book
		_=json.NewDecoder(router.Body).Decode(&book)
		book.ID = params["id"]
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
		return
	}
}
json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router) // Get Params
	//Loop through Books and find with Id
	for index, item := range books {
		if item.ID == params["id"] { 
		books = append(books[:index], books[index+1:]...)
		break
	}
}
json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID:"1", Isbn: "345543", Title: "Book One", Author: &Author {Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID:"2", Isbn: "4646646", Title: "Book Two", Author: &Author {Firstname: "Steve", Lastname: "Smith"}})



	// Route handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router)) 
}