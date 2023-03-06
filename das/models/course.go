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

type Subject struct {
	ID        any
	ShortName any
	Name      any
	Owner     any
	Note      any
	Remark    any
}

type Course struct {
	ID        any
	Subject   any
	Term      any
	Owner     any
	BeginDate time.Time
	FinalDate time.Time
	CloseDate time.Time
	Remark    any
	Note      any
	Topics    []any
}

type Topic struct {
	ID                any
	Title             any
	Subject           any
	Number            any
	Detail            any
	Reference         any
	SolutionIdea      any
	Remark            any
	Tags              []any
	Importance        any
	RequredSupporters any
}

type Term struct {
	ID        any
	Name      any
	StartDate time.Time
	EndDate   time.Time
	Note      any
	Remark    any
}

func GetSubject(id string) Subject {
	return Subject{
		ID:        id,
		ShortName: "EXSB",
		Name:      "Example Subject",
		Owner:     GetTeacher("teacherid"),
		Note:      "Subejct Note",
		Remark:    "Subejct Remark",
	}
}

func GetCourse(id string) Course {
	return Course{
		ID:        id,
		Subject:   GetSubject("subjectid"),
		Term:      GetTerm("termid"),
		Owner:     GetTeacher("teacherid"),
		BeginDate: time.Date(2022, time.January, 01, 00, 00, 00, 0, time.UTC),
		FinalDate: time.Date(2024, time.January, 01, 00, 00, 00, 0, time.UTC),
		CloseDate: time.Date(2025, time.January, 01, 00, 00, 00, 0, time.UTC),
		Remark:    "Course Remark",
		Note:      "Course Note",
		Topics:    []any{"topicid"},
	}
}

func GetTopic(id string) Topic {
	return Topic{
		ID:                id,
		Subject:           "subjectid",
		Title:             "Topic Title",
		Detail:            "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
		Reference:         "Topic Ref",
		SolutionIdea:      "Topic Solution Idea",
		Remark:            "Topic Remark",
		Tags:              []any{"TopicTag1", "TopicTag2"},
		Importance:        IMP_ESSENTIAL,
		RequredSupporters: "More than Half",
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
