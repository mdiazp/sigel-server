package controllers

/*
import (
	"github.com/astaxie/beego"

	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

type AreasBaseController struct {
	beego.Controller
}

func (this *AreasBaseController) get() models.Area {
	pthis := &this.Controller
	id, e := this.GetInt("id")
	if e != nil {
		wre(pthis, 400)
	}

	a, e := AppModel.GetAreaById(id)
	if e != nil {
		if e == models.ErrResultNotFound {
			wre(pthis, 404)
		}
		beego.Error(e.Error())
		wre(pthis, 500)
	}
	return a
}

func (this *AreasBaseController) list() []models.Area {
	pthis := &this.Controller
	qs := AppModel.GetAreaQuerySeter()

	if enable_to_reserve, e := this.GetBool("enable_to_reserve"); e == nil {
		qs.Filter("enable_to_reserve", enable_to_reserve)
	}

	opt := ReadPagAndOrdOptions(pthis)

	qs = qs.Limit(opt.Limit).Offset(opt.Offset)

	if opt.OrderBy != "" {
		qs = qs.OrderBy(fmtorder(&opt))
	}

	var areas []models.Area
	_, e := qs.All(&areas)

	if e != nil {
		beego.Error(e.Error())
		wre(pthis, 500)
	}

	return areas
}
*/
