package handlers

import (
	"dasagilestudieren/utils"

	"net/http"
	"html/template"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	
utils.ShowErrorMsg()
}