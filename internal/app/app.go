package app

import (
	"context"
	"html/template"
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
		e.File("/", "public/index.html")
		e.Static("/css", "static/css")
		e.Static("/js", "static/js")
		e.Static("/img", "static/img")
		e.Renderer = &helpers.Template{Templates: template.Must(template.ParseGlob("public/views/*.html"))}
		e.Validator = &helpers.RequestValidator{Validator: validator.New()}

		// link api routes
		linkRepo := services.NewLinkRepository()
		linkService := services.NewLinkService(linkRepo)
		linkController := controllers.NewLinkController(linkService)
		g := e.Group("/links")
		g.GET("", linkController.GetAllLinks)
		g.GET("/:id", linkController.GetLinkById)
		g.POST("", linkController.CreateLink)
		g.PUT("/:id", linkController.UpdateLink)
		g.DELETE("/:id", linkController.DeleteLink)

		// link htmx routes
		linkHtmxController := controllers.NewLinkHtmxController()
		g2 := e.Group("/htmx")
		g2.GET("/links", linkHtmxController.GetAllLinks)

		// init data
		dbInitializer := helpers.NewDbInitializer(linkRepo)
		dbInitializer.InitData()

		app = &App{
			e: e,
		}
	})

	return app
}

func (app *App) Start() {
	app.e.Logger.Fatal(app.e.Start(":3001"))
}

func (app *App) Stop() {
	err := app.e.Shutdown(context.Background())
	if err != nil {
		app.e.Logger.Error("error shutting down app")
	}
}
