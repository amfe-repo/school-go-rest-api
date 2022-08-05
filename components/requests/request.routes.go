package requests

import (
	"net/http"

	"github.com/school-sys-rest-api/services/routeservice"
)

func RequestRouter() {
	router := routeservice.Routes.PathPrefix("/request-courses").Subrouter()
	router.HandleFunc("/", GetRequestsHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/{id}", GetRequestHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/", PostRequestHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/{id}", UpdateRequestHandler).Methods("PUT", http.MethodOptions)
	router.HandleFunc("/{id}", DeleteRequestHandler).Methods("DELETE", http.MethodOptions)
}
