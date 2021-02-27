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
	

	e.GET("/", handler.Hello)

	e.POST("/", handler.CreateShortURL)

	e.GET("/:shortURL", handler.HandleShortURLRedirect)

	e.Logger.Fatal(e.Start(":8080"))
}
