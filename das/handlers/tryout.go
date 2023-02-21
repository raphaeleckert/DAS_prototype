package handlers

import (
	"html/template"
	"net/http"
)

func Tryout(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/sidebar.html",
			"../resources/templates/tryout.html")
		t.ExecuteTemplate(writer, "base", nil)
	}
}
