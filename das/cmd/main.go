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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func router() {

	server := http.Server{
		Addr:    ":8080",
		Handler: nil, // set to nil to use the default ServeMux
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/start", handlers.Start)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/tryout", handlers.Tryout)

	//topic
	http.HandleFunc("/topic", utils.LoginRequired(handlers.TopicHandler))

	//team
	http.HandleFunc("/team", utils.LoginRequired(handlers.TeamOverviewHandler))
	http.HandleFunc("/team/open", handlers.TeamOpenHandler)
	http.HandleFunc("/team/work", handlers.TeamWorkHandler)
	http.HandleFunc("/team/ready", handlers.TeamReadyHandler)
	http.HandleFunc("/team/done", handlers.TeamDoneHandler)
	http.HandleFunc("/team/review", handlers.TeamReviewHandler)

	//resources
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("../resources"))))

	log.Fatal(server.ListenAndServe())
}

func main() {

	router()
}
