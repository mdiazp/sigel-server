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
		u := m.NewUser()
		u.UserInfo = ui
		e = m.Create(u)
	}
	c.WE(e, 500)
	return u, e
}

// GetUsers ...
func (c *BaseUsersController) GetUsers() *[]*models.User {
	prefixFilter := c.ReadString("prefixFilter")
	limit := c.ReadInt("limit")
	offset := c.ReadInt("offset")
	orderby := c.ReadString("orderby")
	desc := c.ReadBool("ordDesc")

	data, e := app.Model().GetUsers(prefixFilter, limit, offset, orderby, desc)
	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)

	return data
}
