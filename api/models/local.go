package models

type Local struct {
	Id                      int    `json:"id"`
	AreaId                  int    `json:"area_id" valid:"Required"`
	Name                    string `json:"name" valid:"Required;MaxSize(100)"`
	Description             string `json:"description" valid:"Required;MaxSize(1024)"`
	Location                string `json:"location" valid:"Required;MaxSize(1024)"`
	WorkingMonths           string `json:"working_months" valid:"Required;MinSize(12);MaxSize(12)"`
	WorkingWeekDays         string `json:"working_week_days" valid:"Required;MinSize(7);MaxSize(7)"`
	WorkingBeginTimeHours   int    `json:"working_begin_time_hours" valid:"Min(0);Max(23)"`
	WorkingBeginTimeMinutes int    `json:"working_begin_time_minutes" valid:"Min(0);Max(59)"`
	WorkingEndTimeHours     int    `json:"working_end_time_hours" valid:"Min(0);Max(23)"`
	WorkingEndTimeMinutes   int    `json:"working_end_time_minutes" valid:"Min(0);Max(59)"`
	EnableToReserve         bool   `json:"enable_to_reserve"`
}
