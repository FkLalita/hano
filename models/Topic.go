package models

import (
	"database/sql"
	"log"
	"time"
)

type Topic struct {
	TopicID     int
	Title       string
	Description string
	CreatedAt   time.Time
	UserID      int
}

// CreateTopic creates a new topic in the database.
func CreateTopic(db *sql.DB, title string, description string, username string) error {
	var user_id int
	err := db.QueryRow("SELECT user_id FROM Users WHERE username = ?", username).Scan(&user_id)
	if err != nil {
		log.Print("errror getting user_id", err)
	}
	_, err = db.Exec("INSERT INTO Topics (title, description, user_id) VALUES (?, ?, ?)", title, description, user_id)
	if err != nil {
		log.Println("Error creating topic:", err)
		return err
	}
	return nil
}

// GetTopics retrieves all topics from the database.
func GetTopics(db *sql.DB) ([]Topic, error) {
	var topics []Topic

	rows, err := db.Query("SELECT * FROM Topics")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t Topic
		var createdAtStr []uint8 // Temporary variable to store the string from the database.

		if err := rows.Scan(&t.TopicID, &t.Title, &t.Description, &createdAtStr, &t.UserID); err != nil {
			log.Println(err)
			continue
		}

		// Convert the createdAtStr ([]uint8) to a string.
		createdAtString := string(createdAtStr)

		// Parse the createdAtString as time.Time.
		parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAtString)
		if err != nil {
			log.Println(err)
		} else {
			t.CreatedAt = parsedTime
		}

		topics = append(topics, t)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	return topics, nil
}
