package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"dasagilestudieren/handlers"
	"dasagilestudieren/models"
	"dasagilestudieren/utils"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	utils.Store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	utils.Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(models.User{})
	fmt.Println("Start")

}

func router() {

	server := http.Server{
		Addr:    ":8080",
		Handler: nil, // set to nil to use the default ServeMux
	}

	http.HandleFunc("/", handlers.Start)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/tryout", handlers.Tryout)
	http.HandleFunc("/landing", utils.LoginRequired(handlers.LandingHandler))

	//topic
	http.HandleFunc("/topic", utils.TeacherRequired(handlers.TopicHandler))
	http.HandleFunc("/topic/info", utils.TeacherRequired(handlers.TopicInfoHandler))
	http.HandleFunc("/topic/edit", utils.TeacherRequired(handlers.TopicEditHandler))
	http.HandleFunc("/topic/list", utils.TeacherRequired(handlers.TopicListHandler))

	//team
	http.HandleFunc("/team", utils.LoginRequired(handlers.TeamHandler))
	http.HandleFunc("/team/open", utils.LoginRequired(handlers.TeamOpenHandler))
	http.HandleFunc("/team/work", utils.LoginRequired(handlers.TeamWorkHandler))
	http.HandleFunc("/team/ready", utils.LoginRequired(handlers.TeamReadyHandler))
	http.HandleFunc("/team/done", utils.LoginRequired(handlers.TeamDoneHandler))
	http.HandleFunc("/team/review", utils.LoginRequired(handlers.TeamReviewHandler))
	http.HandleFunc("/team/remove", utils.TeacherRequired(handlers.RemoveUserFormHandler))
	http.HandleFunc("/team/add", utils.TeacherRequired(handlers.AddUserFormHandler))

	//solve
	http.HandleFunc("/solve", utils.LoginRequired(handlers.SolveHandler))
	http.HandleFunc("/solve/edit", utils.LoginRequired(handlers.EditProposalHandler))
	http.HandleFunc("/solve/create", utils.LoginRequired(handlers.CreateProposalHandler))

	//subject
	http.HandleFunc("/subject", utils.TeacherRequired(handlers.SubjectHandler))

	//course
	http.HandleFunc("/course", utils.TeacherRequired(handlers.CourseHandler))
	http.HandleFunc("/course/create", utils.TeacherRequired(handlers.CreateCourseHandler))
	http.HandleFunc("/course/validate", utils.TeacherRequired(handlers.ValidateCourseDatesHandler))

	//resources
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("../resources"))))

	log.Fatal(server.ListenAndServe())
}

func main() {

	router()
}
