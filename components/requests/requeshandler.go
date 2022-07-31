package requests

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func GetRequestsHandler(w http.ResponseWriter, r *http.Request) {
	var course []utils.StudentCourseRequests

	response := &httpop.Response{}

	res := db.DB.Find(&course)

	if res.Error != nil || res.RowsAffected < 1 {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, "requests not found")
	} else {
		for counter := range course {
			db.DB.Model(&course[counter]).Association("Course").Find(&course[counter].Course)
			db.DB.Model(&course[counter]).Association("Student").Find(&course[counter].Student)
		}
		response.GenerateOkResponse(&course, "Ok request")
	}

	response.SendResponse(w)
}

func GetRequestHandler(w http.ResponseWriter, r *http.Request) {
	var course utils.StudentCourseRequests
	params := mux.Vars(r)

	response := &httpop.Response{}

	if response.ValidateError(db.DB.First(&course, params["id"]), w, "request not found") {
		db.DB.Model(&course).Association("Course").Find(&course.Course)
		db.DB.Model(&course).Association("Student").Find(&course.Student)
		response.GenerateOkResponse(&course, "Ok request")
	}

	response.SendResponse(w)
}

func PostRequestHandler(w http.ResponseWriter, r *http.Request) {
	var course utils.StudentCourseRequests

	json.NewDecoder(r.Body).Decode(&course)

	res := db.DB.Create(&course)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&course)
}

func UpdateRequestHandler(w http.ResponseWriter, r *http.Request) {
	var course, newCourse utils.StudentCourseRequests
	params := mux.Vars(r)

	response := &httpop.Response{}

	json.NewDecoder(r.Body).Decode(&newCourse)

	if response.ValidateError(db.DB.Find(&course, params["id"]), w, "request not found") {
		if response.ValidateError(db.DB.Model(&course).Updates(newCourse), w, "update error") {
			response.GenerateOkResponse(&course, "Ok request")
		}
	}
	response.SendResponse(w)
}

func DeleteRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete student"))
}
