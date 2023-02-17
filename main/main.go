package main

import (
	"net/http"
	"github.com/baxaa/e-commerce/CRUD"
)

func main() {
	http.HandleFunc("/create", createUser)
	http.ListenAndServe(":8181", nil)
	// CRUD.Indexhandler()
}
