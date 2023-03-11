package handlers

import (
	"html/template"
	"net/http"

	"dasagilestudieren/models"
)

type CoursePage struct {
	Reviews    []models.Clickable
	Teams      []models.Clickable
	Topics     []models.Clickable
	CourseData models.Course
}

func CourseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		courseId := r.URL.Query().Get("courseid")

		course := models.GetCourse(courseId)

		//TODO: Get real list of courses
		reviews := []models.Clickable{models.GetReviewBasic("review1"), models.GetReviewBasic("review2")}
		teams := []models.Clickable{models.GetTeamBasic("team1"), models.GetTeamBasic("team2")}
		topics := []models.Clickable{models.GetTopicBasic("topic1"), models.GetTopicBasic("topic2")}

		p := CoursePage{
			Reviews:    reviews,
			Teams:      teams,
			Topics:     topics,
			CourseData: course,
		}

		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/course/course.html")
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
