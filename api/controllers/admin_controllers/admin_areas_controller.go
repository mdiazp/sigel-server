package admin_controllers

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
	"github.com/mdiazp/sirel-server/api/models/models2"
)

type AdminAreasController struct {
	controllers.BaseAreasController
}

func (this *AdminAreasController) accessControl() {
	var (
		e error
	)

	author := this.GetAuthor()
	if author.HaveRol(models.RolSuperadmin) {
		return
	}

	area_id, e := this.GetInt("area_id")
	this.WE(e, 400)

	var tmp []orm.Params
	_, e = app.Model().Raw("select user_id from area_admin "+
		"where area_id=? and user_id=? limit 1 offset 0",
		area_id, author.ID).Values(&tmp)

	if e != nil {
		beego.Error(e)
		this.WE(e, 500)
	}

	if len(tmp) > 0 {
		return
	}

	this.WE(errors.New("Forbidden"), 403)
}

// @Title Retrieve Area Info
// @Description Get area info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area_id		query	int	true		"Area id"
// @Success 200 {object} models2.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [get]
func (this *AdminAreasController) Get() {
	this.accessControl()

	model := app.Model()
	o := model.NewArea()
	this.BaseAreasController.Show(o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Create new area
// @Description Create new area (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area		body	models.Area	true		"New Area"
// @Success 200 {object} models2.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [post]
func (this *AdminAreasController) Post() {
	o := app.Model().NewArea()
	this.BaseAreasController.Create(o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Update Area
// @Description Edit area (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area_id		query	int	true		"Area id"
// @Param	area		body	models2.Area	true		"Edited Area"
// @Success 200 {object} models2.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [patch]
func (this *AdminAreasController) Patch() {
	this.accessControl()

	o := app.Model().NewArea()
	this.BaseAreasController.Update(o)

	this.Data["json"] = o
	this.ServeJSON()
}

// @Title Delete Area
// @Description Remove area by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area_id		query	string	true		"Area id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [delete]
func (this *AdminAreasController) Remove() {
	this.accessControl()

	this.BaseAreasController.Remove()

	this.Data["json"] = "OK"
	this.ServeJSON()
}

// @Title Get Areas List
// @Description Get areas list (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	orderDirection		query	string	false		"asc or desc"
// @Param	search		query	string	false		"Search in name"
// @Success 200 {object} []models2.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /areas [get]
func (this *AdminAreasController) List() {
	var (
		l []models2.Area
	)

	u := this.GetAuthor()
	if !u.HaveRol(models.RolSuperadmin) {
		this.Ctx.Input.SetParam("ofAdmin", "true")
	}
	this.BaseAreasController.List(&l)

	this.Data["json"] = l
	this.ServeJSON()
}

// @Title Get Admins
// @Description Delete user from admins by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area_id		query	string	true		"Area id"
// @Success 200 {[]models.UserPublicInfo}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area/admins [get]
func (this *AdminAreasController) Admins() {
	this.accessControl()

	var admins []models.UserPublicInfo
	this.BaseAreasController.Admins(&admins)
	this.Data["json"] = admins
	this.ServeJSON()
}

// @Title Delete User from Admins
// @Description Delete user from admins by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area_id		query	int	true		"Area id"
// @Param	user_id		query	int	true		"User id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area/admins [put]
func (this *AdminAreasController) PutAdmin() {
	this.accessControl()

	this.BaseAreasController.AddAdmin()
	this.Data["json"] = "OK"
	this.ServeJSON()
}

// @Title Delete User from Admins
// @Description Delete user from admins by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area_id		query	string	true		"Area id"
// @Param	user_id		query	string	true		"User id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area/admins [delete]
func (this *AdminAreasController) DeleteAdmin() {
	this.accessControl()

	this.BaseAreasController.RemoveAdmin()
	this.Data["json"] = "OK"
	this.ServeJSON()
}
