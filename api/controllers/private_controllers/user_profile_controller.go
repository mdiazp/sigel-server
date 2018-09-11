package private_controllers

import (
	"github.com/astaxie/beego"
	"gitlab.com/manuel.diaz/sirel/server/api/app"
	"gitlab.com/manuel.diaz/sirel/server/api/controllers"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type ProfileController struct {
	controllers.BaseAreasController
}

// @Title Get User Profile
// @Description Get user profile by username
// @Param	authHd		header	string	true		"Authentication token"
// @Success 200 {object} public_controllers.Profile
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /profile [get]
func (this *ProfileController) Get() {
	au := this.GetAuthor()
	this.Data["json"] = toProfile(&au)
	this.ServeJSON()
}

// @Title Edit Profile
// @Description Edit profile
// @Param	authHd		header	string	true		"Authentication token"
// @Param	profile		body	public_controllers.ProfileEdit	true		"Edited Profile"
// @Success 200 {object} public_controllers.Profile
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /profile [put]
func (this *ProfileController) Put() {
	var (
		e error
	)
	au := this.GetAuthor()

	profile := ProfileEdit{}
	this.ReadInputBody(&profile)

	au.Email = profile.Email
	au.SendNotificationsToEmail = profile.SendNotificationsToEmail

	_, e = app.Model().Update(&au)
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
	this.Data["json"] = toProfile(&au)
	this.ServeJSON()
}

func toProfile(u *models.User) Profile {
	return Profile{
		Username: u.Username,
		Name:     u.Name,
		ProfileEdit: ProfileEdit{
			Email: u.Email,
			SendNotificationsToEmail: u.SendNotificationsToEmail,
		},
	}
}

type Profile struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	ProfileEdit
}

type ProfileEdit struct {
	Email                    string `json:"email"`
	SendNotificationsToEmail bool   `json:"send_notifications_to_email"`
}
