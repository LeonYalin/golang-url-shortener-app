package helpers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/LeonYalin/golang-todo-list-app/internal/models"
	"github.com/LeonYalin/golang-todo-list-app/internal/services"
	"github.com/google/uuid"
)

type InitDataConfig struct {
	Links []models.Link `json:"links"`
}

type DbInitializer struct {
	repo services.ILinkRepository
}

func NewDbInitializer(repo services.ILinkRepository) *DbInitializer {
	return &DbInitializer{repo: repo}
}

func (this *DbInitializer) InitData() {
	initData := os.Getenv("INIT_DATA")
	if initData != "" {
		// open file
		file, err := os.Open(initData)
		if err != nil {
			fmt.Println("InitData: error opening file")
		}
		defer file.Close()

		// read file contents
		var config InitDataConfig
		if err := json.NewDecoder(file).Decode(&config); err != nil {
			fmt.Println(fmt.Errorf("InitData: error reading file contents, %v", err))
		}

		// store data to db
		for _, link := range config.Links {
			_, err := this.repo.Create(uuid.NewString(), link.Original, link.Short)
			if err != nil {
				fmt.Println(fmt.Errorf("error storing data to db, %v", err))
			}
		}
	} else {
		fmt.Println(fmt.Errorf("InitData: no file found"))
	}
}
