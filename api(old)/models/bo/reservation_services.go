package bo

import (
	"github.com/astaxie/beego/orm"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

var rzero = models.Reservation{}

func (this *Model) CreateReservation(r models.Reservation) (models.Reservation, error) {
	r.Id = 0

	_, e := this.orm.Insert(&r)
	if e != nil {
		return rzero, e
	}

	return r, nil
}

func (this *Model) GetReservationById(id int) (models.Reservation, error) {
	r := models.Reservation{}
	e := this.orm.QueryTable(&models.Reservation{}).Filter("id", id).Limit(1).One(&r)

	if e == orm.ErrNoRows {
		return rzero, models.ErrResultNotFound
	}
	if e != nil {
		return rzero, e
	}

	return r, nil
}

func (this *Model) UpdateReservation(l models.Reservation) (models.Reservation, error) {
	// Check that reservation exist
	r, e := this.GetReservationById(l.Id)
	if e != nil {
		return rzero, e
	}

	_, e = this.orm.Update(&r)

	if e != nil {
		return rzero, e
	}

	return r, nil
}

func (this *Model) DeleteReservation(id int) error {
	num, e := this.orm.Delete(&models.Reservation{Id: id})
	if e == nil {
		if num == 0 {
			return models.ErrResultNotFound
		}
		return nil
	}
	return e
}

func (this *Model) GetReservationQuerySeter() models.ReservationQuerySeter {
	return this.orm.QueryTable(&models.Reservation{})
}
