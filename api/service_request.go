package api

import (
	"bytes"
	"text/template"
)

type serviceRequest struct {
	Name          string `json:"name"`
	Image         string `json:"image"`
	ContainerPort int    `json:"container_port"`
	Env           []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"env"`
	AutoscalingMetric string `json:"autoscaling_metric"`
	AutoscalingTarget int    `json:"autoscaling_target"`
	MaxScale          int    `json:"max_scale"`
	MinScale          int    `json:"min_scale"`
}

const serviceTemplate = `{
	"apiVersion": "serving.knative.dev/v1",
	"kind": "Service",
	"metadata": {
		"name": "{{.Name}}"
	},
	"spec": {
		"template": {
			"metadata": {
				"annotations": {
					"autoscaling.knative.dev/metric": "{{.AutoscalingMetric}}",
					"autoscaling.knative.dev/target": "{{.AutoscalingTarget}}",
					"autoscaling.knative.dev/min-scale": "{{.MinScale}}",
					"autoscaling.knative.dev/max-scale": "{{.MaxScale}}"
				}
			},
			"spec": {
				"containers": [
					{
						"image": "{{.Image}}",
						"ports": [
							{
								"containerPort": {{.ContainerPort}}
							}
						]
						{{if .Env}}
						,
						"env": [
							{{$env := .Env}}
							{{range $idx, $elem := .Env}}
							{{if $idx}},{{end}}
							{"name":"{{$elem.Name}}","value":"{{$elem.Value}}"}
							{{end}}
						]
						{{end}}
					}
				]
			}
		}
	}
}`

func serviceBuf(request serviceRequest) (*bytes.Buffer, error) {
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, request); err != nil {
		return nil, err
	}

	return buf, nil
}
