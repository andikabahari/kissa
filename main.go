package main

import (
	"github.com/andikabahari/kissa/internal/cluster"
	"github.com/andikabahari/kissa/server"
	"github.com/labstack/echo/v4"
)

func main() {
	cluster.InitClient()
	e := echo.New()
	server.Route(e)
	server.Run(e)
}
