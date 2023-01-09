package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/andikabahari/kissa/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Run(e *echo.Echo) {
	config := config.Get()

	e.Logger.SetLevel(log.INFO)

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", config.HTTPPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
