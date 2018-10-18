package controllers

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/models"
	"github.com/mdiazp/sirel-server/api/models/models2"
)

// BaseAreasController ...
type BaseAreasController struct {
	BaseController
}

// Show ...
func (this *BaseAreasController) Show(container *models2.Area) {
	var (
		e error
	)

	container.ID = *this.ReadInt("area_id", true)
	beego.Debug("area_id = ", container.ID)
	e = container.Load()

	if e == models2.ErrNoRows {
		this.WE(e, 404)
	}
	this.WE(e, 500)
}

func (this *BaseAreasController) Create(container *models2.Area) {
	var (
		e error
	)

	this.ReadInputBody(container)
	this.Validate(container)

	e = app.Model().Create(container)
	this.WE(e, 500)

	author := this.GetAuthor()
	if !author.HaveRol(models2.RolSuperadmin) {
		this.addAdmin(container.ID, author.ID)
	}
}

func (this *BaseAreasController) Update(container *models2.Area) {
	var e error

	id := this.ReadInt("area_id", true)
	this.ReadInputBody(container)
	this.Validate(container)
	container.ID = *id

	e = container.Update()
	if e == models2.ErrNoRows {
		this.WE(e, 404)
	}
	this.WE(e, 500)
}

func (this *BaseAreasController) Remove() {
	var e error

	id := this.ReadInt("area_id", true)

	model := app.Model()
	area := model.NewArea()
	area.ID = *id
	e = area.Load()

	if e == models2.ErrNoRows {
		this.WE(e, 404)
	}
	this.WE(e, 500)
}

func (this *BaseAreasController) List(container *[]models2.Area) {
	var (
		e error
	)

	qb, e := models.GetQueryBuilder()
	if e != nil {
		beego.Debug(e.Error())
		this.WE(e, 500)
	}

	opt := this.ReadPagAndOrdOptions("id", "id", "name")

	qb = qb.Select("area.id",
		"area.name",
		"area.description",
		"area.location").From("area")

	ofAdmin := this.GetString("ofAdmin")
	if ofAdmin == "true" {
		author := this.GetAuthor()
		qb = qb.InnerJoin("area_admin").
			On("area.id=area_admin.area_id").
			Where(fmt.Sprintf("area_admin.user_id=%d", author.ID))
	}

	fname := this.GetString("search")
	if fname != "" {
		qb = qb.Where(fmt.Sprintf("name__icontains=%s", fname))
	}

	if opt.OrderBy != "" {
		qb = qb.OrderBy("area." + opt.OrderBy)
		if opt.orderDirection == "desc" {
			qb = qb.Desc()
		} else {
			qb = qb.Asc()
		}
	}

	qb = qb.Limit(opt.Limit).Offset(opt.Offset)

	beego.Debug(qb.String())
	cnt, e := app.Model().Raw(qb.String()).QueryRows(container)

	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}

	if e == models.ErrResultNotFound || cnt == 0 {
		*container = make([]models2.Area, 0)
	}

	for _, a := range *container {
		beego.Debug(a.ID)
	}
}

func (this *BaseAreasController) Admins(admins *[]models.UserPublicInfo) {
	var (
		e error
	)

	id, e := this.GetInt("area_id")
	this.WE(e, 400)

	query := models.QueryAreaAdmins
	_, e = app.Model().Raw(query, id).QueryRows(admins)
	if e != nil {
		if e == models.ErrResultNotFound {
			*admins = make([]models.UserPublicInfo, 0)
		} else {
			beego.Error(e.Error())
			this.WE(e, 500)
		}
	}
}

func (this *BaseAreasController) AddAdmin() {
	var (
		e               error
		userID, localID int
	)

	userID, e = this.GetInt("user_id")
	if e == nil {
		localID, e = this.GetInt("area_id")
	}
	this.WE(e, 400)

	//Checking for previous existence
	var maps []orm.Params
	_, e = app.Model().Raw("select area_id, user_id from area_admin where area_id=? and user_id=?",
		localID, userID).Values(&maps)

	if e != nil {
		beego.Error(e)
		this.WE(e, 500)
	}

	if len(maps) > 0 {
		//The admin has been added previously
		this.WE(errors.New("The user already is admin"), 400)
	}
	this.addAdmin(localID, userID)
}

func (this *BaseAreasController) addAdmin(areaID, userID int) {
	rp, e := app.Model().Raw("insert into area_admin(area_id,user_id) values(?,?)").Prepare()
	if e == nil {
		_, e = rp.Exec(areaID, userID)
		rp.Close()
	}

	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseAreasController) RemoveAdmin() {
	var (
		e               error
		userID, localID int
	)

	userID, e = this.GetInt("user_id")
	if e == nil {
		localID, e = this.GetInt("area_id")
	}
	this.WE(e, 400)

	rp, e := app.Model().Raw("delete from area_admin where area_id=? and user_id=?").Prepare()
	_, e = rp.Exec(localID, userID)
	rp.Close()

	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}
