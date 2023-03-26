package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"dasagilestudieren/models"
	"dasagilestudieren/utils"
)

type CoursePage struct {
	Reviews    []models.Clickable
	Teams      []models.Clickable
	Topics     []models.Clickable
	CourseData models.Course
}

func CourseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user, ok := r.Context().Value("user").(models.User)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		subjectId := r.URL.Query().Get("subjectid")
		r.ParseForm()
		termId := r.PostFormValue("term")
		beginDate := r.FormValue("begin-date")
		finalDate := r.FormValue("final-date")
		closeDate := r.FormValue("close-date")

		beginDateTime, _ := time.Parse("2006-01-02T15:04", beginDate)
		finalDateTime, _ := time.Parse("2006-01-02T15:04", finalDate)
		closeDateTime, _ := time.Parse("2006-01-02T15:04", closeDate)
		currentTime := time.Now().UTC()
		if beginDateTime.Before(currentTime) || finalDateTime.Before(currentTime) || closeDateTime.Before(currentTime) {
			utils.ShowErrorMsg("All dates must be in the future.", w)
			return
		}
		if beginDateTime.After(finalDateTime) {
			utils.ShowErrorMsg("Begin Date should be before Final Date.", w)
			return
		}
		if finalDateTime.After(closeDateTime) {
			utils.ShowErrorMsg("Final Date should be before Close Date.", w)
			return
		}

		subject, err := models.GetSubject(subjectId)
		term, err := models.GetTerm(termId)
		if err != nil {

		}

		newCourse := models.Course{
			ID:        "newid",
			Subject:   subject,
			Term:      term,
			Owner:     user.Username,
			BeginDate: beginDateTime,
			FinalDate: finalDateTime,
			CloseDate: closeDateTime,
			Remark:    "",
			Note:      "",
			Topics:    []string{},
		}

		//TODO Save to db
		fmt.Printf("%+v", newCourse)

		newAdress := "/course?courseid=" + newCourse.ID
		w.Header().Set("HX-Redirect", newAdress)
	}
	if r.Method == http.MethodGet {
		courseId := r.URL.Query().Get("courseid")

		course, _ := models.GetCourse(courseId)

		//TODO: Get real list of courses
		review, _ := models.GetReviewBasic("review1")
		team, _ := models.GetTeamBasic("team1")
		topics, err := models.GetTopicsByCourseBasic(courseId)

		p := CoursePage{
			Reviews:    []models.Clickable{review, review, review},
			Teams:      []models.Clickable{team, team, team},
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

type CreateCourseForm struct {
	Terms   []models.Clickable
	Subject models.Clickable
}

func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		subjectId := r.URL.Query().Get("subjectid")

		term, _ := models.GetTermBasic("term1")
		subject, _ := models.GetSubjectBasic(subjectId)

		p := CreateCourseForm{
			Terms:   []models.Clickable{term, term, term},
			Subject: subject,
		}

		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/course/course_create.html")
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

type CourseTopicPage struct {
	CourseData   models.Course
	CourseTopics struct {
		Selected   []models.Topic
		Unselected []models.Topic
	}
}

func CourseTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		courseId := r.URL.Query().Get("courseid")
		topicId := r.URL.Query().Get("topicid")

		err := models.AssignTopicToCourse(courseId, topicId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == http.MethodGet || r.Method == http.MethodPost {
		courseId := r.URL.Query().Get("courseid")
		courseData, err := models.GetCourse(courseId)
		courseTopics, err := models.GetTopicsByCourse(courseId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p := CourseTopicPage{
			CourseData:   courseData,
			CourseTopics: courseTopics,
		}

		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/course/course_assign_topic.html")
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
