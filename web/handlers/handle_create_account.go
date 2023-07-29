package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

type CreateAccountErrors struct {
	UsernameError       bool
	PasswordError       bool
	UsernameExistsError bool
	UnexpectedError     bool
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
		queryParams := r.URL.Query()
		var createAccountErrors CreateAccountErrors
		if queryParams.Has("username_error") {
			createAccountErrors.UsernameError, err = strconv.ParseBool(queryParams.Get("username_error"))
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}
		if queryParams.Has("password_error") {
			createAccountErrors.PasswordError, err = strconv.ParseBool(queryParams.Get("password_error"))
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}
		if queryParams.Has("username_exists_error") {
			createAccountErrors.UsernameExistsError, err = strconv.ParseBool(queryParams.Get("username_exists_error"))
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}
		if queryParams.Has("unexpected_error") {
			createAccountErrors.UnexpectedError, err = strconv.ParseBool(queryParams.Get("unexpected_error"))
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}
		if err = tpl.Execute(w, createAccountErrors); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	default:
		http.NotFound(w, r)
	}
}
