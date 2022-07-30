package users

import (
	"github.com/school-sys-rest-api/components/roles"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name               string                    `gorm:"not null;type:varchar(30)" json:"name"`
	Lastname           string                    `gorm:"not null;type:varchar(30)" json:"lastname"`
	Username           string                    `gorm:"not null;type:varchar(30)" json:"username"`
	Password           string                    `gorm:"not null;type:varchar(30)" json:"password"`
	Email              string                    `gorm:"not null;type:varchar(30);unique_index" json:"email"`
	IdRole             int                       `gorm:"not null;type:integer" json:"id_role"`
	AdministrativeRole roles.AdministrativeRoles `gorm:"foreignKey:IdRole" json:"administrative_role"`
	Status             string                    `gorm:"type:bit" json:"status"`
}

type Teachers struct {
	gorm.Model
	RegistrationNumber string  `gorm:"not null;type:varchar(15);unique_index" json:"registration_number"`
	Salary             float32 `gorm:"not null;type:real" json:"name" json:"salary"`
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
