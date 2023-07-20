package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Handler:        mux,
		Addr:           "localhost:8080",
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		IdleTimeout:    15 * time.Minute,
		MaxHeaderBytes: 1 << 10,
	}
	fs := http.FileServer(http.Dir("./html_css/static"))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			switch r.URL.Path {
			case "/", "/index":
				http.ServeFile(w, r, "./html_css/static/index.html")
			case "/create_account":
				http.ServeFile(w, r, "./html_css/static/create_account.html")
			default:
				fs.ServeHTTP(w, r)
			}
		case http.MethodPost:
			switch r.URL.Path {
			default:
				http.NotFound(w, r)
			}
		default:
			http.NotFound(w, r)
		}
	}))
	if err := server.ListenAndServeTLS("./server.crt", "./server.key"); err != nil {
		log.Fatalln(err)
	}
}
