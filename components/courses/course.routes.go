package courses

import "github.com/school-sys-rest-api/services/routeservice"

func CoursesRouter() {
	router := routeservice.Routes.PathPrefix("/courses").Subrouter()
	router.HandleFunc("/", GetCoursesHandler).Methods("GET")
	router.HandleFunc("/{id}", GetCourseHandler).Methods("GET")
	router.HandleFunc("/", PostCourseHandler).Methods("POST")
	router.HandleFunc("/{id}", UpdateCourseHandler).Methods("PUT")
	router.HandleFunc("/{id}", DeleteCourseHandler).Methods("DELETE")
}

func EnrollCoursesRouter() {
	router := routeservice.Routes.PathPrefix("/enroll-courses").Subrouter()
	router.HandleFunc("/", GetEnrollCoursesHandler).Methods("GET")
	router.HandleFunc("/{id}", GetEnrollCourseHandler).Methods("GET")
	router.HandleFunc("/", PostEnrollCourseHandler).Methods("POST")
	router.HandleFunc("/{id}", UpdateEnrollCourseHandler).Methods("PUT")
	router.HandleFunc("/{id}", DeleteEnrollCourseHandler).Methods("DELETE")
}

func EnrollStudentCoursesRouter() {
	router := routeservice.Routes.PathPrefix("/enroll-courses-students").Subrouter()
	router.HandleFunc("/", GetEnrollCoursesStudentsHandler).Methods("GET")
	router.HandleFunc("/{id}", GetEnrollCoursesStudentHandler).Methods("GET")
	router.HandleFunc("/", PostEnrollCoursesStudentHandler).Methods("POST")
	router.HandleFunc("/{id}", UpdateEnrollCoursesStudentHandler).Methods("PUT")
	router.HandleFunc("/{id}", DeleteEnrollCoursesStudentHandler).Methods("DELETE")
}

func EnrollTeacherCoursesRouter() {
	router := routeservice.Routes.PathPrefix("/enroll-courses-teachers").Subrouter()
	router.HandleFunc("/", GetEnrollCoursesTeachersHandler).Methods("GET")
	router.HandleFunc("/{id}", GetEnrollCoursesTeacherHandler).Methods("GET")
	router.HandleFunc("/", PostEnrollCoursesTeacherHandler).Methods("POST")
	router.HandleFunc("/{id}", UpdateEnrollCoursesTeacherHandler).Methods("PUT")
	router.HandleFunc("/{id}", DeleteEnrollCoursesTeacherHandler).Methods("DELETE")
}
