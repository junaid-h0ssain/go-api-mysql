package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	r := mux.NewRouter()
	RegisterBooksStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}