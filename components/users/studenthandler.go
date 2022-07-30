package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
)

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	var students []Students

	response := &httpop.Response{}

	res := db.DB.Find(&students)

	if res.Error != nil || res.RowsAffected < 1 {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, res.Error.Error())
	} else {
		response.GenerateOkResponse(&students, "Ok request")
	}

	response.SendResponse(w)
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Students
	params := mux.Vars(r)

	response := &httpop.Response{}

	if response.ValidateError(db.DB.First(&student, params["id"]), w, "user not found") {
		db.DB.Model(&student).Association("User").Find(&student.User)
		db.DB.Model(&student.User).Association("AdministrativeRole").Find(&student.User.AdministrativeRole)
		db.DB.Model(&student.User.AdministrativeRole).Association("PermissionGroup").Find(&student.User.AdministrativeRole.PermissionGroup)
		response.GenerateOkResponse(&student, "Ok request")
	}

	response.SendResponse(w)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Students

	json.NewDecoder(r.Body).Decode(&student)

	res := db.DB.Create(&student)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&student)
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student, newStudent Students
	params := mux.Vars(r)

	response := &httpop.Response{}

	json.NewDecoder(r.Body).Decode(&newStudent)

	if response.ValidateError(db.DB.Find(&student, params["id"]), w, "student not found") {
		if response.ValidateError(db.DB.Model(&student).Updates(newStudent), w, "update error") {
			response.GenerateOkResponse(&student, "Ok request")
		}
	}
	response.SendResponse(w)
}

func createUser(r *http.Request) (Users, error) {
	var user Users
	json.NewDecoder(r.Body).Decode(&user)

	res := db.DB.Create(&user)

	if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}

func PostUserStudentHandler(w http.ResponseWriter, r *http.Request) {

	var student Students
	user, err := createUser(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewDecoder(r.Body).Decode(&student)

	student.IdUser = int(user.ID)

	res := db.DB.Create(&student)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(res.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteStudentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete student"))
}
