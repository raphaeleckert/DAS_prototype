package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"dasagilestudieren/models"

	"github.com/google/uuid"
)

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
	TopicCourses      []models.Clickable
}

type TopicBasePage struct {
	TopicId string
}

func TopicHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Method == http.MethodDelete {
	//	topicId := r.URL.Query().Get("topicid")
	//
	//}
	if r.Method == http.MethodGet {
		topicId := r.URL.Query().Get("topicid")
		p := TopicBasePage{TopicId: topicId}
		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/topic/topic.html")
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
	if r.Method == http.MethodPost {
		subjectId := r.URL.Query().Get("subjectid")
		subject, err := models.GetSubject(subjectId)
		r.ParseForm()
		title := r.Form.Get("title")
		ref := r.Form.Get("ref")
		tags := r.Form["tags"]
		supporters := r.Form.Get("supporters")
		remark := r.Form.Get("remark")
		importance := r.Form.Get("importance")
		detail := r.Form.Get("detail")
		solutionIdea := r.Form.Get("solutionIdea")

		topicNew := models.Topic{
			ID:                 uuid.New().String(),
			Title:              title,
			Reference:          ref,
			Tags:               tags,
			RequiredSupporters: supporters,
			Remark:             remark,
			Importance:         importance,
			Detail:             detail,
			SolutionIdea:       solutionIdea,
			Subject:            subject,
		}

		err = models.CreateTopic(topicNew)
		//TODO: Give blank Topic a unique id and save to db
		fmt.Printf("%+v", topicNew)
		newAdress := "/topic?topicid=" + topicNew.ID

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Print("hi")
		w.Header().Set("HX-Redirect", newAdress)
		return
	}
	if r.Method == http.MethodPut {
		topicId := r.URL.Query().Get("topicid")
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
		fmt.Printf("%+v", topicId)
	}
	if r.Method == http.MethodGet || r.Method == http.MethodPut {
		topicId := r.URL.Query().Get("topicid")

		topic, err := models.GetTopic(topicId)
		course, err := models.GetCourse("courseid")
		coursesWithSubject, err := models.GetCourseBasicListBySubject(course.Subject.ID)
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
			TopicCourses:      coursesWithSubject,
		}
		fmt.Printf("%+v", topic)
		fmt.Printf("topicid %v", topicId)
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/topic/topic_info.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
	TopicCourses      []models.Clickable
	SupporterOptions  []string
	TagOptions        []string
	ImportanceOptions []string
}

func TopicEditHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		topicId := r.URL.Query().Get("topicid")
		tags := []string{"a tag", "another tag", "yet another tag"}
		importanceOptions := []string{models.IMP_ESSENTIAL, models.IMP_MUST, models.IMP_SHOULD, models.IMP_COULD, models.IMP_VOLUNTARY}
		supporterOptions := []string{models.SUP_NONE, models.SUP_ONE, models.SUP_TWO, models.SUP_HALF, models.SUP_BUT, models.SUP_ALL}
		topic, err := models.GetTopic(topicId)
		coursesWithSubject, err := models.GetCourseBasicListBySubject("course1")
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
			TopicCourses:      coursesWithSubject,
			SupporterOptions:  supporterOptions,
			TagOptions:        tags,
			ImportanceOptions: importanceOptions,
		}
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/topic/topic_edit.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type TopicCreatePage struct {
	TopicSubject      models.Clickable
	TopicCourses      []models.Clickable
	SupporterOptions  []string
	TagOptions        []string
	ImportanceOptions []string
}

func TopicCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		subjectId := r.URL.Query().Get("subjectid")
		subject, err := models.GetSubjectBasic(subjectId)
		tags := []string{"a tag", "another tag", "yet another tag"}
		importanceOptions := []string{models.IMP_ESSENTIAL, models.IMP_MUST, models.IMP_SHOULD, models.IMP_COULD, models.IMP_VOLUNTARY}
		supporterOptions := []string{models.SUP_NONE, models.SUP_ONE, models.SUP_TWO, models.SUP_HALF, models.SUP_BUT, models.SUP_ALL}
		p := TopicCreatePage{
			TopicSubject:      subject,
			SupporterOptions:  supporterOptions,
			TagOptions:        tags,
			ImportanceOptions: importanceOptions,
		}
		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/topic/topic_create.html")
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

type TopicList struct {
	TopicTable []models.Topic
}

func TopicListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		topicId := r.URL.Query().Get("topicid")
		fmt.Printf("%s Deleted", topicId)
	}
	if r.Method == http.MethodGet {
		subjectId := r.URL.Query().Get("subjectid")
		topics, err := models.GetTopicsBySubject(subjectId)

		p := TopicList{
			TopicTable: topics,
		}

		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/topic/topic_list.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
