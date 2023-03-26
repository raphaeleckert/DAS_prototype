/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package models

import (
	"fmt"

	"dasagilestudieren/prototype"
	"dasagilestudieren/repo"
)

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

func GetTopicsByCourseBasic(courseID string) ([]Clickable, error) {
	course, err := GetCourse(courseID)
	clickableTopics := []Clickable{}
	for _, topic := range course.Topics {
		clickableTopic, _ := GetTopicBasic(topic)
		clickableTopics = append(clickableTopics, clickableTopic)
	}
	if err != nil {
		return []Clickable{}, fmt.Errorf("failed to get topics for course %s: %v", course, err)
	}
	return clickableTopics, nil
}

func GetTopicsBySubject(subjectID string) ([]Topic, error) {
	topicRepo := repo.TopicRepo
	subject, err := GetSubject(subjectID)
	if err != nil {
		return []Topic{}, fmt.Errorf("failed to get topics for subject %s: %v", subjectID, err)
	}

	allTopics, err := topicRepo.ListBySubject(subject.ID)
	populatedTopics := []Topic{}
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
	if err != nil {
		return []Topic{}, fmt.Errorf("failed to get topics for subject %s: %v", subjectID, err)
	}
	return populatedTopics, nil
}

func GetTopicsBySubjectBasic(subjectID string) ([]Clickable, error) {
	topicRepo := repo.TopicRepo
	subject, err := GetSubject(subjectID)
	if err != nil {
		return []Clickable{}, fmt.Errorf("failed to get topics for subject %s: %v", subjectID, err)
	}

	allTopics, err := topicRepo.ListBySubject(subject.ID)
	clickableTopics := []Clickable{}
	for _, topic := range allTopics {
		clickableTopics = append(clickableTopics, Clickable{
			ID:   topic.ID,
			Name: topic.Title,
		})
	}
	if err != nil {
		return []Clickable{}, fmt.Errorf("failed to get topics for subject %s: %v", subjectID, err)
	}
	return clickableTopics, nil
}

func CreateTopic(topic Topic) error {
	repo := repo.TopicRepo
	newTopic := prototype.Topic{
		ID:                 topic.ID,
		Title:              topic.Title,
		Subject:            topic.Subject.ID,
		Number:             topic.Number,
		Detail:             topic.Detail,
		Reference:          topic.Reference,
		SolutionIdea:       topic.SolutionIdea,
		Remark:             topic.Remark,
		Tags:               topic.Tags,
		Importance:         topic.Importance,
		RequiredSupporters: topic.RequiredSupporters,
	}
	err := repo.Create(&newTopic)
	if err != nil {
		return fmt.Errorf("failed to create topic: %v", err)
	}
	return nil
}
