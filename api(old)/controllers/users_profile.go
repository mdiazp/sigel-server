package controllers

import (
	"github.com/astaxie/beego"
)

type ProfileController struct {
	beego.Controller
}

// @Title Get User Profile
// @Description Get user profile by username
// @Param	authHd		header	string	true		"Authentication token"
// @Success 200 {object} controllers.Profile
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /profile [get]
func (this *ProfileController) Get() {
	au := GetAuthor(&this.Controller)

	this.Data["json"] = Profile{
		Username: au.Username,
		Name:     au.Name,
		ProfileEdit: ProfileEdit{
			Email: au.Email,
			SendNotificationsToEmail: au.SendNotificationsToEmail,
		},
	}
	this.ServeJSON()
}

// @Title Edit Profile
// @Description Edit profile
// @Param	authHd		header	string	true		"Authentication token"
// @Param	profile		body	controllers.ProfileEdit	true		"Edited Profile"
// @Success 200
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /profile [put]
func (this *ProfileController) Put() {
	pthis := &this.Controller
	au := GetAuthor(pthis)

	profile := ProfileEdit{}
	ReadInputBody(pthis, &profile)

	au.Email = profile.Email
	au.SendNotificationsToEmail = profile.SendNotificationsToEmail
	Validate(pthis, &au)

	_, e := AppModel.UpdateUser(au)
	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	this.ServeJSON()
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
