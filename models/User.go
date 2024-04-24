package models

import (
	"database/sql"
	"log"
)

type User struct {
	username string
}

func CreateUser(db *sql.DB, username string) {
	_, err := db.Exec("INSERT INTO User (user_name) VALUES (?)", username)
	if err != nil {
		log.Fatal(err)
	}
}
