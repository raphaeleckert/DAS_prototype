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
	Courses   []models.Course
	Subjects  []models.Subject
	Teams     []models.Team
	isTeacher bool
}

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user, ok := r.Context().Value("user").(models.User)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//TODO: Get real list of courses
		teams := []models.Team{models.GetTeam("team1"), models.GetTeam("team2")}
		p := LandingPage{}
		if user.IsTeacher {
			courses := []models.Course{models.GetCourse("course1"), models.GetCourse("course2")}
			subjects := []models.Subject{models.GetSubject("sub1"), models.GetSubject("sub2")}
			p = LandingPage{
				Courses:  courses,
				Subjects: subjects,
				Teams:    teams,
			}
		} else {
			p = LandingPage{
				Courses:  []models.Course{},
				Subjects: []models.Subject{},
				Teams:    teams,
			}
		}

		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/landing.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "base", p)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
