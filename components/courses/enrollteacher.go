package courses

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func GetEnrollCoursesTeachersHandler(w http.ResponseWriter, r *http.Request) {
	var enrollCourses []utils.TeacherEnrollCourses

	response := &httpop.Response{}

	res := db.DB.Find(&enrollCourses)

	if res.Error != nil || res.RowsAffected < 1 {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, "enroll courses not found")
	} else {
		for counter := range enrollCourses {
			db.DB.Model(&enrollCourses[counter]).Association("Teacher").Find(&enrollCourses[counter].Teacher)
			db.DB.Model(&enrollCourses[counter]).Association("Enrollment").Find(&enrollCourses[counter].Enrollment)
		}
		response.GenerateOkResponse(&enrollCourses, "Ok request")
	}

	response.SendResponse(w)
}

func GetEnrollCoursesTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var enrollCourse utils.TeacherEnrollCourses
	params := mux.Vars(r)

	response := &httpop.Response{}

	if response.ValidateError(db.DB.First(&enrollCourse, params["id"]), w, "enroll course not found") {
		db.DB.Model(&enrollCourse).Association("Teacher").Find(&enrollCourse.Teacher)
		db.DB.Model(&enrollCourse).Association("Enrollment").Find(&enrollCourse.Enrollment)
		response.GenerateOkResponse(&enrollCourse, "Ok request")
	}

	response.SendResponse(w)
}

func PostEnrollCoursesTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var course utils.TeacherEnrollCourses

	json.NewDecoder(r.Body).Decode(&course)

	res := db.DB.Create(&course)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&course)
}

func UpdateEnrollCoursesTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var course, newCourse utils.TeacherEnrollCourses
	params := mux.Vars(r)

	response := &httpop.Response{}

	json.NewDecoder(r.Body).Decode(&newCourse)

	if response.ValidateError(db.DB.Find(&course, params["id"]), w, "enroll course not found") {
		if response.ValidateError(db.DB.Model(&course).Updates(newCourse), w, "update error") {
			response.GenerateOkResponse(&course, "Ok request")
		}
	}
	response.SendResponse(w)
}

func DeleteEnrollCoursesTeacherHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete student"))
}
