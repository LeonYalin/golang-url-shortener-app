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

func (this *LinkController) GetAll(c echo.Context) error {
	slog.Info("(LinkController:GetAll) enter")
	request := api.GetAllLinksRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}
	links, err := this.service.Get(request)
	if err != nil {
		return err
	}
	slog.Info("(LinkController:GetAll)", slog.Any("links", links))
	return c.JSON(http.StatusOK, links)
}

func (this *LinkController) GetById(c echo.Context) error {
	id := c.Param("id")
	slog.Info("(LinkController:GetById) enter", slog.Any("id", id))
	requestedLink, err := this.service.GetById(id)
	if err != nil {
		return err
	}
	slog.Info("(LinkController:GetById)", slog.Any("requested_link", requestedLink))
	return c.JSON(http.StatusOK, requestedLink)
}

func (this *LinkController) Create(c echo.Context) error {
	slog.Info("(LinkController:Create) enter")
	request := api.CreateLinkRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}
	createdLink, err := this.service.Create(request)
	if err != nil {
		return err
	}
	slog.Info("(LinkController:Create)", slog.Any("created_link", createdLink))
	return c.JSON(http.StatusOK, createdLink)
}

func (this *LinkController) Update(c echo.Context) error {
	id := c.Param("id")
	slog.Info("(LinkController:Update) enter", slog.Any("id", id))
	request := api.UpdateLinkRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}
	updatedLink, err := this.service.Update(request, id)
	if err != nil {
		return err
	}
	slog.Info("(LinkController:Update)", slog.Any("updated_link", updatedLink))
	return c.JSON(http.StatusOK, updatedLink)
}
