package util

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/andikabahari/kissa/constants"
	"github.com/andikabahari/kissa/dto"
	"github.com/andikabahari/kissa/knative"
	"k8s.io/client-go/rest"
)

func MapK8sResult(result rest.Result) (map[string]interface{}, error) {
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

func ServiceBuf(request dto.ServiceRequest) (*bytes.Buffer, error) {
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
