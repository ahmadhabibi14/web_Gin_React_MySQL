package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          int    //   `json:"id" gorm:"primary_key"`
	Title       string // `json:"title"`
	Description string // `json:"description"`
	Author      string // `json:"author"`
	Year        uint16 // `json:"year"`
}

// create a book
func CreateBook(db *gorm.DB, books *Book) (err error) {
	err = db.Create(books).Error
	if err != nil {
		return err
	}
	return nil
}

// get all books
func GetBooks(db *gorm.DB, books *[]Book) (err error) {
	err = db.Find(books).Error
	if err != nil {
		return err
	}
	return nil
}

// get book by id
func GetBook(db *gorm.DB, books *Book, id int) (err error) {
	err = db.Where("id = ?", id).First(books).Error
	if err != nil {
		return err
	}
	return nil
}

// update book
func UpdateBook(db *gorm.DB, books *Book) (err error) {
	db.Save(books)
	return nil
}

// delete book
func DeleteBook(db *gorm.DB, books *Book, id int) (err error) {
	db.Where("id = ?", id).Delete(books)
	return nil
}
