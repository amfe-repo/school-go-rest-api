package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
)

func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {
	var teacher []Teachers

	response := &httpop.Response{}

	res := db.DB.Find(&teacher)

	if res.Error != nil || res.RowsAffected < 1 {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, res.Error.Error())
	} else {
		response.GenerateOkResponse(&teacher, "Ok request")
	}

	response.SendResponse(w)
}

func GetTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var teacher Teachers
	params := mux.Vars(r)

	response := &httpop.Response{}

	if response.ValidateError(db.DB.First(&teacher, params["id"]), w, "user not found") {
		db.DB.Model(&teacher).Association("User").Find(&teacher.User)
		db.DB.Model(&teacher.User).Association("AdministrativeRole").Find(&teacher.User.AdministrativeRole)
		db.DB.Model(&teacher.User.AdministrativeRole).Association("PermissionGroup").Find(&teacher.User.AdministrativeRole.PermissionGroup)
		response.GenerateOkResponse(&teacher, "Ok request")
	}

	response.SendResponse(w)
}

func PostTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var teacher Teachers

	json.NewDecoder(r.Body).Decode(&teacher)

	res := db.DB.Create(&teacher)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&teacher)
}

func UpdateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var teacher, newTeacher Teachers
	params := mux.Vars(r)

	response := &httpop.Response{}

	json.NewDecoder(r.Body).Decode(&newTeacher)

	if response.ValidateError(db.DB.Find(&teacher, params["id"]), w, "teacher not found") {
		if response.ValidateError(db.DB.Model(&teacher).Updates(newTeacher), w, "update error") {
			response.GenerateOkResponse(&teacher, "Ok request")
		}
	}
	response.SendResponse(w)
}

func PostUserTeacherHandler(w http.ResponseWriter, r *http.Request) {

	var teacher Teachers
	user, err := createUser(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewDecoder(r.Body).Decode(&teacher)

	teacher.IdUser = int(user.ID)

	res := db.DB.Create(&teacher)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteTeacherHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete teacher"))
}
