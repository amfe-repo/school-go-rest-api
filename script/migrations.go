package script

import (
	"fmt"

	"github.com/school-sys-rest-api/components/courses"
	"github.com/school-sys-rest-api/components/requests"
	"github.com/school-sys-rest-api/components/roles"
	"github.com/school-sys-rest-api/components/users"
	"github.com/school-sys-rest-api/services/db"
)

type Table interface{}

var (
	user    users.Users
	teacher users.Teachers
	student users.Students

	admin roles.AdministrativeRoles
	per   roles.PermissionGroups

	course    courses.Courses
	enrr      courses.EnrollCourses
	stcourse  courses.StudentEnrollCourses
	tchcourse courses.TeacherEnrollCourses

	reqst requests.StudentCourseRequests
)

func migrateTables(tables []Table) {
	for _, table := range tables {
		err := db.DB.AutoMigrate(table)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func CompleteMigrateTables() {
	migrateTables([]Table{&user, &teacher, &student, &admin, &per, &course, &enrr, &stcourse, &tchcourse, &reqst})
}
