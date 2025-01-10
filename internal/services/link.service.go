package services

import (
	"errors"
	"fmt"
	"sync"

	"github.com/LeonYalin/golang-todo-list-app/api"
	"github.com/LeonYalin/golang-todo-list-app/internal/models"
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
)

var lock = sync.Mutex{}
var links = map[string]*models.Link{}

type ILinkService interface {
	Get(request api.GetAllLinksRequest) (api.GetAllLinksResponse, error)
	Create(request api.CreateLinkRequest) (api.CreateLinkResponse, error)
	Update(request api.UpdateLinkRequest, id string) (api.UpdateLinkResponse, error)
	GetById(id string) (api.GetLinkByIdResponse, error)
	Delete(id string) (api.DeleteLinkResponse, error)
}

type LinkService struct {
}

func NewLinkService() *LinkService {
	return &LinkService{}
}

func (this *LinkService) Get(request api.GetAllLinksRequest) (api.GetAllLinksResponse, error) {
	lock.Lock()
	defer lock.Unlock()
	linksToSend := make([]models.Link, 0)
	for _, v := range links {
		linksToSend = append(linksToSend, *v)
	}
	response := api.GetAllLinksResponse{Links: linksToSend}
	return response, nil
}

func (this *LinkService) Create(request api.CreateLinkRequest) (api.CreateLinkResponse, error) {
	lock.Lock()
	defer lock.Unlock()
	url := request.Original
	if _, exists := links[url]; exists {
		return api.CreateLinkResponse{}, errors.New("link already exists")
	}
	link := &models.Link{Id: uuid.NewString(), Original: url, Short: fmt.Sprintf("/l/%s", shortuuid.New())}
	links[link.Id] = link

	return api.CreateLinkResponse{Link: *link}, nil
}

func (this *LinkService) GetById(id string) (api.GetLinkByIdResponse, error) {
	lock.Lock()
	defer lock.Unlock()
	if link, exists := links[id]; !exists {
		return api.GetLinkByIdResponse{}, errors.New("link does not exist")
	} else {
		return api.GetLinkByIdResponse{Link: *link}, nil
	}
}

func (this *LinkService) Update(request api.UpdateLinkRequest, id string) (api.UpdateLinkResponse, error) {
	lock.Lock()
	defer lock.Unlock()
	if _, exists := links[id]; !exists {
		return api.UpdateLinkResponse{}, errors.New("link does not exist")
	} else {
		links[id].Original = request.Original
		return api.UpdateLinkResponse{Link: *links[id]}, nil
	}
}

func (this *LinkService) Delete(id string) (api.DeleteLinkResponse, error) {
	lock.Lock()
	defer lock.Unlock()
	if linkToDelete, exists := links[id]; !exists {
		return api.DeleteLinkResponse{}, errors.New("link does not exist")
	} else {
		delete(links, id)
		return api.DeleteLinkResponse{Link: *linkToDelete}, nil
	}

}
