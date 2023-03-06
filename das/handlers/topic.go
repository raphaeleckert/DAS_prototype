package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"dasagilestudieren/models"
	"dasagilestudieren/utils"
)

type TopicPage struct {
	User      utils.User
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
