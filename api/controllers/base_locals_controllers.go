package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/mdiazp/sirel-server/api/app"
	"github.com/mdiazp/sirel-server/api/models"
)

type BaseLocalsController struct {
	BaseController
}

func (this *BaseLocalsController) Show(container *models.Local) {
	var (
		e error
	)

	id, e := this.GetInt("local_id")
	this.WE(e, 400)

	qs := app.Model().QueryTable(&models.Local{}).Filter("id", id)
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

func (this *BaseLocalsController) Create(container *models.Local) {
	var (
		e error
	)

	this.ReadInputBody(container)

	this.Validate(container)

	//only users that are admins of an area can
	//insert in it any local
	author := this.GetAuthor()

	if !author.HaveRol(models.RolSuperadmin) {
		var tmp []orm.Params
		_, e = app.Model().Raw("select user_id from area_admin "+
			"where area_id=? and user_id=? limit 1 offset 0",
			container.AreaId, author.Id).Values(&tmp)

		if e != nil {
			beego.Error(e)
			this.WE(e, 500)
		}

		if len(tmp) == 0 {
			this.WE(errors.New("only users that are admins of an area can insert in it any local"), 403)
		}
	}

	_, e = app.Model().Insert(container)
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}

	if !author.HaveRol(models.RolSuperadmin) {
		//Add the user as admin of the local
		this.addAdmin(container.Id, author.Id)
	}
}

func (this *BaseLocalsController) Update(container *models.Local) {
	var (
		e error
	)

	id, e := this.GetInt("local_id")
	this.WE(e, 400)

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

func (this *BaseLocalsController) Remove() {
	var (
		e error
	)

	id, e := this.GetInt("local_id")
	this.WE(e, 400)

	_, e = app.Model().QueryTable(&models.Local{}).Filter("id", id).Limit(1).Delete()
	if e == models.ErrResultNotFound {
		this.WE(e, 404)
	}
	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseLocalsController) List(container *[]models.Local) {
	var (
		e error
	)

	qb, e := models.GetQueryBuilder()
	if e != nil {
		beego.Debug(e.Error())
		this.WE(e, 500)
	}

	opt := this.ReadPagAndOrdOptions("id", "id", "name")

	qb = qb.Select("*").From("local")

	ofAdmin := this.GetString("ofAdmin")
	if ofAdmin == "true" {
		author := this.GetAuthor()
		qb = qb.Where("local.id").In(
			fmt.Sprintf(
				"SELECT local_admin.local_id from local_admin "+
					"WHERE local_admin.user_id=%d", author.Id,
			),
		).Or("local.id").In(
			fmt.Sprintf(
				"SELECT local.id from area_admin "+
					"INNER JOIN local ON area_admin.area_id=local.area_id "+
					"WHERE area_admin.user_id=%d", author.Id,
			),
		)
	}

	fname := this.GetString("search")
	if fname != "" {
		qb = qb.Where(fmt.Sprintf("name__icontains=%s", fname))
	}

	tmp := this.GetString("area_id")
	if tmp != "" {
		area_id, e := strconv.Atoi(tmp)
		this.WE(e, 400)
		qb = qb.Where(fmt.Sprintf("area_id=%d", area_id))
	}

	tmp = this.GetString("enable_to_reserve")
	if tmp != "" {
		enable_to_reserve, e := strconv.ParseBool(tmp)
		this.WE(e, 400)
		qb = qb.Where(fmt.Sprintf("local.enable_to_reserve=%t", enable_to_reserve))
	}

	if opt.OrderBy == "" {
		opt.OrderBy = "id"
	}
	if opt.OrderBy != "" {
		qb = qb.OrderBy("local." + opt.OrderBy)
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
		*container = make([]models.Local, 0)
	}
}

func (this *BaseLocalsController) Admins(admins *[]models.UserPublicInfo) {
	var (
		e error
	)

	id, e := this.GetInt("local_id")
	this.WE(e, 400)

	query := models.QueryLocalAdmins
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

func (this *BaseLocalsController) AddAdmin() {
	var (
		e               error
		userID, localID int
	)

	userID, e = this.GetInt("user_id")
	if e == nil {
		localID, e = this.GetInt("local_id")
	}
	this.WE(e, 400)

	//Checking for previous existence
	var maps []orm.Params
	_, e = app.Model().Raw("select local_id, user_id from local_admin where local_id=? and user_id=? limit 1 offset 0",
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

func (this *BaseLocalsController) addAdmin(localID, userID int) {
	rp, e := app.Model().Raw("insert into local_admin(local_id,user_id) values(?,?)").Prepare()
	if e == nil {
		_, e = rp.Exec(localID, userID)
		rp.Close()
	}

	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}

func (this *BaseLocalsController) RemoveAdmin() {
	var (
		e error
	)

	id, e := this.GetInt("local_id")
	this.WE(e, 400)

	userID, e := this.GetInt("user_id")
	this.WE(e, 400)

	rp, e := app.Model().Raw("delete from local_admin where local_id=? and user_id=?").Prepare()
	_, e = rp.Exec(id, userID)
	rp.Close()

	if e != nil {
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}
