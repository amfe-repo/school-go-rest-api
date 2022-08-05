package users

import (
	"net/http"

	"github.com/school-sys-rest-api/services/routeservice"
)

func UserRouter() {
	router := routeservice.Routes.PathPrefix("/users").Subrouter()
	router.HandleFunc("/", GetUsersHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/{id}", GetUserHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/", PostUserHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/{id}", UpdateUserHandler).Methods("PUT", http.MethodOptions)
	router.HandleFunc("/{id}", DeleteUserHandler).Methods("DELETE", http.MethodOptions)
}

func StudentRouter() {
	router := routeservice.Routes.PathPrefix("/students").Subrouter()
	router.HandleFunc("/", GetStudentsHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/{id}", GetStudentHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/create-user", PostUserStudentHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/", PostStudentHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/{id}", UpdateStudentHandler).Methods("PUT", http.MethodOptions)
	router.HandleFunc("/{id}", DeleteStudentsHandler).Methods("DELETE", http.MethodOptions)
}

func TeacherRouter() {
	router := routeservice.Routes.PathPrefix("/teachers").Subrouter()
	router.HandleFunc("/", GetTeachersHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/{id}", GetTeacherHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/create-user", PostUserTeacherHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/", PostTeacherHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/{id}", UpdateTeacherHandler).Methods("PUT", http.MethodOptions)
	router.HandleFunc("/{id}", DeleteTeacherHandler).Methods("DELETE", http.MethodOptions)
}
