package Controller

import (
	"encoding/json"
	"fmt"
	"github.com/baxaa/e-commerce/Helper"
	"github.com/baxaa/e-commerce/db"
	"github.com/baxaa/e-commerce/model"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a home page")
	fmt.Println("Success")

}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []model.Book
	db.DB.Find(&books)

	if len(books) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&books)
	fmt.Println("Success")

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)
	createdbook := db.DB.Create(&book)
	err := createdbook.Error
	Helper.Check(err)
	w.WriteHeader(http.StatusOK)
	fmt.Println("Success")

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book model.Book
	db.DB.First(&book, params["id"])

	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&book)
	fmt.Println("Success")

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book model.Book
	db.DB.First(&book, params["id"])
	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db.DB.Unscoped().Delete(&book)
	w.WriteHeader(http.StatusOK)
	fmt.Println("Success")

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book model.Book
	db.DB.First(&book, params["id"])
	json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book.Title)
	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db.DB.Model(&book).Where("ID = ?", params["id"]).Update("Title", book.Title)
	db.DB.Model(&book).Where("ID = ?", params["id"]).Update("Description", book.Description)
	//fmt.Println(params)
	db.DB.Save(&book)
	w.WriteHeader(http.StatusOK)
	fmt.Println("Success")

}

func SearchBook(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	//fmt.Println(title)
	var books []model.Book
	db.DB.Find(&books)

	var res []model.Book
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
			res = append(res, book)
		}
	}

	if len(res) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(res)
	fmt.Println("Success")
}

func SortedBooks(w http.ResponseWriter, r *http.Request) {
	var order string
	if r.URL.Query().Get("order") == "asc" {
		order = "asc"
	} else {
		order = "desc"
	}
	//fmt.Println(r.URL.Query().Get("order"))

	var books []model.Book

	db.DB.Order("cost " + order).Find(&books)
	//fmt.Println(db.DB.Order("cost " + order).Find(&books))
	if len(books) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&books)
	w.WriteHeader(http.StatusOK)
}
