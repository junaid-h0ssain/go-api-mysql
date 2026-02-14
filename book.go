package main

import (
	"github.com/jinzhu/gorm"

)


type Book struct {
	gorm.Model
	Name  string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	connect()
	db = GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var Book Book
	db := db.Where("id = ?", id).Find(&Book)
	return &Book, db
}

func DeleteBookByID(id int64) Book {
	var Book Book
	db.Where("id = ?", id).Delete(&Book)
	return Book
}
