package models

import "github.com/astaxie/beego/validation"

// AreaCustomModel ...
type AreaCustomModel interface {
	GetAreas(filter AreaFilter, limit, offset *int, orderby *string,
		desc *bool) (*AreaCollection, error)
	GetAreasCount(filter AreaFilter) (int, error)
}

// AreaFilter ...
type AreaFilter struct {
	Search *string
}

func (m *model) GetAreas(filter AreaFilter, limit, offset *int, orderby *string,
	desc *bool) (*AreaCollection, error) {

	if orderby != nil {
		*orderby = "area." + *orderby
	}

	hf, join := m.MakeAreaHorizontalFilter(filter)

	if orderby == nil {
		tmp := "name"
		orderby = &tmp
		tmp2 := false
		desc = &tmp2
	}

	areas := m.NewAreaCollection()
	e := m.RetrieveCollection(hf, limit, offset, orderby, desc, areas, join...)
	return areas, e
}

func (m *model) GetAreasCount(filter AreaFilter) (int, error) {
	hf, join := m.MakeAreaHorizontalFilter(filter)

	o := m.NewArea()
	count := 0
	e := m.RetrieveCount(hf, o, &count, join...)
	return count, e
}

// MakeAreaHorizontalFilter ...
func (m *model) MakeAreaHorizontalFilter(f AreaFilter) (hf *string, join []*string) {
	where := ""
	if f.Search != nil {
		if where != "" {
			where += " AND "
		}
		where += "area.name ilike '%" + *(f.Search) + "%'"
	}

	if where != "" {
		hf = &where
	}

	return
}

// Valid ...
func (a *AreaInfo) Valid(v *validation.Validation) {
	validateNotEmptyString("name", a.Name, v)
	validateNotEmptyString("description", a.Description, v)
	validateNotEmptyString("location", a.Location, v)
}
