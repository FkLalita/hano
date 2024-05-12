package utils

import (
	"database/sql"
	"fmt"
)

func CreateTable(db *sql.DB) error {
	// Execute SQL statement to create table

	_, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS Users (
            user_id INT PRIMARY KEY AUTO_INCREMENT,
            username VARCHAR(50) NOT NULL
      );    
  `)

	if err != nil {
		fmt.Println("Error creating user table:", err)
		return err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Topics  (
            post_id INT PRIMARY KEY AUTO_INCREMENT,
            title VARCHAR(255),
            description TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            user_id INT,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON UPDATE CASCADE ON DELETE CASCADE
          );
  `)

	if err != nil {
		fmt.Println("Error creating  topic table:", err)
		return err
	}

	_, err = db.Exec(`
       CREATE TABLE IF NOT EXISTS ChatMessage (
            message_id INT PRIMARY KEY AUTO_INCREMENT,
            post_id INT,
            user_id INT,
            message_content TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (post_id) REFERENCES Topics(post_id) ON UPDATE CASCADE ON DELETE CASCADE,
            FOREIGN KEY (user_id) REFERENCES Users(user_id) ON UPDATE CASCADE ON DELETE CASCADE
          );
  `)
	if err != nil {
		fmt.Println("Error creating  chat table:", err)
		return err
	}
	fmt.Println("table created successfully")

	return nil

}
