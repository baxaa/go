package main

import (
	// "database/sql"
	// "fmt"
	_ "github.com/lib/pq"
	"github.com/baxaa/e-commerce/CRUD"
)
const (
	user = "postgres"
	password = "uhupup667"
	dbname = "go"
)

func check(err error){
	if err != nil{
		panic(err)
	}
}

func main() {
	CRUD.Routes()
	// conn := fmt.Sprintf("user = %s password = %s dbname = %s sslmode = disable", user, password, dbname)
	// db, err := sql.Open("postgres", conn)
	// check(err)
	// defer db.Close()
	// var n string
	// fmt.Scanln(&n)
	// var i int
	// fmt.Scanln(&i)
	// var p string
	// fmt.Scanln(&p)
	// result, er := db.Exec("insert into Users (name, id, password) values ($1, $2, $3)", n, i, p)
	// check(er)
	// fmt.Print(result.RowsAffected())
}
