package utils

import (
	"html/template"
	"net/http"
)

type errorMsg struct {
	ErrorText string
}

func ShowErrorMsg(message string, writer http.ResponseWriter) {
	p := errorMsg{ErrorText: message}

	t, _ := template.ParseFiles(
		"../resources/templates/htmx_wrapper.html",
		"../resources/templates/error_msg.html",
	)
	t.ExecuteTemplate(writer, "htmx_wrapper", p)
}