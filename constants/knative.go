package constants

const KnativeServiceTemplate = `{
	"apiVersion": "serving.knative.dev/v1",
	"kind": "Service",
	"metadata": {
		"name": "{{.Name}}"
	},
	"spec": {
		"template": {
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
