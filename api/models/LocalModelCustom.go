package models

import (
	"fmt"
)

// LocalCustomModel ...
type LocalCustomModel interface {
	GetLocals(filter LocalFilter, limit, offset *int, orderby *string,
		desc *bool) (*LocalCollection, error)
	GetLocalsCount(filter LocalFilter) (int, error)
	MakeLocalHorizontalFilter(f LocalFilter) (hf *string, join []*string)
	GetLocalAdmins(localID int) (*UserCollection, error)
}

type LocalFilter struct {
	AreaID          *int
	Search          *string
	EnableToReserve *bool
	AdminID         *int
}

func (m *model) GetLocals(filter LocalFilter, limit, offset *int, orderby *string,
	desc *bool) (*LocalCollection, error) {

	if orderby != nil {
		*orderby = "local." + *orderby
	}

	hf, join := m.MakeLocalHorizontalFilter(filter)

	if orderby == nil {
		tmp := "name"
		orderby = &tmp
		tmp2 := false
		desc = &tmp2
	}

	locals := m.NewLocalCollection()
	e := m.RetrieveCollection(hf, limit, offset, orderby, desc, locals, join...)
	return locals, e
}

func (m *model) GetLocalsCount(filter LocalFilter) (int, error) {
	hf, join := m.MakeLocalHorizontalFilter(filter)

	o := m.NewLocal()
	count := 0
	e := m.RetrieveCount(hf, o, &count, join...)
	return count, e
}

// MakeLocalHorizontalFilter ...
func (m *model) MakeLocalHorizontalFilter(f LocalFilter) (hf *string, join []*string) {
	where := ""
	if f.AdminID != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("local_admin.user_id=%d", *(f.AdminID))
		tmp := "local_admin ON local.id=local_admin.local_id"
		join = append(join, &tmp)
	}

	if f.AreaID != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("local.area_id=%d", *(f.AreaID))
	}

	if f.EnableToReserve != nil {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("local.enable_to_reserve=%t", *(f.EnableToReserve))
	}

	if f.Search != nil {
		if where != "" {
			where += " AND "
		}
		where += "local.name ilike '%" + *(f.Search) + "%'"
	}

	if where != "" {
		hf = &where
	}

	return
}

// GetLocalAdmins ...
func (m *model) GetLocalAdmins(localID int) (*UserCollection, error) {
	where := fmt.Sprintf("local_admin.local_id=%d", localID)
	join := "local_admin ON k_user.id=local_admin.user_id"

	orderby := "k_user.username"
	desc := false

	admins := m.NewUserCollection()
	e := m.RetrieveCollection(&where, nil, nil, &orderby, &desc, admins, &join)
	return admins, e
}
