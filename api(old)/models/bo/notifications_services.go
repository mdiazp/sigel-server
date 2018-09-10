package bo

import (
	"github.com/astaxie/beego/orm"
	"gitlab.com/manuel.diaz/sirel/server/api/models"
)

var nzero = models.Notification{}

func (this *Model) CreateNotification(n models.Notification) (models.Notification, error) {
	n.Id = 0

	_, e := this.orm.Insert(&n)
	if e != nil {
		return nzero, e
	}

	return n, nil
}

func (this *Model) GetNotificationById(id int) (models.Notification, error) {
	l := models.Notification{}
	e := this.orm.QueryTable(&models.Notification{}).Filter("id", id).Limit(1).One(&l)

	if e == orm.ErrNoRows {
		return nzero, models.ErrResultNotFound
	}
	if e != nil {
		return nzero, e
	}

	return l, nil
}

func (this *Model) UpdateNotification(n models.Notification) (models.Notification, error) {
	// Check that notification exist
	n, e := this.GetNotificationById(n.Id)
	if e != nil {
		return nzero, e
	}

	_, e = this.orm.Update(&n)

	if e != nil {
		return nzero, e
	}

	return n, nil
}

func (this *Model) DeleteNotification(id int) error {
	num, e := this.orm.Delete(&models.Notification{Id: id})
	if e == nil {
		if num == 0 {
			return models.ErrResultNotFound
		}
		return nil
	}
	return e
}

func (this *Model) GetNotificationQuerySeter() models.NotificationQuerySeter {
	return this.orm.QueryTable(&models.Notification{})
}
