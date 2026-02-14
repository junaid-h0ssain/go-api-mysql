package main

import (
	"github.com/jinzhu/gorm"
	"go-mysql-api/app"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name  string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	app.connect()
	db = app.GetDB()
	db.AutoMigrate(&Book{})
}


