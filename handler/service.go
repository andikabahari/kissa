package handler

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/andikabahari/kissa/constants"
	"github.com/andikabahari/kissa/dto"
	"github.com/andikabahari/kissa/errors"
	"github.com/andikabahari/kissa/knative"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
)

type ServiceHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type serviceHandler struct {
	knative knative.Knative
}

func NewServiceHandler(kn knative.Knative) ServiceHandler {
	return &serviceHandler{
		knative: kn,
	}
}

func (h *serviceHandler) List(w http.ResponseWriter, r *http.Request) {
	data, err := h.knative.ListMap("services")
	if err != nil {
		log.Err(err)

		if k8serrors.IsNotFound(err) {
			dto.JSONResponse(w, http.StatusNotFound, err.Error(), nil)
		} else {
			dto.JSONResponse(w, http.StatusInternalServerError, errors.InternalServerError.Error(), nil)
		}

		return
	}

	dto.JSONResponse(w, http.StatusOK, "success", data)
}

func (h *serviceHandler) Get(w http.ResponseWriter, r *http.Request) {
	serviceName := chi.URLParam(r, "serviceName")
	data, err := h.knative.GetMap("services", serviceName)
	if err != nil {
		log.Err(err)

		if k8serrors.IsNotFound(err) {
			dto.JSONResponse(w, http.StatusNotFound, err.Error(), nil)
		} else {
			dto.JSONResponse(w, http.StatusInternalServerError, errors.InternalServerError.Error(), nil)
		}

		return
	}

	dto.JSONResponse(w, http.StatusOK, "success", data)
}

func (h *serviceHandler) Create(w http.ResponseWriter, r *http.Request) {
	request := dto.ServiceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Err(err)
		dto.JSONResponse(w, http.StatusInternalServerError, errors.InternalServerError.Error(), nil)
		return
	}

	tmpl, err := template.New("service").Parse(constants.KnativeServiceTemplate)
	if err != nil {
		log.Err(err)
		dto.JSONResponse(w, http.StatusInternalServerError, errors.InternalServerError.Error(), nil)
		return
	}

	obj := knative.ServiceObject{
		Name:          request.Name,
		Image:         request.Image,
		ContainerPort: request.ContainerPort,
		Env:           request.Env,
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, obj); err != nil {
		log.Err(err)
		dto.JSONResponse(w, http.StatusInternalServerError, errors.InternalServerError.Error(), nil)
		return
	}

	data, err := h.knative.CreateMap("services", buf.Bytes())
	if err != nil {
		log.Print(err)
		dto.JSONResponse(w, http.StatusInternalServerError, errors.InternalServerError.Error(), nil)
		return
	}

	dto.JSONResponse(w, http.StatusOK, "success", data)
}
