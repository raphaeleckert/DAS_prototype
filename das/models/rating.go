package models

import (
	"fmt"

	"dasagilestudieren/repo"
)

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