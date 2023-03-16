package repo

import (
	"fmt"

	"dasagilestudieren/prototype"
)

func NewCourseRepository(db *prototype.PrototypeDb) *CourseRepository {
	return &CourseRepository{db: db}
}

func NewSubjectRepository(db *prototype.PrototypeDb) *SubjectRepository {
	return &SubjectRepository{db: db}
}

func NewTopicRepository(db *prototype.PrototypeDb) *CourseRepository {
	return &CourseRepository{db: db}
}

func NewTermRepository(db *prototype.PrototypeDb) *CourseRepository {
	return &CourseRepository{db: db}
}

type SubjectInterface interface {
	Create(subject *prototype.Subject) error
	Read(id string) (*prototype.Subject, error)
	Update(subject *prototype.Subject) error
	Delete(id string) error
	ListByOwner(ownerID string) ([]*prototype.Subject, error)
}

type CourseInterface interface {
	Create(course *prototype.Course) error
	Read(id string) (*prototype.Course, error)
	Update(course *prototype.Course) error
	Delete(id string) error
	ListBySubject(subjectID string) ([]*prototype.Course, error)
	ListByTerm(termID string) ([]*prototype.Course, error)
}

type TopicInterface interface {
	Create(topic *prototype.Topic) error
	Read(id string) (*prototype.Topic, error)
	Update(topic *prototype.Topic) error
	Delete(id string) error
	ListBySubject(subjectID string) ([]*prototype.Topic, error)
}

type TermInterface interface {
	Create(term *prototype.Term) error
	Read(id string) (*prototype.Term, error)
	Update(term *prototype.Term) error
	Delete(id string) error
}

type SubjectRepository struct {
	db *prototype.PrototypeDb
}

type CourseRepository struct {
	db *prototype.PrototypeDb
}

type TopicRepository struct {
	db *prototype.PrototypeDb
}

type TermRepository struct {
	db *prototype.PrototypeDb
}

// Course
func (repo *CourseRepository) Create(course *prototype.Course) error {
	repo.db.Courses[course.ID] = course
	return nil
}

func (repo *CourseRepository) Read(id string) (*prototype.Course, error) {
	course, ok := repo.db.Courses[id]
	if !ok {
		return nil, fmt.Errorf("course not found")
	}
	return course, nil
}

func (repo *CourseRepository) Update(course *prototype.Course) error {
	_, ok := repo.db.Courses[course.ID]
	if !ok {
		return fmt.Errorf("course not found")
	}
	repo.db.Courses[course.ID] = course
	return nil
}

func (repo *CourseRepository) Delete(id string) error {
	_, ok := repo.db.Courses[id]
	if !ok {
		return fmt.Errorf("course not found")
	}
	delete(repo.db.Courses, id)
	return nil
}

func (repo *CourseRepository) ListBySubject(subjectID string) ([]*prototype.Course, error) {
	var courses []*prototype.Course
	for _, course := range repo.db.Courses {
		if course.Subject == subjectID {
			courses = append(courses, course)
		}
	}
	return courses, nil
}

func (repo *CourseRepository) ListByTerm(termID string) ([]*prototype.Course, error) {
	var courses []*prototype.Course
	for _, course := range repo.db.Courses {
		if course.Term == termID {
			courses = append(courses, course)
		}
	}
	return courses, nil
}

// Subject
func (repo *SubjectRepository) Create(subject *prototype.Subject) error {
	repo.db.Subjects[subject.ID] = subject
	return nil
}

func (repo *SubjectRepository) Read(id string) (*prototype.Subject, error) {
	subject, ok := repo.db.Subjects[id]
	if !ok {
		return nil, fmt.Errorf("subject not found")
	}
	return subject, nil
}

func (repo *SubjectRepository) List(owner string) ([]*prototype.Subject, error) {
	subjects := []*prototype.Subject{}
	for _, s := range repo.db.Subjects {
		if s.Owner == owner {
			subjects = append(subjects, s)
		}
	}
	return subjects, nil
}

func (repo *SubjectRepository) Update(subject *prototype.Subject) error {
	_, ok := repo.db.Subjects[subject.ID]
	if !ok {
		return fmt.Errorf("subject not found")
	}
	repo.db.Subjects[subject.ID] = subject
	return nil
}

func (repo *SubjectRepository) Delete(id string) error {
	_, ok := repo.db.Subjects[id]
	if !ok {
		return fmt.Errorf("subject not found")
	}
	delete(repo.db.Subjects, id)
	return nil
}

// Topic
func (repo *TopicRepository) Create(topic *prototype.Topic) error {
	repo.db.Topics[topic.ID] = topic
	return nil
}

func (repo *TopicRepository) Read(id string) (*prototype.Topic, error) {
	topic, ok := repo.db.Topics[id]
	if !ok {
		return nil, fmt.Errorf("topic not found")
	}
	return topic, nil
}

func (repo *TopicRepository) Update(topic *prototype.Topic) error {
	_, ok := repo.db.Topics[topic.ID]
	if !ok {
		return fmt.Errorf("topic not found")
	}
	repo.db.Topics[topic.ID] = topic
	return nil
}

func (repo *TopicRepository) Delete(id string) error {
	_, ok := repo.db.Topics[id]
	if !ok {
		return fmt.Errorf("topic not found")
	}
	delete(repo.db.Topics, id)
	return nil
}

func (repo *TopicRepository) ListBySubject(subjectID string) ([]*prototype.Topic, error) {
	var topics []*prototype.Topic
	for _, topic := range repo.db.Topics {
		if topic.Subject == subjectID {
			topics = append(topics, topic)
		}
	}
	return topics, nil
}

// Term
func (repo *TermRepository) Create(term *prototype.Term) error {
	repo.db.Terms[term.ID] = term
	return nil
}

func (repo *TermRepository) Read(id string) (*prototype.Term, error) {
	term, ok := repo.db.Terms[id]
	if !ok {
		return nil, fmt.Errorf("term not found")
	}
	return term, nil
}

func (repo *TermRepository) Update(term *prototype.Term) error {
	_, ok := repo.db.Terms[term.ID]
	if !ok {
		return fmt.Errorf("term not found")
	}
	repo.db.Terms[term.ID] = term
	return nil
}

func (repo *TermRepository) Delete(id string) error {
	_, ok := repo.db.Terms[id]
	if !ok {
		return fmt.Errorf("term not found")
	}
	delete(repo.db.Terms, id)
	return nil
}
