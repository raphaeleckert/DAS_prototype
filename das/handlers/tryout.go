package handlers

import (
	"html/template"
	"net/http"
)

type TryoutPage struct {
	Test string
}

func Tryout(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p := TryoutPage{Test: `<button class="w3-red">Hi</button>`}

		t, _ := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/sidebar.html",
			"../resources/templates/tryout.html")
		t.ExecuteTemplate(writer, "base", p)
	}
}
