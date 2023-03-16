package handlers

import (
	"html/template"
	"net/http"

	"dasagilestudieren/models"
)

type SubjectPage struct {
	Subject models.Subject
	Courses []models.Clickable
}

func SubjectHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		subjectid := r.URL.Query().Get("subjectid")
		subject, err := models.GetSubject(subjectid)
		courses := []models.Clickable{models.GetCourseBasic("course1"), models.GetCourseBasic("course2")}

		p := SubjectPage{
			Subject: *subject,
			Courses: courses,
		}

		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/subject/subject.html")
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
