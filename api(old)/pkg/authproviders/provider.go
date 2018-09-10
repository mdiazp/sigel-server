package authproviders

type UserRecords struct {
	Username string
	Name     string
	Email    string
}

type Provider interface {
	Authenticate(username, password string) (e error)
	GetUserRecords(username string) (UserRecords, error)
}
