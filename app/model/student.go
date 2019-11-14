package model

import (
	"fmt"
	"time"
)

type Student struct {
	Id        int       `json:"id" form:"id" gorm:"column:id"`
	FirstName string    `json:"first_name" gorm:"column:first_name" form:"first_name"`
	LastName  string    `json:"last_name" gorm:"column:last_name" form:"last_name"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (student Student) All() []Student {
	students := []Student{}
	db.Find(&students)
	return students
}

func (student Student) Store() Student {
	db.Create(&student)
	fmt.Println("stored", student)
	return student
}

func (student Student) Show(id string) Student {
	db.First(&student, id)
	fmt.Println(student)
	return student
}
