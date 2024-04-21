package models

import (
	"github.com/jinzhu/gorm"
	"github.com/murali-muniyan/go-book-store/pkg/config"
)

type Book struct {
	gorm.Model
	Name         string `gorm:"name"`
	Author       string `gorm:"author"`
	Publications string `gorm:"publications"`
}

func Migrate() {
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook(book *Book) (*Book, error) {
	db := config.GetDB()
	db.Save(&book)

	if db.Error != nil {
		return nil, db.Error
	}

	return book, nil
}

func GetBooks() ([]Book, error) {
	db := config.GetDB()

	var books []Book

	db.Find(&books)

	if db.Error != nil {
		return nil, db.Error
	}

	return books, nil
}

func GetBook(id int64) (*Book, error) {
	db := config.GetDB()

	var book Book

	db.Where("id = $1", id).Find(&book)

	if db.Error != nil {
		return nil, db.Error
	}

	return &book, nil
}

func DeleteBook(id int64) error {
	db := config.GetDB()

	var book Book

	db.Where("id = $1", id).Delete(&book)

	if db.Error != nil {
		return db.Error
	}

	return nil
}

func UpdateBook(id int64, book *Book) (*Book, error) {
	db := config.GetDB()

	oldBook, err := GetBook(id)
	if err != nil {
		return nil, err
	}

	if book.Author != "" {
		oldBook.Author = book.Author
	}

	if book.Name != "" {
		oldBook.Name = book.Name
	}

	if book.Publications != "" {
		oldBook.Publications = book.Publications
	}

	db.Save(oldBook)

	if db.Error != nil {
		return nil, db.Error
	}

	return GetBook(id)
}
