package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/akgmage/go-sequel/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("COntent-Type", "pkg/publication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}