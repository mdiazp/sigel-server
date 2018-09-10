package controllers

/*
type PublicLocalsController struct {
	LocalsBaseController
}

// @Title Get public local info
// @Description Get local info by id
// @Param	authHd		header	string	true		"Authentication token"
// @Param	id		path	int	true		"Local id"
// @Success 200 {object} models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /local/:id [get]
func (this *PublicLocalsController) Get() {
	pthis := &this.Controller
	this.Ctx.Input.SetParam("enable_to_reserve", "true")
	area := this.get()
	if !area.EnableToReserve {
		wre(pthis, 404)
	}
	this.Data["json"] = area
	this.ServeJSON()
}

// @Title Get public locals list
// @Description Get locals list (role admin required)
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	bool	false		"Order Desc"
// @Success 200 {object} []models.Local
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 403 Forbidden
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /locals [get]
func (this *PublicLocalsController) List() {
	this.Ctx.Input.SetParam("enable_to_reserve", "true")
	objs := this.list()
	this.Data["json"] = objs
	this.ServeJSON()
}
*/
