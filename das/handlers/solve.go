package handlers

import (
	"html/template"
	"net/http"

	"dasagilestudieren/models"
)

type SolvePage struct {
	CurrentUser   string
	TopicData     models.Topic
	Proposal      *models.Proposal
	Rating        *models.Rating
	RatingOptions [4]string
	Status        string
	CanAccept     bool
	CanEdit       bool
	TeamId        string
	TopicId       string
}

func SolveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//Create Proposal
	}
	if r.Method == http.MethodPatch {
		//Update Proposal
	}
	if r.Method == http.MethodGet || r.Method == http.MethodPost || r.Method == http.MethodPatch {
		user := r.Context().Value("user").(models.User)
		teamid := r.URL.Query().Get("teamid")
		topicid := r.URL.Query().Get("topicid")
		topic := models.GetTopic(topicid)
		solution := models.GetSolutionByTeamAndTopic(teamid, topicid)
		proposal := models.GetProposalBySolution(solution.ID)
		rating := models.GetRatingByProposal(proposal.ID)
		canAccept := true
		canEdit := true

		p := SolvePage{
			CurrentUser:   user.Username,
			TopicData:     topic,
			Proposal:      &proposal,
			Rating:        &rating,
			RatingOptions: models.RATINGS,
			Status:        solution.State,
			CanAccept:     canAccept,
			CanEdit:       canEdit,
			TeamId:        teamid,
			TopicId:       topicid,
		}

		t, err := template.ParseFiles(
			"../resources/templates/base.html",
			"../resources/templates/solve/solve.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.ExecuteTemplate(w, "base", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type CreateProposalForm struct {
	TeamId, TopicId string
}

func CreateProposalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		teamid := r.URL.Query().Get("teamid")
		topicid := r.URL.Query().Get("topicid")

		p := EditProposalForm{
			TeamId:  teamid,
			TopicId: topicid,
		}
		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/solve/create_proposal.html")
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

type EditProposalForm struct {
	TeamId, TopicId string
	CurrentProposal models.Proposal
}

func EditProposalHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		teamid := r.URL.Query().Get("teamid")
		topicid := r.URL.Query().Get("topicid")
		solution := models.GetSolutionByTeamAndTopic(teamid, topicid)
		proposal := models.GetProposalBySolution(solution.ID)

		p := EditProposalForm{
			TeamId:          teamid,
			TopicId:         topicid,
			CurrentProposal: proposal,
		}

		t, err := template.ParseFiles(
			"../resources/templates/htmx_wrapper.html",
			"../resources/templates/solve/edit_proposal.html")
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
