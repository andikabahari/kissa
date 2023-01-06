package server

import (
	"net/http"

	"github.com/andikabahari/kissa/cluster"
	"github.com/andikabahari/kissa/config"
	"github.com/andikabahari/kissa/dto"
	"github.com/andikabahari/kissa/handler"
	"github.com/andikabahari/kissa/knative"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
)

func Route(r *chi.Mux) {
	r.Use(cors.AllowAll().Handler)
	r.Use(customRecoverer)

	kn := knative.NewKnative(cluster.Client, config.Get().ClusterNamespace)

	serviceHandler := handler.NewServiceHandler(kn)
	r.Get("/api/services", serviceHandler.List)
	r.Post("/api/services", serviceHandler.Create)
	r.Put("/api/services/{serviceName}", serviceHandler.Update)
	r.Delete("/api/services/{serviceName}", serviceHandler.Delete)
}

func customRecoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error().Msgf("%v", err)

				// buf := make([]byte, 2048)
				// n := runtime.Stack(buf, false)
				// buf = buf[:n]
				// fmt.Print(string(buf))

				dto.JSONResponse(w, http.StatusInternalServerError, "internal server error", nil)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
