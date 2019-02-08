package public

import (
	"errors"
	"fmt"

	"github.com/mdiazp/sirel-server/api/models"

	"github.com/astaxie/beego"
	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/controllers"
)

// LoginController ...
type LoginController struct {
	controllers.BaseUsersController
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("/login", c.Login)
}

// Login ...
// @Title Login
// @Summary Login
// @Description Open session for user into the system
// @Param	credentials		body	public.Credentials	true		"Credentials for authentication"
// @Success 200 {object} public.Session
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 500 Internal Server Error
// @Accept json
// @router /login [post]
func (c *LoginController) Login() {
	// Read Credentials
	cred := Credentials{}
	c.ReadInputBody(&cred)

	var (
		e error
		u *models.User
	)

	if cred.Username != "SIREL" {
		// Ldap Authentication
		authp := app.AuthProvider(beego.AppConfig.String("AUTH_PROVIDER"))
		e = authp.Authenticate(cred.Username, cred.Password)
		c.WE(e, 401)

		// query for user records in AD and compare usernames to
		// check username diferences
		urecords, e := authp.GetUserRecords(cred.Username)
		c.WE(e, 500)
		if urecords.Username != cred.Username {
			c.WE(fmt.Errorf("Usuario incorrecto"), 400)
		}
		cred.Username = urecords.Username

		// Register user if they is not yet
		u, e = c.Register(
			models.UserInfo{
				Username: urecords.Username,
				Name:     urecords.Name,
				Email:    urecords.Email,
				Rol:      models.RolUser,
				Enable:   true,
				SendNotificationsToEmail: true,
			},
		)

		// Check if user is enable
		if !u.Enable {
			c.WE(errors.New("disabled user"), 401)
		}
	} else {
		if beego.AppConfig.String("SIREL_PASSWORD") != cred.Password {
			c.WE(errors.New("wrong credentials"), 401)
		}
		u = app.Model().NewUser()
		u.UserInfo = models.UserInfo{
			Username: "SIREL",
			Rol:      models.RolSuperadmin,
			Enable:   true,
		}
	}

	// Encrypt credentials
	s, e := app.Crypto().Encrypt(cred.Username)
	c.WE(e, 500)

	// Prepare response
	c.Data["json"] = Session{
		UserID:   u.ID,
		Username: u.Username,
		Rol:      u.Rol,
		Token:    s,
	}

	c.ServeJSON()
}

// Session ...
type Session struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	Rol      string `json:"rol"`
	Token    string `json:"jwtToken"`
}

// Credentials ...
type Credentials struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}
