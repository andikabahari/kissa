package server

import (
	"github.com/andikabahari/kissa/cluster"
	"github.com/andikabahari/kissa/config"
	"github.com/andikabahari/kissa/handler"
	"github.com/andikabahari/kissa/knative"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Route(r *chi.Mux) {
	r.Use(cors.AllowAll().Handler)

	kn := knative.NewKnative(cluster.Client, config.Get().ClusterNamespace)

	serviceHandler := handler.NewServiceHandler(kn)
	r.Get("/api/services", serviceHandler.List)
	r.Post("/api/services", serviceHandler.Create)
	r.Get("/api/services/{serviceName}", serviceHandler.Get)
}
