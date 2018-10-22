package admin

import (
	"github.com/astaxie/beego"
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
)

// LocalsController ...
type LocalsController struct {
	controllers.BaseLocalsController
}

// Get ...
// @Title Retrieve Local Info
// @Description Get local info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local_id		query	int	true		"Local id"
// @Success 200 {object} models.LocalInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [get]
func (c *LocalsController) Get() {
	beego.Debug(models.RolAdmin)
	c.AccessControl(models.RolAdmin)
	c.Data["json"] = c.BaseLocalsController.Show()
	c.ServeJSON()
}

// Post ...
// @Title Create new local
// @Description Create new local (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local		body	models.LocalInfo	true		"New Local"
// @Success 200 {object} models.LocalInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [post]
func (c *LocalsController) Post() {
	c.AccessControl(models.RolSuperadmin)
	c.Data["json"] = c.BaseLocalsController.Create()
	c.ServeJSON()
}

// Patch ...
// @Title Update Local
// @Description Edit local (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local_id		query	int	true		"Local id"
// @Param	local		body	models.LocalInfo	true		"Edited Local"
// @Success 200 {object} models.LocalInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [patch]
func (c *LocalsController) Patch() {
	c.AccessControl(models.RolAdmin)
	c.Data["json"] = c.BaseLocalsController.Update()
	c.ServeJSON()
}

// Delete ...
// @Title Delete Local
// @Description Remove local by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local_id		query	string	true		"Local id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [delete]
func (c *LocalsController) Delete() {
	c.AccessControl(models.RolSuperadmin)
	c.BaseLocalsController.Remove()
	c.Data["json"] = "OK"
	c.ServeJSON()
}

// List ...
// @Title Get Locals List
// @Description Get locals list (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	orderDesc		query	string	false		"true or false"
// @Param	enable_to_reserve		query	string	false		"Local Property (true o false)"
// @Param	area_id		query	int	false		"Local Property"
// @Param	search		query	string	false		"Search in name"
// @Success 200 {object} []models.LocalInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /locals [get]
func (c *LocalsController) List() {
	c.AccessControl(models.RolAdmin)
	if !c.GetAuthor().HaveRol(models.RolSuperadmin) {
		c.Ctx.Input.SetParam("ofAdmin", "true")
	}
	c.Data["json"] = c.BaseLocalsController.List().Locals
	c.ServeJSON()
}

// Admins ...
// @Title Get Admins
// @Description Delete user from admins by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local_id		query	string	true		"Local id"
// @Success 200 {object} []models.UserPublicInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local/admins [get]
func (c *LocalsController) Admins() {
	c.AccessControl(models.RolSuperadmin)
	ladmins := *(c.BaseLocalsController.Admins().Users)
	admins := make([]models.UserPublicInfo, 0)

	for _, a := range ladmins {
		admins = append(admins,
			models.UserPublicInfo{
				ID:       a.ID,
				Username: a.Username,
				Name:     a.Name,
			},
		)
	}

	c.Data["json"] = admins
	c.ServeJSON()
}

// PutAdmin ...
// @Title Add User to Admins
// @Description Delete user from admins by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local_id		query	int	true		"Local id"
// @Param	user_id		query	int	true		"User id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local/admins [put]
func (c *LocalsController) PutAdmin() {
	c.AccessControl(models.RolSuperadmin)
	c.Data["json"] = c.BaseLocalsController.AddAdmin()
	c.ServeJSON()
}

// DeleteAdmin ...
// @Title Delete User from Admins
// @Description Delete user from admins by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local_id		query	string	true		"Local id"
// @Param	user_id		query	string	true		"User id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local/admins [delete]
func (c *LocalsController) DeleteAdmin() {
	c.AccessControl(models.RolSuperadmin)
	c.BaseLocalsController.RemoveAdmin()
	c.Data["json"] = "OK"
	c.ServeJSON()
}
