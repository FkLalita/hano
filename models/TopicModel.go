package models

import (
	"database/sql"
	"time"

	"log"
)

type Topic struct {
	Title       string
	Description string
	CreatedAt   time.Time
}

func CreateTopic(db *sql.DB, title string, description string) {
	_, err := db.Exec("INSERT INTO Topics (title, description) VALUES (?,?) ", title, description)
	if err != nil {
		log.Println(err)
	}

}
