package usersession

import (
	"net/http"

	"github.com/school-sys-rest-api/services/routeservice"
)

func LoginRouter() {
	router := routeservice.Routes.PathPrefix("/login").Subrouter()
	router.HandleFunc("/", LoginHandler).Methods("POST", http.MethodOptions)
}
