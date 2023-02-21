package main

import (
	"fmt"
	"log"
	"net/http"

	"dasagilestudieren/handlers"
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func router() {

	server := http.Server{
		Addr:    ":8080",
		Handler: nil, // set to nil to use the default ServeMux
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/login/", handlers.Login)
	http.HandleFunc("/tryout/", handlers.Tryout)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	log.Fatal(server.ListenAndServe())
}

func main() {

	router()
}
