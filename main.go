package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/xHappyface/social/application/database"
	mysqldb "github.com/xHappyface/social/internal/adapters/mysql_db"
	"github.com/xHappyface/social/web/cmd"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/social_media?parseTime=true", dbUser, dbPass)
	db, err := sql.Open("mysql", dbSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	userRepo := mysqldb.NewUserRepoImpl(db, 15*time.Second)
	userService := database.NewUserService(userRepo)
	cmd.RunWebServer(userService)
}
