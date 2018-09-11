package admin_controllers

import (
	"errors"

	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/app"
	"gitlab.com/manuel.diaz/sirel/server/api/controllers"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type AdminUsersController struct {
	controllers.BaseController
}

// @Title Get User Info
// @Description Get user info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		query	int	true		"User id"
// @Success 200 {object} models.User
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /user [get]
func (this *AdminUsersController) Get() {
	var (
		e error
	)

	id, e := this.GetInt("id")
	this.WE(e, 400)

	u := models.User{}
	e = app.Model().QueryTable(&models.User{}).Filter("id", id).Limit(1).One(&u)
	if e != nil {
		if e == models.ErrResultNotFound {
			this.WE(e, 404)
		}
		this.WE(e, 500)
	}
	this.Data["json"] = u
	this.ServeJSON()
}

// @Title Edit User
// @Description Edit rol and enable properties (role admin required, user can't edit itself)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	int	true		"User id"
// @Param	userEdit		body	admin_controllers.UserEdit	true		"Edited User"
// @Success 200 {object} models.User
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /user [put]
func (this *AdminUsersController) Put() {
	var (
		e error
	)

	id, e := this.GetInt("id")
	this.WE(e, 400)

	// The author can't update itself
	au := this.GetAuthor()

	if au.Id == id {
		this.WE(errors.New("author can't update itself"), 403)
	}

	uedit := UserEdit{}
	this.ReadInputBody(&uedit)

	u := models.User{}
	e = app.Model().QueryTable(&models.User{}).Filter("id", id).Limit(1).One(&u)
	if e != nil {
		if e == models.ErrResultNotFound {
			this.WE(e, 404)
		}
		this.WE(e, 500)
	}

	u.Rol = uedit.Rol
	u.Enable = uedit.Enable
	this.Validate(&u)

	_, e = app.Model().Update(&u)
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}

	this.Data["json"] = u
	this.ServeJSON()
}

// @Title Get Users List
// @Description Get users list (role admin required, user can't edit itself)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	bool	false		"Order Desc"
// @Success 200 {object} []models.User
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /users [get]
func (this *AdminUsersController) List() {

	//Nota: validate orderBy column name

	var (
		e error
	)

	qs := app.Model().QueryTable(&models.User{})

	opt := this.ReadPagAndOrdOptions()
	qs = qs.Limit(opt.Limit).Offset(opt.Offset)
	if opt.OrderBy == "" {
		opt.OrderBy = "id"
	}
	if opt.OrderBy != "" {
		qs = qs.OrderBy(this.Fmtorder(&opt))
	}

	var l []models.User
	_, e = qs.All(&l)

	if e != nil {
		if e == models.ErrResultNotFound {
			this.WE(e, 404)
		}
		beego.Error(e.Error())
		this.WE(e, 500)
	}

	this.Data["json"] = l
	this.ServeJSON()
}

type UserEdit struct {
	Rol    string `json:"rol"`
	Enable bool   `json:"enable"`
}
