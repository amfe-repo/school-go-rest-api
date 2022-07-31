package requests

import "github.com/school-sys-rest-api/services/routeservice"

func RequestRouter() {
	router := routeservice.Routes.PathPrefix("/request-courses").Subrouter()
	router.HandleFunc("/", GetRequestsHandler).Methods("GET")
	router.HandleFunc("/{id}", GetRequestHandler).Methods("GET")
	router.HandleFunc("/", PostRequestHandler).Methods("POST")
	router.HandleFunc("/{id}", UpdateRequestHandler).Methods("PUT")
	router.HandleFunc("/{id}", DeleteRequestHandler).Methods("DELETE")
}
