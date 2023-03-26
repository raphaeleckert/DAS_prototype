package models

import (
	"fmt"

	"dasagilestudieren/repo"
)

func GetSubject(id string) (Subject, error) {
	repo := repo.SubjectRepo
	data, err := repo.Read(id)
	if err != nil {
		return Subject{}, fmt.Errorf("failed to get subject with ID %s: %v", id, err)
	}
	subject := Subject{
		ID:        data.ID,
		ShortName: data.ShortName,
		Name:      data.Name,
		Owner:     data.Owner,
		Note:      data.Note,
		Remark:    data.Remark,
	}
	return subject, nil
}

func GetSubjectListBasic(owner string) ([]Clickable, error) {
	repo := repo.SubjectRepo
	data, err := repo.List(owner)
	clickables := []Clickable{}
	for _, s := range data {
		clickables = append(clickables, Clickable{
			ID:   s.ID,
			Name: s.Name,
		})
	}

	if err != nil {
		return nil, fmt.Errorf("failed retrieve subjects %v", err)
	}
	return clickables, nil
}

func GetSubjectBasic(id string) (Clickable, error) {
	repo := repo.SubjectRepo
	data, err := repo.Read(id)
	if err != nil {
		return Clickable{}, fmt.Errorf("failed to get subject with ID %s: %v", id, err)
	}
	subject := Clickable{
		ID:   data.ID,
		Name: data.Name,
	}
	if err != nil {
		return Clickable{}, fmt.Errorf("failed to get subject with ID %s: %v", id, err)
	}
	return subject, nil
}
