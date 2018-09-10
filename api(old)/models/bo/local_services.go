package bo

import (
	"github.com/astaxie/beego/orm"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

var lzero = models.Local{}

func (this *Model) CreateLocal(l models.Local) (models.Local, error) {
	l.Id = 0

	_, e := this.orm.Insert(&l)
	if e != nil {
		return lzero, e
	}

	return l, nil
}

func (this *Model) GetLocalById(id int) (models.Local, error) {
	l := models.Local{}
	e := this.orm.QueryTable(&models.Local{}).Filter("id", id).Limit(1).One(&l)

	if e == orm.ErrNoRows {
		return lzero, models.ErrResultNotFound
	}
	if e != nil {
		return lzero, e
	}

	return l, nil
}

func (this *Model) UpdateLocal(l models.Local) (models.Local, error) {
	// Check that local exist
	l, e := this.GetLocalById(l.Id)
	if e != nil {
		return lzero, e
	}

	_, e = this.orm.Update(&l)

	if e != nil {
		return lzero, e
	}

	return l, nil
}

func (this *Model) DeleteLocal(id int) error {
	num, e := this.orm.Delete(&models.Local{Id: id})
	if e == nil {
		if num == 0 {
			return models.ErrResultNotFound
		}
		return nil
	}
	return e
}

func (this *Model) GetLocalQuerySeter() models.LocalQuerySeter {
	return this.orm.QueryTable(&models.Local{})
}
