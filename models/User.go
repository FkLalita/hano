package models

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
)

type User struct {
	username string
}

func CreateUser(db *sql.DB, username string) {
	_, err := db.Exec("INSERT INTO Users (username) VALUES (?)", username)
	if err != nil {
		log.Fatal(err)
	}
}
func GetUserID(db *sql.DB, e echo.Context, username string) int {
	var user_id int
	err := db.QueryRow("SELECT user_id FROM  Users WHERE username = ?", username).Scan(&user_id)
	if err != nil {
		log.Println("error getting user_id", err)
	}

	return user_id
}
