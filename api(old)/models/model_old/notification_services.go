package models

type NotificationServices interface {
	CreateNotification(n Notification) (Notification, error)
	GetNotificationById(id int) (Notification, error)
	UpdateNotification(n Notification) (Notification, error)
	DeleteNotification(id int) error
	GetNotificationQuerySeter() NotificationQuerySeter
}
