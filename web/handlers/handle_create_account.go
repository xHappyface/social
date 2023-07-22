package handlers

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/xHappyface/social/core/errors"
	"github.com/xHappyface/social/core/user"
	"golang.org/x/crypto/bcrypt"
)

func (hn *UserHandlerImpl) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	const _URL_CHECK = "https://localhost:8080/create_account"
	referer := r.Header.Get("Referer")
	if referer == _URL_CHECK {
		r.Method = http.MethodGet
	}
	switch r.Method {
	case http.MethodGet:
		tpl, err := template.ParseFiles("./web/html_css/templates/create_account.html")
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if referer == _URL_CHECK {
			if err = tpl.Execute(w, struct{}{}); err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		} else {
			if err = tpl.Execute(w, nil); err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}
	case http.MethodPost:
		const _MIN_LEN = 5
		const _MAX_LEN = 36
		if err := r.ParseForm(); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		exp := `^[a-zA-Z0-9\Q.?!@#$%^*-_=+\E]+$`
		username := r.Form.Get("Username")
		password := r.Form.Get("Password")
		ok, err := regexp.MatchString(exp, username)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if !ok || !(len(username) >= _MIN_LEN && len(username) <= _MAX_LEN) {
			http.Redirect(w, r, "/create_account", http.StatusPermanentRedirect)
			return
		}
		ok, err = regexp.MatchString(exp, password)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if !ok || !(len(password) >= _MIN_LEN && len(password) <= _MAX_LEN) {
			http.Redirect(w, r, "/create_account", http.StatusPermanentRedirect)
			return
		}
		var hash []byte
		hash, err = bcrypt.GenerateFromPassword([]byte(password), 16)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		user := user.NewUser(username, string(hash))
		if _, err = hn.h.ReadByID(user.ID); err != errors.ErrNoRowsRetrieved && err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		} else if err == nil {
			// handle case where userID already exists.
			return
		}
		if err = hn.h.Create(user); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
	default:
		http.NotFound(w, r)
	}
}
