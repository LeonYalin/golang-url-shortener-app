package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LeonYalin/golang-todo-list-app/api"
	"github.com/LeonYalin/golang-todo-list-app/internal/helpers"
	"github.com/LeonYalin/golang-todo-list-app/internal/models"
	"github.com/LeonYalin/golang-todo-list-app/internal/services"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var linkJSON = `{"original": "https://www.google.com"}`
var link = models.Link{Id: uuid.NewString(), Original: "https://www.google.com", Short: "/l/lala"}

func TestLinkController(t *testing.T) {
	suite.Run(t, new(LinkControllerSuite))
}

type LinkControllerSuite struct {
	suite.Suite
	controller *LinkController
	e          *echo.Echo
}

func (this *LinkControllerSuite) SetupSuite() {
	this.e = echo.New()
	this.e.Validator = &helpers.RequestValidator{Validator: validator.New()}
	service := services.NewMockLinkService()
	this.controller = NewLinkController(service)
}

func (this *LinkControllerSuite) TearDownSuite() {
	this.e.Shutdown(context.Background())
}

func (this *LinkControllerSuite) TestGetAllLinks() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := this.e.NewContext(req, rec)
	links := make([]models.Link, 0)
	links = append(links, link)
	this.controller.service.(*services.MockLinkService).GetAllLinksFunc = func(request api.GetAllLinksRequest) (api.GetAllLinksResponse, error) {
		return api.GetAllLinksResponse{Links: links}, nil
	}

	// Act
	err := this.controller.GetAllLinks(c)

	// Assert
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), rec.Code, http.StatusOK)
	var response api.GetAllLinksResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(this.T(), err)
	assert.EqualValues(this.T(), response.Links, links)
}

func (this *LinkControllerSuite) TestGetLinkById() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := this.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(link.Id)
	this.controller.service.(*services.MockLinkService).GetLinkByIdFunc = func(id string) (api.GetLinkByIdResponse, error) {
		return api.GetLinkByIdResponse{Link: link}, nil
	}

	// Act
	err := this.controller.GetLinkById(c)

	// Assert
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), http.StatusOK, rec.Code)
	var response api.GetLinkByIdResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), response.Link, link)
}

func (this *LinkControllerSuite) TestGetLinkByShort() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/:short", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := this.e.NewContext(req, rec)
	c.SetParamNames("short")
	c.SetParamValues(link.Short)
	this.controller.service.(*services.MockLinkService).GetLinkByShortFunc = func(id string) (api.GetLinkByShortResponse, error) {
		return api.GetLinkByShortResponse{Link: link}, nil
	}

	// Act
	err := this.controller.GetLinkByShort(c)

	// Assert
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), http.StatusPermanentRedirect, rec.Code)
	assert.Equal(this.T(), link.Original, rec.Header().Get("Location"))
}

func (this *LinkControllerSuite) TestCreateLink() {
	// Arrange
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(linkJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := this.e.NewContext(req, rec)
	this.controller.service.(*services.MockLinkService).CreateLinkFunc = func(request api.CreateLinkRequest) (api.CreateLinkResponse, error) {
		return api.CreateLinkResponse{Link: link}, nil
	}

	// Act
	err := this.controller.CreateLink(c)

	// Assert
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), http.StatusCreated, rec.Code)
	var response api.CreateLinkResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), response.Link, link)
}

func (this *LinkControllerSuite) TestUpdateLink() {
	// Arrange
	req := httptest.NewRequest(http.MethodPut, "/:id", strings.NewReader(linkJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := this.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(link.Id)
	this.controller.service.(*services.MockLinkService).UpdateLinkFunc = func(request api.UpdateLinkRequest, id string) (api.UpdateLinkResponse, error) {
		return api.UpdateLinkResponse{Link: link}, nil
	}

	// Act
	err := this.controller.UpdateLink(c)

	// Assert
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), http.StatusOK, rec.Code)
	var response api.UpdateLinkResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), response.Link, link)
}

func (this *LinkControllerSuite) TestDeteleLink() {
	// Arrange
	req := httptest.NewRequest(http.MethodDelete, "/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := this.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(link.Id)
	this.controller.service.(*services.MockLinkService).DeleteLinkFunc = func(id string) (api.DeleteLinkResponse, error) {
		return api.DeleteLinkResponse{Link: link}, nil
	}

	// Act
	err := this.controller.DeleteLink(c)

	// Assert
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), http.StatusOK, rec.Code)
	var response api.DeleteLinkResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), response.Link, link)
}
