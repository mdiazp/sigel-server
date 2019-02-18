package controllers

import (
	"fmt"

	"github.com/mdiazp/sigel-server/api/app"
	"github.com/mdiazp/sigel-server/api/models"
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

	c.checkLocalUniqueNameConstraint(l, true)

	e := app.Model().Create(l)
	c.WE(e, 500)
	return l
}

// Update ...
func (c *BaseLocalsController) Update() *models.Local {
	var el EditLocalInfo
	c.ReadObjectInBody("Local", &el, true)

	l := app.Model().NewLocal()
	l.ID = *c.ReadInt("local_id", true)

	l.Load()

	el.copyToLocal(l)
	c.Validate(l)

	c.checkLocalUniqueNameConstraint(l, false)

	e := l.Update()

	fmt.Println(l)

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

// Count ...
func (c *BaseLocalsController) Count() int {
	f := c.ReadLocalsFilter()

	count, e := app.Model().GetLocalsCount(f)

	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
	return count
}

// List ...
func (c *BaseLocalsController) List() *models.LocalCollection {
	limit, offset, orderby, desc := c.ReadPagOrder()

	locals, e := app.Model().GetLocals(
		c.ReadLocalsFilter(),
		limit, offset, orderby, desc,
	)
	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
	return locals
}

// ReadLocalsFilter ...
func (c *BaseLocalsController) ReadLocalsFilter() models.LocalFilter {
	enableToReserve := c.ReadBool("enable_to_reserve")
	areaID := c.ReadInt("area_id")
	search := c.ReadString("search")
	ofAdmin := c.ReadBool("ofAdmin")

	var userID *int

	if ofAdmin != nil && *ofAdmin {
		userID = &c.GetAuthor().ID
	}

	return models.LocalFilter{
		EnableToReserve: enableToReserve,
		AreaID:          areaID,
		Search:          search,
		AdminID:         userID,
	}
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

func (c *BaseLocalsController) checkLocalUniqueNameConstraint(l *models.Local, create bool) {
	// Checking unique name constraint
	lx := app.Model().NewLocal()
	e := app.Model().RetrieveOne(lx, "name=$1", l.Name)
	if e != models.ErrNoRows {
		if e == nil && (create || lx.ID != l.ID) {
			c.WE(fmt.Errorf("Name is already taked"), 400)
		}
		c.WE(e, 500)
	}
}

// EditLocalInfo ...
type EditLocalInfo struct {
	ID                      int
	Name                    string
	Description             string
	Location                string
	WorkingMonths           string
	WorkingWeekDays         string
	WorkingBeginTimeHours   int
	WorkingBeginTimeMinutes int
	WorkingEndTimeHours     int
	WorkingEndTimeMinutes   int
	EnableToReserve         bool
}

func (el *EditLocalInfo) copyToLocal(l *models.Local) {
	l.Name = el.Name
	l.Description = el.Description
	l.Location = el.Location
	l.WorkingMonths = el.WorkingMonths
	l.WorkingWeekDays = el.WorkingWeekDays
	l.WorkingBeginTimeHours = el.WorkingBeginTimeHours
	l.WorkingBeginTimeMinutes = el.WorkingBeginTimeMinutes
	l.WorkingEndTimeHours = el.WorkingEndTimeHours
	l.WorkingEndTimeMinutes = el.WorkingEndTimeMinutes
	l.EnableToReserve = el.EnableToReserve
}
