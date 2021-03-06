package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var SERVER_PORT = "8000"

type Book struct {
	ID		 string  `json:"id"`
	Isbn	 string  `json:"isbn"`
	Title	 string  `json:"title"`
	Author	 *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// slice is variable length array
// init books var as slice Book struct
var books []Book


func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if (item.ID == params["id"]){
			json.NewEncoder(w).Encode(item)
			return;
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusCreated)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, ogBook := range(books){
		if (ogBook.ID == params["id"]){
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			books[index] = book
			json.NewEncoder(w).Encode(book)
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range(books){
		if (item.ID == params["id"]){
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {

	// INIT ROUTER
	// infer type
	router := mux.NewRouter()

	// var age int = 35

	// mock data
	books = append(books, 
		Book{
			ID: "1", 
			Isbn: "448743", 
			Title: "Book One", 
			Author: &Author{
				Firstname: "John",
				Lastname: "Doe",
			},
		},
		Book{
			ID: "2",
			Isbn: "454545",
			Title: "Book Two",
			Author: &Author{
				Firstname: "Steve",
				Lastname: "Smith",
			},
		},
	)


	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books", updateBook).Methods("PUT")
	router.HandleFunc("/api/books", deleteBook).Methods("DELETE")

	println("Server will run on port " + SERVER_PORT)
	log.Fatal(http.ListenAndServe(":" + SERVER_PORT, router))
}