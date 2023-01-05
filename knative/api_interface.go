package knative

type KnativeAPI interface {
	List(resource string) ([]byte, error)
	ListMap(resource string) (map[string]interface{}, error)
	Get(resource, name string) ([]byte, error)
	GetMap(resource, name string) (map[string]interface{}, error)
	Create(resource string, obj interface{}) ([]byte, error)
	CreateMap(resource string, obj interface{}) (map[string]interface{}, error)
}
