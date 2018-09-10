package controllers

/*
type PublicAreasController struct {
	AreasBaseController
}

// @Title Get public area info
// @Description Get area info by id
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	int	true		"Area id"
// @Success 200 {object} models.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area/:id [get]
func (this *PublicAreasController) Get() {
	pthis := &this.Controller
	this.Ctx.Input.SetParam("enable_to_reserve", "true")
	area := this.get()
	if !area.EnableToReserve {
		wre(pthis, 404)
	}
	this.Data["json"] = area
	this.ServeJSON()
}

// @Title Get public areas list
// @Description Get areas list (role admin required)
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	bool	false		"Order Desc"
// @Success 200 {object} []models.Area
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /areas [get]
func (this *PublicAreasController) List() {
	this.Ctx.Input.SetParam("enable_to_reserve", "true")
	areas := this.list()
	this.Data["json"] = areas
	this.ServeJSON()
}
*/
