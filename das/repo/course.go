package repo

import (
	"dasagilestudieren/models"
)

type SubjectRepository interface {
	Create(subject *models.Subject) error
	Read(id string) (*models.Subject, error)
	Update(subject *models.Subject) error
	Delete(id string) error
}

type CourseRepository interface {
	Create(course *models.Course) error
	Read(id string) (*models.Course, error)
	Update(course *models.Course) error
	Delete(id string) error
	ListBySubject(subjectID string) ([]*models.Course, error)
	ListByTerm(termID string) ([]*models.Course, error)
}

type TopicRepository interface {
	Create(topic *models.Topic) error
	Read(id string) (*models.Topic, error)
	Update(topic *models.Topic) error
	Delete(id string) error
	ListBySubject(subjectID string) ([]*models.Topic, error)
}

type TermRepository interface {
	Create(term *models.Term) error
	Read(id string) (*models.Term, error)
	Update(term *models.Term) error
	Delete(id string) error
}
