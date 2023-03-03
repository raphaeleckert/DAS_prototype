package handlers

import (
	"html/template"
	"net/http"
)

type InTableLink struct {
	Text      string
	Reference string
}

func TeamOverviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p := struct{ TeamName string }{TeamName: "Example Team Name"}

		t, _ := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/team/team.html")
		t.ExecuteTemplate(w, "base", p)
	}

}

type TeamOpenTableRow struct {
	Text      string
	Reference string
}

func TeamOpenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		topics := []TeamOpenTableRow{
			{Text: "Topic1", Reference: "Topic1Ref"},
			{Text: "Topic2", Reference: "Topic2Ref"},
			{Text: "Topic3", Reference: "Topic3Ref"},
		}
		p := struct{ TableData []TeamOpenTableRow }{TableData: topics}

		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_open.html")
		t.ExecuteTemplate(w, "htmx_wrapper", p)
	}

}

func TeamWorkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_work.html")
		t.ExecuteTemplate(w, "htmx_wrapper", nil)
	}

}

func TeamReadyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_ready.html")
		t.ExecuteTemplate(w, "htmx_wrapper", nil)
	}

}

func TeamDoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_done.html")
		t.ExecuteTemplate(w, "htmx_wrapper", nil)
	}

}

func TeamReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_review.html")
		t.ExecuteTemplate(w, "htmx_wrapper", nil)
	}

}
