/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package repo

import (
	"dasagilestudieren/prototype"
	"fmt"
)

func NewTeamRepository(db *prototype.PrototypeDb) *TeamRepository {
	return &TeamRepository{db: db}
}

func NewSolutionRepository(db *prototype.PrototypeDb) *SolutionRepository {
	return &SolutionRepository{db: db}
}

func NewProposalRepository(db *prototype.PrototypeDb) *ProposalRepository {
	return &ProposalRepository{db: db}
}

func NewSupporterRepository(db *prototype.PrototypeDb) *SupporterRepository {
	return &SupporterRepository{db: db}
}

func NewReviewRepository(db *prototype.PrototypeDb) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func NewRatingRepository(db *prototype.PrototypeDb) *RatingRepository {
	return &RatingRepository{db: db}
}

type TeamInterface interface {
	Create(team *prototype.Team) error
	Read(id string) (*prototype.Team, error)
	Update(team *prototype.Team) error
	Delete(id string) error
	ListByCourse(courseID string) ([]*prototype.Team, error)
}

type SolutionInterface interface {
	Create(solution *prototype.Solution) error
	Read(id string) (*prototype.Solution, error)
	Update(solution *prototype.Solution) error
	Delete(id string) error
	ListByTopic(topicID string) ([]*prototype.Solution, error)
	ListByTeam(teamID string) ([]*prototype.Solution, error)
}

type ProposalInterface interface {
	Create(proposal *prototype.Proposal) error
	Read(id string) (*prototype.Proposal, error)
	Update(proposal *prototype.Proposal) error
	Delete(id string) error
	ListBySolution(solution string) ([]*prototype.Proposal, error)
}

type SupporterInterface interface {
	Create(supporter *prototype.Supporter) error
	Read(id string) (*prototype.Supporter, error)
	Update(supporter *prototype.Supporter) error
	Delete(id string) error
	ListByProposal(proposalID string) ([]*prototype.Supporter, error)
}

type ReviewInterface interface {
	Create(review *prototype.Review) error
	Read(id string) (*prototype.Review, error)
	Update(review *prototype.Review) error
	Delete(id string) error
	ListByCourse(courseID string) ([]*prototype.Review, error)
}

type RatingInterface interface {
	Create(rating *prototype.Rating) error
	Read(id string) (*prototype.Rating, error)
	Update(rating *prototype.Rating) error
	Delete(id string) error
	ListByProposal(proposalID string) ([]*prototype.Rating, error)
	ListByReview(reviewID string) ([]*prototype.Rating, error)
}

type TeamRepository struct {
	db *prototype.PrototypeDb
}

type SolutionRepository struct {
	db *prototype.PrototypeDb
}

type ProposalRepository struct {
	db *prototype.PrototypeDb
}

type SupporterRepository struct {
	db *prototype.PrototypeDb
}

type RatingRepository struct {
	db *prototype.PrototypeDb
}

type ReviewRepository struct {
	db *prototype.PrototypeDb
}

// Team
func (repo *TeamRepository) Create(team *prototype.Team) error {
	repo.db.Teams[team.ID] = team
	return nil
}

func (repo *TeamRepository) Read(id string) (*prototype.Team, error) {
	team, ok := repo.db.Teams[id]
	if !ok {
		return nil, fmt.Errorf("team not found")
	}
	return team, nil
}

func (repo *TeamRepository) Update(team *prototype.Team) error {
	_, ok := repo.db.Teams[team.ID]
	if !ok {
		return fmt.Errorf("team not found")
	}
	repo.db.Teams[team.ID] = team
	return nil
}

func (repo *TeamRepository) Delete(id string) error {
	_, ok := repo.db.Teams[id]
	if !ok {
		return fmt.Errorf("team not found")
	}
	delete(repo.db.Teams, id)
	return nil
}

func (repo *TeamRepository) ListByCourse(courseID string) ([]*prototype.Team, error) {
	var teams []*prototype.Team
	for _, team := range repo.db.Teams {
		if team.Course == courseID {
			teams = append(teams, team)
		}
	}
	return teams, nil
}

// Solution
func (repo *SolutionRepository) Create(solution *prototype.Solution) error {
	repo.db.Solutions[solution.ID] = solution
	return nil
}

func (repo *SolutionRepository) Read(id string) (*prototype.Solution, error) {
	solution, ok := repo.db.Solutions[id]
	if !ok {
		return nil, fmt.Errorf("solution with ID %s not found", id)
	}
	return solution, nil
}

func (repo *SolutionRepository) Update(solution *prototype.Solution) error {
	_, ok := repo.db.Solutions[solution.ID]
	if !ok {
		return fmt.Errorf("solution with ID %s not found", solution.ID)
	}
	repo.db.Solutions[solution.ID] = solution
	return nil
}

func (repo *SolutionRepository) Delete(id string) error {
	_, ok := repo.db.Solutions[id]
	if !ok {
		return fmt.Errorf("solution with ID %s not found", id)
	}
	delete(repo.db.Solutions, id)
	return nil
}

func (repo *SolutionRepository) ListByTopic(topicID string) ([]*prototype.Solution, error) {
	var solutions []*prototype.Solution
	for _, solution := range repo.db.Solutions {
		if solution.Topic == topicID {
			solutions = append(solutions, solution)
		}
	}
	return solutions, nil
}

func (repo *SolutionRepository) ListByTeam(teamID string) ([]*prototype.Solution, error) {
	var solutions []*prototype.Solution
	for _, solution := range repo.db.Solutions {
		if solution.Team == teamID {
			solutions = append(solutions, solution)
		}
	}
	return solutions, nil
}

