package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/school-sys-rest-api/components/users"
	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI == "/login/" || r.RequestURI == "/students/create-user" || r.RequestURI == "/teachers/create-user" {
			next.ServeHTTP(w, r)
			return
		}

		resp := new(httpop.Response)

		var user utils.Users

		data := r.Header.Get("Session-user-data")

		err := json.Unmarshal([]byte(data), &user)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error data"))
			return
		}

		if user, ok := userExist(user.Email, user.Password); ok {
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusNotFound)
			resp.GenerateErrorAccessDeniedResponse(nil, "user incorrect")
			resp.SendResponse(w)
		}

	})
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
