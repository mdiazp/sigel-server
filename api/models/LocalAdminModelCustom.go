package models

// LocalAdminCustomModel ...
type LocalAdminCustomModel interface {
	GetLocalAdmin(localID, userID int) (*LocalAdmin, error)
	AddLocalAdmin(localID, userID int) (*LocalAdmin, error)
	DeleteLocalAdmin(localID, userID int) error
}

// GetLocalAdmin ...
func (m *model) GetLocalAdmin(localID, userID int) (*LocalAdmin, error) {
	la := m.NewLocalAdmin()
	e := m.RetrieveOne(la,
		"local_admin.local_id=$1 AND local_admin.user_id=$2", localID, userID)
	return la, e
}

// AddLocalAdmin ...
func (m *model) AddLocalAdmin(localID, userID int) (*LocalAdmin, error) {
	la, e := m.GetLocalAdmin(localID, userID)
	if e != ErrNoRows {
		if e == nil {
			return la, e
		}
		return nil, e
	}
	la = m.NewLocalAdmin()
	la.LocalID = localID
	la.UserID = userID
	e = m.Create(la)
	return la, e
}

// DeleteLocalAdmin ...
func (m *model) DeleteLocalAdmin(localID, userID int) error {
	la, e := m.GetLocalAdmin(localID, userID)
	if e == nil {
		e = m.Delete(la)
	}
	return e
}
