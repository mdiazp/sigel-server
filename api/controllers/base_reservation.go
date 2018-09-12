package controllers

import (
	"strconv"

	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/app"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type ReservationBaseController struct {
	BaseController
}

func (this *ReservationBaseController) Show(container models.Reservation) {
	var (
		e error
	)

	id, e := this.GetInt("id")
	this.WE(e, 400)

	qs := app.Model().QueryTable(&models.Reservation{}).Filter("id", id)
	if e != nil {
		if e == models.ErrResultNotFound {
			this.WE(e, 404)
		}
		beego.Error(e.Error())
		this.WE(e, 500)
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

func (this *ReservationBaseController) List(container *[]models.Reservation) {
	var (
		e   error
		tmp string
	)

	qs := app.Model().QueryTable(&models.Reservation{})

	opt := this.ReadPagAndOrdOptions()
	qs = qs.Limit(opt.Limit).Offset(opt.Offset)
	if opt.OrderBy == "" {
		opt.OrderBy = "id"
	}
	if opt.OrderBy != "" {
		qs = qs.OrderBy(this.Fmtorder(&opt))
	}

	tmp = this.GetString("user_id")
	if tmp != "" {
		user_id, e := strconv.Atoi(tmp)
		this.WE(e, 400)
		qs = qs.Filter("user_id", user_id)
	}

	tmp = this.GetString("local_id")
	if tmp != "" {
		local_id, e := strconv.Atoi(tmp)
		this.WE(e, 400)
		qs = qs.Filter("local", local_id)
	}

	activity_name := this.GetString("activity_name")
	if activity_name != "" {
		qs = qs.Filter("activity_name__icontains", activity_name)
	}

	tmp = this.GetString("confirmed")
	if tmp != "" {
		confirmed, e := strconv.ParseBool(tmp)
		this.WE(e, 400)
		qs = qs.Filter("confirmed", confirmed)
	}

	tmp = this.GetString("pending")
	if tmp != "" {
		pending, e := strconv.ParseBool(tmp)
		this.WE(e, 400)
		qs = qs.Filter("pending", pending)
	}

	fname := this.GetString("fname")
	if fname != "" {
		qs = qs.Filter("name__icontains", fname)
	}

	_, e = qs.All(container)

	if e != nil {
		if e == models.ErrResultNotFound {
			this.WE(e, 404)
		}
		beego.Error(e.Error())
		this.WE(e, 500)
	}
}
