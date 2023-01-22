package knative

type ObjectEnv struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ServiceObject struct {
	Name          string      `json:"name"`
	Image         string      `json:"image"`
	ContainerPort int         `json:"container_port"`
	Env           []ObjectEnv `json:"env"`
}
