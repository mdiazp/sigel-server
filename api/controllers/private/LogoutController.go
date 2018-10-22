package private

import (
	"github.com/mdiazp/sirel-server/api/controllers"
)

type LogoutController struct {
	controllers.BaseController
}

func (c *LogoutController) URLMapping() {
	c.Mapping("/logout", c.Logout)
}

// @Title Logout
// @Summary Logout
// @Description Close session of the user in the system
// @Param	authHd		header	string	true		"Authorized Token"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 500 Internal Server Error
// @router /logout [delete]
func (this *LogoutController) Logout() {
	this.ServeJSON()
}
