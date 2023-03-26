/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

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
