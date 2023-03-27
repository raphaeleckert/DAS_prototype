/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package handlers

import (
	"dasagilestudieren/models"
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

// This file contains all the handlers for the team resource

type TeamTabPage struct {
	Team      struct{ Name, ID string }
	TableData []models.Clickable
}

type TeamBasePage struct {
	Name, ID, Subject, Term string
	Member                  []string
	IsTeacher               bool
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	// POST: Creates a new Team
	if r.Method == http.MethodPost {
		courseId := r.URL.Query().Get("courseid")

		course, err := models.GetCourse(courseId)
		newTeam := models.Team{
			ID:       uuid.New().String(),
			Course:   course,
			Number:   99,
			Member:   []string{},
			ReadOnly: false,
			Note:     "",
			Remark:   "",
		}
		//TODO Add User
		fmt.Printf("created Team %s", newTeam.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("HX-Redirect", "/team?teamid=team1")

	}
	// PATCH: Adds or removes users from the team
	if r.Method == http.MethodPatch {
		teamid := r.URL.Query().Get("teamid")
		action := r.URL.Query().Get("action")
		team, err := models.GetTeam(teamid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r.ParseForm()
		user := r.FormValue("user")
		if action == "add" {
			//TODO Add User
			fmt.Printf("%s added to Team %d", user, team.Number)
		}
		if action == "remove" {
			//TODO Remove User
			fmt.Printf("%s removed from Team %d", user, team.Number)
		}
	}
	// GET: Returns the base page for a team
	if r.Method == http.MethodGet || r.Method == http.MethodPatch {
		user := r.Context().Value("user").(models.User)
		teamid := r.URL.Query().Get("teamid")
		team, err := models.GetTeam(teamid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		p := TeamBasePage{
			Name:      fmt.Sprintf("Team #%d", team.Number),
			ID:        team.ID,
			Subject:   team.Course.Subject.Name,
			Term:      team.Course.Term.Name,
			Member:    team.Member,
			IsTeacher: user.IsTeacher,
		}

		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/team/team.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "base", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// DELETE: Deletes a team
	} else if r.Method == http.MethodDelete {
		teamid := r.URL.Query().Get("teamid")
		//TODO Delete Team
		fmt.Printf("%s Deleted", teamid)

		w.Header().Set("HX-Redirect", "/landing")
		return
	}

}

func TeamOpenHandler(w http.ResponseWriter, r *http.Request) {
	// Returns open topics for a team
	if r.Method == http.MethodGet {
		teamid := r.URL.Query().Get("teamid")

		//TODO: Add Real Topics
		topics := []models.Clickable{
			func() models.Clickable {
				t, _ := models.GetTopicBasic("topic1")
				return t
			}(),
			func() models.Clickable {
				t, _ := models.GetTopicBasic("topic2")
				return t
			}(),
			func() models.Clickable {
				t, _ := models.GetTopicBasic("topic3")
				return t
			}(),
		}
		team, err := models.GetTeamBasic(teamid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p := TeamTabPage{TableData: topics, Team: team}

		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_open.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

}

func TeamWorkHandler(w http.ResponseWriter, r *http.Request) {
	// Returns in work topics for a team
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_work.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func TeamReadyHandler(w http.ResponseWriter, r *http.Request) {
	// returns ready topics for a team
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_ready.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func TeamDoneHandler(w http.ResponseWriter, r *http.Request) {
	// returns done topics for a team
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_done.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func TeamReviewHandler(w http.ResponseWriter, r *http.Request) {
	// returns reviews for a team
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/team_review.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

type AddUserForm struct {
	TeamId string
}

func AddUserFormHandler(w http.ResponseWriter, r *http.Request) {
	// GET: Returns the form to add a user
	if r.Method == http.MethodGet {
		teamid := r.URL.Query().Get("teamid")

		p := RemoveUserForm{
			TeamId: teamid,
		}

		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/add_user.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type RemoveUserForm struct {
	TeamId string
	Member []string
}

func RemoveUserFormHandler(w http.ResponseWriter, r *http.Request) {
	// GET: Returns the form do delete a user
	if r.Method == http.MethodGet {
		teamid := r.URL.Query().Get("teamid")
		team, err := models.GetTeam(teamid)

		p := RemoveUserForm{
			TeamId: teamid,
			Member: team.Member,
		}

		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/team/remove_user.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "htmx_wrapper", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
