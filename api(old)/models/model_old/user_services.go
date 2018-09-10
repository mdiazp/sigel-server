package models

type UserServices interface {
	CreateUser(u User) (User, error)
	GetUserById(id int) (User, error)
	GetUserByUsername(username string) (User, error)
	UpdateUser(nu User) (User, error)
	DeleteUser(id int) error
	GetUserQuerySeter() UserQuerySeter
}
