package public

import (
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
)

// UsersController ...
type UsersController struct {
	controllers.BaseUsersController
}

// GetUsersPublicInfo ...
// @Title Get usernames list
// @Description Get user's public info list
// @Param	prefixFilter		query	string	false		"prefixFilter"
// @Param	limit		query	int	false		"Limit (10 or 50 or 100)"
// @Param	offset		query	int	false		"Offset"
// @Param	orderby		query	string	false		"OrderBy (property name)"
// @Param	ordDesc		query	string	false		"true or false"
// @Success 200 {object} []models.UserPublicInfo
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /users/publicinfo [get]
func (c *UsersController) GetUsersPublicInfo() {
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
