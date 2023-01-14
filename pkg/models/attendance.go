package models

import (
	"github.com/aditya/go-students/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Student struct {
	gorm.Model
	Name      string `gorm:""json:"studentName"`
	StudentId string `json:"studentId"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}
func (s *Student) CreateStudents() *Student {
	db.NewRecord(s)
	db.Create(&s)
	return s
}
func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}
func GetStudentById(Id int64) (*Student, *gorm.DB) {
	var getStudent Student
	db := db.Where("ID=?", Id).Find(&getStudent)
	return &getStudent, db
}
func DeleteStudent(ID int64) Student {
	var student Student
	db.Where("ID=?", ID).Delete(student)
	return student
}
