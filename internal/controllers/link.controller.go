package controllers

import (
	"log/slog"
	"net/http"

	"github.com/LeonYalin/golang-todo-list-app/api"
	"github.com/LeonYalin/golang-todo-list-app/internal/services"
	"github.com/labstack/echo/v4"
)

type LinkController struct {
	service services.ILinkService
}

func NewLinkController(service services.ILinkService) *LinkController {
	return &LinkController{service: service}
}

func (this *LinkController) GetAllLinks(c echo.Context) error {
	slog.Info("(LinkController:GetAllLinks) enter")
	request := api.GetAllLinksRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}
	links, err := this.service.GetAllLinks(request)
	if err != nil {
		return err
	}
	slog.Info("(LinkController:GetAllLinks)", slog.Any("links", links))
	return c.JSON(http.StatusOK, links)
}

func (this *LinkController) GetLinkById(c echo.Context) error {
	id := c.Param("id")
	slog.Info("(LinkController:GetLinkById) enter", slog.Any("id", id))
	requestedLink, err := this.service.GetLinkById(id)
	if err != nil {
		return err
	}
	slog.Info("(LinkController:GetLinkById)", slog.Any("requested_link", requestedLink))
	return c.JSON(http.StatusOK, requestedLink)
}

func (this *LinkController) GetLinkByShort(c echo.Context) error {
	short := c.Param("short")
	slog.Info("(LinkController:GetLinkByShort) enter", slog.Any("short", short))
	res, err := this.service.GetLinkByShort(short)
	if err != nil {
		return err
	}
	slog.Info("(LinkController:GetLinkByShort)", slog.Any("requested_link", res))
	return c.Redirect(http.StatusPermanentRedirect, res.Link.Original)
}

func (this *LinkController) CreateLink(c echo.Context) error {
	slog.Info("(LinkController:CreateLink) enter")
	request := api.CreateLinkRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	createdLink, err := this.service.CreateLink(request)
	if err != nil {
		return err
	}
	slog.Info("(LinkController:CreateLink)", slog.Any("created_link", createdLink))
	return c.JSON(http.StatusCreated, createdLink)
}

func (this *LinkController) UpdateLink(c echo.Context) error {
	id := c.Param("id")
	slog.Info("(LinkController:UpdateLink) enter", slog.Any("id", id))
	request := api.UpdateLinkRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	updatedLink, err := this.service.UpdateLink(request, id)
	if err != nil {
		return err
	}
	slog.Info("(LinkController:UpdateLink)", slog.Any("updated_link", updatedLink))
	return c.JSON(http.StatusOK, updatedLink)
}

func (this *LinkController) DeleteLink(c echo.Context) error {
	id := c.Param("id")
	slog.Info("(LinkController:DeleteLink) enter", slog.Any("id", id))
	deletedLink, err := this.service.DeleteLink(id)
	if err != nil {
		return err
	}
	slog.Info("(LinkController:DeleteLink)", slog.Any("deleted_link", deletedLink))
	return c.JSON(http.StatusOK, deletedLink)
}
