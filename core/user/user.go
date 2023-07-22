package user

import (
	"time"

	"github.com/google/uuid"
)

// define a core object

type User struct {
	ID          string
	Username    string
	Password    string
	DateCreated time.Time
}

func NewUser(username, password string) *User {
	return &User{
		ID:          uuid.NewString(),
		Username:    username,
		Password:    password,
		DateCreated: time.Now(),
	}
}
