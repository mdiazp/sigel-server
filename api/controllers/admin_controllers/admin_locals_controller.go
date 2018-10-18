package admin_controllers

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
)

type AdminLocalsController struct {
	controllers.BaseLocalsController
}

func (this *AdminLocalsController) accessControl() {
	var (
		e error
	)

	author := this.GetAuthor()
	if author.HaveRol(models.RolSuperadmin) {
		return
	}

	local_id, e := this.GetInt("local_id")
	this.WE(e, 400)

	var tmp []orm.Params

	//checking local_admin
	_, e = app.Model().Raw("select user_id from local_admin "+
		"where local_id=? and user_id=? limit 1 offset 0",
		local_id, author.ID).Values(&tmp)

	if e != nil {
		beego.Error(e)
		this.WE(e, 500)
	}

	if len(tmp) > 0 {
		return
	}

	//checking area_admin
	_, e = app.Model().Raw("select user_id from area_admin "+
		"join area on area.id=area_admin.area_id join local on local.area_id=area.id "+
		"where local.id=? and area_admin.user_id=? limit 1 offset 0",
		local_id, author.ID).Values(&tmp)

	if e != nil {
		beego.Error(e)
		this.WE(e, 500)
	}

	if len(tmp) > 0 {
		return
	}

	this.WE(errors.New("Forbidden"), 403)
}

// @Title Retrieve Local Info
// @Description Get local info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local_id		query	int	true		"Local id"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [get]
func (this *AdminLocalsController) Get() {
	this.accessControl()

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
// @Param	local_id		query	int	true		"Local id"
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
	this.accessControl()

	o := models.Local{}
	this.BaseLocalsController.Update(&o)

	this.Data["json"] = o
	this.ServeJSON()
}

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
func (this *AdminLocalsController) Delete() {
	this.accessControl()

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
	u := this.GetAuthor()
	if !u.HaveRol(models.RolSuperadmin) {
		this.Ctx.Input.SetParam("ofAdmin", "true")
	}

	var l []models.Local
	this.BaseLocalsController.List(&l)
	this.Data["json"] = l
	this.ServeJSON()
}

// @Title Get Admins
// @Description Delete user from admins by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	local_id		query	string	true		"Local id"
// @Success 200 {[]models.UserPublicInfo}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local/admins [get]
func (this *AdminLocalsController) Admins() {
	this.accessControl()

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
	this.accessControl()

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
	this.accessControl()

	this.BaseLocalsController.RemoveAdmin()
	this.Data["json"] = "OK"
	this.ServeJSON()
}
