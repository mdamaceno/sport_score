package helpers

import (
	"encoding/json"
	"io"
)

func DecodeRawJson(body io.ReadCloser) (map[string]interface{}, error) {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(body).Decode(&json_map)

	if err != nil {
		return nil, err
	}

	return json_map, nil
}
