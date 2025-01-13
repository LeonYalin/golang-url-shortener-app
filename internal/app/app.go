package app

import (
	"sync"

	"github.com/LeonYalin/golang-todo-list-app/internal/controllers"
	"github.com/LeonYalin/golang-todo-list-app/internal/helpers"
	"github.com/LeonYalin/golang-todo-list-app/internal/services"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type App struct {
	e *echo.Echo
}

var once sync.Once
var app *App

func NewApp() *App {
	once.Do(func() {
		e := echo.New()
		e.File("/", "static/index.html")
		e.Static("/js", "static/js")
		e.Validator = &helpers.RequestValidator{Validator: validator.New()}

		// link routes
		linkRepo := services.NewLinkRepository()
		linkService := services.NewLinkService(linkRepo)
		linkController := controllers.NewLinkController(linkService)
		g := e.Group("/links")
		g.GET("", linkController.GetAllLinks)
		g.GET("/:id", linkController.GetLinkById)
		g.POST("", linkController.CreateLink)
		g.PUT("/:id", linkController.UpdateLink)
		g.DELETE("/:id", linkController.DeleteLink)

		app = &App{
			e: e,
		}
	})

	return app
}

func (app *App) Start() {
	app.e.Logger.Fatal(app.e.Start(":3001"))
}
