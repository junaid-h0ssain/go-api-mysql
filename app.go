package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func connect(){
	d, err := gorm.Open("mysql", "root:secret123@tcp(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Failed to connect to database! Error: ", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}