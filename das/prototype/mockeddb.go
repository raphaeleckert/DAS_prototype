package prototype

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

// Solution States
const (
	STATE_UNKNOWN  = "Unknown"
	STATE_INWORK   = "In Work"
	STATE_READY    = "Ready to Review"
	STATE_INREVIEW = "In Review"
	STATE_REJECTED = "Rejected"
	STATE_DONE     = "Done"
	STATE_FINAL    = "Final"
)

type PrototypeDb struct {
	Courses    map[string]*Course
	Subjects   map[string]*Subject
	Terms      map[string]*Term
	Teams      map[string]*Team
	Topics     map[string]*Topic
	Solutions  map[string]*Solution
	Proposals  map[string]*Proposal
	Supporters map[string]*Supporter
	Reviews    map[string]*Review
	Ratings    map[string]*Rating
}

func PopulatePrototypeDb() PrototypeDb {
	mockedDb := PrototypeDb{
		Courses:    make(map[string]*Course),
		Subjects:   make(map[string]*Subject),
		Terms:      make(map[string]*Term),
		Teams:      make(map[string]*Team),
		Topics:     make(map[string]*Topic),
		Solutions:  make(map[string]*Solution),
		Proposals:  make(map[string]*Proposal),
		Supporters: make(map[string]*Supporter),
		Reviews:    make(map[string]*Review),
		Ratings:    make(map[string]*Rating),
	}
	// Subject
	sub1 := &Subject{ID: "subject1", ShortName: "sub1_short", Name: "Projektmanagement", Owner: "teacher", Note: "Note 1", Remark: "Remark 1"}
	sub2 := &Subject{ID: "subject2", ShortName: "sub1_short", Name: "Einführung in die Programmierung", Owner: "teacher", Note: "Note 2", Remark: "Remark 2"}
	mockedDb.Subjects[sub1.ID] = sub1
	mockedDb.Subjects[sub2.ID] = sub2

	// Topic
	topic1 := &Topic{ID: "topic1", Title: "Topic 1", Subject: sub1.ID, Number: "1", Detail: "Details for Topic 1", Reference: "BTXT-1", SolutionIdea: "Solution idea for Topic 1", Remark: "Remark 1", Tags: []string{"tag1", "tag2"}, Importance: IMP_ESSENTIAL, RequiredSupporters: SUP_BUT}
	topic2 := &Topic{ID: "topic2", Title: "Topic 2", Subject: sub1.ID, Number: "2", Detail: "Details for Topic 2", Reference: "BTXT-2", SolutionIdea: "Solution idea for Topic 2", Remark: "Remark 2", Tags: []string{"tag2", "tag3"}, Importance: IMP_SHOULD, RequiredSupporters: SUP_HALF}
	topic3 := &Topic{ID: "topic3", Title: "Topic 3", Subject: sub1.ID, Number: "3", Detail: "Details for Topic 3", Reference: "BTXT-3", SolutionIdea: "Solution idea for Topic 3", Remark: "Remark 3", Tags: []string{"tag2", "tag3"}, Importance: IMP_SHOULD, RequiredSupporters: SUP_HALF}
	topic4 := &Topic{ID: "topic4", Title: "Topic 4", Subject: sub1.ID, Number: "4", Detail: "Details for Topic 4", Reference: "BTXT-4", SolutionIdea: "Solution idea for Topic 4", Remark: "Remark 4", Tags: []string{"tag2", "tag3"}, Importance: IMP_SHOULD, RequiredSupporters: SUP_HALF}
	topic5 := &Topic{ID: "topic5", Title: "Topic 5", Subject: sub1.ID, Number: "5", Detail: "Details for Topic 5", Reference: "BTXT-5", SolutionIdea: "Solution idea for Topic 5", Remark: "Remark 5", Tags: []string{"tag2", "tag3"}, Importance: IMP_SHOULD, RequiredSupporters: SUP_HALF}
	topic6 := &Topic{ID: "topic6", Title: "Topic 6", Subject: sub1.ID, Number: "6", Detail: "Details for Topic 6", Reference: "BTXT-6", SolutionIdea: "Solution idea for Topic 6", Remark: "Remark 6", Tags: []string{"tag2", "tag3"}, Importance: IMP_SHOULD, RequiredSupporters: SUP_HALF}
	mockedDb.Topics[topic1.ID] = topic1
	mockedDb.Topics[topic2.ID] = topic2
	mockedDb.Topics[topic3.ID] = topic3
	mockedDb.Topics[topic4.ID] = topic4
	mockedDb.Topics[topic5.ID] = topic5
	mockedDb.Topics[topic6.ID] = topic6

	// Term
	term1 := &Term{ID: "term1", Name: "Term 1", StartDate: time.Now(), EndDate: time.Now().Add(time.Hour * 24 * 90), Note: "note1", Remark: "remark1"}
	term2 := &Term{ID: "term2", Name: "Term 2", StartDate: time.Now().Add(time.Hour * 24 * 90), EndDate: time.Now().Add(time.Hour * 24 * 180), Note: "note2", Remark: "remark2"}
	mockedDb.Terms[term1.ID] = term1
	mockedDb.Terms[term2.ID] = term2

	// Course
	course1 := &Course{ID: "course1", Subject: sub1.ID, Term: term1.ID, Owner: "teacher", BeginDate: time.Now(), FinalDate: time.Now().Add(time.Hour * 24 * 7), CloseDate: time.Now().Add(time.Hour * 24 * 14), Remark: "remark1", Note: "note1", Topics: []string{"topic1", "topic2"}}
	course2 := &Course{ID: "course2", Subject: sub1.ID, Term: term1.ID, Owner: "teacher", BeginDate: time.Now().Add(time.Hour * 24 * 7), FinalDate: time.Now().Add(time.Hour * 24 * 14), CloseDate: time.Now().Add(time.Hour * 24 * 21), Remark: "remark2", Note: "note2", Topics: []string{topic1.ID, topic2.ID, topic3.ID}}
	mockedDb.Courses[course1.ID] = course1
	mockedDb.Courses[course2.ID] = course2

	// Team
	team1 := &Team{ID: "team1", Course: course1.ID, Number: 1, Member: []string{"Alice", "Bob"}, ReadOnly: false, Note: "Note 1", Remark: "Remark 1"}
	team2 := &Team{ID: "team2", Course: course2.ID, Number: 2, Member: []string{"Charlie", "David"}, ReadOnly: false, Note: "Note 2", Remark: "Remark 2"}
	mockedDb.Teams[team1.ID] = team1
	mockedDb.Teams[team2.ID] = team2

	// Solution
	solution1 := &Solution{ID: "solution1", Team: team1.ID, Topic: topic1.ID, State: STATE_INWORK, Remark: "Remark 1"}
	solution2 := &Solution{ID: "solution2", Team: team1.ID, Topic: topic2.ID, State: STATE_INREVIEW, Remark: "Remark 2"}
	solution3 := &Solution{ID: "solution3", Team: team1.ID, Topic: topic3.ID, State: STATE_READY, Remark: "Remark 2"}
	solution4 := &Solution{ID: "solution4", Team: team1.ID, Topic: topic4.ID, State: STATE_READY, Remark: "Remark 2"}
	solution5 := &Solution{ID: "solution5", Team: team1.ID, Topic: topic5.ID, State: STATE_INREVIEW, Remark: "Remark 2"}
	solution6 := &Solution{ID: "solution6", Team: team1.ID, Topic: topic6.ID, State: STATE_FINAL, Remark: "Remark 2"}
	mockedDb.Solutions[solution1.ID] = solution1
	mockedDb.Solutions[solution2.ID] = solution2
	mockedDb.Solutions[solution3.ID] = solution3
	mockedDb.Solutions[solution4.ID] = solution4
	mockedDb.Solutions[solution5.ID] = solution5
	mockedDb.Solutions[solution6.ID] = solution6

	// Proposal
	proposal1 := &Proposal{ID: "proposal1", Solution: solution1.ID, ModifiedBy: "Alice", Detail: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et"}
	proposal2 := &Proposal{ID: "proposal2", Solution: solution2.ID, ModifiedBy: "Charlie", Detail: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et"}
	mockedDb.Proposals[proposal1.ID] = proposal1
	mockedDb.Proposals[proposal2.ID] = proposal2

	// Supporter
	supporter1 := &Supporter{ID: "supporter1", Proposal: proposal1.ID, Member: "Bob"}
	supporter2 := &Supporter{ID: "supporter2", Proposal: proposal1.ID, Member: "David"}
	mockedDb.Supporters[supporter1.ID] = supporter1
	mockedDb.Supporters[supporter2.ID] = supporter2

	// Review
	review1 := &Review{ID: "r1", Course: "course1", ReviewDate: time.Now(), MaxReviews: 10, BeginDate: time.Now(), EndDate: time.Now().Add(time.Hour * 24), Note: "Note 1", Remark: "Remark 1"}
	review2 := &Review{ID: "r2", Course: "course2", ReviewDate: time.Now(), MaxReviews: 5, BeginDate: time.Now(), EndDate: time.Now().Add(time.Hour * 24), Note: "Note 2", Remark: "Remark 2"}
	mockedDb.Reviews[review1.ID] = review1
	mockedDb.Reviews[review2.ID] = review2

	// Rating
	rating1 := &Rating{ID: "ra1", Proposal: proposal1.ID, Review: "r1", Rating: 2, Remark: "Remark 1"}
	rating2 := &Rating{ID: "ra2", Proposal: proposal2.ID, Review: "r2", Rating: 3, Remark: "Remark 2"}
	mockedDb.Ratings[rating1.ID] = rating1
	mockedDb.Ratings[rating2.ID] = rating2

	return mockedDb
}
