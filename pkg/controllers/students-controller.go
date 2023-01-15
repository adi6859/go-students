package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aditya/go-students/pkg/models"
	"github.com/aditya/go-students/pkg/utils"
	"github.com/gorilla/mux"
)

var NewStudent models.Student
var CreateSTudentAttendance models.StudentAttendance

//func getStudent(w http.ResponseWriter, r *http.Request) {
// 	newStudents := models.GetAllStudents()
// 	res, _ := json.Marshal(newStudents)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func GetStudentById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	studentId := vars["studentId"]
// 	ID, err := strconv.ParseInt(studentId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error with parsing")
// 	}
// 	studentDetails, _ := models.GetStudentById(ID)
// 	res, _ := json.Marshal(studentDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func CreateStudents(w http.ResponseWriter, r *http.Request) {
	CreateStudent := &models.Student{}
	utils.ParseBody(r, CreateStudent)
	b := CreateStudent.CreateStudents()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	NewTeacher := &models.Teacher{}
	utils.ParseBody(r, NewTeacher)
	b := NewTeacher.CreateTeacher()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error with parsing")
	}
	student := models.DeleteStudent(ID)
	res, _ := json.Marshal(student)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teacherId := vars["teacherId"]
	ID, err := strconv.ParseInt(teacherId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	teacher := models.DeleteTeacher(ID)
	res, _ := json.Marshal(teacher)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStudentAttendance(w http.ResponseWriter, r *http.Request) {
	GetSAtt := &models.StudentAttendance{}
	vars := mux.Vars(r)
	studentId := vars["Id"]
	thisMonth := vars["month"]
	thisYear := vars["year"]

	ID, err := strconv.ParseInt(studentId, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}
	MONTH, err := strconv.ParseInt(thisMonth, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}
	YEAR, err := strconv.ParseInt(thisYear, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}
	studentDetails := GetSAtt.CreateStudentAttend(ID, time.Month(MONTH), int(YEAR))
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTeacherAttendance(w http.ResponseWriter, r *http.Request) {
	GetTAtt := &models.TeacherAttendance{}
	vars := mux.Vars(r)
	teacherId := vars["Id"]
	thisMonth := vars["month"]
	thisYear := vars["year"]

	ID, err := strconv.ParseInt(teacherId, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}
	MONTH, err := strconv.ParseInt(thisMonth, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}
	YEAR, err := strconv.ParseInt(thisYear, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}
	teacherDetails := GetTAtt.CreateTeacherAttend(ID, time.Month(MONTH), int(YEAR))
	res, _ := json.Marshal(teacherDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetAttendanceOfClass(w http.ResponseWriter, r *http.Request) {
	AttendanceOfThisClass := &models.StudentAttendance{}
	vars := mux.Vars(r)
	classIs := vars["class"]
	CLASS, err := strconv.ParseInt(classIs, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}
	classDetails := AttendanceOfThisClass.GetAttendanceOfClass(int(CLASS))
	res, _ := json.Marshal(classDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func StudentPunchIn(w http.ResponseWriter, r *http.Request) {
	CreateSTudentAttendance := &models.StudentAttendance{}
	vars := mux.Vars(r)
	studentId := vars["Id"]
	studentClass := vars["class"]

	ID, err := strconv.ParseInt(studentId, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}
	CLASS, err := strconv.ParseInt(studentClass, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}
	CreateSTudentAttendance.StudentSId = uint(ID)
	CreateSTudentAttendance.Class = int(CLASS)
	CreateSTudentAttendance.Year = time.Now().Year()
	CreateSTudentAttendance.Month = time.Now().Month()
	CreateSTudentAttendance.Day = time.Now().YearDay()
	CreateSTudentAttendance.PunchIn = time.Now()
	s := CreateSTudentAttendance.StudentPunchIn()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func teacherPunchIn(w http.ResponseWriter, r *http.Request) {
	CreateTeacherAttendance := &models.TeacherAttendance{}
	vars := mux.Vars(r)
	teacherId := vars["Id"]

	ID, err := strconv.ParseInt(teacherId, 10, 32)
	if err != nil {
		fmt.Println("error while parsing")
	}

	if err != nil {
		fmt.Println("error while parsing")
	}
	CreateTeacherAttendance.TeacherTId = uint(ID)
	CreateTeacherAttendance.Year = time.Now().Year()
	CreateTeacherAttendance.Month = time.Now().Month()
	CreateTeacherAttendance.Day = time.Now().YearDay()
	CreateTeacherAttendance.PunchIn = time.Now()
	s := CreateSTudentAttendance.StudentPunchIn()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// func UpdateStudent(w http.ResponseWriter, r *http.Request) {
// 	var updateStudent = &models.Student{}
// 	utils.ParseBody(r, updateStudent)
// 	vars := mux.Vars(r)
// 	studentId := vars["studentId"]
// 	ID, err := strconv.ParseInt(studentId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error with parsing")
// 	}
// 	studentDetails, db := models.GetStudentById(ID)
// 	if updateStudent.Name != "" {
// 		studentDetails.Name = updateStudent.Name
// 	}
// 	db.Save(&studentDetails)
// 	res, _ := json.Marshal(studentDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }
