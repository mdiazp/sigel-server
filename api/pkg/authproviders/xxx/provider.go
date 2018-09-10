package xxx

import (
	"errors"

	"gitlab.com/manuel.diaz/sirel/server/api/pkg/authproviders"
)

type provider struct{}

func (this *provider) Authenticate(username, password string) error {
	if users[username] == zero {
		return errors.New("User not found")
	}
	if password != "123" {
		return errors.New("Fail Authentication")
	}
	return nil
}

func (this *provider) GetUserRecords(username string) (authproviders.UserRecords, error) {
	if users[username] == zero {
		return zero, errors.New("User not found")
	}
	return users[username], nil
}

func GetProvider() authproviders.Provider {
	return &provider{}
}

var zero authproviders.UserRecords

var users map[string]authproviders.UserRecords = map[string]authproviders.UserRecords{
	"manuel.diaz": authproviders.UserRecords{
		Username: "manuel.diaz",
		Name:     "Manuel Alejandro Diaz Perez",
		Email:    "manueldiazp92@gmail.com",
	},
	"claupd": authproviders.UserRecords{
		Username: "claupd",
		Name:     "Claudia Permuy Diaz",
		Email:    "clau.permuy@gmail.com",
	},
	"ale": authproviders.UserRecords{
		Username: "ale",
		Name:     "Alejandra Diaz Permuy",
		Email:    "ale@gmail.com",
	},
}
