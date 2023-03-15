package prototype

import (
	"dasagilestudieren/models"
	"time"
)

type PrototypeDb struct {
	courses    map[string]*Course
	subjects   map[string]*Subject
	terms      map[string]*Term
	teams      map[string]*Team
	topics     map[string]*Topic
	solutions  map[string]*Solution
	proposals  map[string]*Proposal
	supporters map[string]*Supporter
	reviews    map[string]*Review
	ratings    map[string]*Rating
}

func PopulatePrototypeDb() PrototypeDb {
	mockedDb := PrototypeDb{
		courses:    make(map[string]*Course),
		subjects:   make(map[string]*Subject),
		terms:      make(map[string]*Term),
		teams:      make(map[string]*Team),
		topics:     make(map[string]*Topic),
		solutions:  make(map[string]*Solution),
		proposals:  make(map[string]*Proposal),
		supporters: make(map[string]*Supporter),
		reviews:    make(map[string]*Review),
		ratings:    make(map[string]*Rating),
	}
	// Subject
	sub1 := &Subject{ID: "subject1", ShortName: "sub1_short", Name: "Projektmanagement", Owner: "teacher", Note: "Note 1", Remark: "Remark 1"}
	sub2 := &Subject{ID: "subject2", ShortName: "sub1_short", Name: "Einführung in die Programmierung", Owner: "teacher", Note: "Note 2", Remark: "Remark 2"}
	mockedDb.subjects[sub1.ID] = sub1
	mockedDb.subjects[sub2.ID] = sub2

	// Topic
	topic1 := &Topic{ID: "topic1", Title: "Topic 1", Subject: sub1.ID, Number: "1", Detail: "Details for Topic 1", Reference: "BTXT-1", SolutionIdea: "Solution idea for Topic 1", Remark: "Remark 1", Tags: []string{"tag1", "tag2"}, Importance: models.IMP_ESSENTIAL, RequiredSupporters: models.SUP_BUT}
	topic2 := &Topic{ID: "topic2", Title: "Topic 2", Subject: sub1.ID, Number: "2", Detail: "Details for Topic 2", Reference: "BTXT-2", SolutionIdea: "Solution idea for Topic 2", Remark: "Remark 2", Tags: []string{"tag2", "tag3"}, Importance: models.IMP_SHOULD, RequiredSupporters: models.SUP_HALF}
	topic3 := &Topic{ID: "topic3", Title: "Topic 3", Subject: sub1.ID, Number: "3", Detail: "Details for Topic 3", Reference: "BTXT-3", SolutionIdea: "Solution idea for Topic 3", Remark: "Remark 3", Tags: []string{"tag2", "tag3"}, Importance: models.IMP_SHOULD, RequiredSupporters: models.SUP_HALF}
	topic4 := &Topic{ID: "topic4", Title: "Topic 4", Subject: sub1.ID, Number: "4", Detail: "Details for Topic 4", Reference: "BTXT-4", SolutionIdea: "Solution idea for Topic 4", Remark: "Remark 4", Tags: []string{"tag2", "tag3"}, Importance: models.IMP_SHOULD, RequiredSupporters: models.SUP_HALF}
	topic5 := &Topic{ID: "topic5", Title: "Topic 5", Subject: sub1.ID, Number: "5", Detail: "Details for Topic 5", Reference: "BTXT-5", SolutionIdea: "Solution idea for Topic 5", Remark: "Remark 5", Tags: []string{"tag2", "tag3"}, Importance: models.IMP_SHOULD, RequiredSupporters: models.SUP_HALF}
	topic6 := &Topic{ID: "topic6", Title: "Topic 6", Subject: sub1.ID, Number: "6", Detail: "Details for Topic 6", Reference: "BTXT-6", SolutionIdea: "Solution idea for Topic 6", Remark: "Remark 6", Tags: []string{"tag2", "tag3"}, Importance: models.IMP_SHOULD, RequiredSupporters: models.SUP_HALF}
	mockedDb.topics[topic1.ID] = topic1
	mockedDb.topics[topic2.ID] = topic2
	mockedDb.topics[topic3.ID] = topic3
	mockedDb.topics[topic4.ID] = topic4
	mockedDb.topics[topic5.ID] = topic5
	mockedDb.topics[topic6.ID] = topic6

	// Course
	course1 := &Course{ID: "course1", Subject: sub1.ID, Term: term1.ID, Owner: "teacher", BeginDate: time.Now(), FinalDate: time.Now().Add(time.Hour * 24 * 7), CloseDate: time.Now().Add(time.Hour * 24 * 14), Remark: "remark1", Note: "note1", Topics: []string{"topic1", "topic2"}}
	course2 := &Course{ID: "course2", Subject: sub1.ID, Term: term1.ID, Owner: "teacher", BeginDate: time.Now().Add(time.Hour * 24 * 7), FinalDate: time.Now().Add(time.Hour * 24 * 14), CloseDate: time.Now().Add(time.Hour * 24 * 21), Remark: "remark2", Note: "note2", Topics: []string{topic1.ID, topic2.ID, topic3.ID}}
	mockedDb.courses[course1.ID] = course1
	mockedDb.courses[course2.ID] = course2

	// Term
	term1 := &Term{ID: "term1", Name: "Term 1", StartDate: time.Now(), EndDate: time.Now().Add(time.Hour * 24 * 90), Note: "note1", Remark: "remark1"}
	term2 := &Term{ID: "term2", Name: "Term 2", StartDate: time.Now().Add(time.Hour * 24 * 90), EndDate: time.Now().Add(time.Hour * 24 * 180), Note: "note2", Remark: "remark2"}

	// Team
	team1 := &Team{ID: "team1", Course: "course1", Number: 1, Member: []string{"Alice", "Bob"}, ReadOnly: false, Note: "Note 1", Remark: "Remark 1"}
	team2 := &Team{ID: "team2", Course: "course2", Number: 2, Member: []string{"Charlie", "David"}, ReadOnly: false, Note: "Note 2", Remark: "Remark 2"}
	mockedDb.teams[team1.ID] = team1
	mockedDb.teams[team2.ID] = team2

	// Solution
	solution1 := &Solution{ID: "solution1", Team: team1.ID, Topic: "topic1", State: models.STATE_INWORK, Remark: "Remark 1"}
	solution2 := &Solution{ID: "solution2", Team: team1.ID, Topic: "topic2", State: models.STATE_INREVIEW, Remark: "Remark 2"}
	solution3 := &Solution{ID: "solution3", Team: team1.ID, Topic: "topic2", State: models.STATE_READY, Remark: "Remark 2"}
	solution4 := &Solution{ID: "solution4", Team: team1.ID, Topic: "topic2", State: models.STATE_READY, Remark: "Remark 2"}
	solution5 := &Solution{ID: "solution5", Team: team1.ID, Topic: "topic2", State: models.STATE_INREVIEW, Remark: "Remark 2"}
	solution6 := &Solution{ID: "solution6", Team: team1.ID, Topic: "topic2", State: models.STATE_FINAL, Remark: "Remark 2"}
	mockedDb.solutions[solution1.ID] = solution1
	mockedDb.solutions[solution2.ID] = solution2
	mockedDb.solutions[solution3.ID] = solution3
	mockedDb.solutions[solution4.ID] = solution4
	mockedDb.solutions[solution5.ID] = solution5
	mockedDb.solutions[solution6.ID] = solution6

	// Proposal
	proposal1 := &Proposal{ID: "proposal1", Solution: solution1.ID, ModifiedBy: "Alice", Detail: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et"}
	proposal2 := &Proposal{ID: "proposal2", Solution: solution2.ID, ModifiedBy: "Charlie", Detail: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et"}
	mockedDb.proposals[proposal1.ID] = proposal1
	mockedDb.proposals[proposal2.ID] = proposal2

	// Supporter
	supporter1 := &Supporter{ID: "supporter1", Proposal: proposal1.ID, Member: "Bob"}
	supporter2 := &Supporter{ID: "supporter2", Proposal: proposal1.ID, Member: "David"}
	mockedDb.supporters[supporter1.ID] = supporter1
	mockedDb.supporters[supporter2.ID] = supporter2

	// Review
	review1 := &Review{ID: "r1", Course: "course1", ReviewDate: time.Now(), MaxReviews: 10, BeginDate: time.Now(), EndDate: time.Now().Add(time.Hour * 24), Note: "Note 1", Remark: "Remark 1"}
	review2 := &Review{ID: "r2", Course: "course2", ReviewDate: time.Now(), MaxReviews: 5, BeginDate: time.Now(), EndDate: time.Now().Add(time.Hour * 24), Note: "Note 2", Remark: "Remark 2"}
	mockedDb.reviews[review1.ID] = review1
	mockedDb.reviews[review2.ID] = review2

	// Rating
	rating1 := &Rating{ID: "ra1", Proposal: "p1", Review: "r1", Rating: 4, Remark: "Remark 1"}
	rating2 := &Rating{ID: "ra2", Proposal: "p2", Review: "r2", Rating: 5, Remark: "Remark 2"}
	mockedDb.ratings[rating1.ID] = rating1
	mockedDb.ratings[rating2.ID] = rating2

	return mockedDb
}
