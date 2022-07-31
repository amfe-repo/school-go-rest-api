package usersession

import (
	"encoding/json"
	"net/http"

	"github.com/school-sys-rest-api/components/users"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	resp := new(httpop.Response)
	var user utils.Users

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error data"))
		return
	}

	if user, ok := userExist(user.Email, user.Password); ok {
		resp.GenerateOkResponse(user, "ok request")
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
