package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func login(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("/templates/registration/login.html")
	t.Execute(writer, nil)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login/", login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
