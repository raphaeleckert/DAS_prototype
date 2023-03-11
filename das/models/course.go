package models

import (
	"time"
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

func GetSubject(id string) Subject {
	return Subject{
		ID:        id,
		ShortName: "EXSB",
		Name:      "Example Subject",
		Owner:     "teacherid",
		Note:      "Subejct Note",
		Remark:    "Subejct Remark",
	}
}

func GetCourse(id string) Course {
	return Course{
		ID:        id,
		Subject:   GetSubject("subjectid"),
		Term:      GetTerm("termid"),
		Owner:     "teacherid",
		BeginDate: time.Date(2022, time.January, 01, 00, 00, 00, 0, time.UTC),
		FinalDate: time.Date(2024, time.January, 01, 00, 00, 00, 0, time.UTC),
		CloseDate: time.Date(2025, time.January, 01, 00, 00, 00, 0, time.UTC),
		Remark:    "Course Remark",
		Note:      "Course Note",
		Topics:    []string{"topicid"},
	}
}

func GetTopic(id string) Topic {
	if id == "blankid" {
		return Topic{
			ID:                 "blankid",
			Subject:            GetSubject("subjectid"),
			Title:              "-",
			Detail:             "-",
			Reference:          "-",
			SolutionIdea:       "-",
			Remark:             "-",
			Tags:               []string{},
			Importance:         IMP_SHOULD,
			RequiredSupporters: SUP_HALF,
		}
	}
	return Topic{
		ID:                 id,
		Subject:            GetSubject("subjectid"),
		Title:              "Ich sage das Alphabet rückwärts auf",
		Detail:             "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
		Reference:          "BT-1",
		SolutionIdea:       "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea",
		Remark:             "Topic Remark",
		Tags:               []string{"buchstaben", "reihenfolge"},
		Importance:         IMP_ESSENTIAL,
		RequiredSupporters: SUP_HALF,
	}
}

func GetTopicBasic(id string) Clickable {
	return Clickable{
		ID:   id,
		Name: "BT-1: Ich sage das Alphabet rückwärts auf",
	}
}

func GetTerm(id string) Term {
	return Term{
		ID:        id,
		Name:      "Example Term",
		StartDate: time.Date(2022, time.January, 01, 00, 00, 00, 0, time.UTC),
		EndDate:   time.Date(2024, time.January, 01, 00, 00, 00, 0, time.UTC),
		Note:      "Term Note",
		Remark:    "Term Remark",
	}
}

func GetTermBasic(id string) Clickable {
	return Clickable{
		ID:   id,
		Name: "Example Term",
	}
}

func GetCourseBasic(id string) Clickable {
	return Clickable{
		ID:   id,
		Name: "Bachelorthesis Wintersemester 2023",
	}
}

func GetSubjectBasic(id string) Clickable {
	return Clickable{
		ID:   id,
		Name: "Bachelorthesis",
	}
}

func GetTopicsBySubject(subject string) []Topic {
	topic := Topic{
		ID:                 "topicId",
		Subject:            GetSubject(subject),
		Title:              "Ich sage das Alphabet rückwärts auf",
		Detail:             "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
		Reference:          "BT-1",
		SolutionIdea:       "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea",
		Remark:             "Topic Remark",
		Tags:               []string{"buchstaben", "reihenfolge"},
		Importance:         IMP_ESSENTIAL,
		RequiredSupporters: SUP_HALF,
	}
	return []Topic{topic, topic, topic}
}
