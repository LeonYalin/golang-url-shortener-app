package services

import (
	"errors"
	"sync"

	"github.com/LeonYalin/golang-todo-list-app/internal/models"
)

type ILinkRepository interface {
	GetAll() (map[string]*models.Link, error)
	Create(id string, original string, short string) (*models.Link, error)
	Update(id string, original string) (*models.Link, error)
	GetById(id string) (*models.Link, error)
	Delete(id string) (*models.Link, error)
}

type LinkRepository struct {
	lock  sync.Mutex
	links map[string]*models.Link
}

func NewLinkRepository() *LinkRepository {
	return &LinkRepository{
		lock:  sync.Mutex{},
		links: make(map[string]*models.Link),
	}
}

func (this *LinkRepository) GetAll() (map[string]*models.Link, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.links, nil
}

func (this *LinkRepository) Create(id string, original string, short string) (*models.Link, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, exists := this.links[id]; exists {
		return nil, errors.New("link already exists")
	}
	link := &models.Link{Id: id, Original: original, Short: short}
	this.links[link.Id] = link
	return this.links[link.Id], nil
}

func (this *LinkRepository) GetById(id string) (*models.Link, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if link, exists := this.links[id]; !exists {
		return nil, errors.New("link does not exist")
	} else {
		return link, nil
	}
}

func (this *LinkRepository) Update(id string, original string) (*models.Link, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, exists := this.links[id]; !exists {
		return nil, errors.New("link does not exist")
	}
	this.links[id].Original = original
	return this.links[id], nil
}

func (this *LinkRepository) Delete(id string) (*models.Link, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if linkToDelete, exists := this.links[id]; !exists {
		return nil, errors.New("link does not exist")
	} else {
		delete(this.links, id)
		return linkToDelete, nil
	}
}
