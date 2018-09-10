package controllers

/*
import (
	"github.com/astaxie/beego"

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
func (this *AdminUsersController) Show(id int) {
	pthis := &this.Controller
	u, e := AppModel.GetUserById(id)
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
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /user/:id [put]
func (this *AdminUsersController) Edit(id int) {
	pthis := &this.Controller
	au := GetAuthor(pthis)

	if au.Id == id {
		wre(pthis, 403)
	}

	uedit := UserEdit{}
	ReadInputBody(pthis, &uedit)

	u, e := AppModel.GetUserById(id)
	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		wre(pthis, 500)
	}

	u.Rol = uedit.Rol
	u.Enable = uedit.Enable
	Validate(pthis, &u)

	_, e = AppModel.UpdateUser(u)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}

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
	pthis := &this.Controller
	GetAuthor(pthis)

	qs := AppModel.GetUserQuerySeter()

	opt := ReadPagAndOrdOptions(pthis)

	qs = qs.Limit(opt.Limit).Offset(opt.Offset)

	if opt.OrderBy != "" {
		qs = qs.OrderBy(fmtorder(&opt))
	}

	var users []models.User
	_, e := qs.All(&users)

	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}

	this.Data["json"] = users
	this.ServeJSON()
}

type UserEdit struct {
	Rol    string `json:"rol"`
	Enable bool   `json:"enable"`
}
*/
