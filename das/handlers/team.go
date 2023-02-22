package handlers

import (
	"html/template"
	"net/http"
)

type TeamOverviewPage struct {
	TeamName string
}

func TeamOverviewHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p := TeamOverviewPage{TeamName: "Example Team Name"}

		t, _ := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/team/team.html")
		t.ExecuteTemplate(writer, "base", p)
	}

}

func TeamOpenHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_open.html")
		t.ExecuteTemplate(writer, "htmx_wrapper", nil)
	}

}

func TeamWorkHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/team/team_work.html")
		t.ExecuteTemplate(writer, "base", nil)
	}

}

func TeamReadyHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/team/team_ready.html")
		t.ExecuteTemplate(writer, "base", nil)
	}

}

func TeamDoneHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/team/team_done.html")
		t.ExecuteTemplate(writer, "base", nil)
	}

}

func TeamReviewHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/team/team_review.html")
		t.ExecuteTemplate(writer, "base", nil)
	}

}
