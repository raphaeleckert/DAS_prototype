package handlers

import (
	"fmt"
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
	if request.Method == http.MethodPost {
		session, err := utils.Store.Get(request, "cookie-name")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		user := utils.GetUser(session)
		if user.Authenticated {
			writer.Header().Set("HX-Redirect", "/team")
		} else {
			utils.ShowErrorMsg("Wrong username or password!", writer)
		}

	} else if request.Method == http.MethodGet {
		fmt.Println("post")

		request.ParseForm()
		username := request.FormValue("username")
		password := request.FormValue("password")

		// autentication based only on input
		// authenticated := username != "" && password != "" && password != "wrong"
		authenticated := true
		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(authenticated)

		isTeacher := username == "teacher"

		user := &utils.User{
			Username:      username,
			Authenticated: authenticated,
			IsTeacher:     isTeacher,
		}

		session, err := utils.Store.Get(request, "cookie-name")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["user"] = user

		err = session.Save(request, writer)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
