package app

import (
	"sync"

	"github.com/LeonYalin/golang-todo-list-app/internal/controllers"
	"github.com/LeonYalin/golang-todo-list-app/internal/services"
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

		// link routes
		linkService := services.NewLinkService()
		linkController := controllers.NewLinkController(linkService)
		g := e.Group("/links")
		g.GET("", linkController.GetAll)
		g.GET("/:id", linkController.GetById)
		g.POST("", linkController.Create)
		g.PUT("/:id", linkController.Update)

		app = &App{
			e: e,
		}
	})

	return app
}

func (app *App) Start() {
	app.e.Logger.Fatal(app.e.Start(":3001"))
}
