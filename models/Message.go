package models

import (
	"database/sql"
	"log"
	"time"
)

type Message struct {
	TopicID        int
	UserID         int
	UserName       string
	MessageContent string
	CreatedAt      time.Time
}

func CreateMessage(db *sql.DB, topic_id int, user_id int, username string, message_content string) error {
	_, err := db.Exec("INSERT INTO ChatMessage (post_id, user_id,username, message_content) VALUES (?, ?, ?, ?)", topic_id, user_id, username, message_content)
	if err != nil {
		log.Println("Error creating message:", err)
		return err
	}
	return nil
}

// GetMessages retrieves all topics from the database.
func GetMessages(db *sql.DB, post_id int) ([]Message, error) {
	var messages []Message

	rows, err := db.Query("SELECT  post_id, user_id,username, message_content, created_at FROM ChatMessage WHERE post_id = ?", post_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var m Message
		var createdAtStr []uint8 // Temporary variable to store the string from the database.

		if err := rows.Scan(&m.TopicID, &m.UserID, &m.UserName, &m.MessageContent, &createdAtStr); err != nil {
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
			m.CreatedAt = parsedTime
		}

		messages = append(messages, m)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	return messages, nil
}
