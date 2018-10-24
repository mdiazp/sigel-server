package models

import (
	"time"

	"github.com/mdiazp/kmodel"
)

///////////////////////////////////////////////////////////////////////////////////

// NotificationInfo ...
type NotificationInfo struct {
	ID           int
	UserID       int
	Message      string
	CreationTime time.Time
	Readed       bool
}

// Notification ...
type Notification struct {
	NotificationInfo
	model Model

	user *User
}

/////////////////////////////////////////////////////

// TableName ...
func (n *Notification) TableName() string {
	return "notification"
}

// AutoPKey ...
func (n *Notification) AutoPKey() bool {
	return true
}

// PkeyName ...
func (n *Notification) PkeyName() string {
	return "id"
}

// PkeyValue ...
func (n *Notification) PkeyValue() interface{} {
	return n.ID
}

// PkeyPointer ...
func (n *Notification) PkeyPointer() interface{} {
	return &n.ID
}

// ColumnNames ...
func (n *Notification) ColumnNames() []string {
	return []string{
		"user_id",
		"message",
		"creation_time",
		"readed",
	}
}

// ColumnValues ...
func (n *Notification) ColumnValues() []interface{} {
	return []interface{}{
		n.UserID,
		n.Message,
		n.CreationTime,
		n.Readed,
	}
}

// ColumnPointers ...
func (n *Notification) ColumnPointers() []interface{} {
	return []interface{}{
		&n.UserID,
		&n.Message,
		&n.CreationTime,
		&n.Readed,
	}
}

/////////////////////////////////////////////////////

// Update ...
func (n *Notification) Update() error {
	return n.model.Update(n)
}

// Load ...
func (n *Notification) Load() error {
	return n.model.Retrieve(n)
}

// User ...
func (n *Notification) User() (*User, error) {
	var e error
	if n.user == nil {
		n.user = n.model.NewUser()
		n.user.ID = n.UserID
		e = n.model.Retrieve(n.user)
	}
	return n.user, e
}

///////////////////////////////////////////////////////////////////////////////////

// NotificationCollection ...
type NotificationCollection struct {
	model         Model
	Notifications *[]*Notification
}

// NewObjectModel ...
func (c *NotificationCollection) NewObjectModel() kmodel.ObjectModel {
	return c.model.NewNotification()
}

// Add ...
func (c *NotificationCollection) Add() kmodel.ObjectModel {
	n := c.model.NewNotification()
	*(c.Notifications) = append(*(c.Notifications), n)
	return n
}

///////////////////////////////////////////////////////////////////////////////////

// NotificationModel ...
type NotificationModel interface {
	NewNotification() *Notification
	NewNotificationCollection() *NotificationCollection
	Notifications(limit, offset *int, orderby *string,
		orderDesc *bool) (*NotificationCollection, error)

	NotificationCustomModel
}

/////////////////////////////////////////////////////

// NewNotification ...
func (m *model) NewNotification() *Notification {
	n := &Notification{
		model: m,
	}
	return n
}

// NewNotificationCollection ...
func (m *model) NewNotificationCollection() *NotificationCollection {
	kk := make([]*Notification, 0)
	return &NotificationCollection{
		model:         m,
		Notifications: &kk,
	}
}

func (m *model) Notifications(limit, offset *int, orderby *string,
	orderDesc *bool) (*NotificationCollection, error) {

	collection := m.NewNotificationCollection()
	e := m.RetrieveCollection(nil, limit, offset, orderby, orderDesc, collection)
	return collection, e
}
