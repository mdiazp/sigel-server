package models

import (
	"fmt"
)

// LocalCustomModel ...
type LocalCustomModel interface {
	GetLocals(
		areaID *int, search *string, enableToReserve *bool, adminID *int,
		limit, offset *int, orderby *string, desc *bool) (*LocalCollection, error)
	GetLocalAdmins(localID int) (*UserCollection, error)
}

// GetLocals ...
func (m *model) GetLocals(
	areaID *int, search *string, enableToReserve *bool, adminID *int,
	limit, offset *int, orderby *string, desc *bool) (*LocalCollection, error) {

	var join []*string
	where := ""
	if adminID != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("local_admin.user_id=%d", *adminID)
		tmp := "local_admin ON local.id=local_admin.local_id"
		join = append(join, &tmp)
	}

	if areaID != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("local.area_id=%d", *areaID)
	}

	if enableToReserve != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("local.enable_to_reserve=%t", *enableToReserve)
	}

	if search != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("local.name like '%s'", *search+"%")
	}

	if orderby != nil {
		*orderby = "local." + *orderby
	}

	var hf *string
	if where != "" {
		hf = &where
	}

	locals := m.NewLocalCollection()
	e := m.RetrieveCollection(hf, limit, offset, orderby, desc, locals, join...)
	return locals, e
}

// GetLocalAdmins ...
func (m *model) GetLocalAdmins(localID int) (*UserCollection, error) {
	where := fmt.Sprintf("local_admin.local_id=%d", localID)
	join := "local_admin ON k_user.id=local_admin.user_id"

	admins := m.NewUserCollection()
	e := m.RetrieveCollection(&where, nil, nil, nil, nil, admins, &join)
	return admins, e
}
