package models

import "time"

type Reservation struct {
	Id                  int       `json:"id"`
	UserId              int       `json:"user_id" valid:"Required"`
	LocalId             int       `json:"local_id" valid:"Required"`
	ActivityName        string    `json:"activity_name" valid:"Required;MaxSize(100)"`
	ActivityDescription string    `json:"activity_description" valid:"Required;MaxSize(1024)"`
	BeginTime           time.Time `json:"begin_time" valid:"Required;"`
	EndTime             time.Time `json:"end_time" valid:"Rquired"`
	Confirmed           bool      `json:"confirmed"`
	Pending             bool      `json:"pending"`
}
