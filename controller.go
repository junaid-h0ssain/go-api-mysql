package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

var NewBook Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := GetAllBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return 
	}
}
