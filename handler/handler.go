package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/shin-iji/go-shorten-url/shortener"

	"github.com/shin-iji/go-shorten-url/store"
)

type URLRequest struct {
	URL string `json:"url" validate:"required"`
}

type User struct {
	Name  string `json:"name"  validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age"   validate:"gte=0,lte=80"`
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func CreateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, true)
}

func CreateShortURL(c echo.Context) error {
	urlRep := new(URLRequest)
	if err := c.Bind(urlRep); err != nil {
		return err
	}
	if err := c.Validate(urlRep); err != nil {
		return err
	}

	var Content struct {
		Link string `json:"link"`
	}

	shortURL := shortener.GenerateShortLink(urlRep.URL)
	store.SaveURLMapping(shortURL, urlRep.URL)

	host := "http://localhost:8080/"
	Content.Link = host + shortURL

	return c.JSON(http.StatusOK, &Content)
}

func HandleShortURLRedirect(c echo.Context) error {
	shortURL := c.Param("shortURL")
	initialURL := store.RetrieveInitialURL(shortURL)
	return c.Redirect(302, initialURL)
}
