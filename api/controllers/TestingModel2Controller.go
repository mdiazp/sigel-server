package controllers

import (
	"database/sql"

	"github.com/astaxie/beego"
	"github.com/mdiazp/sirel-server/api/app"
)

// TestingModel2Controller ...
type TestingModel2Controller struct {
	BaseController
}

// CreateArea ...
// @Title Create Area
// @Description Create area
// @Param	area		body	models.AreaInfo	true		"Area"
// @Success 200 {object} models.AreaInfo
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [post]
func (c *TestingModel2Controller) CreateArea() {
	var (
		e error
	)
	model := app.Model()
	a := model.NewArea()

	c.ReadInputBody(a)

	e = model.Create(a)
	if e != nil {
		beego.Debug(e.Error())
		c.WE(e, 500)
	}
	c.Data["json"] = a
	c.ServeJSON()
}

// GetArea ...
// @Title Retrieve Area list
// @Description Get area
// @Param	area_id		query	int	true		"area id"
// @Success 200 {object} models.AreaInfo
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [get]
func (c *TestingModel2Controller) GetArea() {
	var (
		e error
	)
	model := app.Model()

	a := model.NewArea()
	areaID, e := c.GetInt("area_id")
	a.ID = areaID

	e = model.Retrieve(a)

	if e != nil {
		if e == sql.ErrNoRows {
			c.WE(e, 404)
		}
		beego.Debug(e.Error())
		c.WE(e, 500)
	}

	c.Data["json"] = *a
	c.ServeJSON()
}

// UpdateArea ...
// @Title Update Area
// @Description Update area
// @Param	area		body	models.AreaInfo	true		"Area"
// @Success 200 {object} models.AreaInfo
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [patch]
func (c *TestingModel2Controller) UpdateArea() {
	var (
		e error
	)
	model := app.Model()
	a := model.NewArea()

	c.ReadInputBody(a)

	e = a.Update()
	if e != nil {
		if e == sql.ErrNoRows {
			c.WE(e, 404)
		}
		beego.Debug(e.Error())
		c.WE(e, 500)
	}

	c.Data["json"] = a
	c.ServeJSON()
}

// DeleteArea ...
// @Title Delete Area
// @Description Delete area
// @Param	area_id		query	int	true		"area id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [delete]
func (c *TestingModel2Controller) DeleteArea() {
	var (
		e error
	)
	model := app.Model()
	a := model.NewArea()

	a.ID, e = c.GetInt("area_id")
	c.WE(e, 400)

	e = model.Delete(a)
	c.WE(e, 500)

	c.Data["json"] = "OK"
	c.ServeJSON()
}

// Areas ...
// @Title Areas
// @Description Areas
// @Param	searchInName		query	string	false		"searchInName"
// @Param	limit		query	int	false		"limit"
// @Param	offset		query	int	false		"offset"
// @Param	orderby		query	string	false		"orderby"
// @Param	orderDesc		query	string	false		"orderDesc (true or false)"
// @Success 200 {object} []models.AreaInfo
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /areas [get]
func (c *TestingModel2Controller) Areas() {
	var (
		e error
	)
	limit := c.ReadInt("limit")
	offset := c.ReadInt("offset")
	orderby := c.ReadString("orderby")
	desc := c.ReadBool("orderDesc")
	// searchInName := c.ReadString("searchInName")

	model := app.Model()
	// areas, e := model.Areas(limit, offset, orderby, desc)
	areas := model.NewAreaCollection()
	e = model.RetrieveCollection(nil, limit, offset, orderby, desc, areas)
	c.WE(e, 500)

	c.Data["json"] = areas.Areas
	c.ServeJSON()
}
