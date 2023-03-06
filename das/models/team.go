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

func GetTeam(id string) Team {
	return Team{
		ID:       id,
		Course:   "courseid",
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
		Team:   "teamid",
		Topic:  "topicidea",
		State:  STATE_INWORK,
		Remark: "Solution Remark",
	}
}
