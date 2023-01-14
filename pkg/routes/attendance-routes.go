package routes

import (
	"github.com/aditya/go-students/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterStudentsroutes = func(router *mux.Router) {
	router.HandleFunc("/student/", controllers.CreateStudents).Methods("POST")
	router.HandleFunc("/student/", controllers.CreateStudents).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.CreateStudents).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.CreateStudents).Methods("PUT")
	router.HandleFunc("/student/{studentId}", controllers.CreateStudents).Methods("DELETE")
}
