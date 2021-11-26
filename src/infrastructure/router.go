package infrastructure

import (
	"net/http"

	"github.com/dj-hirrot/ca_with_go/src/interface/controller"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Clean Architecture with Go
// @version 1.0
// @description This is a Go API server template.

// @contact.name Hiroto Shibutani
// @contact.url https://github.com/dj-hirrot
// @contact.email hirotoshibutani@gmail.com

// @host localhost:8080
// @BasePath /v1
func Init() {
	e := echo.New()

	sqlHandler := NewSqlHandler()
	userController := controller.NewUserController(sqlHandler)

	e.GET("/", healthCheck)

	v1 := e.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", func(c echo.Context) error { return userController.Index(c) })
			users.GET("/:id", func(c echo.Context) error { return userController.Show(c) })
			users.POST("", func(c echo.Context) error { return userController.Create(c) })
			users.PUT("/:id", func(c echo.Context) error { return userController.Save(c) })
			users.DELETE("/:id", func(c echo.Context) error { return userController.Delete(c) })
		}
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "**Running**")
}
