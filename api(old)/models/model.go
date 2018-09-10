package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

var (
	ErrResultNotFound                   = errors.New("Result not found")
	ErrUserWithSameUsernameAlreadyExist = errors.New("User with same username already exist")
)

func ErrNotImplementet(e error) bool {
	switch e {
	case ErrUserWithSameUsernameAlreadyExist:
		return false
	case ErrResultNotFound:
		return false
	case nil:
		return false
	default:
		return true
	}
}

type Model interface {
	GetUserServices() IdentityServices
	GetAreaServices() IdentityServices
	GetLocalServices() IdentityServices
	GetReservationServices() IdentityServices
	GetNotificationServices() IdentityServices
}

type IdentityServices interface {
}

type QuerySeter interface {
	orm.QuerySeter
}
