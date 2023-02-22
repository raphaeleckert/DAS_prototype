package handlers

import (
	"html/template"
	"net/http"
)

type InTableLink struct {
	Text      string
	Reference string
}

func TeamOverviewHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p := struct{ TeamName string }{TeamName: "Example Team Name"}

		t, _ := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/team/team.html")
		t.ExecuteTemplate(writer, "base", p)
	}

}

type TeamOpenTableRow struct {
	Text      string
	Reference string
}

func TeamOpenHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		topics := []TeamOpenTableRow{
			{Text: "Topic1", Reference: "Topic1Ref"},
			{Text: "Topic2", Reference: "Topic2Ref"},
			{Text: "Topic3", Reference: "Topic3Ref"},
		}
		p := struct{ TableData []TeamOpenTableRow }{TableData: topics}

		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_open.html")
		t.ExecuteTemplate(writer, "htmx_wrapper", p)
	}

}

func TeamWorkHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_work.html")
		t.ExecuteTemplate(writer, "htmx_wrapper", nil)
	}

}

func TeamReadyHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_ready.html")
		t.ExecuteTemplate(writer, "htmx_wrapper", nil)
	}

}

func TeamDoneHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_done.html")
		t.ExecuteTemplate(writer, "htmx_wrapper", nil)
	}

}

func TeamReviewHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_review.html")
		t.ExecuteTemplate(writer, "htmx_wrapper", nil)
	}

}
