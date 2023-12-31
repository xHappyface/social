package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/xHappyface/social/application/database"
	"github.com/xHappyface/social/application/web"
	"github.com/xHappyface/social/web/handlers"
)

func RunWebServer(userService database.UserService) {
	mux := http.NewServeMux()
	server := http.Server{
		Handler:        mux,
		Addr:           "localhost:8080",
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		IdleTimeout:    15 * time.Minute,
		MaxHeaderBytes: 4 << 10,
	}
	fs := http.FileServer(http.Dir("./web/html_css/static"))
	userHandler := handlers.NewUserHandlerImpl(web.NewUserHandler(userService))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/", "/index", "/login":
			handlers.HandleLogin(w, r)
		case "/create_account":
			handlers.HandleCreateAccount(w, r)
		case "/user/create":
			userHandler.HandleUserCreate(w, r)
		case "/user/read":
			userHandler.HandleUserRead(w, r)
		default:
			fs.ServeHTTP(w, r)
		}
	}))
	if err := server.ListenAndServeTLS("./web/server.crt", "./web/server.key"); err != nil {
		log.Fatalln(err)
	}
}
