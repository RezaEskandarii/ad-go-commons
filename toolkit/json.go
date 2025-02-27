package toolkit

import (
	"bytes"
	"encoding/json"
)

func ToJSON(data interface{}) (string, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
