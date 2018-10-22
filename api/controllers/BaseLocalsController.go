package controllers

import (
	"fmt"

	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/models"
)

// BaseLocalsController ...
type BaseLocalsController struct {
	BaseController
}

// Show ...
func (c *BaseLocalsController) Show() *models.Local {
	enableToReserve := c.ReadBool("enable_to_reserve")

	l := app.Model().NewLocal()
	l.ID = *c.ReadInt("local_id", true)
	e := l.Load()

	if e == models.ErrNoRows ||
		(enableToReserve != nil && l.EnableToReserve != *enableToReserve) {
		c.WE(e, 404)
	}
	c.WE(e, 500)

	return l
}

// Create ...
func (c *BaseLocalsController) Create() *models.Local {
	l := app.Model().NewLocal()
	c.ReadInputBody(l)
	c.Validate(l)

	c.checkLocalUniqueNameConstraint(l)

	e := app.Model().Create(l)
	c.WE(e, 500)
	return l
}

// Update ...
func (c *BaseLocalsController) Update() *models.Local {
	l := app.Model().NewLocal()
	c.ReadInputBody(l)
	c.Validate(l)
	l.ID = *c.ReadInt("local_id", true)

	c.checkLocalUniqueNameConstraint(l)

	e := l.Update()
	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)
	return l
}

// Remove ...
func (c *BaseLocalsController) Remove() {
	l := app.Model().NewLocal()
	l.ID = *c.ReadInt("local_id", true)
	e := l.Load()

	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)

	app.Model().Delete(l)
}

// List ...
func (c *BaseLocalsController) List() *models.LocalCollection {
	limit := c.ReadInt("limit")
	offset := c.ReadInt("offset")
	orderby := c.ReadString("orderby")
	desc := c.ReadBool("orderDesc")
	enableToReserve := c.ReadBool("enable_to_reserve")
	areaID := c.ReadInt("area_id")
	search := c.ReadString("search")
	ofAdmin := c.ReadBool("ofAdmin")
	var userID *int

	if ofAdmin != nil && *ofAdmin {
		userID = &c.GetAuthor().ID
	}

	locals, e := app.Model().GetLocals(
		areaID, search, enableToReserve, userID,
		limit, offset, orderby, desc)
	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
	return locals
}

// Admins ...
func (c *BaseLocalsController) Admins() *models.UserCollection {
	localID := c.ReadInt("local_id", true)
	admins, e := app.Model().GetLocalAdmins(*localID)
	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)
	return admins
}

// AddAdmin ...
func (c *BaseLocalsController) AddAdmin() *models.LocalAdmin {
	userID := c.ReadInt("user_id", true)
	localID := c.ReadInt("local_id", true)
	la, e := app.Model().AddLocalAdmin(*localID, *userID)
	c.WE(e, 500)
	return la
}

// RemoveAdmin ...
func (c *BaseLocalsController) RemoveAdmin() {
	userID := c.ReadInt("user_id", true)
	localID := c.ReadInt("local_id", true)
	e := app.Model().DeleteLocalAdmin(*localID, *userID)
	if e == models.ErrNoRows {
		return
	}
	c.WE(e, 500)
}

func (c *BaseLocalsController) checkLocalUniqueNameConstraint(l *models.Local) {
	// Checking unique name constraint
	lx := app.Model().NewLocal()
	e := app.Model().RetrieveOne(lx, "name=$1", l.Name)
	if e != models.ErrNoRows {
		if e == nil && lx.ID != l.ID {
			c.WE(fmt.Errorf("Name is already taked"), 400)
		}
		c.WE(e, 500)
	}
}
