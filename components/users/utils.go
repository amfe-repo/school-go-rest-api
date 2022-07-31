package users

import (
	"net/http"

	"github.com/school-sys-rest-api/services/db"
	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func VerifyCreateUsersPermission(p utils.Users) bool {
	return p.AdministrativeRole.PermissionGroup.CreateUsersPermission == "1"
}

func VerifyEditUserPermission(p utils.Users) bool {
	return p.AdministrativeRole.PermissionGroup.EditUserPermission == "1"
}

func ValidateGetUsersPermission(p utils.Users) bool {
	return VerifyCreateUsersPermission(p) || VerifyEditUserPermission(p)
}

// REFACTOR
func PostPutAuth(loggedUser utils.Users, fn func(p utils.Users) bool, response *httpop.Response, w http.ResponseWriter) bool {
	if !fn(loggedUser) {
		response.GenerateErrorAccessDeniedResponse(nil, "access denied")
		response.SendResponse(w)
		return false
	}
	return true
}

//REFACTOR

func authorizationUsers(loggedUser utils.Users, IUser interface{}, response *httpop.Response, w http.ResponseWriter, id string) bool {
	var res bool
	switch IUser.(type) {
	case *[]utils.Users:
		{
			data := IUser.(*[]utils.Users)
			if ValidateGetUsersPermission(loggedUser) {
				res = response.ValidateError(db.DB.Find(&data), w, "not users")
			} else {
				res = response.ValidateError(db.DB.Omit("AdministrativeRole", "Password", "Email").Find(&data), w, "not users")
			}
		}

	case *utils.Users:
		{
			data := IUser.(*utils.Users)
			if ValidateGetUsersPermission(loggedUser) {
				res = response.ValidateError(db.DB.Find(&data, id), w, "not users")
			} else {
				res = response.ValidateError(db.DB.Omit("AdministrativeRole", "Password", "Email").Find(&data, id), w, "not users")
			}
		}
	}

	return res
}

func authorizationStudents(loggedUser utils.Users, IUser interface{}, response *httpop.Response, w http.ResponseWriter, id string) bool {
	var res bool
	switch IUser.(type) {
	case *[]utils.Students:
		{
			data := IUser.(*[]utils.Students)
			if ValidateGetUsersPermission(loggedUser) {
				res = response.ValidateError(db.DB.Find(&data), w, "not students")
				dt := *data
				for counter := range dt {
					db.DB.Model(&dt[counter]).Association("User").Find(&dt[counter].User)
				}
			} else {
				res = response.ValidateError(db.DB.Omit("User").Find(&data), w, "not students")
			}
		}

	case *utils.Students:
		{
			data := IUser.(*utils.Students)
			if ValidateGetUsersPermission(loggedUser) {
				res = response.ValidateError(db.DB.Find(&data, id), w, "not student")
				db.DB.Model(&data).Association("User").Find(&data.User)
				db.DB.Model(&data.User).Association("AdministrativeRole").Find(&data.User.AdministrativeRole)
				db.DB.Model(&data.User.AdministrativeRole).Association("PermissionGroup").Find(&data.User.AdministrativeRole.PermissionGroup)
			} else {
				res = response.ValidateError(db.DB.Omit("User").Find(&data, id), w, "not student")
			}
		}
	}

	return res
}

func authorizationTeacher(loggedUser utils.Users, IUser interface{}, response *httpop.Response, w http.ResponseWriter, id string) bool {
	var res bool
	switch IUser.(type) {
	case *[]utils.Teachers:
		{
			data := IUser.(*[]utils.Teachers)
			if ValidateGetUsersPermission(loggedUser) {
				res = response.ValidateError(db.DB.Find(&data), w, "not students")
				for counter := range *(data) {
					db.DB.Model(&(*data)[counter]).Association("User").Find(&(*data)[counter].User)
				}
			} else {
				res = response.ValidateError(db.DB.Omit("User").Find(&data), w, "not teachers")
			}
		}

	case *utils.Teachers:
		{
			data := IUser.(*utils.Teachers)
			if ValidateGetUsersPermission(loggedUser) {
				res = response.ValidateError(db.DB.Find(&data, id), w, "not teacher")
				db.DB.Model(&data).Association("User").Find(&data.User)
				db.DB.Model(&data.User).Association("AdministrativeRole").Find(&data.User.AdministrativeRole)
				db.DB.Model(&data.User.AdministrativeRole).Association("PermissionGroup").Find(&data.User.AdministrativeRole.PermissionGroup)
			} else {
				res = response.ValidateError(db.DB.Omit("User").Find(&data, id), w, "not teacher")
			}
		}
	}

	return res
}
