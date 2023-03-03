package utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

type User struct {
	Username      string
	Authenticated bool
	IsTeacher     bool
}

// getUser returns a user from session s
// on error returns an empty user
func GetUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {

		return User{Authenticated: false}
	}
	return user
}

func LoginRequired(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "das-session")
		user := GetUser(session)
		if err != nil {
			fmt.Println("Bad Session")
			http.SetCookie(w, &http.Cookie{Name: "das-session", MaxAge: -1, Path: "/"})
			http.Redirect(w, r, "/start", http.StatusSeeOther)
			return
		}
		if user.Authenticated == false {
			http.Redirect(w, r, "/start", http.StatusSeeOther)
		}
		r = r.WithContext(context.WithValue(r.Context(), "session", session))
		h(w, r)
	}
}
