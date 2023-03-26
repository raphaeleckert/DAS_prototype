/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package utils

import (
	"html/template"
	"net/http"
)

type errorMsg struct {
	ErrorText string
}

func ShowErrorMsg(message string, w http.ResponseWriter) {
	p := errorMsg{ErrorText: message}

	t, _ := template.ParseFiles(
		"../resources/templates/htmx_wrapper.html",
		"../resources/templates/error_msg.html",
	)
	t.ExecuteTemplate(w, "htmx_wrapper", p)
}