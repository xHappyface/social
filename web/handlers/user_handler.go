package handlers

import (
	"github.com/xHappyface/social/application/web"
)

type UserHandlerImpl struct {
	h *web.UserHandler
}

func NewUserHandlerImpl(h *web.UserHandler) *UserHandlerImpl {
	return &UserHandlerImpl{
		h: h,
	}
}
