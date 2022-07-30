package requests

import (
	"github.com/school-sys-rest-api/components/courses"
	"github.com/school-sys-rest-api/components/users"
	"gorm.io/gorm"
)

type StudentCourseRequests struct {
	gorm.Model
	IdStudent int             `gorm:"not null;type:integer"`
	Student   users.Students  `gorm:"foreignKey:IdStudent"`
	IdCourse  int             `gorm:"not null;type:integer"`
	Course    courses.Courses `gorm:"foreignKey:IdCourse"`
	SendDate  string          `gorm:"not null;type:date"`
	Status    int             `gorm:"type:bit"`
}
