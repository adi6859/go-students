package models

import (
	"time"

	"github.com/aditya/go-students/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Student struct {
	Name        string              `json:"studentname"`
	SId         uint                `gorm:"primary_key;auto_increment" json:"id"`
	Class       string              `json:"class"`
	SAttendance []StudentAttendance //this is used to connect two table
}
type Teacher struct {
	Name        string `json:"name"`
	TId         uint   `gorm:"primary_key;auto_increment" json:"id"`
	TAttendance []TeacherAttendance
}
type StudentAttendance struct {
	StudentSId uint       `json:"studentId"`
	Class      int        `json:"class"`
	PunchIn    time.Time  `json:"punchin"`
	PunchOut   time.Time  `json:"punchout"`
	Day        int        `json:"day"`
	Month      time.Month `json:"month"`
	Year       int        `json:"year"`
}
type TeacherAttendance struct {
	TeacherTId uint       `json:"teacherId"`
	PunchIn    time.Time  `json:"punchin"`
	PunchOut   time.Time  `json:"punchout"`
	Day        int        `json:"day"`
	Month      time.Month `json:"month"`
	Year       int        `json:"year"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&StudentAttendance{})
	db.AutoMigrate(&TeacherAttendance{})

}
func (s Student) CreateStudents() Student { //*Student here pointer is not required because it points the previous entry always
	db.NewRecord(s)
	db.Create(&s)
	return s
}
func (s Teacher) CreateTeacher() Teacher {
	db.NewRecord(s)
	db.Create(&s)
	return s
}
func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}
func DeleteStudent(ID int64) Student {
	var student Student
	db.Where("ID=?", uint(ID)).Delete(student)
	return student
}
func DeleteTeacher(ID int64) Teacher {
	var teacher Teacher
	db.Where("ID=?", uint(ID)).Delete(teacher)
	return teacher
}

// func GetStudentById(Id int64) (*Student, *gorm.DB) {
// 	var getStudent Student
// 	db := db.Where("ID=?", Id).Find(&getStudent)
// 	return &getStudent, db
// }

func (s *StudentAttendance) CreateStudentAttend(Id int64, month time.Month, year int) *[]StudentAttendance {
	var SAttendances []StudentAttendance
	db.Where("student_s_id=? AND month=? AND year=?", uint(Id), month, year).Find(&SAttendances)
	return &SAttendances
}

//	func (s *TeacherAttendance) CreateTeacherAttend() *TeacherAttendance {
//		db.NewRecord(s)
//		db.Create(&s)
//		return s
//	}
func (s *TeacherAttendance) CreateTeacherAttend(Id int64, month time.Month, year int) *[]TeacherAttendance {
	var TAttendances []TeacherAttendance
	db.Where("teacher_t_id=? AND month=? AND year=?", uint(Id), month, year).Find(&TAttendances)
	return &TAttendances
}

func (a *StudentAttendance) GetAttendanceOfClass(class int) *[]StudentAttendance {
	var getClassAt []StudentAttendance
	db.Where("student_class=?", class).Find(&getClassAt)
	return &getClassAt
}

func (a StudentAttendance) StudentPunchIn() StudentAttendance {
	db.NewRecord(a)
	db.Create(&a)
	return a
}
func (a TeacherAttendance) TeacherPunchIn() TeacherAttendance {
	db.NewRecord(a)
	db.Create(&a)
	return a
}
