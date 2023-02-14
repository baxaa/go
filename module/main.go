package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type User struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	password string `json:"password"`
}

//var database *sql.DB

func Indexhandler() {
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
	Users := []User{}
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Name, &u.Id, &u.password)
		if err != nil {
			panic(err)
			continue
		}
		Users = append(Users, u)
	}
	for _, p := range Users {
		fmt.Println(p.Id, p.Name)
	}
	//tmpl, _ := template.ParseFiles("templates/index.html")
	//tmpl.Execute(w, Users)
}

func createUser() {
	conn := "user=baxa password=Uhupup667 dbname=go sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into Users (name, id, password) values ('a', $1, $2)", 111, 123)
	if err != nil {
		panic(err)
	}
	fmt.Print(result.RowsAffected())
}

func main() {
	createUser()
	//conn := "user=baxa password=Uhupup667 dbname=go sslmode=disable"
	//db, err := sql.Open("postgres", conn)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//database = db
	//
	//http.HandleFunc("/magazine", Indexhandler)
	//fmt.Println("wait a secund")
	//http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	return
	//}
	//Indexhandler();
	Indexhandler()
}
