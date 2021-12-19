package main

import (
	"encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
) 
// books struct . 
type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author *Author `json:"author"`
	Year string `json:"year"`
	ISBN string `json:"isbn"`
}

// AUTHOR STRUCT
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
// create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
// update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main() {
	// fmt.Println("harshit sharma is good")
	// Init router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}, Year: "2018", ISBN: "1234567890"})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}, Year: "2019", ISBN: "1234567890"})
	books = append(books, Book{ID: "3", Title: "Book Three", Author: &Author{Firstname: "Jane", Lastname: "Doe"}, Year: "2020", ISBN: "1234567890"})
	books = append(books, Book{ID: "4", Title: "Book Four", Author: &Author{Firstname: "John", Lastname: "Doe"}, Year: "2018", ISBN: "1234567890"})
	books = append(books, Book{ID: "5", Title: "Book Five", Author: &Author{Firstname: "Steve", Lastname: "Smith"}, Year: "2019", ISBN: "1234567890"})

 
	// router handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))


}
