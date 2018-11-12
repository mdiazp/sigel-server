package public

import (
	"github.com/mdiazp/sirel-server/api/controllers"
)

// LocalsController ...
type LocalsController struct {
	controllers.BaseLocalsController
}

// Get ...
// @Title Retrieve public local info
// @Description Get local info by id
// @Param	local_id		query	int	true		"Local id"
// @Success 200 {object} models.LocalInfo
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local [get]
func (c *LocalsController) Get() {
	c.Ctx.Input.SetParam("enable_to_reserve", "true")
	c.Data["json"] = c.BaseLocalsController.Show()
	c.ServeJSON()
}

// List ...
// @Title Get public locals list
// @Description Get public locals list
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	string	false		"true or false"
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
	c.Ctx.Input.SetParam("enable_to_reserve", "true")
	c.Ctx.Input.SetParam("ofAdmin", "false")
	c.Data["json"] = c.BaseLocalsController.List().Locals
	c.ServeJSON()
}
