package utils

import (
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