package Router

import (
	"fmt"
	"github.com/baxaa/e-commerce/Controller"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", Controller.HomePage)
	myRouter.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if r.URL.Query().Get("title") != "" {
				Controller.SearchBook(w, r)
			} else if r.URL.Query().Get("order") != "" {
				Controller.SortedBooks(w, r)
			} else {
				Controller.GetAllBooks(w, r)
			}
		} else if r.Method == "POST" {
			Controller.CreateBook(w, r)
		}
	})
	myRouter.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			Controller.DeleteBook(w, r)
		} else if r.Method == "PUT" {
			Controller.UpdateBook(w, r)
		} else {
			Controller.GetBook(w, r)
		}
	})
	//myRouter.HandleFunc("/books", Controller.GetAllBooks).Methods("GET")
	//myRouter.HandleFunc("/books", Controller.CreateBook).Methods("POST")
	//myRouter.HandleFunc("/books/{id}", Controller.DeleteBook).Methods("DELETE")
	//myRouter.HandleFunc("/books/{id}", Controller.GetBook).Methods("GET")
	//myRouter.HandleFunc("/books/{id}", Controller.UpdateBook).Methods("PUT")
	//myRouter.HandleFunc("/books", Controller.SearchBook).Methods("GET")
	//myRouter.HandleFunc("/books", Controller.SortedBooks).Methods("GET")
	err := http.ListenAndServe(":8080", myRouter)
	if err != nil {
		fmt.Println(err)
	}
}
