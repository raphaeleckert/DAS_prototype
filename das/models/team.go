package models

import (
	"fmt"

	"dasagilestudieren/repo"
)

// Team
func GetTeam(id string) (Team, error) {
	teamRepo := repo.TeamRepo
	team, err := teamRepo.Read(id)
	if err != nil {
		return Team{}, fmt.Errorf("failed to get team with ID %s: %v", id, err)
	}
	course, err := GetCourse(team.Course)
	if err != nil {
		return Team{}, fmt.Errorf("failed to get course for team with ID %s: %v", id, err)
	}
	return Team{
		ID:       team.ID,
		Course:   course,
		Number:   team.Number,
		Member:   team.Member,
		ReadOnly: team.ReadOnly,
		Note:     team.Note,
		Remark:   team.Remark,
	}, nil
}

func GetTeamBasic(id string) (Clickable, error) {
	team, err := GetTeam(id)
	if err != nil {
		return Clickable{}, fmt.Errorf("failed to get team for ID %s: %v", id, err)
	}
	return Clickable{
		ID:   team.ID,
		Name: fmt.Sprintf("%s | Team %d", team.Course.Subject.Name, team.Number),
	}, nil
}
