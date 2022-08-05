package routes

import (
	"github.com/gorilla/mux"
	"github.com/school-sys-rest-api/components/courses"
	"github.com/school-sys-rest-api/components/requests"
	"github.com/school-sys-rest-api/components/roles"
	"github.com/school-sys-rest-api/components/users"
	"github.com/school-sys-rest-api/components/usersession"
	"github.com/school-sys-rest-api/services/auth"
	"github.com/school-sys-rest-api/services/routeservice"
)

func AllRoutes() *mux.Router {

	users.UserRouter()
	users.StudentRouter()
	users.TeacherRouter()

	roles.RolesRouter()
	roles.PermissionsRouter()

	courses.CoursesRouter()
	courses.EnrollCoursesRouter()
	courses.EnrollStudentCoursesRouter()
	courses.EnrollTeacherCoursesRouter()

	requests.RequestRouter()

	usersession.LoginRouter()

	routeservice.Routes.Use(auth.AuthenticationMiddleware)
	routeservice.Routes.Use(mux.CORSMethodMiddleware(routeservice.Routes))

	return routeservice.Routes
}
