package roles

import (
	"net/http"

	"github.com/school-sys-rest-api/services/httpop"
	"github.com/school-sys-rest-api/utils"
)

func VerifyCreateRolePermission(p utils.Users) bool {
	return p.AdministrativeRole.PermissionGroup.CreateUsersPermission == "1"
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
