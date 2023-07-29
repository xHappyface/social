package handlers

import (
	"github.com/xHappyface/social/application/web"
	"github.com/xHappyface/social/core/user"
)

type UserHandler interface {
	Create(*user.User) error
	ReadByID(string) (*user.User, error)
	ReadByName(string) (*user.User, error)
	Update(*user.User) error
	DeleteByID(string) error
}

type UserHandlerImpl struct {
	h *web.UserHandler
}

func NewUserHandlerImpl(h *web.UserHandler) *UserHandlerImpl {
	return &UserHandlerImpl{
		h: h,
	}
}

func (hn *UserHandlerImpl) Create(user *user.User) error {
	return hn.h.Create(user)
}

func (hn *UserHandlerImpl) ReadByID(userID string) (*user.User, error) {
	return hn.h.ReadByID(userID)
}

func (hn *UserHandlerImpl) ReadByName(userName string) (*user.User, error) {
	return hn.h.ReadByName(userName)
}

func (hn UserHandlerImpl) Update(user *user.User) error {
	return hn.h.Update(user)
}

func (hn *UserHandlerImpl) DeleteByID(userID string) error {
	return hn.h.DeleteByID(userID)
}
