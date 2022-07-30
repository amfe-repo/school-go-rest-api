package users

import "github.com/school-sys-rest-api/services/routeservice"

func UserRouter() {
	router := routeservice.Routes.PathPrefix("/users").Subrouter()
	router.HandleFunc("/", GetUsersHandler).Methods("GET")
	router.HandleFunc("/{id}", GetUserHandler).Methods("GET")
	router.HandleFunc("/", PostUserHandler).Methods("POST")
	router.HandleFunc("/{id}", UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/{id}", DeleteUserHandler).Methods("DELETE")
}

func StudentRouter() {
	router := routeservice.Routes.PathPrefix("/students").Subrouter()
	router.HandleFunc("/", GetStudentsHandler).Methods("GET")
	router.HandleFunc("/{id}", GetStudentHandler).Methods("GET")
	router.HandleFunc("/create-user", PostUserStudentHandler).Methods("POST")
	router.HandleFunc("/", PostStudentHandler).Methods("POST")
	router.HandleFunc("/{id}", UpdateStudentHandler).Methods("PUT")
	router.HandleFunc("/{id}", DeleteStudentsHandler).Methods("DELETE")
}

func TeacherRouter() {
	router := routeservice.Routes.PathPrefix("/teachers").Subrouter()
	router.HandleFunc("/", GetTeachersHandler).Methods("GET")
	router.HandleFunc("/{id}", GetTeacherHandler).Methods("GET")
	router.HandleFunc("/create-user", PostUserTeacherHandler).Methods("POST")
	router.HandleFunc("/", PostTeacherHandler).Methods("POST")
	router.HandleFunc("/{id}", UpdateTeacherHandler).Methods("PUT")
	router.HandleFunc("/{id}", DeleteTeacherHandler).Methods("DELETE")
}
