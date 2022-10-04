package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm err: %v", err)
		return
	}
	fmt.Fprintf(writer, "POST request successful")
	name := request.FormValue("name")
	age := request.FormValue("age")
	fmt.Fprintf(writer, "name: %s", name)
	fmt.Fprintf(writer, "age: %s", age)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 Not Found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, "Hello World!!")
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Stating Server at port: 8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalln(err)
	}
}
