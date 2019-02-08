package public

import (
	"github.com/mdiazp/sigel-server/api/controllers"
)

// AreasController ...
type AreasController struct {
	controllers.BaseAreasController
}

// Get ...
// @Title Retrieve public area info
// @Description Get area info by id
// @Param	area_id		query	int	true		"Area id"
// @Success 200 {object} models.AreaInfo
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [get]
func (c *AreasController) Get() {
	c.Data["json"] = c.BaseAreasController.Show()
	c.ServeJSON()
}

// List ...
// @Title Get public areas list
// @Description Get public areas list
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	string	false		"true or false"
// @Param	search		query	string	false		"Search in name"
// @Success 200 {object} []models.AreaInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /areas [get]
func (c *AreasController) List() {
	c.Data["json"] = c.BaseAreasController.List().Areas
	c.ServeJSON()
}
