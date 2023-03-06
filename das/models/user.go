package models

type User struct {
	Username      string
	Authenticated bool
	IsTeacher     bool
}

func GetUser(name string) User {
	return User{
		Username:      name,
		Authenticated: true,
		IsTeacher:     false,
	}
}

func GetTeacher(name string) User {
	return User{
		Username:      name,
		Authenticated: true,
		IsTeacher:     true,
	}
}
