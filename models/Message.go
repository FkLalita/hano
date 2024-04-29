package models

import (
	"database/sql"
	"log"
	"time"
)

type Message struct {
	TopicID        int
	UserID         int
	MessageContent string
	CreatedAt      time.Time
}

func CreateMessage(db *sql.DB, topic_id int, user_id int, message_content string) error {
	_, err := db.Exec("INSERT INTO Topics (post_id, user_id, message_content) VALUES (?, ?, ?)", topic_id, user_id, message_content)
	if err != nil {
		log.Println("Error creating message:", err)
		return err
	}
	return nil
}

// GetTopics retrieves all topics from the database.
func GetMessages(db *sql.DB) ([]Message, error) {
	var messages []Message

	rows, err := db.Query("SELECT  post_id, user_id, message_content, created_at) FROM ChatMessage")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var m Message
		var createdAtStr []uint8 // Temporary variable to store the string from the database.

		if err := rows.Scan(&m.TopicID, &m.UserID, &m.MessageContent, &createdAtStr); err != nil {
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
