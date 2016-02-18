package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Get("/aa", func(c *echo.Context) error {
		c.String(http.StatusOK, "route /aa")
		return nil
	})
	e.Get("/ab", func(c *echo.Context) error {
		c.String(http.StatusOK, "route /ab")
		return nil
	})
	e.Static("/", "./")

	e.Run(":12345")
}
