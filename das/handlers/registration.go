package handlers

import (
	"html/template"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"C:/Users/rapha/Desktop/Bachelorarbeit/DAS_prototype/das/templates/base.html",
			"C:/Users/rapha/Desktop/Bachelorarbeit/DAS_prototype/das/templates/registration/login.html")
		t.ExecuteTemplate(writer, "base", nil)
	}

}
