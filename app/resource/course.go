package resource

import (
	"pro/restapi/app/model"
)

type CourseResource struct {
	Id      int
	Title   string
	Student uint
}

func (CourseResource) Collection(courses []model.Course) []CourseResource {

	var collection []CourseResource

	for _, item := range courses {
		collection = append(collection, CourseResource{
			Id:      item.Id,
			Title:   item.Title,
			Student: item.StudentId,
		})
	}

	return collection

}

func (CourseResource) One(course model.Course) CourseResource {

	var one CourseResource

	one.Id = course.Id
	one.Title = course.Title
	one.Student = course.StudentId

	return one
}
