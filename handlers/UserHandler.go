package handlers

import (
	"database/sql"
	"net/http"

	"github.com/FkLalita/hano/models"
	"github.com/FkLalita/hano/utils"

	"github.com/FkLalita/hano/templates"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(db *sql.DB, e echo.Context) error {
	if e.Request().Method == http.MethodPost {
		username := e.FormValue("username")
		models.CreateUser(db, username)
		return e.Redirect(http.StatusSeeOther, "/")
	}

	return utils.Render(e, http.StatusOK, templates.Username())

}
