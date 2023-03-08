package models

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

type Team struct {
	ID       any
	Course   any
	Number   any
	Member   []any
	ReadOnly bool
	Note     any
	Remark   any
}

type Solution struct {
	ID     any
	Team   any
	Topic  any
	State  any
	Remark any
}

type Proposal struct {
	ID         any
	Solution   Solution
	ModifiedBy User
	Detail     any
}

func GetTeam(id string) Team {
	return Team{
		ID:       id,
		Course:   GetCourse("courseid"),
		Number:   1,
		Member:   []any{"student1", "student2"},
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

func GetProposal(id string) Proposal {
	return Proposal{
		ID:         id,
		Solution:   GetSolution("solutionid"),
		ModifiedBy: GetUser("userid"),
		Detail:     "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
	}
}

func GetTeamBasic(id string) Clickable {
	return Clickable{
		ID:   id,
		Name: "Bachelorthesis @ Wintersemester 2023, #1, reckert",
	}
}
