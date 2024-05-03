package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"github.com/FkLalita/hano/handlers"
	"github.com/FkLalita/hano/utils"
)

func main() {
	mySqlConnect := os.Getenv("HANOENV")
	db, err := sql.Open("mysql", mySqlConnect)
	if err != nil {
		log.Println(err)
		//	panic(err.Error())

	}

	defer db.Close()

	if err := utils.CreateTable(db); err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Connection successful!")
	}
	e := echo.New()

	// Serve static files from the node_modules directory
	e.Static("/node_modules", "node_modules")

	e.POST("/username", func(e echo.Context) error {
		return handlers.CreateUserHandler(db, e)
	})
	e.GET("username", func(e echo.Context) error {
		return handlers.CreateUserHandler(db, e)
	})

	e.GET("/", func(e echo.Context) error {
		return handlers.GetTopicsHandler(db, e)

	})
	e.POST("/create", func(e echo.Context) error {
		return handlers.CreateTopicHandler(db, e)
	})

	e.GET("/create", func(e echo.Context) error {
		return handlers.CreateTopicHandler(db, e)
	})

	e.GET("/topics/:id/messages", func(e echo.Context) error {
		return handlers.GetMessagesHandlers(db, e)
	})
	e.POST("/topics/:id/messages", func(e echo.Context) error {
		return handlers.GetMessagesHandlers(db, e)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
