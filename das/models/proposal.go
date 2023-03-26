/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package models

import (
	"fmt"

	"dasagilestudieren/repo"
)

// Proposal
func GetProposal(id string) (Proposal, error) {
	repo := repo.ProposalRepo
	proposal, err := repo.Read(id)
	solution, err := GetSolution(proposal.Solution)
	if err != nil {
		return Proposal{}, fmt.Errorf("failed to get proposal for ID %s: %v", id, err)
	}
	return Proposal{
		ID:         proposal.ID,
		Solution:   solution,
		ModifiedBy: proposal.ModifiedBy,
		Detail:     proposal.Detail,
	}, nil
}

func GetProposalBySolution(solutionID string) (Proposal, error) {
	repo := repo.ProposalRepo
	proposal, err := repo.ListBySolution(solutionID)
	if err != nil || len(proposal) == 0 {
		return Proposal{}, fmt.Errorf("failed to get proposal for solution %s: %v", solutionID, err)
	}
	solution, err := GetSolution(proposal[0].Solution)
	if err != nil {
		return Proposal{}, fmt.Errorf("failed to get proposal for solution %s: %v", solutionID, err)
	}
	return Proposal{
		ID:         proposal[0].ID,
		Solution:   solution,
		ModifiedBy: proposal[0].ModifiedBy,
		Detail:     proposal[0].Detail,
	}, nil
}