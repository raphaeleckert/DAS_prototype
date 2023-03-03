package models

type Topic struct {
	ID                string
	Title             string
	Detail            string
	Reference         string
	SolutionIdea      string
	Remark            string
	Tags              []string
	Importance        string
	RequredSupporters string
}

func GetTopic(id string) Topic {
	return Topic{
		ID:                id,
		Title:             "Topic Title",
		Detail:            "Topic Detail",
		Reference:         "Topic Ref",
		SolutionIdea:      "Topic Solution Idea",
		Remark:            "Topic Remark",
		Tags:              []string{"TopicTag1", "TopicTag2"},
		Importance:        "Very Important",
		RequredSupporters: "More than Half",
	}
}
