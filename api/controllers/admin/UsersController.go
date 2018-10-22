package admin

import (
	"fmt"

	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
)

// UsersController ...
type UsersController struct {
	controllers.BaseUsersController
}

// Get ...
// @Title Get User Info
// @Description Get user info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	user_id		query	int	true		"User id"
// @Success 200 {object} models.UserInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /user [get]
func (c *UsersController) Get() {
	c.AccessControl(models.RolSuperadmin)
	u := c.GetUser()
	c.Data["json"] = u
	c.ServeJSON()
}

// Patch ...
// @Title Edit User
// @Description Edit rol and enable properties (role admin required, user can't edit itself)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	user_id		query	int	true		"User id"
// @Param	userEdit		body	admin.UserEdit	true		"Edited User"
// @Success 200 {object} models.UserInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /user [patch]
func (c *UsersController) Patch() {
	c.AccessControl(models.RolSuperadmin)
	u := c.GetUser()

	uedit := UserEdit{}
	c.ReadInputBody(&uedit)

	if uedit.Rol == models.RolSuperadmin && c.GetAuthor().Username != "SIREL" {
		c.WE(fmt.Errorf("Only user SIREl is enabled to create superadmin users"), 403)
	}

	u.Rol = uedit.Rol
	u.Enable = uedit.Enable
	c.Validate(u)

	e := u.Update()
	c.WE(e, 500)

	c.Data["json"] = u
	c.ServeJSON()
}

// List ...
// @Title Get Users List
// @Description Get users list (role admin required, user can't edit itself)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	prefixFilter		query	string	false		"prefixFilter"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	ordDesc		query	string	false		"true or false"
// @Success 200 {object} []models.UserInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /users [get]
func (c *UsersController) List() {
	c.Data["json"] = c.BaseUsersController.GetUsers()
	c.ServeJSON()
}

// UserEdit ...
type UserEdit struct {
	Rol    string `json:"rol"`
	Enable bool   `json:"enable"`
}
