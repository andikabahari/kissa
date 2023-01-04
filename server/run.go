package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andikabahari/kissa/config"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func Run(r *chi.Mux) {
	config := config.Get()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.HTTPPort),
		Handler: r,
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal().Msg("graceful shutdown timed out...forcing exit.")
			}
		}()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Error().Err(err)
		}
		serverStopCtx()
	}()

	log.Info().Msgf("server running on port %d", config.HTTPPort)

	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Error().Err(err)
	}

	<-serverCtx.Done()
	log.Print("server stopped")
}
