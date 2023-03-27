/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"

	"dasagilestudieren/models"
)

// Helper functions for passing the user to the handlers and redirect if required

var Store *sessions.CookieStore

// GetUser returns a user from session s
// on error returns an empty user
func GetUser(s *sessions.Session) models.User {
	val := s.Values["user"]
	var user = models.User{}
	user, ok := val.(models.User)
	if !ok {

		return models.User{Authenticated: false}
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
			http.Redirect(w, r, "/start", http.StatusForbidden)
		}
		r = r.WithContext(context.WithValue(r.Context(), "user", user))
		h(w, r)
	}
}

func TeacherRequired(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "das-session")
		user := GetUser(session)
		if err != nil {
			fmt.Println("Bad Session")
			http.SetCookie(w, &http.Cookie{Name: "das-session", MaxAge: -1, Path: "/"})
			http.Redirect(w, r, "/start", http.StatusSeeOther)
			return
		}
		if user.IsTeacher == false {
			http.Redirect(w, r, "/start", http.StatusForbidden)
		}
		r = r.WithContext(context.WithValue(r.Context(), "user", user))
		h(w, r)
	}
}
