package server

import (
	"github.com/andikabahari/kissa/config"
	"github.com/andikabahari/kissa/handler"
	"github.com/andikabahari/kissa/internal/cluster"
	"github.com/andikabahari/kissa/knative"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	kn := knative.NewKnative(cluster.Client, config.Get().ClusterNamespace)

	api := e.Group("/api")

	serviceHandler := handler.NewServiceHandler(kn)
	services := api.Group("/services")
	services.GET("", serviceHandler.List)
	services.POST("", serviceHandler.Create)
	services.GET("/:service_name", serviceHandler.Get)
	services.PATCH("/:service_name", serviceHandler.Update)
	services.DELETE("/:service_name", serviceHandler.Delete)

	revisionHandler := handler.NewRevisionHandler(kn)
	revisions := api.Group("/revisions")
	revisions.GET("", revisionHandler.List)
	revisions.GET("/:revision_name", revisionHandler.Get)
	revisions.DELETE("/:revision_name", revisionHandler.Delete)
}
