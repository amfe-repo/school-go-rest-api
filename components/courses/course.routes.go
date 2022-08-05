package courses

import (
	"net/http"

	"github.com/school-sys-rest-api/services/routeservice"
)

func CoursesRouter() {
	router := routeservice.Routes.PathPrefix("/courses").Subrouter()
	router.HandleFunc("/", GetCoursesHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/{id}", GetCourseHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/", PostCourseHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/{id}", UpdateCourseHandler).Methods("PUT", http.MethodOptions)
	router.HandleFunc("/{id}", DeleteCourseHandler).Methods("DELETE", http.MethodOptions)
}

func EnrollCoursesRouter() {
	router := routeservice.Routes.PathPrefix("/enroll-courses").Subrouter()
	router.HandleFunc("/", GetEnrollCoursesHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/{id}", GetEnrollCourseHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/", PostEnrollCourseHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/{id}", UpdateEnrollCourseHandler).Methods("PUT", http.MethodOptions)
	router.HandleFunc("/{id}", DeleteEnrollCourseHandler).Methods("DELETE", http.MethodOptions)
}

func EnrollStudentCoursesRouter() {
	router := routeservice.Routes.PathPrefix("/enroll-courses-students").Subrouter()
	router.HandleFunc("/", GetEnrollCoursesStudentsHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/{id}", GetEnrollCoursesStudentHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/", PostEnrollCoursesStudentHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/{id}", UpdateEnrollCoursesStudentHandler).Methods("PUT", http.MethodOptions)
	router.HandleFunc("/{id}", DeleteEnrollCoursesStudentHandler).Methods("DELETE", http.MethodOptions)
}

func EnrollTeacherCoursesRouter() {
	router := routeservice.Routes.PathPrefix("/enroll-courses-teachers").Subrouter()
	router.HandleFunc("/", GetEnrollCoursesTeachersHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/{id}", GetEnrollCoursesTeacherHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/", PostEnrollCoursesTeacherHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/{id}", UpdateEnrollCoursesTeacherHandler).Methods("PUT", http.MethodOptions)
	router.HandleFunc("/{id}", DeleteEnrollCoursesTeacherHandler).Methods("DELETE", http.MethodOptions)
}
