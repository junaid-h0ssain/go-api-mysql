package main

import (
	//. "go-mysql-api/controller"

	"github.com/gorilla/mux"
)

var RegisterBooksStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books/", GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", DeleteBook).Methods("DELETE")
}