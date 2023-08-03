package api

import (
	"github.com/andikabahari/kissa/knative"
)

type Handler struct {
	knative knative.Knative
}

func NewHandler(kn knative.Knative) *Handler {
	return &Handler{
		knative: kn,
	}
}
