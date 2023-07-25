package helper

import (
	"encoding/json"
)

func BuildQueryString(query map[string]interface{}) (string, error) {
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return "", err
	}

	return string(queryJSON), nil
}
