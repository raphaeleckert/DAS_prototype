package models

import (
	"fmt"

	"dasagilestudieren/repo"
)

// Course
func GetCourse(id string) (Course, error) {
	repo := repo.CourseRepo
	data, err := repo.Read(id)
	if err != nil {
		return Course{}, fmt.Errorf("failed to get term for course with ID %s: %v", id, err)
	}
	subject, err := GetSubject(data.Subject)
	term, err := GetTerm(data.Term)
	if err != nil {
		return Course{}, fmt.Errorf("failed to get term for course with ID %s: %v", id, err)
	}
	course := Course{
		ID:        data.ID,
		Subject:   subject,
		Term:      term,
		Owner:     data.Owner,
		BeginDate: data.BeginDate,
		FinalDate: data.FinalDate,
		CloseDate: data.CloseDate,
		Remark:    data.Remark,
		Note:      data.Note,
		Topics:    data.Topics,
	}
	return course, nil
}

func GetCourseBasic(id string) (Clickable, error) {
	repo := repo.CourseRepo
	data, err := repo.Read(id)
	subject, err := GetSubject(data.Subject)
	term, err := GetTerm(data.Term)
	if err != nil {
		return Clickable{}, fmt.Errorf("failed to get course with ID %s: %v", id, err)
	}
	course := Clickable{
		ID:   data.ID,
		Name: fmt.Sprintf("%s | %s", subject.Name, term.Name),
	}
	return course, nil
}

func GetCourseBasicListBySubject(subjectID string) ([]Clickable, error) {
	repo := repo.CourseRepo
	data, err := repo.ListBySubject(subjectID)
	clickables := []Clickable{}
	for _, s := range data {
		subject, err := GetSubjectBasic(s.Subject)
		term, err := GetTermBasic(s.Subject)
		if err != nil {
			return []Clickable{}, fmt.Errorf("failed to get courses from subject %s: %v", subjectID, err)
		}
		clickables = append(clickables, Clickable{
			ID:   s.ID,
			Name: fmt.Sprintf("%s | %s", subject.Name, term.Name),
		})
	}
	if err != nil {
		return []Clickable{}, fmt.Errorf("failed to get courses from subject %s: %v", subjectID, err)
	}
	return clickables, nil
}

func UpdateTopicListForCourse(courseId string, topicList []string) error {
	repo := repo.CourseRepo
	course, err := repo.Read(courseId)
	if err != nil {
		return fmt.Errorf("failed to get course with ID %s: %v", courseId, err)
	}
	course.Topics = topicList
	err = repo.Update(course)
	if err != nil {
		return fmt.Errorf("failed to update course with ID %s: %v", courseId, err)
	}
	return nil
}
