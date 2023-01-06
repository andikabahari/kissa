package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/andikabahari/kissa/constants"
	"github.com/andikabahari/kissa/dto"
	"github.com/andikabahari/kissa/knative"
	"github.com/go-chi/chi/v5"
	"k8s.io/client-go/rest"
)

type ServiceHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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

	buf, err := serviceBuf(request)
	if err != nil {
		panic(err)
	}

	result := h.knative.Create("services", buf.Bytes())
	writeK8sResponse(w, result)
}

func (h *serviceHandler) Update(w http.ResponseWriter, r *http.Request) {
	request := dto.ServiceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		panic(err)
	}

	serviceName := chi.URLParam(r, "serviceName")
	request.Name = serviceName

	buf, err := serviceBuf(request)
	if err != nil {
		panic(err)
	}

	result := h.knative.Update("services/"+serviceName, buf.Bytes())
	writeK8sResponse(w, result)
}

func (h *serviceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	serviceName := chi.URLParam(r, "serviceName")
	result := h.knative.Delete("services/" + serviceName)
	writeK8sResponse(w, result)
}

func writeK8sResponse(w http.ResponseWriter, result rest.Result) {
	code := 0
	result.StatusCode(&code)
	if code >= 400 {
		dto.JSONResponse(w, code, result.Error().Error(), nil)
		return
	}

	data, err := mapK8sResult(result)
	if err != nil {
		panic(err)
	}

	dto.JSONResponse(w, http.StatusOK, "success", data)
}

func mapK8sResult(result rest.Result) (map[string]interface{}, error) {
	raw, err := result.Raw()
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func serviceBuf(request dto.ServiceRequest) (*bytes.Buffer, error) {
	tmpl, err := template.New("service").Parse(constants.KnativeServiceTemplate)
	if err != nil {
		return nil, err
	}

	obj := knative.ServiceObject{
		Name:          request.Name,
		Image:         request.Image,
		ContainerPort: request.ContainerPort,
		Env:           request.Env,
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}
