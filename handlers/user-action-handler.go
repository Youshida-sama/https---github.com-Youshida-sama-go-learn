package handlers

import (
	"main/repositories"
	"main/requests"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func UserActionHandler(c echo.Context) error {
	userActionRequest := requests.UserActionRequest{}

	if err := c.Bind(&userActionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "Ошибка привязки модели").Error())
	}

	if err := c.Validate(&userActionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, errors.Wrap(err, "Ошибка валидации").Error())
	}

	userAction, err := userActionRequest.Map()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userAction.ID += 1

	updatedUserAction, err := repositories.CreateUserAction(*userAction)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, updatedUserAction)
}
