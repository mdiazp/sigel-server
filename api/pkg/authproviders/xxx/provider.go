package xxx

import (
	"errors"
	"strings"

	"github.com/mdiazp/sirel-server/api/pkg/authproviders"
)

type provider struct{}

// Authenticate ...
func (p *provider) Authenticate(username, password string) error {
	if users[username] == zero {
		users[username] = authproviders.UserRecords{
			Username: username,
			Email:    username + "@upr.edu.cu",
			Name:     strings.Split(username, ".")[0] + " " + strings.Split(username, ".")[1],
		}
		return nil
		// return errors.New("User not found")
	}
	if password != "123" {
		return errors.New("Fail Authentication")
	}
	return nil
}

func (p *provider) GetUserRecords(username string) (authproviders.UserRecords, error) {
	if users[username] == zero {
		return zero, errors.New("User not found")
	}
	return users[username], nil
}

// GetProvider ...
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
