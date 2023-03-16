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
	fmt.Println("1")
	repo := repo.SubjectRepo
	fmt.Printf("2")
	fmt.Printf("%+v", repo)
	fmt.Printf("2")
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
		return Course{}, fmt.Errorf("failed to get course with ID %s: %v", id, err)
	}
	subject, err := GetSubject(data.Subject)
	if err != nil {
		return Course{}, fmt.Errorf("failed to get subject for course with ID %s: %v", id, err)
	}
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

func GetCourseBasic(id string) Clickable {
	return Clickable{
		ID:   id,
		Name: "Bachelorthesis Wintersemester 2023",
	}
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

func GetTopicsByCourse(course string) struct {
	Selected   []Topic
	Unselected []Topic
} {
	topic := Topic{
		ID: "topicId",
		Subject: Subject{
			ID:        "id",
			ShortName: "EXSB",
			Name:      "Example Subject",
			Owner:     "teacherid",
			Note:      "Subejct Note",
			Remark:    "Subejct Remark",
		},
		Title:              "Ich sage das Alphabet rückwärts auf",
		Detail:             "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
		Reference:          "BT-1",
		SolutionIdea:       "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea",
		Remark:             "Topic Remark",
		Tags:               []string{"buchstaben", "reihenfolge"},
		Importance:         IMP_ESSENTIAL,
		RequiredSupporters: SUP_HALF,
	}
	return struct {
		Selected   []Topic
		Unselected []Topic
	}{
		Selected:   []Topic{topic, topic, topic},
		Unselected: []Topic{topic, topic, topic},
	}
}
