package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/FkLalita/hano/models"
	"github.com/FkLalita/hano/templates"
	"github.com/FkLalita/hano/utils"

	"github.com/labstack/echo/v4"
)

func GetMessagesHandlers(db *sql.DB, e echo.Context) error {
	// Retrieve all messages from the database

	post_id, _ := strconv.Atoi(e.Param("id"))
	// the creation has been moved to websocket
	//if e.Request().Method == http.MethodPost {

	//	user_id := 1
	//	content := e.FormValue("message")

	//	err := models.CreateMessage(db, post_id, user_id, content)
	//	if err != nil {
	//		return e.String(http.StatusInternalServerError, "Failed to send message")
	//	}
	//	broadcast(content, e)

	//}
	messages, err := models.GetMessages(db, post_id)
	if err != nil {
		e.Logger().Error(err)
		// If there's an error retrieving messages, return an internal server error response
		return e.String(http.StatusInternalServerError, "Failed to retrieve messages from the database")
	}

	return utils.Render(e, http.StatusOK, templates.GetMessage(messages, post_id))

}
