package handlers

import (
	"database/sql"
	"net/http"

	"github.com/FkLalita/hano/models"
	"github.com/FkLalita/hano/templates"
	"github.com/FkLalita/hano/utils"
	"github.com/labstack/echo/v4"
)

func GetMessagesHandlers(db *sql.DB, e echo.Context) error {
	// Retrieve all messages from the database
	messages, err := models.GetMessages(db)
	if err != nil {
		// If there's an error retrieving messages, return an internal server error response
		return e.String(http.StatusInternalServerError, "Failed to retrieve messages from the database")
	}

	return utils.Render(e, http.StatusOK, templates.Home(messages))

}
