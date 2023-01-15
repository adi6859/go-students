package routes

import (
	"github.com/aditya/go-students/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterStudentsroutes = func(router *mux.Router) {
	router.HandleFunc("/api/student", controllers.CreateStudents).Methods("POST")
	router.HandleFunc("/api/teacher", controllers.CreateTeacher).Methods("POST")
	router.HandleFunc("/api/student/{studentId}/{month}/{year}", controllers.GetStudentAtt).Methods("GET")
	router.HandleFunc("/api/teacher/{teacherId}/{month}/{year}", controllers.GetTeacherAtt).Methods("GET")
	router.HandleFunc("/api/student/punchin", controllers.StudentPunchIn).Methods("POST")
	router.HandleFunc("/api/student/punchout", controllers.StudentPunchOut).Methods("POST")
	router.HandleFunc("/api/teacher/punchin", controllers.TeacherPunchIn).Methods("POST")
	router.HandleFunc("/api/teacher/punchout", controllers.TeacherPunchOut).Methods("POST")
	router.HandleFunc("/api/student/{class}", controllers.GetClassAtt).Methods("GET")

	// router.HandleFunc("/student/", controllers.CreateStudents).Methods("POST")
	// router.HandleFunc("/student/", controllers.CreateStudents).Methods("GET")
	// router.HandleFunc("/student/{studentId}", controllers.CreateStudents).Methods("GET")
	// router.HandleFunc("/student/{studentId}", controllers.CreateStudents).Methods("PUT")
	// router.HandleFunc("/student/{studentId}", controllers.CreateStudents).Methods("DELETE")
}
