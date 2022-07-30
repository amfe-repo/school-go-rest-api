package roles

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
)

func GetPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	var pg []PermissionGroups

	response := &httpop.Response{}

	res := db.DB.Find(&pg)

	if res.Error != nil || res.RowsAffected < 1 {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, res.Error.Error())
	} else {
		response.GenerateOkResponse(&pg, "Ok request")
	}

	response.SendResponse(w)
}

func GetPermissionHandler(w http.ResponseWriter, r *http.Request) {
	var pg PermissionGroups
	params := mux.Vars(r)

	response := &httpop.Response{}

	if response.ValidateError(db.DB.First(&pg, params["id"]), w, "user not found") {
		response.GenerateOkResponse(&pg, "Ok request")
	}

	response.SendResponse(w)
}

func PostPermissionHandler(w http.ResponseWriter, r *http.Request) {
	var pg PermissionGroups

	json.NewDecoder(r.Body).Decode(&pg)

	res := db.DB.Create(&pg)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&pg)
}

func UpdatePermissionHandler(w http.ResponseWriter, r *http.Request) {
	var pg, newpg PermissionGroups
	params := mux.Vars(r)

	response := &httpop.Response{}

	json.NewDecoder(r.Body).Decode(&newpg)

	if response.ValidateError(db.DB.Find(&pg, params["id"]), w, "permission group not found") {
		if response.ValidateError(db.DB.Model(&pg).Updates(newpg), w, "update error") {
			response.GenerateOkResponse(&pg, "Ok request")
		}
	}
	response.SendResponse(w)
}

func DeletePermissionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete student"))
}
