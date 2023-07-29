package handlers

import (
	"net/http"

	"github.com/xHappyface/social/core/errors"
	"golang.org/x/crypto/bcrypt"
)

func (hn *UserHandlerImpl) HandleUserRead(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		username := r.Form.Get("Username")
		password := r.Form.Get("Password")
		user, err := hn.ReadByName(username)
		if err == errors.ErrNoRowsRetrieved {
			http.Redirect(w, r, "/login?invalid_username_error=true", http.StatusTemporaryRedirect)
			return
		} else if err != nil {
			http.Redirect(w, r, "/login?unexpected_error=true", http.StatusTemporaryRedirect)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			http.Redirect(w, r, "/login?invalid_password_error=true", http.StatusTemporaryRedirect)
			return
		}
	default:
		http.NotFound(w, r)
	}
}
