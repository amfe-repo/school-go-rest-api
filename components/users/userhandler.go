package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func BuildUser(user *utils.Users) {
	db.DB.Model(&user).Association("AdministrativeRole").Find(&user.AdministrativeRole)
	db.DB.Model(&user.AdministrativeRole).Association("PermissionGroup").Find(&user.AdministrativeRole.PermissionGroup)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []utils.Users

	var response *httpop.Response = new(httpop.Response)

	loggedUser := r.Context().Value("user").(utils.Users)

	if authorizationUsers(loggedUser, &users, response, w, "") {
		response.GenerateOkResponse(&users, "Ok request")
	}

	response.SendResponse(w)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user utils.Users
	params := mux.Vars(r)

	var response *httpop.Response = new(httpop.Response)

	loggedUser := r.Context().Value("user").(utils.Users)

	if authorizationUsers(loggedUser, &user, response, w, params["id"]) {
		response.GenerateOkResponse(&user, "Ok request")
	}

	//fmt.Println(num)
	response.SendResponse(w)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	response := &httpop.Response{}
	loggedUser := r.Context().Value("user").(utils.Users)

	if !PostPutAuth(loggedUser, VerifyCreateUsersPermission, response, w) {
		return
	}

	var user utils.Users
	json.NewDecoder(r.Body).Decode(&user)

	res := db.DB.Create(&user)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, "user not inserted")
	}

	response.SendResponse(w)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	response := &httpop.Response{}
	loggedUser := r.Context().Value("user").(utils.Users)

	if !VerifyEditUserPermission(loggedUser) {
		response.GenerateErrorAccessDeniedResponse(nil, "access denied")
		response.SendResponse(w)
		return
	}

	var user, newUser utils.Users
	params := mux.Vars(r)

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
