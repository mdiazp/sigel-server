package private

import (
	"fmt"

	"github.com/mdiazp/sigel-server/api/controllers"
	"github.com/mdiazp/sigel-server/api/models"
)

// ProfileController ...
type ProfileController struct {
	controllers.BaseAreasController
}

// Get ...
// @Title Get User Profile
// @Description Get user profile by username
// @Param	authHd		header	string	true		"Authentication token"
// @Success 200 {object} models.UserProfile
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/profile [get]
func (c *ProfileController) Get() {
	au := c.GetAuthor()
	c.Data["json"] = toProfile(au)
	c.ServeJSON()
}

// Patch ...
// @Title Edit Profile
// @Description Edit profile
// @Param	authHd		header	string	true		"Authentication token"
// @Param	profile		body	models.ProfileEdit	true		"Edited Profile"
// @Success 200 {object} models.UserProfile
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /session/profile [patch]
func (c *ProfileController) Patch() {
	profile := models.ProfileEdit{}
	c.ReadInputBody(&profile)

	au := c.GetAuthor()
	if au.Username == "SIREL" {
		c.WE(fmt.Errorf("User SIREL can't be updated"), 403)
	}

	au.Email = profile.Email
	au.SendNotificationsToEmail = profile.SendNotificationsToEmail
	c.Validate(au)

	e := au.Update()
	c.WE(e, 500)

	c.Data["json"] = toProfile(au)
	c.ServeJSON()
}

func toProfile(u *models.User) models.UserProfile {
	return models.UserProfile{
		Username: u.Username,
		Name:     u.Name,
		Email:    u.Email,
		SendNotificationsToEmail: u.SendNotificationsToEmail,
	}
}
