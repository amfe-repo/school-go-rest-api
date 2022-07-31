package usersession

import "github.com/school-sys-rest-api/services/routeservice"

func LoginRouter() {
	router := routeservice.Routes.PathPrefix("/login").Subrouter()
	router.HandleFunc("/", LoginHandler).Methods("POST")
}
