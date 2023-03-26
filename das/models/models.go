/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package models

import (
	"time"
)

// Topic Importances
const (
	IMP_ESSENTIAL = "Prerequisite for other topics"
	IMP_MUST      = "Basic knowledge"
	IMP_SHOULD    = "Average knowledge"
	IMP_COULD     = "Special knowledge"
	IMP_VOLUNTARY = "Voluntary topic"
)

// Topic Required Supporters
const (
	SUP_NONE = "No Supporters"
	SUP_ONE  = "One Supporter"
	SUP_TWO  = "Two Supporters"
	SUP_HALF = "Half of the Team"
	SUP_BUT  = "All but One"
	SUP_ALL  = "All Team Members"
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

type Subject struct {
	ID        string
	ShortName string
	Name      string
	Owner     string
	Note      string
	Remark    string
}

type Course struct {
	ID        string
	Subject   Subject
	Term      Term
	Owner     string
	BeginDate time.Time
	FinalDate time.Time
	CloseDate time.Time
	Remark    string
	Note      string
	Topics    []string
}

type Topic struct {
	ID                 string
	Title              string
	Subject            Subject
	Number             string
	Detail             string
	Reference          string
	SolutionIdea       string
	Remark             string
	Tags               []string
	Importance         string
	RequiredSupporters string
}

type Term struct {
	ID        string
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Note      string
	Remark    string
}

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
