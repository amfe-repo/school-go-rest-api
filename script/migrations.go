package script

import (
	"fmt"

	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/utils"
)

type Table interface{}

var (
	user    utils.Users
	teacher utils.Teachers
	student utils.Students

	admin utils.AdministrativeRoles
	per   utils.PermissionGroups

	course    utils.Courses
	enrr      utils.EnrollCourses
	stcourse  utils.StudentEnrollCourses
	tchcourse utils.TeacherEnrollCourses

	reqst utils.StudentCourseRequests
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
