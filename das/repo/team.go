package repo

import (
	"dasagilestudieren/models"
)

type TeamRepository interface {
	Create(team *models.Team) error
	Read(id string) (*models.Team, error)
	Update(team *models.Team) error
	Delete(id string) error
	ListByCourse(courseID string) ([]*models.Team, error)
}

type SolutionRepository interface {
	Create(solution *models.Solution) error
	Read(id string) (*models.Solution, error)
	Update(solution *models.Solution) error
	Delete(id string) error
	ListByTopic(topicID string) ([]*models.Solution, error)
	ListByTeam(teamID string) ([]*models.Solution, error)
}

type ProposalRepository interface {
	Create(proposal *models.Proposal) error
	Read(id string) (*models.Proposal, error)
	Update(proposal *models.Proposal) error
	Delete(id string) error
}

type SupporterRepository interface {
	Create(supporter *models.Supporter) error
	Read(id string) (*models.Supporter, error)
	Update(supporter *models.Supporter) error
	Delete(id string) error
	ListByProposal(proposalID string) ([]*models.Supporter, error)
}

type ReviewRepository interface {
	Create(review *models.Review) error
	Read(id string) (*models.Review, error)
	Update(review *models.Review) error
	Delete(id string) error
	ListByCourse(courseID string) ([]*models.Review, error)
}

type RatingRepository interface {
	Create(rating *models.Rating) error
	Read(id string) (*models.Rating, error)
	Update(rating *models.Rating) error
	Delete(id string) error
	ListByProposal(proposalID string) ([]*models.Rating, error)
	ListByReview(reviewID string) ([]*models.Rating, error)
}