// Proposal
func (repo *ProposalRepository) Create(proposal *prototype.Proposal) error {
	repo.db.Proposals[proposal.ID] = proposal
	return nil
}

func (repo *ProposalRepository) Read(id string) (*prototype.Proposal, error) {
	proposal, ok := repo.db.Proposals[id]
	if !ok {
		return nil, fmt.Errorf("proposal with ID %s not found", id)
	}
	return proposal, nil
}

func (repo *ProposalRepository) Update(proposal *prototype.Proposal) error {
	_, ok := repo.db.Proposals[proposal.ID]
	if !ok {
		return fmt.Errorf("proposal with ID %s not found", proposal.ID)
	}
	repo.db.Proposals[proposal.ID] = proposal
	return nil
}

func (repo *ProposalRepository) Delete(id string) error {
	_, ok := repo.db.Proposals[id]
	if !ok {
		return fmt.Errorf("proposal with ID %s not found", id)
	}
	delete(repo.db.Proposals, id)
	return nil
}

func (repo *ProposalRepository) ListBySolution(solution string) ([]*prototype.Proposal, error) {
	var result []*prototype.Proposal
	for _, proposal := range repo.db.Proposals {
		if proposal.Solution == solution {
			result = append(result, proposal)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no proposals with solution %s found", solution)
	}
	return result, nil
}

// Supporter
func (repo *SupporterRepository) Create(supporter *prototype.Supporter) error {
	repo.db.Supporters[supporter.ID] = supporter
	return nil
}

func (repo *SupporterRepository) Read(id string) (*prototype.Supporter, error) {
	supporter, ok := repo.db.Supporters[id]
	if !ok {
		return nil, fmt.Errorf("supporter with ID %s not found", id)
	}
	return supporter, nil
}

func (repo *SupporterRepository) Update(supporter *prototype.Supporter) error {
	_, ok := repo.db.Supporters[supporter.ID]
	if !ok {
		return fmt.Errorf("supporter with ID %s not found", supporter.ID)
	}
	repo.db.Supporters[supporter.ID] = supporter
	return nil
}

func (repo *SupporterRepository) Delete(id string) error {
	_, ok := repo.db.Supporters[id]
	if !ok {
		return fmt.Errorf("supporter with ID %s not found", id)
	}
	delete(repo.db.Supporters, id)
	return nil
}

func (repo *SupporterRepository) ListByProposal(proposalID string) ([]*prototype.Supporter, error) {
	supporters := []*prototype.Supporter{}
	for _, supporter := range repo.db.Supporters {
		if supporter.Proposal == proposalID {
			supporters = append(supporters, supporter)
		}
	}
	return supporters, nil
}

// Review
func (repo *ReviewRepository) Create(review *prototype.Review) error {
	repo.db.Reviews[review.ID] = review
	return nil
}

func (repo *ReviewRepository) Read(id string) (*prototype.Review, error) {
	review, ok := repo.db.Reviews[id]
	if !ok {
		return nil, fmt.Errorf("review with ID %s not found", id)
	}
	return review, nil
}

func (repo *ReviewRepository) Update(review *prototype.Review) error {
	_, ok := repo.db.Reviews[review.ID]
	if !ok {
		return fmt.Errorf("review with ID %s not found", review.ID)
	}
	repo.db.Reviews[review.ID] = review
	return nil
}

func (repo *ReviewRepository) Delete(id string) error {
	_, ok := repo.db.Reviews[id]
	if !ok {
		return fmt.Errorf("review with ID %s not found", id)
	}
	delete(repo.db.Reviews, id)
	return nil
}

func (repo *ReviewRepository) ListByCourse(courseID string) ([]*prototype.Review, error) {
	reviews := []*prototype.Review{}
	for _, review := range repo.db.Reviews {
		if review.Course == courseID {
			reviews = append(reviews, review)
		}
	}
	return reviews, nil
}

// Rating
func (repo *RatingRepository) Create(rating *prototype.Rating) error {
	repo.db.Ratings[rating.ID] = rating
	return nil
}

func (repo *RatingRepository) Read(id string) (*prototype.Rating, error) {
	rating, ok := repo.db.Ratings[id]
	if !ok {
		return nil, fmt.Errorf("rating with ID %s not found", id)
	}
	return rating, nil
}

func (repo *RatingRepository) Update(rating *prototype.Rating) error {
	_, ok := repo.db.Ratings[rating.ID]
	if !ok {
		return fmt.Errorf("rating with ID %s not found", rating.ID)
	}
	repo.db.Ratings[rating.ID] = rating
	return nil
}

func (repo *RatingRepository) Delete(id string) error {
	_, ok := repo.db.Ratings[id]
	if !ok {
		return fmt.Errorf("rating with ID %s not found", id)
	}
	delete(repo.db.Ratings, id)
	return nil
}

func (repo *RatingRepository) ListByProposal(proposalID string) ([]*prototype.Rating, error) {
	ratings := []*prototype.Rating{}
	for _, rating := range repo.db.Ratings {
		if rating.Proposal == proposalID {
			ratings = append(ratings, rating)
		}
	}
	return ratings, nil
}

func (repo *RatingRepository) ListByReview(reviewID string) ([]*prototype.Rating, error) {
	ratings := []*prototype.Rating{}
	for _, rating := range repo.db.Ratings {
		if rating.Review == reviewID {
			ratings = append(ratings, rating)
		}
	}
	return ratings, nil
}
