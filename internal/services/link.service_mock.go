package services

import "github.com/LeonYalin/golang-todo-list-app/api"

type MockLinkService struct {
}

func (this *MockLinkService) GetAllLinks(request api.GetAllLinksRequest) (api.GetAllLinksResponse, error) {
	return api.GetAllLinksResponse{}, nil
}
func (this *MockLinkService) CreateLink(request api.CreateLinkRequest) (api.CreateLinkResponse, error) {
	return api.CreateLinkResponse{}, nil
}
func (this *MockLinkService) UpdateLink(request api.UpdateLinkRequest, id string) (api.UpdateLinkResponse, error) {
	return api.UpdateLinkResponse{}, nil
}
func (this *MockLinkService) GetLinkById(id string) (api.GetLinkByIdResponse, error) {
	return api.GetLinkByIdResponse{}, nil
}
func (this *MockLinkService) DeleteLink(id string) (api.DeleteLinkResponse, error) {
	return api.DeleteLinkResponse{}, nil
}
