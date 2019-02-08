package public

import (
	"github.com/mdiazp/sigel-server/api/controllers"
	"github.com/mdiazp/sigel-server/api/models"
)

// UsersController ...
type UsersController struct {
	controllers.BaseUsersController
}

// GetUserPublicInfo ...
// @Title Get user's public info
// @Description Get user's public info
// @Param	user_id		query	int	false		"User ID"
// @Success 200 {object} models.UserPublicInfo
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /user/publicinfo [get]
func (c *UsersController) GetUserPublicInfo() {
	u := c.GetUser()
	upinfo := models.UserPublicInfo{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
	}
	c.Data["json"] = upinfo
	c.ServeJSON()
}

// GetUsersPublicInfo ...
// @Title Get usernames list
// @Description Get user's public info list
// @Param	username		query	string	false		"prefixFilter"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	desc		query	string	false		"true or false"
// @Success 200 {object} []models.UserPublicInfo
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /users/publicinfo [get]
func (c *UsersController) GetUsersPublicInfo() {
	c.Ctx.Input.SetParam("name", "")
	c.Ctx.Input.SetParam("email", "")
	c.Ctx.Input.SetParam("rol", "")
	c.Ctx.Input.SetParam("enable", "")

	users := c.GetUsers()
	upinfo := make([]*models.UserPublicInfo, 0)

	for _, u := range *users {
		up := models.UserPublicInfo{
			ID:       u.ID,
			Username: u.Username,
			Name:     u.Name,
		}
		upinfo = append(upinfo, &up)
	}

	c.Data["json"] = upinfo
	c.ServeJSON()
}
