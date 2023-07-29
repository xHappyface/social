package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

type LoginErrors struct {
	InvalidUsernameError bool
	InvalidPasswordError bool
	UnexpectedError      bool
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == "https://localhost:8080/create_account" {
		r.Method = http.MethodGet
	}
	tpl, err := template.ParseFiles("./web/html_css/templates/index.html")
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
		var loginErrors LoginErrors
		if queryParams.Has("invalid_username_error") {
			loginErrors.InvalidUsernameError, err = strconv.ParseBool(queryParams.Get("invalid_username_error"))
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}
		if queryParams.Has("invalid_password_error") {
			loginErrors.InvalidPasswordError, err = strconv.ParseBool(queryParams.Get("invalid_password_error"))
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}
		if queryParams.Has("unexpected_error") {
			loginErrors.UnexpectedError, err = strconv.ParseBool(queryParams.Get("unexpected_error"))
			if err != nil {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
		}
		if err = tpl.Execute(w, loginErrors); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	default:
		http.NotFound(w, r)
	}
}
