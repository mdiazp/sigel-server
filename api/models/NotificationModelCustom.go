package models

import (
	"fmt"
	"time"
)

// NotificationCustomModel ...
type NotificationCustomModel interface {
	NotificateToUser(userID int, message string) error
	GetNotification(nID, userID int) (*Notification, error)
	GetNotifications(limit, offset *int, orderby *string, desc *bool,
		userID *int, date *Date, readed *bool) (*[]*Notification, error)
	GetNotificationsCount(userID *int, date *Date, readed *bool) (int, error)
	SetUserNotificationsAsReaded(userID int) error
}

func (m *model) NotificateToUser(userID int, message string) error {
	n := m.NewNotification()
	n.UserID = userID
	n.Message = message
	n.CreationTime = time.Now()
	n.Readed = false

	e := m.Create(n)
	return e
}

func (m *model) GetNotification(nID, userID int) (*Notification, error) {
	n := m.NewNotification()
	e := m.RetrieveOne(n, "id=$1 and user_id=$2", nID, userID)
	return n, e
}

func (m *model) GetNotifications(limit, offset *int, orderby *string, desc *bool,
	userID *int, date *Date, readed *bool) (*[]*Notification, error) {

	hf := getNotificationsHorizontalFilter(userID, date, readed)

	ns := m.NewNotificationCollection()
	e := m.RetrieveCollection(hf, limit, offset, orderby, desc, ns)
	return ns.Notifications, e
}

func (m *model) GetNotificationsCount(userID *int, date *Date,
	readed *bool) (int, error) {
	hf := getNotificationsHorizontalFilter(userID, date, readed)

	count := 0
	n := m.NewNotification()
	e := m.RetrieveCount(hf, n, &count)
	return count, e
}

func getNotificationsHorizontalFilter(userID *int, date *Date,
	readed *bool) *string {
	where := ""
	if userID != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("notification.user_id=%d", *userID)
	}
	if readed != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("notification.readed=%t", *readed)
	}
	if date != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("extract(year from reservation.begin_time)=%d AND ", date.Year) +
			fmt.Sprintf("extract(month from reservation.begin_time)=%d AND ", date.Month) +
			fmt.Sprintf("extract(day from reservation.begin_time)=%d", date.Day)
	}

	hf := &where
	if where == "" {
		hf = nil
	}
	return hf
}

func (m *model) SetUserNotificationsAsReaded(userID int) error {
	o := m.NewNotification()
	q := "UPDATE " + o.TableName() + " SET readed=$1 WHERE user_id=$2"

	stmt, e := m.DB().Prepare(q)
	if e == nil {
		defer stmt.Close()
		_, e = stmt.Exec(true, userID)
	}
	return e
}
