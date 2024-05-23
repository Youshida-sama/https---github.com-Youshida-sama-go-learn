package main

import (
	"main/handlers"
	"main/storage"
	"main/validations"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	v, err := validations.NewValidator()

	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Validator = v

	e.GET("/", handlers.BlankHandler)
	e.POST("/action", handlers.UserActionHandler)

	storage.InitDB()

	e.Logger.Fatal(e.Start(":8080"))
}
