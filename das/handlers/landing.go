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

// This file contains handlers for the landing page (main page)

type LandingPage struct {
	Courses   []models.Clickable
	Subjects  []models.Clickable
	Teams     []models.Clickable
	IsTeacher bool
}

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	// GET: Returns the landing page
	if r.Method == http.MethodGet {
		user, ok := r.Context().Value("user").(models.User)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//TODO: Get real list of courses
		team, _ := models.GetTeamBasic("team1")
		p := LandingPage{}
		if user.IsTeacher {
			course, err := models.GetCourseBasic("course1")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			subjects, _ := models.GetSubjectListBasic(user.Username)
			p = LandingPage{
				Courses:   []models.Clickable{course, course},
				Subjects:  subjects,
				Teams:     []models.Clickable{team, team},
				IsTeacher: true,
			}
		} else {
			p = LandingPage{
				Courses:   []models.Clickable{},
				Subjects:  []models.Clickable{},
				Teams:     []models.Clickable{team, team},
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
