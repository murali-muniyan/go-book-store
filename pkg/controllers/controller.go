package controllers

import (
	"net/http"

	"github.com/murali-muniyan/go-book-store/pkg/models"
	"github.com/murali-muniyan/go-book-store/pkg/utils"
)

func CreateBook(wtr http.ResponseWriter, req *http.Request) {
	wtr.Header().Set("content-type", "application/json")
	var book models.Book

	err := utils.ParseRequestBody(req, &book)
	if err != nil {
		utils.HandleErr(err, wtr)

		return
	}

	created, err := models.CreateBook(&book)
	if err != nil {
		utils.HandleErr(err, wtr)

		return
	}

	utils.WiteToResponse(wtr, http.StatusCreated, created)
}

func GetBooks(wtr http.ResponseWriter, req *http.Request) {
	wtr.Header().Set("content-type", "application/json")
	books, err := models.GetBooks()
	if err != nil {
		utils.HandleErr(err, wtr)

		return
	}

	utils.WiteToResponse(wtr, http.StatusCreated, &books)

}

func GetBook(wtr http.ResponseWriter, req *http.Request) {
	wtr.Header().Set("content-type", "application/json")
	bookID := utils.GetID(req)

	book, err := models.GetBook(bookID)
	if err != nil {
		utils.HandleErr(err, wtr)

		return
	}

	utils.WiteToResponse(wtr, http.StatusCreated, &book)
}

func UpdateBook(wtr http.ResponseWriter, req *http.Request) {
	wtr.Header().Set("content-type", "application/json")
	var book models.Book
	bookID := utils.GetID(req)

	err := utils.ParseRequestBody(req, &book)
	if err != nil {
		utils.HandleErr(err, wtr)

		return
	}

	updated, err := models.UpdateBook(bookID, &book)
	if err != nil {
		utils.HandleErr(err, wtr)

		return
	}

	utils.WiteToResponse(wtr, http.StatusCreated, updated)
}

func DeleteBook(wtr http.ResponseWriter, req *http.Request) {
	wtr.Header().Set("content-type", "application/json")
	bookID := utils.GetID(req)

	err := models.DeleteBook(bookID)
	if err != nil {
		utils.HandleErr(err, wtr)

		return
	}

	utils.WiteToResponse(wtr, http.StatusCreated, nil)
}
