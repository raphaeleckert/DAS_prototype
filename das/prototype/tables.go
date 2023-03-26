/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package prototype

import (
	"time"
)

// mocks db tables, not meant to be present in a production version.
// Attention: NEVER use these in the DAS-Project. Always use the versions in the models folder

type Team struct {
	ID       string
	Course   string
	Number   int
	Member   []string
	ReadOnly bool
	Note     string
	Remark   string
}

type Solution struct {
	ID     string
	Team   string
	Topic  string
	State  string
	Remark string
}

type Proposal struct {
	ID         string
	Solution   string
	ModifiedBy string
	Detail     string
}

type Supporter struct {
	ID       string
	Proposal string
	Member   string
}

type Review struct {
	ID         string
	Course     string
	ReviewDate time.Time
	MaxReviews int
	BeginDate  time.Time
	EndDate    time.Time
	Note       string
	Remark     string
}

type Rating struct {
	ID       string
	Proposal string
	Review   string
	Rating   int
	Remark   string
}

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
	Subject   string
	Term      string
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
	Subject            string
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
