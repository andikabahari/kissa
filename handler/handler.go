package handler

import (
	"github.com/andikabahari/kissa/knative"
)

type Handler struct {
	knative knative.Knative
}

func New(kn knative.Knative) *Handler {
	return &Handler{
		knative: kn,
	}
}
