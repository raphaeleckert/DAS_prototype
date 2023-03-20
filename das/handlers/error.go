package handlers

import (
	"dasagilestudieren/utils"

	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {

	utils.ShowErrorMsg("test", w)
}
