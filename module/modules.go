package module

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

var database *sql.DB

func createUser() {
	conn := "user=baxa password=Uhupup667 dbname=go sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into Users (name, id, password) values ('a', $1, $2)", 123, 12)
	if err != nil {
		panic(err)
	}
	fmt.Print(result.RowsAffected())
}

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
	Users := []User{}
	for rows.Next() {
		u := User{}
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
