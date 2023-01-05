package handler

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/andikabahari/kissa/constants"
	"github.com/andikabahari/kissa/dto"
	"github.com/andikabahari/kissa/knative"
	"k8s.io/client-go/rest"
)

type ServiceHandler interface {
	List(w http.ResponseWriter, r *http.Request)
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
	resource := "services"

	serviceName := r.URL.Query().Get("service_name")
	if serviceName != "" {
		resource += "/" + serviceName
	}

	result := h.knative.Get(resource)
	writeK8sResponse(w, result)
}

func (h *serviceHandler) Create(w http.ResponseWriter, r *http.Request) {
	request := dto.ServiceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		panic(err)
	}

	tmpl, err := template.New("service").Parse(constants.KnativeServiceTemplate)
	if err != nil {
		panic(err)
	}

	obj := knative.ServiceObject{
		Name:          request.Name,
		Image:         request.Image,
		ContainerPort: request.ContainerPort,
		Env:           request.Env,
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, obj); err != nil {
		panic(err)
	}

	result := h.knative.Create("services", buf.Bytes())
	writeK8sResponse(w, result)
}

func writeK8sResponse(w http.ResponseWriter, result rest.Result) {
	code := 0
	result.StatusCode(&code)
	if code >= 400 {
		dto.JSONResponse(w, code, result.Error().Error(), nil)
		return
	}

	raw, err := result.Raw()
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	if err := json.Unmarshal(raw, &data); err != nil {
		panic(err)
	}

	dto.JSONResponse(w, http.StatusOK, "success", data)
}
