package models

import (
	"time"
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

func GetTeam(id string) Team {
	return Team{
		ID:       id,
		Course:   GetCourse("courseid"),
		Number:   1,
		Member:   []string{"student1", "student2"},
		ReadOnly: false,
		Note:     "Team Note",
		Remark:   "Team Remark",
	}
}

func GetSolution(id string) Solution {
	return Solution{
		ID:     id,
		Team:   GetTeam("teamid"),
		Topic:  GetTopic("topicid"),
		State:  STATE_INWORK,
		Remark: "Solution Remark",
	}
}

func GetSolutionByTeamAndTopic(team string, topic string) Solution {
	return Solution{
		ID:     "solutionid",
		Team:   GetTeam(team),
		Topic:  GetTopic(topic),
		State:  STATE_INWORK,
		Remark: "Solution Remark",
	}
}

func GetProposal(id string) Proposal {
	return Proposal{
		ID:         id,
		Solution:   GetSolution("solutionid"),
		ModifiedBy: "student1",
		Detail:     "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
	}
}

func GetProposalBySolution(solution string) Proposal {
	return Proposal{
		ID:         "solutionid",
		Solution:   GetSolution(solution),
		ModifiedBy: "student1",
		Detail:     "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
	}
}

func GetTeamBasic(id string) Clickable {
	return Clickable{
		ID:   id,
		Name: "Bachelorthesis @ Wintersemester 2023, #1, reckert",
	}
}

func GetReview(id string) Review {
	return Review{
		ID:         id,
		Course:     GetCourse("courseid"),
		ReviewDate: time.Date(2024, time.January, 01, 00, 00, 00, 0, time.UTC),
		MaxReviews: 5,
		BeginDate:  time.Date(2024, time.January, 01, 00, 00, 00, 0, time.UTC),
		EndDate:    time.Date(2022, time.January, 01, 00, 00, 00, 0, time.UTC),
		Note:       "Review Note",
		Remark:     "Review Remark",
	}
}

func GetReviewBasic(id string) Clickable {
	review := GetReview(id)
	date := review.ReviewDate.Format("2000-01-01 12:00")
	return Clickable{
		ID:   id,
		Name: date,
	}
}

func GetRatingByProposal(proposal string) Rating {
	return Rating{
		ID:       "ratingid",
		Proposal: GetProposal(proposal),
		Review:   GetReview("reviewid"),
		Rating:   3,
		Remark:   "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut",
	}
}
