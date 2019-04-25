package main

import (
	_ "encoding/json"
	"log"
	_ "math/rand"
	"net/http"
	_ "strconv"

	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	id     string  `json: "id"`
	isbn   string  `json: "isbn"`
	title  string  `json: "title"`
	author *Author `json: "author"`
}

type Author struct {
	firstName string `json: "firstName"`
	lastName  string `json: "lastName"`
}

// index
func index(w http.ResponseWriter, r *http.Request) {

}

// show
func show(w http.ResponseWriter, r *http.Request) {

}

// create
func create(w http.ResponseWriter, r *http.Request) {

}

// update
func update(w http.ResponseWriter, r *http.Request) {

}

// delete
func delete(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init router
	r := mux.NewRouter()

	// Create handlers / REST Endpoints
	// * Make similar to ruby methods for clarity sake.
	r.HandleFunc("/api/books", index).Methods("GET")
	r.HandleFunc("/api/books/{id}", show).Methods("GET")
	r.HandleFunc("/api/books", create).Methods("POST")
	r.HandleFunc("/api/books/{id}", update).Methods("PUT")
	r.HandleFunc("/api/books/{id}", delete).Methods("DELETE")

	log.Println("Starting server ...")
	err := http.ListenAndServe(":3000", r)
	log.Fatalln(err)
}
