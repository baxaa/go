package main

import (
	"awesomeProject/module"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

var database *sql.DB

func Indexhandler(w http.ResponseWriter, r *http.Request) {
	conn := "user=baxa password=Uhupup667 dbname=go sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	Users := []module.User{}
	for rows.Next() {
		u := module.User{}
		err := rows.Scan(&u.Name, &u.Id)
		if err != nil {
			panic(err)
			continue
		}
		Users = append(Users, u)
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, Users)
}

func main() {
	//createUser()
	db, err := sql.Open("postgresql", "baxa:Uhupup667/go")
	if err != nil {
		panic(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", Indexhandler)
	fmt.Println("wait a secund")
	http.ListenAndServe(":8181", nil)
}
