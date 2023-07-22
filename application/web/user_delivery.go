package web

import (
	"github.com/xHappyface/social/application/database"
	"github.com/xHappyface/social/core/user"
)

type UserDelivery interface {
	Create(*user.User) error
	ReadByID(string) (*user.User, error)
	Update(*user.User) error
	DeleteByID(string) error
}

type UserHandler struct {
	userService database.UserService
}

func NewUserHandler(userService database.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Create(user *user.User) error {
	return h.userService.Create(user)
}

func (h *UserHandler) ReadByID(userID string) (*user.User, error) {
	return h.userService.ReadByID(userID)
}

func (h *UserHandler) Update(user *user.User) error {
	return h.userService.Update(user)
}

func (h *UserHandler) DeleteByID(userID string) error {
	return h.userService.DeleteByID(userID)
}
