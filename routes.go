package main

import (
	"go-mysql-api/controller"

	"github.com/gorilla/mux"
)

var RegisterBooksStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/books/", controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
}