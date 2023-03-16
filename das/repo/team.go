package repo

import (
	"dasagilestudieren/prototype"
)

type TeamRepository interface {
	Create(team *prototype.Team) error
	Read(id string) (*prototype.Team, error)
	Update(team *prototype.Team) error
	Delete(id string) error
	ListByCourse(courseID string) ([]*prototype.Team, error)
}

type SolutionRepository interface {
	Create(solution *prototype.Solution) error
	Read(id string) (*prototype.Solution, error)
	Update(solution *prototype.Solution) error
	Delete(id string) error
	ListByTopic(topicID string) ([]*prototype.Solution, error)
	ListByTeam(teamID string) ([]*prototype.Solution, error)
}

type ProposalRepository interface {
	Create(proposal *prototype.Proposal) error
	Read(id string) (*prototype.Proposal, error)
	Update(proposal *prototype.Proposal) error
	Delete(id string) error
}

type SupporterRepository interface {
	Create(supporter *prototype.Supporter) error
	Read(id string) (*prototype.Supporter, error)
	Update(supporter *prototype.Supporter) error
	Delete(id string) error
	ListByProposal(proposalID string) ([]*prototype.Supporter, error)
}

type ReviewRepository interface {
	Create(review *prototype.Review) error
	Read(id string) (*prototype.Review, error)
	Update(review *prototype.Review) error
	Delete(id string) error
	ListByCourse(courseID string) ([]*prototype.Review, error)
}

type RatingRepository interface {
	Create(rating *prototype.Rating) error
	Read(id string) (*prototype.Rating, error)
	Update(rating *prototype.Rating) error
	Delete(id string) error
	ListByProposal(proposalID string) ([]*prototype.Rating, error)
	ListByReview(reviewID string) ([]*prototype.Rating, error)
}
