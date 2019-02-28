package controllers

import (
	"fmt"

	"github.com/mdiazp/sigel-server/api/app"
	"github.com/mdiazp/sigel-server/api/models"
)

// BaseAreasController ...
type BaseAreasController struct {
	BaseController
}

// Show ...
func (c *BaseAreasController) Show() *models.Area {
	a := app.Model().NewArea()
	a.ID = *c.ReadInt("area_id", true)
	e := a.Load()

	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)

	return a
}

// Create ...
func (c *BaseAreasController) Create() *models.Area {
	a := app.Model().NewArea()
	c.ReadInputBody(a)
	c.Validate(a)

	c.checkAreaUniqueNameConstraint(a, true)

	e := app.Model().Create(a)
	c.WE(e, 500)
	return a
}

// Update ...
func (c *BaseAreasController) Update() *models.Area {
	a := app.Model().NewArea()
	c.ReadInputBody(a)
	c.Validate(a)
	a.ID = *c.ReadInt("area_id", true)

	c.checkAreaUniqueNameConstraint(a, false)

	e := a.Update()
	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)
	return a
}

// Remove ...
func (c *BaseAreasController) Remove() {
	a := app.Model().NewArea()
	a.ID = *c.ReadInt("area_id", true)
	e := a.Load()

	if e == models.ErrNoRows {
		c.WE(e, 404)
	}
	c.WE(e, 500)

	app.Model().Delete(a)
}

// List ...
func (c *BaseAreasController) List() *models.AreaCollection {
	limit, offset, orderby, desc := c.ReadPagOrder()

	areas, e := app.Model().GetAreas(
		c.ReadAreasFilter(),
		limit, offset, orderby, desc,
	)
	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
	return areas
}

// Count ...
func (c *BaseAreasController) Count() int {
	f := c.ReadAreasFilter()

	count, e := app.Model().GetAreasCount(f)

	if e != models.ErrNoRows {
		c.WE(e, 500)
	}
	return count
}

// ReadAreasFilter ...
func (c *BaseAreasController) ReadAreasFilter() models.AreaFilter {
	search := c.ReadString("search")

	return models.AreaFilter{
		Search: search,
	}
}

func (c *BaseAreasController) checkAreaUniqueNameConstraint(a *models.Area, create bool) {
	// Checking unique name constraint
	ax := app.Model().NewArea()
	e := app.Model().RetrieveOne(ax, "name=$1", a.Name)

	if e != models.ErrNoRows {
		if e == nil && (create || ax.ID != a.ID) {
			c.WE(fmt.Errorf("Name is already taked"), 400)
		}
		c.WE(e, 500)
	}
}
