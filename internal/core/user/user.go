package user

import (
	"time"
)

// define a core object

type User struct {
	ID          string
	Username    string
	Password    string
	DateCreated time.Time
}

func NewUser(id, username, password string) *User {
	return &User{
		ID:          id,
		Username:    username,
		Password:    password,
		DateCreated: time.Now(),
	}
}

// define a port for that object w/ functionality that needs fulfilling (CRUD functionality in this case)

type UserRepository interface {
	Create(*User) error
	ReadByID(string) (*User, error)
	Update(*User) error
	DeleteByID(string) error
}
