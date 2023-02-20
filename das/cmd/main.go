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
	t, _ := template.ParseFiles(
		"C:/Users/rapha/Desktop/Bachelorarbeit/DAS_prototype/das/templates/base.html",
		"C:/Users/rapha/Desktop/Bachelorarbeit/DAS_prototype/das/templates/registration/login.html")
	t.ExecuteTemplate(writer, "base", nil)
}

func router() {

	server := http.Server{
		Addr:    ":8080",
		Handler: nil, // set to nil to use the default ServeMux
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/login/", login)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	log.Fatal(server.ListenAndServe())
}

func main() {

	router()
}