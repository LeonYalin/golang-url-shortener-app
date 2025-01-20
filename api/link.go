package api

import (
	"github.com/LeonYalin/golang-todo-list-app/internal/models"
)

// Get all links
type GetAllLinksRequest struct {
	Page     int `query:"page" validate:"number,gte=0"`     // page number
	PageSize int `query:"pageSize" validate:"number,gte=0"` // results per page
}
type GetAllLinksResponse struct {
	Links    []models.Link `json:"links"`    // links results array
	Page     int           `json:"page"`     // page number
	PageSize int           `json:"pageSize"` // results per page
	Total    int           `json:"total"`    // total results
}

// Get link by Short
type GetLinkByShortResponse struct {
	Link models.Link `json:"link"` // requested link
}

// Get link by ID
type GetLinkByIdResponse struct {
	Link models.Link `json:"link"` // requested link
}

// Create link
type CreateLinkRequest struct {
	Original string `json:"original" validate:"required,url"` // original url of the link
}
type CreateLinkResponse struct {
	Link models.Link `json:"link"` // created link
}

// Update link
type UpdateLinkRequest struct {
	Original string `json:"original" validate:"required,url"` // original url of the link
}
type UpdateLinkResponse struct {
	Link models.Link `json:"link"` // created link
}

// Delete link
type DeleteLinkResponse struct {
	Link models.Link `json:"link"` // deleted link
}
