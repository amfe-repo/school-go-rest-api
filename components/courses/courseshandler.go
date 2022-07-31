package courses

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func GetCoursesHandler(w http.ResponseWriter, r *http.Request) {
	var course []utils.Courses

	response := &httpop.Response{}

	res := db.DB.Find(&course)

	if res.Error != nil || res.RowsAffected < 1 {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, "courses not found")
	} else {
		response.GenerateOkResponse(&course, "Ok request")
	}

	response.SendResponse(w)
}

func GetCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course utils.Courses
	params := mux.Vars(r)

	response := &httpop.Response{}

	if response.ValidateError(db.DB.First(&course, params["id"]), w, "user not found") {
		response.GenerateOkResponse(&course, "Ok request")
	}

	response.SendResponse(w)
}

func PostCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course utils.Courses

	json.NewDecoder(r.Body).Decode(&course)

	res := db.DB.Create(&course)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&course)
}

func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course, newCourse utils.Courses
	params := mux.Vars(r)

	response := &httpop.Response{}

	json.NewDecoder(r.Body).Decode(&newCourse)

	if response.ValidateError(db.DB.Find(&course, params["id"]), w, "role not found") {
		if response.ValidateError(db.DB.Model(&course).Updates(newCourse), w, "update error") {
			response.GenerateOkResponse(&course, "Ok request")
		}
	}
	response.SendResponse(w)
}

func DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete student"))
}
