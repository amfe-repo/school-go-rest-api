package utils

import (
	"github.com/school-sys-rest-api/services/db"
)

func userExist(email string, password string) (Users, bool) {
	var user Users

	result := db.DB.Where("email = ? AND password = ?", email, password).First(&user)

	if result.Error != nil || result.RowsAffected < 1 {
		return user, false
	}

	return user, true
}
