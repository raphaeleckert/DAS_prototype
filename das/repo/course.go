package repo

import (
	"fmt"

	"dasagilestudieren/models"
	"dasagilestudieren/prototype"
)

type SubjectRepository struct {
	db prototype.PrototypeDb
}

type CourseRepository struct {
	db prototype.PrototypeDb
}

type TopicRepository struct {
	db prototype.PrototypeDb
}

type TermRepository struct {
	db prototype.PrototypeDb
}

type SubjectInterface interface {
	Create(subject *models.Subject) error
	Read(id string) (*models.Subject, error)
	Update(subject *models.Subject) error
	Delete(id string) error
}

type CourseInterface interface {
	Create(course *models.Course) error
	Read(id string) (*models.Course, error)
	Update(course *models.Course) error
	Delete(id string) error
	ListBySubject(subjectID string) ([]*models.Course, error)
	ListByTerm(termID string) ([]*models.Course, error)
}

type TopicInterface interface {
	Create(topic *models.Topic) error
	Read(id string) (*models.Topic, error)
	Update(topic *models.Topic) error
	Delete(id string) error
	ListBySubject(subjectID string) ([]*models.Topic, error)
}

type TermInterface interface {
	Create(term *models.Term) error
	Read(id string) (*models.Term, error)
	Update(term *models.Term) error
	Delete(id string) error
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
