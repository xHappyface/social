package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/xHappyface/social/core/errors"
	"github.com/xHappyface/social/core/user"
	"golang.org/x/crypto/bcrypt"
)

func (hn *UserHandlerImpl) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		username := r.Form.Get("Username")
		password := r.Form.Get("Password")
		validUsername := regexp.MustCompile(`^[a-zA-Z0-9_]{5,36}$`).MatchString(username)
		var validPassword bool
		if (len(password) >= 5) && (len(password) <= 36) {
			validPassword = true
		}
		if !validUsername || !validPassword {
			redirectURL := fmt.Sprintf("/create_account?username_error=%t&password_error=%t", !validUsername, !validPassword)
			http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		user := user.NewUser(username, string(hash))
		if _, err = hn.ReadByID(user.ID); err != errors.ErrNoRowsRetrieved && err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		} else if err == nil {
			http.Redirect(w, r, "/create_account?unexpected_error=true", http.StatusTemporaryRedirect)
			return
		}
		if _, err = hn.ReadByName(user.Username); err != errors.ErrNoRowsRetrieved && err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		} else if err == nil {
			http.Redirect(w, r, "/create_account?username_exists_error=true", http.StatusTemporaryRedirect)
			return
		}
		if err = hn.Create(user); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
	default:
		http.NotFound(w, r)
	}
}
