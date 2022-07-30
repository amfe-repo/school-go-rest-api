package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
)

func buildUser(user *Users) {
	db.DB.Model(&user).Association("AdministrativeRole").Find(&user.AdministrativeRole)
	db.DB.Model(&user.AdministrativeRole).Association("PermissionGroup").Find(&user.AdministrativeRole.PermissionGroup)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []Users

	var resp *httpop.Response = new(httpop.Response)

	if resp.ValidateError(db.DB.Find(&users), w, "not users") {
		resp.GenerateOkResponse(&users, "Ok request")
	}

	resp.SendResponse(w)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user Users
	params := mux.Vars(r)

	response := &httpop.Response{}

	res := db.DB.First(&user, params["id"])

	if res.Error != nil || res.RowsAffected < 1 {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, "error")
	} else {
		buildUser(&user)
		response.GenerateOkResponse(&user, "Ok request")
	}

	response.SendResponse(w)
}

//USER NOT STUDENT OR TEACHER
//TEACHERS
//STUDENTS
// USER STUDENT
// USER TEACHER

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user Users
	json.NewDecoder(r.Body).Decode(&user)

	res := db.DB.Create(&user)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user, newUser Users
	params := mux.Vars(r)

	response := &httpop.Response{}

	json.NewDecoder(r.Body).Decode(&newUser)

	if response.ValidateError(db.DB.Find(&user, params["id"]), w, "user not found") {
		if response.ValidateError(db.DB.Model(&user).Updates(newUser), w, "update error") {
			response.GenerateOkResponse(&user, "Ok request")
		}
	}
	response.SendResponse(w)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
