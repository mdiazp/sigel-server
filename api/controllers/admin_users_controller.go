package controllers

import (
	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/app"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type AdminUsersController struct {
	beego.Controller
}

// @Title Get User Info
// @Description Get user info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	int	true		"User id"
// @Success 200 {object} models.User
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /user/:id [get]
func (this *AdminUsersController) Get() {
	var (
		e     error
		pthis = &this.Controller
	)

	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}

	u := models.User{}
	e = app.Model().QueryTable(&models.User{}).Filter("id", id).Limit(1).One(&u)
	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		wre(pthis, 500)
	}
	this.Data["json"] = u
	this.ServeJSON()
}

// @Title Edit User
// @Description Edit rol and enable properties (role admin required, user can't edit itself)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	int	true		"User id"
// @Param	userEdit		body	controllers.UserEdit	true		"Edited User"
// @Success 200 {object} models.User
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /user/:id [put]
func (this *AdminUsersController) Put() {
	var (
		e     error
		pthis = &this.Controller
	)

	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}

	// The author can't update itself
	au := GetAuthor(pthis)

	if au.Id == id {
		wre(pthis, 403)
	}

	uedit := UserEdit{}
	ReadInputBody(pthis, &uedit)

	u := models.User{}
	e = app.Model().QueryTable(&models.User{}).Filter("id", id).Limit(1).One(&u)
	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		wre(pthis, 500)
	}

	u.Rol = uedit.Rol
	u.Enable = uedit.Enable
	Validate(pthis, &u)

	_, e = app.Model().Update(&u)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
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
		e     error
		pthis = &this.Controller
	)

	qs := app.Model().QueryTable(&models.User{})

	opt := ReadPagAndOrdOptions(pthis)
	qs = qs.Limit(opt.Limit).Offset(opt.Offset)
	if opt.OrderBy == "" {
		opt.OrderBy = "id"
	}
	if opt.OrderBy != "" {
		qs = qs.OrderBy(fmtorder(&opt))
	}

	var l []models.User
	_, e = qs.All(&l)

	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		beego.Error(e.Error())
		wre(pthis, 500)
	}

	this.Data["json"] = l
	this.ServeJSON()
}

type UserEdit struct {
	Rol    string `json:"rol"`
	Enable bool   `json:"enable"`
}
