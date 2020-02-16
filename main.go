// frkntony 16, Feb 2020
package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Book struct (Model)
type Book struct {
	ID 			string `json:"id"`
	Isbn 		string `json:"isbn"`
	Title 		string `json:"title"`
	Author 		*Author `json:"author"`
}

type Author struct {
	Firstname 	string `json:"firstname"`
	Lastname 	string `json:"lastname"`
}

// init books slice
var books []Book

// all books
func getBooks(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// one book
func getBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// get book with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// new book
func createBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

// update a book
func updateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
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

// delete a book
func deleteBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main()  {
	// Init router
	r := mux.NewRouter()

	// Mock Data
	books = append(books,
		Book{
			ID:     "1",
			Isbn:   "23421541243",
			Title:  "Book One",
			Author: &Author{
				Firstname: "Tony",
				Lastname:  "Elistratov",
			},
		})
	books = append(books,
		Book{
			ID:     "2",
			Isbn:   "44421541123",
			Title:  "Book Two",
			Author: &Author{
				Firstname: "Adam",
				Lastname:  "Loveless",
			},
		})

	books = append(books,
		Book{
			ID:     "3",
			Isbn:   "63421541299",
			Title:  "Book Three",
			Author: &Author{
				Firstname: "Carol",
				Lastname:  "Brox",
			},
		})

	// Router Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
