package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/LeonYalin/golang-todo-list-app/api"
	"github.com/LeonYalin/golang-todo-list-app/internal/models"
	"github.com/carlmjohnson/requests"
	"github.com/labstack/echo/v4"
)

type LinkHtmxController struct{}

func NewLinkHtmxController() *LinkHtmxController {
	return &LinkHtmxController{}
}

func (this LinkHtmxController) GetAllLinks(c echo.Context) error {
	links := api.GetAllLinksResponse{}
	err := requests.URL("http://localhost:3001/links").ToJSON(&links).Fetch(context.Background())
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "links.html", links)
}

func (this *LinkHtmxController) GetLinkById(c echo.Context) error {
	return nil
}

func (this *LinkHtmxController) CreateLinkClick(c echo.Context) error {
	return c.Render(http.StatusOK, "create_link.html", nil)
}

func (this *LinkHtmxController) CreateLinkClickConfirmClick(c echo.Context) error {
	original := c.FormValue("original")
	if original == "" {
		return errors.New("missing original param")
	}
	req := api.CreateLinkRequest{Original: original}
	var link *models.Link
	err := requests.URL("http://localhost:3001/links").Method("POST").BodyJSON(req).ToJSON(&link).Fetch(context.Background())
	if err != nil {
		return err
	}
	return this.GetAllLinks(c)
}

func (this *LinkHtmxController) DeleteLinkClick(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("missing link id param")
	}
	return c.Render(http.StatusOK, "delete_link.html", id)
}

func (this *LinkHtmxController) DeleteLinkConfirmClick(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("missing link id param")
	}
	var link *models.Link
	err := requests.URL(fmt.Sprintf("http://localhost:3001/links/%s", id)).Method("DELETE").ToJSON(&link).Fetch(context.Background())
	if err != nil {
		return err
	}
	return this.GetAllLinks(c)
}

func (this *LinkHtmxController) EditLinkClick(c echo.Context) error {
	var res api.CreateLinkResponse
	err := requests.URL(fmt.Sprintf("http://localhost:3001/links/%v", c.Param("id"))).ToJSON(&res).Fetch(context.Background())
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "edit_link.html", res.Link)
}

func (this *LinkHtmxController) EditLinkConfirmClick(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return errors.New("missing link id param")
	}
	original := c.FormValue("original")
	if original == "" {
		return errors.New("missing original param")
	}
	req := api.UpdateLinkRequest{Original: original}
	var link *models.Link
	err := requests.URL(fmt.Sprintf("http://localhost:3001/links/%s", id)).Method("PUT").BodyJSON(req).ToJSON(&link).Fetch(context.Background())
	if err != nil {
		return err
	}
	return this.GetAllLinks(c)
}
