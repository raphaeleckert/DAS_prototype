package handlers

import (
	"html/template"
	"net/http"

	"dasagilestudieren/models"
	"dasagilestudieren/utils"
)

func Start(w http.ResponseWriter, r *http.Request) {
	//GET
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles(
			"../resources/templates/registration/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		session, err := utils.Store.Get(r, "das-session")

		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		// TODO: implement real user checking
		authenticated := username != "" && password != "" && password != "wrong"
		isTeacher := username == "teacher"

		user := models.User{
			Username:      username,
			Authenticated: authenticated,
			IsTeacher:     isTeacher,
		}
		session.Values["user"] = user
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user = utils.GetUser(session)

		if user.Authenticated {
			w.Header().Set("HX-Redirect", "/team")
		} else {
			utils.ShowErrorMsg("Wrong username or password!", w)
		}
	}
}
