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

		fmt.Printf("%+v", time.Now())
		fmt.Printf("%+v", beginDateTime)
		fmt.Printf("%+v", finalDateTime)
		fmt.Printf("%+v", closeDateTime)
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

		newCourse := models.Course{
			ID:        "newid",
			Subject:   models.GetSubject(subjectId),
			Term:      models.GetTerm(termId),
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

type CreateCourseForm struct {
	Terms   []models.Clickable
	Subject models.Clickable
}

func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		subjectId := r.URL.Query().Get("subjectid")

		terms := []models.Clickable{models.GetTermBasic("term1"), models.GetTermBasic("term2")}
		subject := models.GetSubjectBasic(subjectId)

		p := CreateCourseForm{
			Terms:   terms,
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

func ValidateCourseDatesHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	beginDate := r.FormValue("begin-date")
	finalDate := r.FormValue("final-date")
	closeDate := r.FormValue("close-date")

	// Use the values as needed
	fmt.Println("Begin Date:", beginDate)
	fmt.Println("Final Date:", finalDate)
	fmt.Println("Close Date:", closeDate)

}
