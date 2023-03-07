package handlers

import (
	"html/template"
	"net/http"

	"dasagilestudieren/models"
)

type Clickable struct {
	Text      string
	Reference string
}

type LandingPage struct {
	Courses   Clickable
	Subjects  Clickable
	Teams     Clickable
	IsTeacher bool
}

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user, ok := r.Context().Value("user").(models.User)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if user.IsTeacher {
			// Handle the case where the user is a teacher.
		}

		t, err := template.ParseFiles(
			"../resources/templates/registration/login.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
