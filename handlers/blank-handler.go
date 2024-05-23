package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BlankHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Ну и что ты тут забыл?")
}
