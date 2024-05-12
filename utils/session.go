package utils

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CreateSession(e echo.Context, username string) error {
	sess, err := session.Get("user-session", e)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	sess.Values["username"] = username
	if err := sess.Save(e.Request(), e.Response()); err != nil {
		return err
	}
	return e.NoContent(http.StatusOK)

}

func GetSession(e echo.Context) (string, error) {
	sess, err := session.Get("user-session", e)
	if err != nil {
		return "", err
	}
	values := sess.Values["username"]
	if values == nil {
		return "", fmt.Errorf("user session  not found or has a nil value")
	}
	// convert the value to a string
	strValue, ok := values.(string)
	if !ok {
		return "", fmt.Errorf("user session  value cannot be converted to string")
	}
	return strValue, nil
}
