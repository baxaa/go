package main

import (
	"fmt"
	"io"
	"net/http"
)

func getroot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got/request\n")
	io.WriteString(w, "This is my website\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got/ hello request\n")
	io.WriteString(w, "Hello, Http\n")
}

func main() {
	//http.HandleFunc("/root", getroot)
	//http.HandleFunc("/hello", getHello)
	//
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	return
	//}
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])

	}
}
