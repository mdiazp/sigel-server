package public_controllers

import (
	"github.com/astaxie/beego"
	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
)

type PublicUsersController struct {
	controllers.BaseController
}

// @Title Get usernames list
// @Description Get usernames list
// @Param	prefixFilter		query	string	true		"prefixFilter"
// @Param	limit		query	int	true		"Limit"
// @Param	offset		query	int	true		"Offset"
// @Success 200 {object} models.UserOnlyUsernamesAndId
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /users/usernames [get]
func (this *PublicUsersController) GetUsernamesList() {
	var (
		e            error
		prefixFilter string
		limit        int
		offset       int
	)
	prefixFilter = this.GetString("prefixFilter")
	limit, e = this.GetInt("limit")
	if e == nil {
		offset, e = this.GetInt("offset")
	}

	this.WE(e, 400)

	var data []models.UserPublicInfo
	_, e = app.Model().Raw(
		"SELECT id, username FROM k_user WHERE username like ? limit ? offset ?",
		prefixFilter+"%", limit, offset).QueryRows(&data)

	beego.Debug(e)
	this.WE(e, 500)

	this.Data["json"] = data
	this.ServeJSON()
}
