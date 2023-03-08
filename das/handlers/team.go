package handlers

import (
	"dasagilestudieren/models"
	"fmt"
	"html/template"
	"net/http"
)

type TeamTabPage struct {
	Team      struct{ Name, ID string }
	TableData []models.Clickable
}

type TeamBasePage struct {
	Name, ID, Subject, Term string
	Member                  []string
	IsTeacher               bool
}

func TeamOverviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user := r.Context().Value("user").(models.User)
		teamid := r.URL.Query().Get("teamid")
		team := models.GetTeam(teamid)

		p := TeamBasePage{
			Name:      fmt.Sprintf("Team #%d", team.Number),
			ID:        team.ID,
			Subject:   team.Course.Subject.Name,
			Term:      team.Course.Term.Name,
			Member:    team.Member,
			IsTeacher: user.IsTeacher,
		}

		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/team/team.html")
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

//type TeamOpenTableRow struct {
//	Text      string
//	Reference string
//}

func TeamOpenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		teamid := r.URL.Query().Get("teamid")
		topics := []models.Clickable{
			models.GetTopicBasic("topic1"),
			models.GetTopicBasic("topic2"),
			models.GetTopicBasic("topic3"),
		}
		p := TeamTabPage{TableData: topics, Team: models.GetTeamBasic(teamid)}

		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_open.html")
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

func TeamWorkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_work.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func TeamReadyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_ready.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func TeamDoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_done.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func TeamReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_review.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
