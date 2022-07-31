package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {
	var teacher []utils.Teachers

	response := &httpop.Response{}

	loggedUser := r.Context().Value("user").(utils.Users)

	if authorizationTeacher(loggedUser, &teacher, response, w, "") {
		response.GenerateOkResponse(&teacher, "Ok request")
	}

	response.SendResponse(w)
}

func GetTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var teacher utils.Teachers
	params := mux.Vars(r)

	response := &httpop.Response{}

	loggedUser := r.Context().Value("user").(utils.Users)

	if authorizationTeacher(loggedUser, &teacher, response, w, params["id"]) {
		response.GenerateOkResponse(&teacher, "Ok request")
	}

	response.SendResponse(w)
}

func PostTeacherHandler(w http.ResponseWriter, r *http.Request) {
	response := &httpop.Response{}
	loggedUser := r.Context().Value("user").(utils.Users)

	if !PostPutAuth(loggedUser, VerifyCreateUsersPermission, response, w) {
		return
	}

	var teacher utils.Teachers

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
	response := &httpop.Response{}
	loggedUser := r.Context().Value("user").(utils.Users)

	if !PostPutAuth(loggedUser, VerifyEditUserPermission, response, w) {
		return
	}

	var teacher, newTeacher utils.Teachers
	params := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&newTeacher)

	if response.ValidateError(db.DB.Find(&teacher, params["id"]), w, "teacher not found") {
		if response.ValidateError(db.DB.Model(&teacher).Updates(newTeacher), w, "update error") {
			response.GenerateOkResponse(&teacher, "Ok request")
		}
	}
	response.SendResponse(w)
}

func PostUserTeacherHandler(w http.ResponseWriter, r *http.Request) {

	/*
		if !PostPutAuth(loggedUser, VerifyCreateUsersPermission, response, w) {
			return
		}*/

	var teacher utils.Teachers
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
