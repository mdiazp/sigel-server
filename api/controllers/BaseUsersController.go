package controllers

import (
	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/models"
)

// BaseUsersController ...
type BaseUsersController struct {
	BaseController
}

// GetUser ...
func (c *BaseUsersController) GetUser() *models.User {
	userID := c.ReadInt("user_id", true)

	u := app.Model().NewUser()
	u.ID = *userID
	e := u.Load()
	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)
	return u
}

// Register ...
func (c *BaseUsersController) Register(ui models.UserInfo) (*models.User, error) {
	m := app.Model()
	u, e := m.GetUser(ui.Username)

	// Register if user don't exits
	if e == models.ErrNoRows {
		u = m.NewUser()
		u.UserInfo = ui
		e = m.Create(u)
	}
	c.WE(e, 500)
	return u, e
}

// GetUsers ...
func (c *BaseUsersController) GetUsers() *[]*models.User {
	username := c.ReadString("username")
	name := c.ReadString("name")
	email := c.ReadString("email")
	rol := c.ReadString("rol")
	enable := c.ReadBool("enable")
	limit, offset, orderby, desc := c.ReadPagOrder()

	data, e := app.Model().GetUsers(username, name, email, rol, enable, limit, offset, orderby, desc)
	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)

	return data
}
