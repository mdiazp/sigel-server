package controllers

import (
	"github.com/astaxie/beego"
	"gitlab.com/manuel.diaz/sirel/server/api/app"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) URLMapping() {
	c.Mapping("/login", c.Login)
}

// @Title Login
// @Summary Login
// @Description Open session for user into the system
// @Param	credentials		body	controllers.Credentials	true		"Credentials for authentication"
// @Success 200 {object} controllers.Session
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 500 Internal Server Error
// @Accept json
// @router /login [post]
func (this *LoginController) Login() {
	var (
		e     error
		u     models.User
		pthis = &this.Controller
	)
	// Read Credentials
	cred := Credentials{}
	ReadInputBody(pthis, &cred)

	if cred.Username != "SIREL" {
		// Ldap Authentication
		authp := app.AuthProvider(beego.AppConfig.String("AUTH_PROVIDER"))
		e = authp.Authenticate(cred.Username, cred.Password)
		if e != nil {
			beego.Info("authprovider.Authenticate() error: ", e.Error())
			wre(pthis, 401)
		}

		// query for user records in AD and compare usernames to
		// check username diferences
		urecords, e := authp.GetUserRecords(cred.Username)
		if e != nil {
			beego.Error("ldap.FullRecordError: ", e.Error())
			wre(pthis, 500)
		}
		if urecords.Username != cred.Username {
			wre(pthis, 400)
		}
		cred.Username = urecords.Username

		// Register user if they is not yet
		e = app.Model().QueryTable(models.User{}).Filter("username", cred.Username).Limit(1).One(&u)
		if e == models.ErrResultNotFound {
			//Then register the new user
			u = models.User{
				Username: urecords.Username,
				Name:     urecords.Name,
				Email:    urecords.Email,
				Rol:      models.RolUser,
				Enable:   true,
				SendNotificationsToEmail: true,
			}

			_, e = app.Model().Insert(&u)
			if e != nil {
				beego.Error("Unexpected error occurred: ", e.Error())
				wre(&this.Controller, 500)
			}
		} else if e != nil {
			beego.Error("Unexpected error occurred: ", e.Error())
			wre(&this.Controller, 500)
		}

		// Check user enable
		if !u.Enable {
			wre(&this.Controller, 401)
		}
	} else {
		if beego.AppConfig.String("SIREL_PASSWORD") != cred.Password {
			wre(&this.Controller, 401)
		}
		u = models.User{
			Username: "SIREL",
			Rol:      models.RolSuperadmin,
			Enable:   true,
		}
	}

	// Encrypt credentials
	s, e := app.Crypto().Encrypt(cred.Username)
	if e != nil {
		beego.Error("crypto.Encrypt(", cred.Username, "):", e.Error())
		wre(&this.Controller, 500)
	}

	// Prepare response
	this.Data["json"] = Session{
		Username: u.Username,
		Rol:      u.Rol,
		Token:    s,
	}

	this.ServeJSON()
}

type LogoutController struct {
	beego.Controller
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

type Session struct {
	Username string `json:"username"`
	Rol      string `json:"rol"`
	Token    string `json:"jwtToken"`
}

type Credentials struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}
