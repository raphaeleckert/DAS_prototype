package handlers

import (
	"html/template"
	"net/http"
)

func Tryout(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../static/templates/base.html",
			"../static/templates/sidebar.html",
			"../static/templates/tryout.html")
		t.ExecuteTemplate(writer, "base", nil)
	}
}
