package roles

import "github.com/school-sys-rest-api/services/routeservice"

func RolesRouter() {
	router := routeservice.Routes.PathPrefix("/roles").Subrouter()
	router.HandleFunc("/", GetRolesHandler).Methods("GET")
	router.HandleFunc("/{id}", GetRoleHandler).Methods("GET")
	router.HandleFunc("/", PostRoleHandler).Methods("POST")
	router.HandleFunc("/{id}", UpdateRoleHandler).Methods("PUT")
	router.HandleFunc("/{id}", DeleteRoleHandler).Methods("DELETE")
}

func PermissionsRouter() {
	router := routeservice.Routes.PathPrefix("/permissions").Subrouter()
	router.HandleFunc("/", GetPermissionsHandler).Methods("GET")
	router.HandleFunc("/{id}", GetPermissionHandler).Methods("GET")
	router.HandleFunc("/", PostPermissionHandler).Methods("POST")
	router.HandleFunc("/{id}", UpdatePermissionHandler).Methods("PUT")
	router.HandleFunc("/{id}", DeletePermissionHandler).Methods("DELETE")
}
