package bo

import (
	"github.com/astaxie/beego/orm"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

var azero = models.Area{}

func (this *Model) CreateArea(a models.Area) (models.Area, error) {
	a.Id = 0

	_, e := this.orm.Insert(&a)
	if e != nil {
		return azero, e
	}

	return a, nil
}

func (this *Model) GetAreaById(id int) (models.Area, error) {
	a := models.Area{}
	e := this.orm.QueryTable(&models.Area{}).Filter("id", id).Limit(1).One(&a)

	if e == orm.ErrNoRows {
		return azero, models.ErrResultNotFound
	}
	if e != nil {
		return azero, e
	}

	return a, nil
}

func (this *Model) UpdateArea(a models.Area) (models.Area, error) {
	// Check that area exist
	a, e := this.GetAreaById(a.Id)
	if e != nil {
		return azero, e
	}

	_, e = this.orm.Update(&a)

	if e != nil {
		return azero, e
	}

	return a, nil
}

func (this *Model) DeleteArea(id int) error {
	num, e := this.orm.Delete(&models.Area{Id: id})
	if e == nil {
		if num == 0 {
			return models.ErrResultNotFound
		}
		return nil
	}
	return e
}

func (this *Model) GetAreaQuerySeter() models.AreaQuerySeter {
	return this.orm.QueryTable(&models.Area{})
}
