package courses

import (
	"github.com/school-sys-rest-api/components/users"
	"gorm.io/gorm"
)

type TeacherEnrollCourses struct {
	gorm.Model
	IdTeacher    int            `gorm:"not null;type:integer"`
	Teacher      users.Teachers `gorm:"foreignKey:IdTeacher"`
	IdEnrollment int            `gorm:"not null;type:integer"`
	Enrollment   EnrollCourses  `gorm:"foreignKey:IdEnrollment"`
	Status       []byte         `gorm:"type:bit"`
}

type StudentEnrollCourses struct {
	gorm.Model
	IdStudent    int            `gorm:"not null;type:integer"`
	Student      users.Students `gorm:"foreignKey:IdStudent"`
	IdEnrollment int            `gorm:"not null;type:integer"`
	Enrollment   EnrollCourses  `gorm:"foreignKey:IdEnrollment"`
	Status       []byte         `gorm:"type:bit"`
}

type EnrollCourses struct {
	gorm.Model
	ID       int     `gorm:"not null;type:integer;unique_index"`
	IdCourse int     `gorm:"not null;type:integer;references:IdCourse"`
	Course   Courses `gorm:"foreignKey:IdCourse"`
	Status   []byte  `gorm:"type:bit"`
}

type Courses struct {
	gorm.Model
	Name       string `gorm:"not null;type:varchar(40)"`
	BeginDate  string `gorm:"not null;type:date"`
	FinishDate string `gorm:"not null;type:date"`
	Stock      int    `gorm:"not null;type:integer"`
}
