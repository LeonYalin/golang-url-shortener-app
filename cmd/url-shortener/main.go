package main

import (
	"log"

	"github.com/LeonYalin/golang-todo-list-app/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// run app
	a := app.NewApp()
	a.Start()
	defer a.Stop()
}
