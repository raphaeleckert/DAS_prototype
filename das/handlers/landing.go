package handlers

import (
	"html/template"
	"net/http"

	"dasagilestudieren/models"
)

type LandingPage struct {
	Courses   []models.Clickable
	Subjects  []models.Clickable
	Teams     []models.Clickable
	IsTeacher bool
}

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user, ok := r.Context().Value("user").(models.User)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//TODO: Get real list of courses
		teams := []models.Clickable{models.GetTeamBasic("team1"), models.GetTeamBasic("team2")}
		p := LandingPage{}
		if user.IsTeacher {
			courses := []models.Clickable{models.GetCourseBasic("course1"), models.GetCourseBasic("course2")}
			subjects, _ := models.GetSubjectListBasic(user.Username)
			p = LandingPage{
				Courses:   courses,
				Subjects:  subjects,
				Teams:     teams,
				IsTeacher: true,
			}
		} else {
			p = LandingPage{
				Courses:   []models.Clickable{},
				Subjects:  []models.Clickable{},
				Teams:     teams,
				IsTeacher: false,
			}
		}

		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/landing.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "base", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
