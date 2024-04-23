package utils

import (
  "database/sql"
  "fmt"
)

func CreateTable(db *sql.DB) {
  // Execute SQL statement to create table

  _, err := db.Exec(`
        CREATE TABLE Topics (
            post_id INT PRIMARY KEY AUTO_INCREMENT,
            topic_title VARCHAR(255),
            topic_content TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `)


    if err != nil {
        fmt.Println("Error creating  topic table:", err)
        return
    }


    _, err = db.Exec(`
       CREATE TABLE User (
            user_id INT PRIMARY KEY AUTO_INCREMENT,
            user_name VARCHAR(50)
      );    
  `)

    if err != nil {
        fmt.Println("Error creating user table:", err)
        return
    } 



    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS ChatMessage (
            message_id INT PRIMARY KEY AUTO_INCREMENT,
            user_id INT,
            message_content TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `)
    if err != nil {
        fmt.Println("Error creating  chat table:", err)
        return
    }

  }
