package models2

import "github.com/astaxie/beego/validation"

// Valid ...
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

// HaveRol ...
func (u *User) HaveRol(rol string) bool {
	p1 := rolPriority(rol)
	p2 := rolPriority(u.Rol)

	return p1 <= p2
}

const (
	// RolSuperadmin ...
	RolSuperadmin string = "Superadmin"
	// RolAdmin ...
	RolAdmin string = "Admin"
	// RolUser ...
	RolUser string = "User"
)

var roltypes = []string{RolUser, RolAdmin, RolSuperadmin}

// RolPriority Return RolPrioity value or -1 if the rol is invalid
func rolPriority(rol string) int {
	for i, r := range roltypes {
		if rol == r {
			return i
		}
	}
	return -1
}
