package controllers

/*
import (
	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type ReservationBaseController struct {
	beego.Controller
}

func (this *ReservationBaseController) get() models.Reservation {
	pthis := &this.Controller
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}

	r, e := AppModel.GetReservationById(id)
	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	return r
}

func (this *ReservationBaseController) list() []models.Reservation {
	pthis := &this.Controller
	qs := AppModel.GetReservationQuerySeter()

	opt := ReadPagAndOrdOptions(pthis)

	qs = qs.Limit(opt.Limit).Offset(opt.Offset)
	if opt.OrderBy != "" {
		qs = qs.OrderBy(fmtorder(&opt))
	}

	var rvs []models.Reservation
	_, e := qs.All(&rvs)

	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}

	return rvs
}
*/
