package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/shin-iji/go-shorten-url/handler"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	//store.InitializeStore()

	e.GET("/", handler.Hello)

	e.POST("/link", handler.CreateShortURL)

	e.GET("/l/:shortURL", handler.HandleShortURLRedirect)

	e.GET("/l/:shortURL/stats", handler.GetLinkCount)

	e.Logger.Fatal(e.Start(":8080"))
}
