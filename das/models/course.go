package models

import (
	"fmt"
	"time"

	"dasagilestudieren/repo"
)

// Topic Importances
const (
	IMP_ESSENTIAL = "Prerequisite for other topics"
	IMP_MUST      = "Basic knowledge"
	IMP_SHOULD    = "Average knowledge"
	IMP_COULD     = "Special knowledge"
	IMP_VOLUNTARY = "Voluntary topic"
)

// Topic Required Supporters
const (
	SUP_NONE = "No Supporters"
	SUP_ONE  = "One Supporter"
	SUP_TWO  = "Two Supporters"
	SUP_HALF = "Half of the Team"
	SUP_BUT  = "All but One"
	SUP_ALL  = "All Team Members"
)

type Subject struct {
	ID        string
	ShortName string
	Name      string
	Owner     string
	Note      string
	Remark    string
}

type Course struct {
	ID        string
	Subject   Subject
	Term      Term
	Owner     string
	BeginDate time.Time
	FinalDate time.Time
	CloseDate time.Time
	Remark    string
	Note      string
	Topics    []string
}

type Topic struct {
	ID                 string
	Title              string
	Subject            Subject
	Number             string
	Detail             string
	Reference          string
	SolutionIdea       string
	Remark             string
	Tags               []string
	Importance         string
	RequiredSupporters string
}

type Term struct {
	ID        string
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Note      string
	Remark    string
}

// Subject
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
	subject := Clickable{
		ID:   data.ID,
		Name: data.Name,
	}
	if err != nil {
		return Clickable{}, fmt.Errorf("failed to get subject with ID %s: %v", id, err)
	}
	return subject, nil
}

// Course
func GetCourse(id string) (Course, error) {
	repo := repo.CourseRepo
	data, err := repo.Read(id)
	if err != nil {
		return Course{}, fmt.Errorf("failed to get term for course with ID %s: %v", id, err)
	}
	fmt.Printf(" subject %v ", data.Subject)
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

// Topic
func GetTopic(id string) (Topic, error) {
	repo := repo.TopicRepo
	data, err := repo.Read(id)
	if err != nil {
		return Topic{}, fmt.Errorf("failed to get topic with ID %s: %v", id, err)
	}
	subject, err := GetSubject(data.Subject)
	if err != nil {
		return Topic{}, fmt.Errorf("failed to get subject for course with ID %s: %v", id, err)
	}
	topic := Topic{
		ID:                 data.ID,
		Title:              data.Title,
		Subject:            subject,
		Number:             data.Number,
		Detail:             data.Detail,
		Reference:          data.Reference,
		SolutionIdea:       data.SolutionIdea,
		Remark:             data.Remark,
		Tags:               data.Tags,
		Importance:         data.Importance,
		RequiredSupporters: data.RequiredSupporters,
	}
	return topic, nil
}

func GetTopicBasic(id string) (Clickable, error) {
	repo := repo.TopicRepo
	data, err := repo.Read(id)
	if err != nil {
		return Clickable{}, fmt.Errorf("failed to get topic with ID %s: %v", id, err)
	}
	topic := Clickable{
		ID:   data.ID,
		Name: data.Title,
	}
	return topic, nil
}

func GetTopicsByCourse(courseID string) (struct {
	Selected   []Topic
	Unselected []Topic
}, error) {
	topicRepo := repo.TopicRepo
	course, err := GetCourse(courseID)
	subject, err := GetSubject(course.Subject.ID)
	if err != nil {
		return struct {
			Selected   []Topic
			Unselected []Topic
		}{}, fmt.Errorf("failed to get topics for course %s: %v", course, err)
	}

	allTopics, err := topicRepo.ListBySubject(subject.ID)
	populatedTopics := []Topic{}

	var selected = []Topic{}
	var unselected = []Topic{}

	for _, topic := range allTopics {
		populatedTopics = append(populatedTopics, Topic{
			ID:                 topic.ID,
			Title:              topic.Title,
			Subject:            subject,
			Number:             topic.Number,
			Detail:             topic.Detail,
			Reference:          topic.Reference,
			SolutionIdea:       topic.SolutionIdea,
			Remark:             topic.Remark,
			Tags:               topic.Tags,
			Importance:         topic.Importance,
			RequiredSupporters: topic.RequiredSupporters,
		})
	}

	for _, subjectTopic := range populatedTopics {
		found := false
		for _, courseTopic := range course.Topics {
			if subjectTopic.ID == courseTopic {
				selected = append(selected, subjectTopic)
				found = true
				break
			}
		}
		if !found {
			unselected = append(unselected, subjectTopic)
		}
	}

	return struct {
		Selected   []Topic
		Unselected []Topic
	}{
		Selected:   selected,
		Unselected: unselected,
	}, nil
}

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
