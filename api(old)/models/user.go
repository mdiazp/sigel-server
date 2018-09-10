package models

import (
	"github.com/astaxie/beego/validation"
)

type User struct {
	Id                       int    `json:"id"`
	Username                 string `json:"username" valid:"Required;MaxSize(100)"`
	Name                     string `json:"name" valid:"Required;MaxSize(100)"`
	Email                    string `json:"email" valid:"Required;MaxSize(100)"`
	SendNotificationsToEmail bool   `json:"send_notifications_to_email"`
	Rol                      string `json:"rol"`
	Enable                   bool   `json:"enable"`
}

func (this *User) Valid(v *validation.Validation) {
	// Validation is only doed to post or put user
	// then only roles of Admin and Superadmin are valids,
	// because Superadmin never have to be created
	rp := rolPriority(this.Rol)
	if !(0 <= rp && rp <= 1) {
		// Set error messages of Name by SetError and HasErrors will return true
		v.SetError("Rol", "Rol contain invalid value")
	}
}

func (this *User) HaveRol(rol string) bool {
	p1 := rolPriority(rol)
	p2 := rolPriority(this.Rol)

	return p1 <= p2
}

const (
	RolSuperadmin string = "Superadmin"
	RolAdmin      string = "Admin"
	RolUser       string = "User"
)

var roltypes []string = []string{RolUser, RolAdmin, RolSuperadmin}

// RolPriority Return RolPrioity value or -1 if the rol is invalid
func rolPriority(rol string) int {
	for i, r := range roltypes {
		if rol == r {
			return i
		}
	}
	return -1
}
