package roles

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func GetRolesHandler(w http.ResponseWriter, r *http.Request) {
	var role []utils.AdministrativeRoles

	var response *httpop.Response = new(httpop.Response)

	res := db.DB.Find(&role)

	if res.Error != nil || res.RowsAffected < 1 {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, "permissions not found")
	} else {
		for counter := range role {
			db.DB.Model(&role[counter]).Association("PermissionGroup").Find(&role[counter].PermissionGroup)
		}
		response.GenerateOkResponse(&role, "Ok request")
	}

	response.SendResponse(w)
}

func GetRoleHandler(w http.ResponseWriter, r *http.Request) {
	var role utils.AdministrativeRoles
	params := mux.Vars(r)

	response := &httpop.Response{}

	if response.ValidateError(db.DB.First(&role, params["id"]), w, "role not found") {
		response.GenerateOkResponse(&role, "Ok request")
	}

	response.SendResponse(w)
}

func PostRoleHandler(w http.ResponseWriter, r *http.Request) {
	response := &httpop.Response{}
	loggedUser := r.Context().Value("user").(utils.Users)

	if !PostPutAuth(loggedUser, VerifyCreateRolePermission, response, w) {
		return
	}

	var role utils.AdministrativeRoles

	json.NewDecoder(r.Body).Decode(&role)

	res := db.DB.Create(&role)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&role)
}

func UpdateRoleHandler(w http.ResponseWriter, r *http.Request) {

	response := &httpop.Response{}
	loggedUser := r.Context().Value("user").(utils.Users)

	if !PostPutAuth(loggedUser, VerifyCreateRolePermission, response, w) {
		return
	}

	var role, newRole utils.AdministrativeRoles
	params := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&newRole)

	if response.ValidateError(db.DB.Find(&role, params["id"]), w, "role not found") {
		if response.ValidateError(db.DB.Model(&role).Updates(newRole), w, "update error") {
			response.GenerateOkResponse(&role, "Ok request")
		}
	}
	response.SendResponse(w)
}

func DeleteRoleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete student"))
}
