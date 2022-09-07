package main

import (
	"strconv"
	"net/http"
	"github.com/drhodes/golorem"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		words, _ := strconv.Atoi(c.QueryParam("words"))
		if words < 1 {
			words = 1
		}
		text := lorem.Sentence(words, words)
		return c.String(http.StatusOK, text)
	})
	e.Logger.Fatal(e.Start(":3000"))
}