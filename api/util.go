package api

import (
	"encoding/json"

	"k8s.io/client-go/rest"
)

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
