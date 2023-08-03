package knative

import "time"

type ServiceItem struct {
	ApiVersion string
	Kind       string
	Metadata   struct {
		Annotations       map[string]interface{}
		CreationTimestamp time.Time
		Generation        int
		ManagedFields     []map[string]interface{}
		Name              string
		Namespace         string
		ResourceVersion   string
		Uid               string
	}
	Spec struct {
		Template map[string]interface{}
		Traffic  []map[string]interface{}
	}
	Status struct {
		Address                   map[string]interface{}
		Conditions                []map[string]interface{}
		LatestCreatedRevisionName string
		LatestReadyRevisionName   string
		ObservedGeneration        int
		Traffic                   []map[string]interface{}
		Url                       string
	}
}

type ServiceList struct {
	ApiVersion string
	Items      []ServiceItem
	Kind       string
	Metadata   map[string]interface{}
}
