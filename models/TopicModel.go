package models

import (
    "database/sql"
    "log"
    "time"
)

type Topic struct {
    Title       string
    Description string
    CreatedAt   time.Time
}

// CreateTopic creates a new topic in the database.
func CreateTopic(db *sql.DB, title string, description string) error {
    _, err := db.Exec("INSERT INTO Topics (title, description) VALUES (?, ?)", title, description)
    if err != nil {
        log.Println("Error creating topic:", err)
        return err
    }
    return nil
}

// GetTopics retrieves all topics from the database.
func GetTopics(db *sql.DB) ([]Topic, error) {
    rows, err := db.Query("SELECT title, description, created_at FROM Topics")
    if err != nil {
        log.Println("Error retrieving topics:", err)
        return nil, err
    }
    defer rows.Close()

    var topics []Topic

    for rows.Next() {
        var topic Topic
        err := rows.Scan(&topic.Title, &topic.Description, &topic.CreatedAt)
        if err != nil {
            log.Println("Error scanning topic row:", err)
            return nil, err
        }
        topics = append(topics, topic)
    }

    if err := rows.Err(); err != nil {
        log.Println("Error iterating over topic rows:", err)
        return nil, err
    }

    return topics, nil
}

