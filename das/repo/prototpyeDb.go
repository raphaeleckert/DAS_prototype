package repo

import (
	"dasagilestudieren/prototype"
	"fmt"
)

var CourseRepo *CourseRepository
var SubjectRepo *SubjectRepository

func init() {
	db := prototype.PopulatePrototypeDb()
	CourseRepo = NewCourseRepository(&db)
	SubjectRepo = NewSubjectRepository(&db)
	fmt.Printf("%+v", CourseRepo)
	fmt.Printf("%+v", SubjectRepo)
}
