package models

type Local struct {
	Id              int    `json:"id"`
	AreaId          int    `json:"area_id" valid:"Required"`
	Name            string `json:"name" valid:"Required;MaxSize(100)"`
	Description     string `json:"description" valid:"Required;MaxSize(1024)"`
	Location        string `json:"location" valid:"Required;MaxSize(1024)"`
	EnableToReserve bool   `json:"enable_to_reserve"`
}
