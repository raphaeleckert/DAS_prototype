package router

import (
	"net/http"
	"log"
)

func router() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/login/", login)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8080", mux))
}