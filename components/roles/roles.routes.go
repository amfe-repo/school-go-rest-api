package roles

import (
	"net/http"

	"github.com/school-sys-rest-api/services/routeservice"
)

func RolesRouter() {
	router := routeservice.Routes.PathPrefix("/roles").Subrouter()
	router.HandleFunc("/", GetRolesHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/{id}", GetRoleHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/", PostRoleHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/{id}", UpdateRoleHandler).Methods("PUT", http.MethodOptions)
	router.HandleFunc("/{id}", DeleteRoleHandler).Methods("DELETE", http.MethodOptions)
}

func PermissionsRouter() {
	router := routeservice.Routes.PathPrefix("/permissions").Subrouter()
	router.HandleFunc("/", GetPermissionsHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/{id}", GetPermissionHandler).Methods("GET", http.MethodOptions)
	router.HandleFunc("/", PostPermissionHandler).Methods("POST", http.MethodOptions)
	router.HandleFunc("/{id}", UpdatePermissionHandler).Methods("PUT", http.MethodOptions)
	router.HandleFunc("/{id}", DeletePermissionHandler).Methods("DELETE", http.MethodOptions)
}
