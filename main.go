package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"go-mysql-api/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}