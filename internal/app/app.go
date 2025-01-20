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
		e.File("/", "public/views/index.html")
		e.Static("/css", "static/css")
		e.Static("/js", "static/js")
		e.Static("/img", "static/img")
		e.Renderer = &helpers.Template{Templates: template.Must(template.ParseGlob("public/views/*.html"))}
		e.Validator = &helpers.RequestValidator{Validator: validator.New()}

		// link api routes
		linkRepo := services.NewLinkRepository()
		linkService := services.NewLinkService(linkRepo)
		linkController := controllers.NewLinkController(linkService)
		gLinks := e.Group("/links")
		gLinks.GET("", linkController.GetAllLinks)
		gLinks.GET("/:id", linkController.GetLinkById)
		gLinks.POST("", linkController.CreateLink)
		gLinks.PUT("/:id", linkController.UpdateLink)
		gLinks.DELETE("/:id", linkController.DeleteLink)

		// htmx routes
		gHtmx := e.Group("/htmx")

		// htmx link routes
		gHtmxLinks := gHtmx.Group("/links")
		linkHtmxController := controllers.NewLinkHtmxController()
		gHtmxLinks.GET("", linkHtmxController.GetAllLinks)
		gHtmxLinks.POST("/create_link_click", linkHtmxController.CreateLinkClick)
		gHtmxLinks.POST("/create_link_confirm_click", linkHtmxController.CreateLinkClickConfirmClick)
		gHtmxLinks.POST("/:id/delete_link_click", linkHtmxController.DeleteLinkClick)
		gHtmxLinks.POST("/:id/delete_link_confirm_click", linkHtmxController.DeleteLinkConfirmClick)
		gHtmxLinks.POST("/:id/edit_link_click", linkHtmxController.EditLinkClick)
		gHtmxLinks.POST("/:id/edit_link_confirm_click", linkHtmxController.EditLinkConfirmClick)

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
