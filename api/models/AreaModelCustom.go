package models

// AreaCustomModel ...
type AreaCustomModel interface {
	GetAreas(search *string, limit, offset *int, orderby *string,
		orderDesc *bool) (*AreaCollection, error)
}

// GetAreas ...
func (m *model) GetAreas(search *string, limit, offset *int, orderby *string,
	orderDesc *bool) (*AreaCollection, error) {

	if search != nil {
		*search = "name ilike '%" + *search + "%'"
	}

	collection := m.NewAreaCollection()
	e := m.RetrieveCollection(search, limit, offset, orderby, orderDesc, collection)
	return collection, e
}
