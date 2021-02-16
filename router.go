package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func endPoint() {
	// init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getAll(Book{})).Methods("GET")
	r.HandleFunc("/api/books/{id}", getOne(Book{})).Methods("GET")
	r.HandleFunc("/api/books", create(Book{})).Methods("POST")
	r.HandleFunc("/api/books/{id}", update(Book{})).Methods("PUT")
	r.HandleFunc("/api/books/{id}", delete(Book{})).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8888", r))
}
