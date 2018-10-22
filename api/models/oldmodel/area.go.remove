package models

type Area struct {
	Id          int    `json:"id"`
	Name        string `json:"name" valid:"Required;MaxSize(100)"`
	Description string `json:"description" valid:"Required;MaxSize(1024)"`
	Location    string `json:"location" valid:"Required;MaxSize(1024)"`
}
