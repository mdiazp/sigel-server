package private

import (
	"github.com/mdiazp/sirel-server/api/controllers"
)

// LogoutController ...
type LogoutController struct {
	controllers.BaseController
}

// URLMapping ...
func (c *LogoutController) URLMapping() {
	c.Mapping("/logout", c.Logout)
}

// Logout ...
// @Title Logout
// @Summary Logout
// @Description Close session of the user in the system
// @Param	authHd		header	string	true		"Authorized Token"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 500 Internal Server Error
// @router /logout [delete]
func (c *LogoutController) Logout() {
	c.ServeJSON()
}
