package services

import (
	"fmt"

	"github.com/LeonYalin/golang-todo-list-app/api"
	"github.com/LeonYalin/golang-todo-list-app/internal/models"
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
)

type ILinkService interface {
	GetAllLinks(request api.GetAllLinksRequest) (api.GetAllLinksResponse, error)
	CreateLink(request api.CreateLinkRequest) (api.CreateLinkResponse, error)
	UpdateLink(request api.UpdateLinkRequest, id string) (api.UpdateLinkResponse, error)
	GetLinkById(id string) (api.GetLinkByIdResponse, error)
	DeleteLink(id string) (api.DeleteLinkResponse, error)
}

type LinkService struct {
	repo ILinkRepository
}

func NewLinkService(repo ILinkRepository) *LinkService {
	return &LinkService{repo: repo}
}

func (this *LinkService) GetAllLinks(request api.GetAllLinksRequest) (api.GetAllLinksResponse, error) {
	links, err := this.repo.GetAll()
	if err != nil {
		return api.GetAllLinksResponse{}, err
	}
	response := api.GetAllLinksResponse{Links: make([]models.Link, 0)}
	for _, v := range links {
		response.Links = append(response.Links, *v)
	}
	return response, nil
}

func (this *LinkService) CreateLink(request api.CreateLinkRequest) (api.CreateLinkResponse, error) {
	createdLink, err := this.repo.Create(uuid.NewString(), request.Original, fmt.Sprintf("/l/%s", shortuuid.New()))
	if err != nil {
		return api.CreateLinkResponse{}, err
	}
	return api.CreateLinkResponse{Link: *createdLink}, nil
}

func (this *LinkService) GetLinkById(id string) (api.GetLinkByIdResponse, error) {
	link, err := this.repo.GetById(id)
	if err != nil {
		return api.GetLinkByIdResponse{}, err
	}
	return api.GetLinkByIdResponse{Link: *link}, nil
}

func (this *LinkService) UpdateLink(request api.UpdateLinkRequest, id string) (api.UpdateLinkResponse, error) {
	link, err := this.repo.Update(id, request.Original)
	if err != nil {
		return api.UpdateLinkResponse{}, err
	}
	return api.UpdateLinkResponse{Link: *link}, nil
}

func (this *LinkService) DeleteLink(id string) (api.DeleteLinkResponse, error) {
	deletedLink, err := this.repo.Delete(id)
	if err != nil {
		return api.DeleteLinkResponse{}, err
	}
	return api.DeleteLinkResponse{Link: *deletedLink}, nil

}
