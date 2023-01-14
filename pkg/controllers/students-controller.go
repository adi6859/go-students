package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aditya/go-students/pkg/models"
	"github.com/aditya/go-students/pkg/utils"
	"github.com/gorilla/mux"
)

var NewStudent models.Student

func getStudent(w http.ResponseWriter, r *http.Request) {
	newStudents := models.GetAllStudents()
	res, _ := json.Marshal(newStudents)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error with parsing")
	}
	studentDetails, _ := models.GetStudentById(ID)
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateStudents(w http.ResponseWriter, r *http.Request) {
	CreateStudent := &models.Student{}
	utils.ParseBody(r, CreateStudent)
	b := CreateStudent.CreateStudents()
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

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var updateStudent = &models.Student{}
	utils.ParseBody(r, updateStudent)
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error with parsing")
	}
	studentDetails, db := models.GetStudentById(ID)
	if updateStudent.Name != "" {
		studentDetails.Name = updateStudent.Name
	}
	db.Save(&studentDetails)
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
