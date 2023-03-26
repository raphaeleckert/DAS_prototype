package models

import (
	"fmt"

	"dasagilestudieren/repo"
)

// Solution
func GetSolution(id string) (Solution, error) {
	repo := repo.SolutionRepo
	solution, err := repo.Read(id)
	team, err := GetTeam(solution.Team)
	topic, err := GetTopic(solution.Topic)
	if err != nil {
		return Solution{}, fmt.Errorf("failed to get solution for ID %s: %v", id, err)
	}
	return Solution{
		ID:     solution.ID,
		Team:   team,
		Topic:  topic,
		State:  solution.State,
		Remark: solution.Remark,
	}, nil
}

func GetSolutionByTeamAndTopic(teamID string, topicID string) (Solution, error) {
	repo := repo.SolutionRepo
	solutionsForTeam, err := repo.ListByTeam(teamID)
	team, err := GetTeam(teamID)
	if err != nil {
		return Solution{}, fmt.Errorf("failed to get solution for team %s: %v", teamID, err)
	}

	for _, sol := range solutionsForTeam {
		if sol.Topic == topicID {
			topic, err := GetTopic(sol.Topic)
			if err != nil {
				return Solution{}, fmt.Errorf("failed to get solution for team %s: %v", teamID, err)
			}
			return Solution{
				ID:     sol.ID,
				Team:   team,
				Topic:  topic,
				State:  sol.State,
				Remark: sol.Remark,
			}, nil
		}
	}
	return Solution{}, fmt.Errorf("failed to get solution for team %s: %v", teamID, err)
}