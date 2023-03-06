package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"dasagilestudieren/models"
)

type TopicPage struct {
	User      models.User
	TopicData models.Topic
}

func TopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user := r.Context().Value("user")
		fmt.Println(user)
		p := TopicPage{TopicData: models.GetTopic("topicid")}

		t, _ := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/topic/topic_student.html")
		t.ExecuteTemplate(w, "base", p)
	}

}

type TopicTeacherPage struct {
	TopicId           string
	TopicTitle        string
	TopicRef          string
	TopicDetail       string
	TopicSolutionIdea string
	TopicTags         []string
	TopicSupporters   string
	TopicRemark       string
	TopicImportance   string
	TopicSubject      string
	TopicCourses      []struct {
		CourseName string
		CourseId   string
	}
}

func TopicTeacherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		topic := models.GetTopic("topicid")
		coursesWithSubject := []models.Course{models.GetCourse("courseid")}
		p := TopicTeacherPage{
			TopicId:           topic.ID,
			TopicTitle:        topic.Title,
			TopicRef:          topic.Reference,
			TopicDetail:       topic.Detail,
			TopicSolutionIdea: topic.SolutionIdea,
			TopicTags:         topic.Tags,
			TopicSupporters:   topic.RequiredSupporters,
			TopicRemark:       topic.Remark,
			TopicImportance:   topic.Importance,
			TopicSubject:      topic.Subject.Name,
			TopicCourses: []struct {
				CourseName string
				CourseId   string
			}{
				{
					CourseName: "Some Course",
					CourseId:   coursesWithSubject[0].ID,
				},
			},
		}
		t, _ := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/topic/topic_teacher.html")
		t.ExecuteTemplate(w, "base", p)
	}
}

func TopicEditHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		topic := models.GetTopic("topicid")
		coursesWithSubject := []models.Course{models.GetCourse("courseid")}
		p := TopicTeacherPage{
			TopicId:           topic.ID,
			TopicTitle:        topic.Title,
			TopicRef:          topic.Reference,
			TopicDetail:       topic.Detail,
			TopicSolutionIdea: topic.SolutionIdea,
			TopicTags:         topic.Tags,
			TopicSupporters:   topic.RequiredSupporters,
			TopicRemark:       topic.Remark,
			TopicImportance:   topic.Importance,
			TopicSubject:      topic.Subject.Name,
			TopicCourses: []struct {
				CourseName string
				CourseId   string
			}{
				{
					CourseName: "Some Course",
					CourseId:   coursesWithSubject[0].ID,
				},
			},
		}
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/topic/topic_edit.html")
		t.ExecuteTemplate(w, "htmx_wrapper", p)
	}
}
