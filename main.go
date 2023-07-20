package main

import (
	"fmt"
	"log"

	"github.com/xHappyface/social/internal/core/session"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/social_media?parseTime=true", dbUser, dbPass)
	// db, err := sql.Open("mysql", dbSourceName)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// var sesh *session.Session
	sesh, err := session.NewSession("b63ed5a1-58b9-444e-b89d-9550dcc64c92")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sesh.ID, len(sesh.ID))
	fmt.Println(sesh.CsrfToken, len(sesh.CsrfToken))
}
