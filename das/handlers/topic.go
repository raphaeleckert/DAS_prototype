package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"dasagilestudieren/models"
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
	TopicCourses      []struct {
		CourseName string
		CourseId   string
	}
}

type TopicBasePage struct {
	TopicId string
}

func TopicHandler(w http.ResponseWriter, r *http.Request) {
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
		blankTopic := models.Topic{
			ID:                 "blankid",
			Subject:            models.GetSubject(subjectId),
			Title:              "-",
			Detail:             "-",
			Reference:          "-",
			SolutionIdea:       "-",
			Remark:             "-",
			Tags:               []string{},
			Importance:         models.IMP_SHOULD,
			RequiredSupporters: models.SUP_HALF,
		}

		//TODO: Give blank Topic a unique id and save to db
		fmt.Printf("%+v", blankTopic)
		newAdress := "/topic?topicid=" + blankTopic.ID
		w.Header().Set("HX-Redirect", newAdress)
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
		fmt.Printf("%+v", topicId)

		topic := models.GetTopic(topicId)
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
		topicId := r.URL.Query().Get("topicid")
		tags := []string{"a tag", "another tag", "yet another tag"}
		importanceOptions := []string{models.IMP_ESSENTIAL, models.IMP_MUST, models.IMP_SHOULD, models.IMP_COULD, models.IMP_VOLUNTARY}
		supporterOptions := []string{models.SUP_NONE, models.SUP_ONE, models.SUP_TWO, models.SUP_HALF, models.SUP_BUT, models.SUP_ALL}
		topic := models.GetTopic(topicId)
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

type TopicList struct {
	TopicTable []models.Topic
}

func TopicListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		subjectid := r.URL.Query().Get("subjectid")
		fmt.Printf("%s Deleted", subjectid)

	}
	if r.Method == http.MethodGet || r.Method == http.MethodDelete {
		subjectid := r.URL.Query().Get("subjectid")
		topics := models.GetTopicsBySubject(subjectid)

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
