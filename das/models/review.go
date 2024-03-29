/*
This file is part of the DAS_prototype and can be used under the GNU Affero General License

https://www.gnu.org/licenses/agpl-3.0.html

*/

package models

import (
	"fmt"

	"dasagilestudieren/repo"
)

// Review
func GetReview(id string) (Review, error) {
	repo := repo.ReviewRepo
	review, err := repo.Read(id)
	course, err := GetCourse(review.Course)
	if err != nil {
		return Review{}, fmt.Errorf("failed to get review for ID %s: %v", id, err)
	}
	return Review{
		ID:         review.ID,
		Course:     course,
		ReviewDate: review.ReviewDate,
		MaxReviews: review.MaxReviews,
		BeginDate:  review.BeginDate,
		EndDate:    review.EndDate,
		Note:       review.Note,
		Remark:     review.Remark,
	}, nil
}

func GetReviewBasic(id string) (Clickable, error) {
	repo := repo.ReviewRepo
	review, err := repo.Read(id)
	course, err := GetCourse(review.Course)
	if err != nil {
		return Clickable{}, fmt.Errorf("failed to get review for ID %s: %v", id, err)
	}
	return Clickable{
		ID:   review.ID,
		Name: fmt.Sprintf("%s | %s", course.Subject.Name, review.ReviewDate.Format("Jan 2, 2006 at 3:04pm (MST)")),
	}, nil
}
