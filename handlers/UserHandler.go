package handlers

import (
	"database/sql"
	"net/http"

	"github.com/FkLalita/hano/models"
	"github.com/FkLalita/hano/utils"

	"github.com/FkLalita/hano/templates"

	"github.com/labstack/echo/v4"
	"log"
)

func CreateUserHandler(db *sql.DB, e echo.Context) error {
	if e.Request().Method == http.MethodPost {
		username := e.FormValue("username")

		models.CreateUser(db, username)

		err := utils.CreateSession(e, username)
		if err != nil {
			log.Println("error creating session")
		}
		return e.Redirect(http.StatusSeeOther, "/")
	}

	return utils.Render(e, http.StatusOK, templates.Username())

}
