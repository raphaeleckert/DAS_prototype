package models

import (
	"fmt"
	"time"

	"dasagilestudieren/repo"
)

// Solution States
const (
	STATE_UNKNOWN  = "Unknown"
	STATE_INWORK   = "In Work"
	STATE_READY    = "Ready to Review"
	STATE_INREVIEW = "In Review"
	STATE_REJECTED = "Rejected"
	STATE_DONE     = "Done"
	STATE_FINAL    = "Final"
)

var RATINGS = [4]string{"Solved", "Minor Issues", "Major Issues", "Unsolved"}

type Team struct {
	ID       string
	Course   Course
	Number   int
	Member   []string
	ReadOnly bool
	Note     string
	Remark   string
}

type Solution struct {
	ID     string
	Team   Team
	Topic  Topic
	State  string
	Remark string
}

type Proposal struct {
	ID         string
	Solution   Solution
	ModifiedBy string
	Detail     string
}

type Supporter struct {
	ID       string
	Proposal Proposal
	Member   string
}

type Review struct {
	ID         string
	Course     Course
	ReviewDate time.Time
	MaxReviews int
	BeginDate  time.Time
	EndDate    time.Time
	Note       string
	Remark     string
}

type Rating struct {
	ID       string
	Proposal Proposal
	Review   Review
	Rating   int
	Remark   string
}

// Team
func GetTeam(id string) (Team, error) {
	teamRepo := repo.TeamRepo
	team, err := teamRepo.Read(id)

	course, err := GetCourse(team.Course)
	if err != nil {
		return Team{}, fmt.Errorf("failed to get team with ID %s: %v", id, err)
	}
	return Team{
		ID:       team.ID,
		Course:   course,
		Number:   team.Number,
		Member:   team.Member,
		ReadOnly: team.ReadOnly,
		Note:     team.Note,
		Remark:   team.Remark,
	}, nil
}

func GetTeamBasic(id string) (Clickable, error) {
	team, err := GetTeam(id)
	if err != nil {
		return Clickable{}, fmt.Errorf("failed to get team for ID %s: %v", id, err)
	}
	return Clickable{
		ID:   team.ID,
		Name: fmt.Sprintf("%s | Team %d", team.Course.Subject.Name, team.Number),
	}, nil
}

// Solution
func GetSolution(id string) (Solution, error) {
	repo := repo.SolutionRepo
	solution, err := repo.Read(id)
	team, err := GetTeam(solution.Team)
	topic, err := GetTopic(solution.Topic)
	if err != nil {
		return Solution{}, fmt.Errorf("failed to get solution for ID %s: %v", id, err)
	}
	return Solution{
		ID:     solution.ID,
		Team:   team,
		Topic:  topic,
		State:  solution.State,
		Remark: solution.Remark,
	}, nil
}

func GetSolutionByTeamAndTopic(teamID string, topicID string) (Solution, error) {
	repo := repo.SolutionRepo
	solutionsForTeam, err := repo.ListByTeam(teamID)
	team, err := GetTeam(teamID)
	if err != nil {
		return Solution{}, fmt.Errorf("failed to get solution for team %s: %v", teamID, err)
	}

	for _, sol := range solutionsForTeam {
		if sol.Topic == topicID {
			topic, err := GetTopic(sol.Topic)
			if err != nil {
				return Solution{}, fmt.Errorf("failed to get solution for team %s: %v", teamID, err)
			}
			return Solution{
				ID:     sol.ID,
				Team:   team,
				Topic:  topic,
				State:  sol.State,
				Remark: sol.Remark,
			}, nil
		}
	}
	return Solution{}, fmt.Errorf("failed to get solution for team %s: %v", teamID, err)
}

// Proposal
func GetProposal(id string) (Proposal, error) {
	repo := repo.ProposalRepo
	proposal, err := repo.Read(id)
	solution, err := GetSolution(proposal.Solution)
	if err != nil {
		return Proposal{}, fmt.Errorf("failed to get proposal for ID %s: %v", id, err)
	}
	return Proposal{
		ID:         proposal.ID,
		Solution:   solution,
		ModifiedBy: proposal.ModifiedBy,
		Detail:     proposal.Detail,
	}, nil
}

func GetProposalBySolution(solutionID string) (Proposal, error) {
	repo := repo.ProposalRepo
	proposal, err := repo.ListBySolution(solutionID)
	if err != nil || len(proposal) == 0 {
		return Proposal{}, fmt.Errorf("failed to get proposal for solution %s: %v", solutionID, err)
	}
	solution, err := GetSolution(proposal[0].Solution)
	if err != nil {
		return Proposal{}, fmt.Errorf("failed to get proposal for solution %s: %v", solutionID, err)
	}
	return Proposal{
		ID:         proposal[0].ID,
		Solution:   solution,
		ModifiedBy: proposal[0].ModifiedBy,
		Detail:     proposal[0].Detail,
	}, nil
}

// Review
func GetReview(id string) (Review, error) {
	repo := repo.ReviewRepo
	review, err := repo.Read(id)
	course, err := GetCourse(review.Course)
	if err != nil {
		return Review{}, fmt.Errorf("failed to get review for ID %s: %v", id, err)
	}
	return Review{
		ID:         review.ID,
		Course:     course,
		ReviewDate: review.ReviewDate,
		MaxReviews: review.MaxReviews,
		BeginDate:  review.BeginDate,
		EndDate:    review.EndDate,
		Note:       review.Note,
		Remark:     review.Remark,
	}, nil
}

func GetReviewBasic(id string) (Clickable, error) {
	repo := repo.ReviewRepo
	review, err := repo.Read(id)
	course, err := GetCourse(review.Course)
	if err != nil {
		return Clickable{}, fmt.Errorf("failed to get review for ID %s: %v", id, err)
	}
	return Clickable{
		ID:   review.ID,
		Name: fmt.Sprintf("%s | %s", course.Subject.Name, review.ReviewDate),
	}, nil
}

// Rating
func GetRatingByProposal(proposalID string) (Rating, error) {
	repo := repo.RatingRepo
	fmt.Printf(" proposalid %v ", proposalID)

	rating, err := repo.ListByProposal(proposalID)
	if len(rating) == 0 {
		return Rating{}, fmt.Errorf("failed to get rating for proposal %s: %v", proposalID, err)
	}
	proposal, err := GetProposal(rating[0].Proposal)

	review, err := GetReview(rating[0].Review)
	if err != nil {
		return Rating{}, fmt.Errorf("failed to get rating for proposal %s: %v", proposalID, err)
	}
	return Rating{
		ID:       rating[0].ID,
		Proposal: proposal,
		Review:   review,
		Rating:   rating[0].Rating,
		Remark:   rating[0].Remark,
	}, nil
}
