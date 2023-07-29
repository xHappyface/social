package ports

import "github.com/xHappyface/social/core/user"

// define a port for that object w/ functionality that needs fulfilling (CRUD functionality in this case)

type UserRepository interface {
	Create(*user.User) error
	ReadByID(string) (*user.User, error)
	ReadByName(string) (*user.User, error)
	Update(*user.User) error
	DeleteByID(string) error
}
