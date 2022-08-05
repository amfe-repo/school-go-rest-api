package usersession

import (
	"encoding/json"
	"net/http"

	"github.com/school-sys-rest-api/components/users"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

type UserWithCondition struct {
	utils.Users
	Student interface{}
	Teacher interface{}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userCompose UserWithCondition
	resp := new(httpop.Response)
	var user utils.Users

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error data"))
		return
	}

	if user, ok := userExist(user.Email, user.Password); ok {

		student, _ := studentExist(int(user.ID))
		teacher, _ := teacherExist(int(user.ID))

		userCompose = UserWithCondition{
			user,
			student,
			teacher,
		}
		resp.GenerateOkResponse(userCompose, "ok request")
	} else {
		w.WriteHeader(http.StatusNotFound)
		resp.GenerateErrorAccessDeniedResponse(nil, "user incorrect")
	}

	resp.SendResponse(w)
}

func userExist(email string, password string) (utils.Users, bool) {
	var user utils.Users

	result := db.DB.Where("email = ? AND password = ?", email, password).First(&user)

	if result.Error != nil || result.RowsAffected < 1 {
		return user, false
	}

	users.BuildUser(&user)

	return user, true
}

func studentExist(id int) (utils.Students, bool) {
	var user utils.Students

	result := db.DB.Where("id_user = ?", id).First(&user)

	if result.Error != nil || result.RowsAffected < 1 {
		return user, false
	}

	//users.BuildUser(&user)

	return user, true
}

func teacherExist(id int) (utils.Teachers, bool) {
	var user utils.Teachers

	result := db.DB.Where("id_user = ?", id).First(&user)

	if result.Error != nil || result.RowsAffected < 1 {
		return user, false
	}

	//users.BuildUser(&user)

	return user, true
}
