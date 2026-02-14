package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

//var db *gorm.DB

func connect(){
	d, err := gorm.Open("mysql", "")
	if err != nil {
		panic("Failed to connect to database!")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}