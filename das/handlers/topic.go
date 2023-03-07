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

//func TopicHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodGet {
//		user := r.Context().Value("user")
//		fmt.Println(user)
//		p := TopicPage{TopicData: models.GetTopic("topicid")}
//
//		t, _ := template.ParseFiles(
//			"../resources/templates/base.html",
//			"../resources/templates/topic/topic_student.html")
//		t.ExecuteTemplate(w, "base", p)
//	}
//
//}

type TopicInfoPage struct {
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

func TopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/topic/topic.html")
		t.ExecuteTemplate(w, "base", nil)
	}
}

type FormData struct {
	Title        string
	Ref          string
	Tags         []string
	Supporters   string
	Remark       string
	Importance   string
	Detail       string
	SolutionIdea string
}

func TopicInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		r.ParseForm()
		title := r.Form.Get("title")
		ref := r.Form.Get("ref")
		tags := r.Form["tags"]
		supporters := r.Form.Get("supporters")
		remark := r.Form.Get("remark")
		importance := r.Form.Get("importance")
		detail := r.Form.Get("detail")
		solutionIdea := r.Form.Get("solutionIdea")

		topicUpdate := FormData{
			Title:        title,
			Ref:          ref,
			Tags:         tags,
			Supporters:   supporters,
			Remark:       remark,
			Importance:   importance,
			Detail:       detail,
			SolutionIdea: solutionIdea,
		}

		// TODO: update data in db
		fmt.Printf("%+v", topicUpdate)
	}
	if r.Method == http.MethodGet || r.Method == http.MethodPut {
		topic := models.GetTopic("topicid")
		coursesWithSubject := []models.Course{models.GetCourse("courseid")}
		p := TopicInfoPage{
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
			"../resources/templates/topic/topic_info.html")
		t.ExecuteTemplate(w, "htmx_wrapper", p)
	}
}

type TopicEditPage struct {
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
	SupporterOptions  []string
	TagOptions        []string
	ImportanceOptions []string
}

func TopicEditHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tags := []string{"a tag", "another tag", "yet another tag"}
		importanceOptions := []string{models.IMP_ESSENTIAL, models.IMP_MUST, models.IMP_SHOULD, models.IMP_COULD, models.IMP_VOLUNTARY}
		supporterOptions := []string{models.SUP_NONE, models.SUP_ONE, models.SUP_TWO, models.SUP_HALF, models.SUP_BUT, models.SUP_ALL}
		topic := models.GetTopic("topicid")
		coursesWithSubject := []models.Course{models.GetCourse("courseid")}
		p := TopicEditPage{
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
			SupporterOptions:  supporterOptions,
			TagOptions:        tags,
			ImportanceOptions: importanceOptions,
		}
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/topic/topic_edit.html")
		t.ExecuteTemplate(w, "htmx_wrapper", p)
	}
}
