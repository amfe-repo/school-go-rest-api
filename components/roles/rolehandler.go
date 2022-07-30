package roles

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
)

func GetRolesHandler(w http.ResponseWriter, r *http.Request) {
	var role []AdministrativeRoles

	response := &httpop.Response{}

	res := db.DB.Find(&role)

	if res.Error != nil || res.RowsAffected < 1 {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, res.Error.Error())
	} else {
		response.GenerateOkResponse(&role, "Ok request")
	}

	response.SendResponse(w)
}

func GetRoleHandler(w http.ResponseWriter, r *http.Request) {
	var role AdministrativeRoles
	params := mux.Vars(r)

	response := &httpop.Response{}

	if response.ValidateError(db.DB.First(&role, params["id"]), w, "user not found") {
		db.DB.Model(&role).Association("PermissionGroup").Find(&role.PermissionGroup)
		response.GenerateOkResponse(&role, "Ok request")
	}

	response.SendResponse(w)
}

func PostRoleHandler(w http.ResponseWriter, r *http.Request) {
	var role AdministrativeRoles

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
	var role, newRole AdministrativeRoles
	params := mux.Vars(r)

	response := &httpop.Response{}

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
