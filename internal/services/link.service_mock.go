package services

import "github.com/LeonYalin/golang-todo-list-app/api"

type MockLinkService struct {
	GetAllLinksFunc    func(request api.GetAllLinksRequest) (api.GetAllLinksResponse, error)
	CreateLinkFunc     func(request api.CreateLinkRequest) (api.CreateLinkResponse, error)
	UpdateLinkFunc     func(request api.UpdateLinkRequest, id string) (api.UpdateLinkResponse, error)
	GetLinkByIdFunc    func(id string) (api.GetLinkByIdResponse, error)
	GetLinkByShortFunc func(id string) (api.GetLinkByShortResponse, error)
	DeleteLinkFunc     func(id string) (api.DeleteLinkResponse, error)
}

func NewMockLinkService() *MockLinkService {
	return &MockLinkService{}
}

func (this *MockLinkService) GetAllLinks(request api.GetAllLinksRequest) (api.GetAllLinksResponse, error) {
	return this.GetAllLinksFunc(request)
}
func (this *MockLinkService) CreateLink(request api.CreateLinkRequest) (api.CreateLinkResponse, error) {
	return this.CreateLinkFunc(request)
}
func (this *MockLinkService) UpdateLink(request api.UpdateLinkRequest, id string) (api.UpdateLinkResponse, error) {
	return this.UpdateLinkFunc(request, id)
}
func (this *MockLinkService) GetLinkById(id string) (api.GetLinkByIdResponse, error) {
	return this.GetLinkByIdFunc(id)
}
func (this *MockLinkService) GetLinkByShort(id string) (api.GetLinkByShortResponse, error) {
	return this.GetLinkByShortFunc(id)
}
func (this *MockLinkService) DeleteLink(id string) (api.DeleteLinkResponse, error) {
	return this.DeleteLinkFunc(id)
}
