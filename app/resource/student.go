package resource

import (
	"pro/restapi/app/model"
)

type StudentResource struct {
	Id        int
	FirstName string
	LastName  string
	TestData  string
}

func (StudentResource) Collection(students []model.Student) []StudentResource {

	var collection []StudentResource

	for _, item := range students {
		collection = append(collection, StudentResource{
			Id:        item.Id,
			FirstName: item.FirstName,
			LastName:  item.LastName,
		})
	}

	return collection

}

func (StudentResource) One(student model.Student) StudentResource {

	var one StudentResource

	one.Id = student.Id
	one.FirstName = student.FirstName
	one.LastName = student.LastName

	return one
}
