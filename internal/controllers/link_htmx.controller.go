package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LinkHtmxController struct{}

func NewLinkHtmxController() *LinkHtmxController {
	return &LinkHtmxController{}
}

func (this LinkHtmxController) GetAllLinks(c echo.Context) error {
	return c.Render(http.StatusOK, "links.html", "World")
}
