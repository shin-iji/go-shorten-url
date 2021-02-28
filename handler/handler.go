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

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
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

	host := "http://a7.tnpl.me:8080/"
	//host := "http://localhost:8080/"
	Content.Link = host + shortURL

	return c.JSON(http.StatusOK, &Content)
}

func HandleShortURLRedirect(c echo.Context) error {
	shortURL := c.Param("shortURL")
	initialURL := store.RetrieveInitialURL(shortURL)
	return c.Redirect(302, initialURL)
}

func GetLinkCount (c echo.Context) error {
	shortURL := c.Param("shortURL")
	count := store.GetLinkCount(shortURL)
	var Content struct {
		Visit int `json:"visit"`
	}
	Content.Visit = count
	return c.JSON(http.StatusOK, &Content)
}