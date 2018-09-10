package controllers

import (
	"github.com/astaxie/beego"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController) URLMapping() {
	c.Mapping("/login", c.Login)
	c.Mapping("/logout", c.Logout)
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
func (this *SessionController) Login() {
	var (
		e error
		u models.User
	)
	// Read Credentials
	cred := Credentials{}
	ReadInputBody(&this.Controller, &cred)

	if cred.Username != "SIREL" {
		// Ldap Authentication
		authp := GetAuthProvider()
		e = authp.Authenticate(cred.Username, cred.Password)
		if e != nil {
			beego.Info("authprovider.Authenticate() error: ", e.Error())
			wre(&this.Controller, 401)
		}

		// query for user records in AD and compare usernames to
		// check username diferences
		urecords, e := authp.GetUserRecords(cred.Username)
		if e != nil {
			beego.Error("ldap.FullRecordError: ", e.Error())
			wre(&this.Controller, 500)
		}
		if urecords.Username != cred.Username {
			wre(&this.Controller, 400)
		}
		cred.Username = urecords.Username

		// Register user if they is not yet
		u, e = AppModel.GetUserByUsername(cred.Username)
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

			u, e = AppModel.CreateUser(u)
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
	crypto := GetCrypto()
	s, e := crypto.Encrypt(cred.Username)
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

// @Title Logout
// @Summary Logout
// @Description Close session of the user in the system
// @Param	authHd		header	string	true		"Authorized Token"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 500 Internal Server Error
// @router /logout [delete]
func (this *SessionController) Logout() {
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

/*
func (this *SessionController) Login() {
	var (
		e error
		u models.User
	)
	// Read Credentials
	cred := GetCredentials()
	e = json.Unmarshal(this.Ctx.Input.RequestBody, &cred)
	if e != nil {
		WriteResponseError(&this.Controller, http.StatusBadRequest)
		return
	}

	if cred.User != "SIREL" {
		// Ldap Authentication
		ldap := GetLdap()
		e = ldap.Authenticate(cred.User, cred.Pass)
		if e != nil {
			beego.Info("ldap.Authenticate() error: ", e.Error())
			WriteResponseError(&this.Controller, http.StatusUnauthorized)
			return
		}

		// query for user records in AD and compare usernames to
		// check username diferences
		aduser, e := ldap.FullRecordAcc(cred.User)
		if e != nil {
			beego.Error("ldap.FullRecordError: ", e.Error())
			WriteResponseError(&this.Controller, http.StatusInternalServerError)
			return
		}
		if aduser["sAMAccountName"][0] != cred.User {
			WriteResponseError(&this.Controller, http.StatusBadRequest)
			return
		}
		cred.User = aduser["sAMAccountName"][0]

		// Register user if they is not yet
		u, e = AppModel.GetUserByUsername(cred.User)
		if models.ErrNotImplementet(e) {
			beego.Error("AppModel not implementet error: ", e.Error())
			WriteResponseError(&this.Controller, http.StatusInternalServerError)
			return
		} else if models.ErrUserNotFound == e {
			//Then register the new user
			u = models.User{
				Username: aduser["sAMAccountName"][0],
				Name:     aduser["displayName"][0],
				Email:    aduser["mail"][0],
				Rol:      models.RolUser,
				Enable:   true,
				SendNotifications_to_email: true,
			}

			u, e = AppModel.CreateUser(u)
			if e != nil {
				beego.Error("Unexpected error occurred: ", e.Error())
				WriteResponseError(&this.Controller, http.StatusInternalServerError)
				return
			}
		} else if e != nil {
			beego.Error("Unexpected error occurred: ", e.Error())
			WriteResponseError(&this.Controller, http.StatusInternalServerError)
			return
		}

	} else {
		if beego.AppConfig.String("SIREL_PASSWORD") != cred.Pass {
			WriteResponseError(&this.Controller, http.StatusUnauthorized)
			return
		}

		u = models.User{
			Username: "SIREL",
			Rol:      models.RolSuperadmin,
		}
	}

	// Encrypt credentials
	crypto := GetCrypto()
	s, e := crypto.Encrypt(&cred)
	if e != nil {
		beego.Error("crypto.Encrypt(", cred.User, "):", e.Error())
		WriteResponseError(&this.Controller, http.StatusInternalServerError)
		return
	}

	this.Data["json"] = Session{
		Username: u.Username,
		Rol:      u.Rol,
		Token:    s,
	}

	this.ServeJSON()
}
*/
