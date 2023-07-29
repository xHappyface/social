package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

type CreateAccountErrors struct {
	UsernameError bool
	PasswordError bool
}

func HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./web/html_css/templates/create_account.html")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		if err = tpl.Execute(w, nil); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var usernameError, passwordError bool
		queryParams := r.URL.Query()
		var boolean bool
		if queryParams.Has("username_error") {
			boolean, err = strconv.ParseBool(queryParams.Get("username_error"))
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			usernameError = boolean
		}
		if queryParams.Has("password_error") {
			boolean, err = strconv.ParseBool(queryParams.Get("password_error"))
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			passwordError = boolean
		}
		if err = tpl.Execute(w, CreateAccountErrors{
			UsernameError: usernameError,
			PasswordError: passwordError,
		}); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	default:
		http.NotFound(w, r)
	}
}
