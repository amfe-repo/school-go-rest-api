package utils

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name               string              `gorm:"not null;type:varchar(30)" json:"name"`
	Lastname           string              `gorm:"not null;type:varchar(30)" json:"lastname"`
	Username           string              `gorm:"not null;type:varchar(30)" json:"username"`
	Password           string              `gorm:"not null;type:varchar(30)" json:"password"`
	Email              string              `gorm:"not null;type:varchar(30);unique_index" json:"email"`
	IdRole             int                 `gorm:"not null;type:integer" json:"id_role"`
	AdministrativeRole AdministrativeRoles `gorm:"foreignKey:IdRole" json:"administrative_role"`
	Status             string              `gorm:"type:bit" json:"status"`
}

type Teachers struct {
	gorm.Model
	RegistrationNumber string  `gorm:"not null;type:varchar(15);unique_index" json:"registration_number"`
	Salary             float32 `gorm:"not null;type:real" json:"salary"`
	IdUser             int     `gorm:"type:integer" json:"id_user"`
	User               Users   `gorm:"foreignKey:IdUser" json:"user"`
	Experience         string  `gorm:"type:text" json:"experience"`
}

type Students struct {
	gorm.Model
	RegistrationNumber string  `gorm:"not null;type:varchar(15);unique_index" json:"registration_number"`
	AverageScore       float32 `gorm:"not null;type:real" json:"average_score"`
	IdUser             int     `gorm:"type:integer" json:"id_user"`
	User               Users   `gorm:"foreignKey:IdUser" json:"user"`
	Experience         string  `gorm:"type:text" json:"experience"`
}

type AdministrativeRoles struct {
	gorm.Model
	Name              string           `gorm:"not null;type:varchar(25)"`
	IdPermissionGroup int              `gorm:"not null;type:integer"`
	PermissionGroup   PermissionGroups `gorm:"foreignKey:IdPermissionGroup"`
}

type PermissionGroups struct {
	gorm.Model
	GroupName                       string `gorm:"not null;type:varchar(40)"`
	CreateCoursePermission          string `gorm:"not null;type:bit"`
	CreateRolesPermission           string `gorm:"not null;type:bit"`
	CreateUsersPermission           string `gorm:"not null;type:bit"`
	EnrollUserPermission            string `gorm:"not null;type:bit"`
	EditUserPermission              string `gorm:"not null;type:bit"`
	DeleteUserPermission            string `gorm:"not null;type:bit"`
	AcceptAcademicRequestPermission string `gorm:"not null;type:bit"`
}

type StudentCourseRequests struct {
	gorm.Model
	IdStudent int      `gorm:"not null;type:integer"`
	Student   Students `gorm:"foreignKey:IdStudent"`
	IdCourse  int      `gorm:"not null;type:integer"`
	Course    Courses  `gorm:"foreignKey:IdCourse"`
	SendDate  string   `gorm:"not null;type:date"`
	Status    int      `gorm:"type:bit"`
}

type TeacherEnrollCourses struct {
	gorm.Model
	IdTeacher    int           `gorm:"not null;type:integer"`
	Teacher      Teachers      `gorm:"foreignKey:IdTeacher"`
	IdEnrollment int           `gorm:"not null;type:integer"`
	Enrollment   EnrollCourses `gorm:"foreignKey:IdEnrollment"`
	Status       string        `gorm:"type:bit"`
}

type StudentEnrollCourses struct {
	gorm.Model
	IdStudent    int           `gorm:"not null;type:integer"`
	Student      Students      `gorm:"foreignKey:IdStudent"`
	IdEnrollment int           `gorm:"not null;type:integer"`
	Enrollment   EnrollCourses `gorm:"foreignKey:IdEnrollment"`
	Status       string        `gorm:"type:bit"`
}

type EnrollCourses struct {
	gorm.Model
	ID       int     `gorm:"not null;type:integer;unique_index"`
	IdCourse int     `gorm:"not null;type:integer;references:IdCourse"`
	Course   Courses `gorm:"foreignKey:IdCourse"`
	Status   string  `gorm:"type:bit"`
}

type Courses struct {
	gorm.Model
	Name       string `gorm:"not null;type:varchar(40)"`
	BeginDate  string `gorm:"not null;type:date"`
	FinishDate string `gorm:"not null;type:date"`
	Stock      int    `gorm:"not null;type:integer"`
}
