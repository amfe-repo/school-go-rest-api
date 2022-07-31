package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	var students []utils.Students

	response := &httpop.Response{}

	loggedUser := r.Context().Value("user").(utils.Users)

	if authorizationStudents(loggedUser, &students, response, w, "") {
		response.GenerateOkResponse(&students, "Ok request")
	}

	response.SendResponse(w)
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student utils.Students
	params := mux.Vars(r)

	response := &httpop.Response{}

	loggedUser := r.Context().Value("user").(utils.Users)

	if authorizationStudents(loggedUser, &student, response, w, params["id"]) {
		response.GenerateOkResponse(&student, "Ok request")
	}

	response.SendResponse(w)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	response := &httpop.Response{}
	loggedUser := r.Context().Value("user").(utils.Users)

	if !PostPutAuth(loggedUser, VerifyCreateUsersPermission, response, w) {
		return
	}

	var student utils.Students
	json.NewDecoder(r.Body).Decode(&student)

	res := db.DB.Create(&student)

	if res.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.GenerateErrorResponse(nil, "student not inserted")
	}

	response.SendResponse(w)
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student, newStudent utils.Students
	params := mux.Vars(r)

	response := &httpop.Response{}

	loggedUser := r.Context().Value("user").(utils.Users)

	if !PostPutAuth(loggedUser, VerifyEditUserPermission, response, w) {
		return
	}

	json.NewDecoder(r.Body).Decode(&newStudent)

	if response.ValidateError(db.DB.Find(&student, params["id"]), w, "student not found") {
		if response.ValidateError(db.DB.Model(&student).Updates(newStudent), w, "update error") {
			response.GenerateOkResponse(&student, "Ok request")
		}
	}
	response.SendResponse(w)
}

func createUser(r *http.Request) (utils.Users, error) {
	var user utils.Users
	json.NewDecoder(r.Body).Decode(&user)

	user.IdRole = 1

	res := db.DB.Create(&user)

	if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}

func PostUserStudentHandler(w http.ResponseWriter, r *http.Request) {

	/*if !PostPutAuth(loggedUser, VerifyCreateUsersPermission, response, w) {
		return
	}*/

	var student utils.Students
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
