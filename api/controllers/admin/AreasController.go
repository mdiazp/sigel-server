package admin

import (
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
)

// AreasController ...
type AreasController struct {
	controllers.BaseAreasController
}

// Get ...
// @Title Retrieve Area Info
// @Description Get area info by id (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area_id		query	int	true		"Area id"
// @Success 200 {object} models.AreaInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [get]
func (c *AreasController) Get() {
	c.Data["json"] = c.BaseAreasController.Show()
	c.ServeJSON()
}

// Post ...
// @Title Create new area
// @Description Create new area (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area		body	models.AreaInfo	true		"New Area"
// @Success 200 {object} models.AreaInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [post]
func (c *AreasController) Post() {
	c.AccessControl(models.RolSuperadmin)
	c.Data["json"] = c.BaseAreasController.Create()
	c.ServeJSON()
}

// Patch ...
// @Title Update Area
// @Description Edit area (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
// @Param	area_id		query	int	true		"Area id"
// @Param	area		body	models.AreaInfo	true		"Edited Area"
// @Success 200 {object} models.AreaInfo
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [patch]
func (c *AreasController) Patch() {
	c.AccessControl(models.RolSuperadmin)
	c.Data["json"] = c.BaseAreasController.Update()
	c.ServeJSON()
}

// Remove ...
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
func (c *AreasController) Remove() {
	c.AccessControl(models.RolSuperadmin)
	c.BaseAreasController.Remove()
	c.Data["json"] = "OK"
	c.ServeJSON()
}

// List ...
// @Title Get Areas List
// @Description Get areas list (role admin required)
// @Param	authHd		header	string	true		"Authentication token"
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
