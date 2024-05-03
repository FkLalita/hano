package handlers

import (
	"database/sql"
	"net/http"

	"github.com/FkLalita/hano/models"
	"github.com/FkLalita/hano/templates"
	"github.com/FkLalita/hano/utils"
	"github.com/labstack/echo/v4"
)

func GetTopicsHandler(db *sql.DB, e echo.Context) error {

	// Retrieve all topics from the database
	topics, err := models.GetTopics(db)
	if err != nil {
		// If there's an error retrieving topics, return an internal server error response
		return e.String(http.StatusInternalServerError, "Failed to retrieve topics from the database")
	}

	return utils.Render(e, http.StatusOK, templates.Home(topics))

}

func CreateTopicHandler(db *sql.DB, e echo.Context) error {
	if e.Request().Method == http.MethodPost {
		title := e.FormValue("title")
		description := e.FormValue("description")
		err := models.CreateTopic(db, title, description)
		if err != nil {
			return e.String(http.StatusInternalServerError, "Failed to create topic")
		}
		e.Redirect(http.StatusSeeOther, "/")

	}
	return utils.Render(e, http.StatusOK, templates.CreateTopic())

}