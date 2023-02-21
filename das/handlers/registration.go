package handlers

import (
	"html/template"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../static/templates/base.html",
			"../static/templates/registration/login.html")
		t.ExecuteTemplate(writer, "base", nil)
	}

}
