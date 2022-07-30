package routes

import (
	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/components/roles"
	"github.com/school-sys-rest-api/components/users"
	"github.com/school-sys-rest-api/services/routeservice"
)

func AllRoutes() *mux.Router {

	users.UserRouter()
	users.StudentRouter()
	users.TeacherRouter()

	roles.RolesRouter()
	roles.PermissionsRouter()

	return routeservice.Routes
}
