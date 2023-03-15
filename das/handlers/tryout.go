package handlers

import (
	"html/template"
	"net/http"
)

type TryoutPage struct {
	Test string
}

func Tryout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p := TryoutPage{Test: `<button class="w3-red">Hi</button>`}

		t, _ := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/tryout.html")
		t.ExecuteTemplate(w, "base", p)
	}
}

func RessourceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		// Update Data
	}
	if r.Method == http.MethodGet || r.Method == http.MethodPatch {
		// Return Page
	} 
}
