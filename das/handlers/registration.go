package handlers

import (
	"html/template"
	"net/http"

	"dasagilestudieren/utils"
)

func Start(writer http.ResponseWriter, request *http.Request) {
	//GET
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles(
			"../resources/templates/registration/login.html")
		t.ExecuteTemplate(writer, "login.html", nil)
	}
}

func Login(writer http.ResponseWriter, request *http.Request) {

	session, err := utils.Store.Get(request, "das-session")

	request.ParseForm()
	username := request.FormValue("username")
	password := request.FormValue("password")
	// autentication based only on input
	authenticated := username != "" && password != "" && password != "wrong"
	isTeacher := username == "teacher"

	user := utils.User{
		Username:      username,
		Authenticated: authenticated,
		IsTeacher:     isTeacher,
	}
	session.Values["user"] = user
	err = session.Save(request, writer)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	user = utils.GetUser(session)

	if user.Authenticated {
		writer.Header().Set("HX-Redirect", "/team")
	} else {
		utils.ShowErrorMsg("Wrong username or password!", writer)
	}
}
