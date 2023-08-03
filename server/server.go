package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andikabahari/kissa/api"
	"github.com/andikabahari/kissa/knative"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type server struct {
	echo    *echo.Echo
	handler *api.Handler
}

func New(kn knative.Knative) *server {
	srv := &server{
		echo:    echo.New(),
		handler: api.NewHandler(kn),
	}

	srv.setupMiddlewares()
	srv.routes()

	return srv
}

func (s *server) Run(port int) {
	s.echo.Logger.SetLevel(log.INFO)

	go func() {
		if err := s.echo.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
			s.echo.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.echo.Shutdown(ctx); err != nil {
		s.echo.Logger.Fatal(err)
	}
}

func (s *server) setupMiddlewares() {
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.CORS())
}

func (s *server) routes() {
	api := s.echo.Group("/api")

	services := api.Group("/services")
	services.GET("", s.handler.ListService)
	services.POST("", s.handler.CreateService)
	services.GET("/:service_name", s.handler.GetService)
	services.PATCH("/:service_name", s.handler.UpdateService)
	services.DELETE("/:service_name", s.handler.DeleteService)

	revisions := api.Group("/revisions")
	revisions.GET("", s.handler.ListRevision)
	revisions.GET("/:revision_name", s.handler.GetRevision)
	revisions.DELETE("/:revision_name", s.handler.DeleteRevision)
}
