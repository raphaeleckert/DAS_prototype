/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package repo

import (
	"dasagilestudieren/prototype"
)

// initializes mocked database on application start

var (
	CourseRepo    *CourseRepository
	SubjectRepo   *SubjectRepository
	TopicRepo     *TopicRepository
	TermRepo      *TermRepository
	TeamRepo      *TeamRepository
	SolutionRepo  *SolutionRepository
	ProposalRepo  *ProposalRepository
	SupporterRepo *SupporterRepository
	ReviewRepo    *ReviewRepository
	RatingRepo    *RatingRepository
)

func init() {
	db := prototype.PopulatePrototypeDb()
	CourseRepo = NewCourseRepository(&db)
	SubjectRepo = NewSubjectRepository(&db)
	TopicRepo = NewTopicRepository(&db)
	TermRepo = NewTermRepository(&db)
	TeamRepo = NewTeamRepository(&db)
	SolutionRepo = NewSolutionRepository(&db)
	ProposalRepo = NewProposalRepository(&db)
	SupporterRepo = NewSupporterRepository(&db)
	ReviewRepo = NewReviewRepository(&db)
	RatingRepo = NewRatingRepository(&db)
}
