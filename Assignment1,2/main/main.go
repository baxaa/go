package main

import (
	"github.com/baxaa/e-commerce/Assignment1,2/CRUD"
	// "database/sql"
	// "fmt"
	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "uhupup667"
	dbname   = "go"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	CRUD.Routes()

}
