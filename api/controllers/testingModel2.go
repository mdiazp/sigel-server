package controllers

import (
	"database/sql"

	"github.com/astaxie/beego"
	"github.com/mdiazp/sirel-server/api/app"
)

type TestingModel2Controller struct {
	BaseController
}

// @Title Create Area
// @Description Create area
// @Param	area		body	models2.Area	true		"Area"
// @Success 200 {object} models2.Area
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [post]
func (this *TestingModel2Controller) CreateArea() {
	var (
		e error
	)
	model := app.Model()
	a := model.NewArea()

	this.ReadInputBody(a)

	e = model.Create(a)
	if e != nil {
		beego.Debug(e.Error())
		this.WE(e, 500)
	}
	this.Data["json"] = a
	this.ServeJSON()
}

// @Title Retrieve Area list
// @Description Get area
// @Param	area_id		query	int	true		"area id"
// @Success 200 {object} models2.Area
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [get]
func (this *TestingModel2Controller) GetArea() {
	var (
		e error
	)
	model := app.Model()

	a := model.NewArea()
	areaID, e := this.GetInt("area_id")
	a.ID = areaID

	e = model.Retrieve(a)

	if e != nil {
		if e == sql.ErrNoRows {
			this.WE(e, 404)
		}
		beego.Debug(e.Error())
		this.WE(e, 500)
	}

	this.Data["json"] = *a
	this.ServeJSON()
}

// @Title Update Area
// @Description Update area
// @Param	area		body	models2.Area	true		"Area"
// @Success 200 {object} models2.Area
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [patch]
func (this *TestingModel2Controller) UpdateArea() {
	var (
		e error
	)
	model := app.Model()
	a := model.NewArea()

	this.ReadInputBody(a)

	e = a.Update()
	if e != nil {
		if e == sql.ErrNoRows {
			this.WE(e, 404)
		}
		beego.Debug(e.Error())
		this.WE(e, 500)
	}

	this.Data["json"] = a
	this.ServeJSON()
}

// @Title Delete Area
// @Description Delete area
// @Param	area_id		query	int	true		"area id"
// @Success 200 {string}
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /area [delete]
func (this *TestingModel2Controller) DeleteArea() {
	var (
		e error
	)
	model := app.Model()
	a := model.NewArea()

	a.ID, e = this.GetInt("area_id")
	this.WE(e, 400)

	e = model.Delete2(a)
	this.WE(e, 500)

	this.Data["json"] = "OK"
	this.ServeJSON()
}

// @Title Areas
// @Description Areas
// @Param	searchInName		query	string	false		"searchInName"
// @Param	limit		query	int	false		"limit"
// @Param	offset		query	int	false		"offset"
// @Param	orderby		query	string	false		"orderby"
// @Param	orderDesc		query	string	false		"orderDesc (true or false)"
// @Success 200 {object} []Models2.Area
// @Failure 400 Bad request
// @Failure 404 Not Found
// @Failure 500 Internal Server Error
// @Accept json
// @router /areas [get]
func (this *TestingModel2Controller) Areas() {
	var (
		e error
	)
	limit := this.ReadInt("limit")
	offset := this.ReadInt("offset")
	orderby := this.ReadString("orderby")
	desc := this.ReadBool("orderDesc")
	// searchInName := this.ReadString("searchInName")

	model := app.Model()
	// areas, e := model.Areas(limit, offset, orderby, desc)
	areas := model.NewAreaCollection()
	e = model.RetrieveCollection(nil, limit, offset, orderby, desc, areas)
	this.WE(e, 500)

	this.Data["json"] = areas.Areas
	this.ServeJSON()
}
