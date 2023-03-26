package models

import (
	"fmt"

	"dasagilestudieren/repo"
)

// Term
func GetTerm(id string) (Term, error) {
	repo := repo.TermRepo
	data, err := repo.Read(id)
	if err != nil {
		return Term{}, fmt.Errorf("failed to get term with ID %s: %v", id, err)
	}
	term := Term{
		ID:        data.ID,
		Name:      data.Name,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		Note:      data.Note,
		Remark:    data.Remark,
	}
	return term, nil
}

func GetTermBasic(id string) (Clickable, error) {
	repo := repo.TermRepo
	data, err := repo.Read(id)
	if err != nil {
		return Clickable{}, fmt.Errorf("failed to get term with ID %s: %v", id, err)
	}
	term := Clickable{
		ID:   data.ID,
		Name: data.Name,
	}
	return term, nil
}