package utils

import (
  "database/sql"
  "fmt"
)

func CreateTable(db *sql.DB) error {
  // Execute SQL statement to create table

  _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Topics  (
            post_id INT PRIMARY KEY AUTO_INCREMENT,
            topic_title VARCHAR(255),
            topic_content TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `)


    if err != nil {
        fmt.Println("Error creating  topic table:", err)
        return err
    }


    _, err = db.Exec(`
       CREATE TABLE IF NOT EXISTS User (
            user_id INT PRIMARY KEY AUTO_INCREMENT,
            user_name VARCHAR(50)
      );    
  `)

    if err != nil {
        fmt.Println("Error creating user table:", err)
        return err
    } 



    _, err = db.Exec(`
       CREATE TABLE IF NOT EXISTS ChatMessage (
            message_id INT PRIMARY KEY AUTO_INCREMENT,
            post_id INT,
            user_id INT,
            message_content TEXT  ,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (post_id) REFERENCES Topics(post_id),
            FOREIGN KEY (user_id) REFERENCES User(user_id)
      );
  `)
    if err != nil {
        fmt.Println("Error creating  chat table:", err)
        return err
    }
    fmt.Println("table created successfully")

    return nil    
    
}
