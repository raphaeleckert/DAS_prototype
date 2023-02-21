package handlers

import (
	"html/template"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/registration/login.html")
		t.ExecuteTemplate(writer, "base", nil)
	}

}
