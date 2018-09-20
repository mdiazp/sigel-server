package public_controllers

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/controllers"
	"github.com/mdiazp/sirel-server/api/models"
)

type LoginController struct {
	controllers.BaseController
}

func (c *LoginController) URLMapping() {
	c.Mapping("/login", c.Login)
}

// @Title Login
// @Summary Login
// @Description Open session for user into the system
// @Param	credentials		body	public_controllers.Credentials	true		"Credentials for authentication"
// @Success 200 {object} public_controllers.Session
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 500 Internal Server Error
// @Accept json
// @router /login [post]
func (this *LoginController) Login() {
	var (
		e error
		u models.User
	)
	// Read Credentials
	cred := Credentials{}
	this.ReadInputBody(&cred)

	if cred.Username != "SIREL" {
		// Ldap Authentication
		authp := app.AuthProvider(beego.AppConfig.String("AUTH_PROVIDER"))
		e = authp.Authenticate(cred.Username, cred.Password)
		if e != nil {
			beego.Info("authprovider.Authenticate() error: ", e.Error())
			this.WE(e, 401)
		}

		// query for user records in AD and compare usernames to
		// check username diferences
		urecords, e := authp.GetUserRecords(cred.Username)
		if e != nil {
			beego.Error("ldap.FullRecordError: ", e.Error())
			this.WE(e, 500)
		}
		if urecords.Username != cred.Username {
			this.WE(e, 400)
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
				this.WE(e, 500)
			}
		} else if e != nil {
			beego.Error("Unexpected error occurred: ", e.Error())
			this.WE(e, 500)
		}

		// Check user enable
		if !u.Enable {
			this.WE(errors.New("401: user disabled"), 401)
		}
	} else {
		if beego.AppConfig.String("SIREL_PASSWORD") != cred.Password {
			this.WE(errors.New("401: wrong credentials"), 401)
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
		this.WE(e, 500)
	}

	// Prepare response
	this.Data["json"] = Session{
		Username: u.Username,
		Rol:      u.Rol,
		Token:    s,
	}

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
