package services_test

import (
	// "errors"
	"testing"

	"github.com/LeonYalin/golang-todo-list-app/api"
	"github.com/LeonYalin/golang-todo-list-app/internal/models"
	"github.com/LeonYalin/golang-todo-list-app/internal/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var link = models.Link{Id: uuid.NewString(), Original: "https://www.google.com", Short: "http://c/lala"}

func TestLinkService(t *testing.T) {
	suite.Run(t, new(LinkServiceTestSuite))
}

type LinkServiceTestSuite struct {
	suite.Suite
	service *services.LinkService
}

func (this *LinkServiceTestSuite) SetupSuite() {
	repo := services.NewLinkRepository()
	this.service = services.NewLinkService(repo)
}

func (this *LinkServiceTestSuite) TestGetAllLinks() {
	req := api.GetAllLinksRequest{}

	// Check there are no links
	res, err := this.service.GetAllLinks(req)
	assert.NoError(this.T(), err)
	assert.Len(this.T(), res.Links, 0)

	// Add 1 link
	createLinkReq := api.CreateLinkRequest{Original: link.Original}
	createLinkRes, err := this.service.CreateLink(createLinkReq)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), createLinkRes.Link.Original, link.Original)

	// Check get all links again
	res, err = this.service.GetAllLinks(req)
	assert.NoError(this.T(), err)
	assert.Len(this.T(), res.Links, 1)
	assert.Equal(this.T(), res.Links[0].Original, link.Original)
}

func (this *LinkServiceTestSuite) TestGetLinkById() {
	// No links, should fail
	_, err := this.service.GetLinkById("123")
	assert.Error(this.T(), err)
	assert.EqualError(this.T(), err, "link does not exist")

	// Insert a link
	createLinkReq := api.CreateLinkRequest{Original: link.Original}
	createLinkRes, err := this.service.CreateLink(createLinkReq)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), createLinkRes.Link.Original, link.Original)

	// Get link by id
	res, err := this.service.GetLinkById(createLinkRes.Link.Id)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), res.Link.Original, link.Original)
}

func (this *LinkServiceTestSuite) TestUpdateLink() {
	// Create link
	createLinkRes, err := this.service.CreateLink(api.CreateLinkRequest{Original: link.Original})
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), createLinkRes.Link.Original, link.Original)

	// Update link
	updOriginal := "http://www.facebook.com"
	res, err := this.service.UpdateLink(api.UpdateLinkRequest{Original: updOriginal}, createLinkRes.Link.Id)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), res.Link.Original, updOriginal)

	// Update wrong id, should fail
	res, err = this.service.UpdateLink(api.UpdateLinkRequest{Original: updOriginal}, "lala")
	assert.Error(this.T(), err)
	assert.EqualError(this.T(), err, "link does not exist")
}

func (this *LinkServiceTestSuite) TestDeleteLink() {
	// Create link
	createLinkRes, err := this.service.CreateLink(api.CreateLinkRequest{Original: link.Original})
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), createLinkRes.Link.Original, link.Original)

	// Delete wrong id, should fail
	_, err = this.service.DeleteLink("123")
	assert.Error(this.T(), err)
	assert.EqualError(this.T(), err, "link does not exist")

	// Delete created link
	res, err := this.service.DeleteLink(createLinkRes.Link.Id)
	assert.NoError(this.T(), err)
	assert.Equal(this.T(), res.Link.Id, createLinkRes.Link.Id)
}
