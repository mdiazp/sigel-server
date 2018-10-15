package controllers

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/models"
)

type BaseAreasController struct {
	BaseController
}

func (this *BaseAreasController) Show(container *models.Area) {
	var (
		e error
	)

	id, e := this.GetInt("area_id")
	this.WE(e, 400)

	qs := app.Model().QueryTable(&models.Area{}).Filter("id", id)
	enableToReserve := this.GetString("enable_to_reserve")
	if enableToReserve != "" {
		qs = qs.Filter("enable_to_reserve", enableToReserve)
	}
	e = qs.Limit(1).One(container)
	if e != nil {
		if e == models.ErrResultNotFound {
			this.WE(e, 404)
		}
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseAreasController) Create(container *models.Area) {
	var (
		e error
	)

	this.ReadInputBody(container)

	this.Validate(container)

	_, e = app.Model().Insert(container)
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}

	author := this.GetAuthor()
	if !author.HaveRol(models.RolSuperadmin) {
		this.addAdmin(container.Id, author.Id)
	}
}

func (this *BaseAreasController) Update(container *models.Area) {
	var (
		e error
	)

	id, e := this.GetInt("area_id")
	this.WE(e, 400)

	//save id to prevent that id in body and in path be diferents
	this.ReadInputBody(container)
	container.Id = id
	this.Validate(container)

	_, e = app.Model().Update(container)
	if e == models.ErrResultNotFound {
		this.WE(e, 404)
	}
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseAreasController) Remove() {
	var (
		e error
	)

	id, e := this.GetInt("area_id")
	this.WE(e, 400)

	_, e = app.Model().QueryTable(&models.Area{}).Filter("id", id).Limit(1).Delete()
	if e == models.ErrResultNotFound {
		this.WE(e, 404)
	}
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseAreasController) List(container *[]models.Area) {
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
			Where(fmt.Sprintf("area_admin.user_id=%d", author.Id))
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
		*container = make([]models.Area, 0)
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
