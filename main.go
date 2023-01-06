package main

import (
	"github.com/andikabahari/kissa/internal/cluster"
	"github.com/andikabahari/kissa/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	cluster.InitClient()
	r := chi.NewRouter()
	server.Route(r)
	server.Run(r)
}
