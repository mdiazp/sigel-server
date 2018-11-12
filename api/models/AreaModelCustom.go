package models

// AreaCustomModel ...
type AreaCustomModel interface {
	GetAreas(search *string, limit, offset *int, orderby *string,
		desc *bool) (*AreaCollection, error)
}

// GetAreas ...
func (m *model) GetAreas(search *string, limit, offset *int, orderby *string,
	desc *bool) (*AreaCollection, error) {

	if search != nil {
		*search = "name ilike '%" + *search + "%'"
	}
	if orderby == nil {
		tmp := "name"
		orderby = &tmp
		tmp2 := false
		desc = &tmp2
	}

	collection := m.NewAreaCollection()
	e := m.RetrieveCollection(search, limit, offset, orderby, desc, collection)
	return collection, e
}
