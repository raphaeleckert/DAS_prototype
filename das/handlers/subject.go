/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package handlers

import (
	"html/template"
	"net/http"

	"dasagilestudieren/models"
)

// This file contains all the handlers for the subject resource

type SubjectPage struct {
	Subject models.Subject
	Courses []models.Clickable
}

func SubjectHandler(w http.ResponseWriter, r *http.Request) {
	// GET: Retruns the subject page
	if r.Method == http.MethodGet {
		subjectid := r.URL.Query().Get("subjectid")
		subject, err := models.GetSubject(subjectid)
		courses, err := models.GetCourseBasicListBySubject(subjectid)

		p := SubjectPage{
			Subject: subject,
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
