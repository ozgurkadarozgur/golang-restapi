package model

import (
	"fmt"
)

type Course struct {
	Id        int     `gorm:"column:id"`
	Title     string  `gorm:"column:title" form:"title"`
	StudentId uint    `gorm:"column:studentId" form:"studentId"`
	Student   Student `gorm:"foreign_key:StudentId"`
}

func GetStudent() {

}

func (course Course) All() []Course {
	var courses []Course
	db.Find(&courses)
	return courses
}

func (course Course) Store() Course {
	db.Create(&course)
	fmt.Println("stored", course)
	return course
}
