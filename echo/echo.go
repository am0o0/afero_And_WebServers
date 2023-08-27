package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.File(c.QueryParam("filePath"))
	})

	e.GET("/attachment", func(c echo.Context) error {
		return c.Attachment(c.QueryParam("filePath"), "file name in response")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
