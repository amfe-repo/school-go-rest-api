package roles

import (
	"gorm.io/gorm"
)

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
