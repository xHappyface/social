package handlers

import (
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == "https://localhost:8080/create_account" {
		r.Method = http.MethodGet
	}
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "./web/html_css/static/index.html")
	default:
		http.NotFound(w, r)
	}
}
