package admin_controllers

import (
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
)

type AdminLocalsController struct {
	controllers.BaseLocalsController
}

// @Title Retrieve Local Info
// @Description Get local info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	int	true		"Local id"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [get]
func (this *AdminLocalsController) Get() {
	o := models.Local{}

	this.BaseLocalsController.Show(&o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Create new local
// @Description Create new local (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local		body	models.Local	true		"New Local"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [post]
func (this *AdminLocalsController) Post() {
	o := models.Local{}

	this.BaseLocalsController.Create(&o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Update Local
// @Description Edit local (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	int	true		"Local id"
// @Param	local		body	models.Local	true		"Edited Local"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [patch]
func (this *AdminLocalsController) Patch() {
	o := models.Local{}

	this.BaseLocalsController.Update(&o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Delete Local
// @Description Remove local by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	string	true		"Local id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [delete]
func (this *AdminLocalsController) Delete() {
	this.BaseLocalsController.Remove()
	this.ServeJSON()
}

// @Title Get Locals List
// @Description Get locals list (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	sortorder		query	string	false		"asc or desc"
// @Param	enable_to_reserve		query	string	false		"Local Property (true o false)"
// @Param	area_id		query	int	false		"Local Property"
// @Param	search		query	string	false		"Search in name"
// @Success 200 {object} []models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /locals [get]
func (this *AdminLocalsController) List() {
	var l []models.Local
	this.BaseLocalsController.List(&l)
	this.Data["json"] = l
	this.ServeJSON()
}

// @Title Get Admins
// @Description Delete user from admins by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	string	true		"Local id"
// @Success 200 {[]models.UserPublicInfo}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local/admins [get]
func (this *AdminLocalsController) Admins() {
	var admins []models.UserPublicInfo
	this.BaseLocalsController.Admins(&admins)
	this.Data["json"] = admins
	this.ServeJSON()
}

// @Title Delete User from Admins
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
func (this *AdminLocalsController) PutAdmin() {
	this.BaseLocalsController.AddAdmin()
	this.Data["json"] = "OK"
	this.ServeJSON()
}

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
func (this *AdminLocalsController) DeleteAdmin() {
	this.BaseLocalsController.RemoveAdmin()
	this.Data["json"] = "OK"
	this.ServeJSON()
}
