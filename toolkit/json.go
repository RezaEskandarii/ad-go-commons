package toolkit

import (
	"bytes"
	"encoding/json"
)

func ToJSON(data interface{}) string {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return ""
	}

	return buf.String()
}

func JSONEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(b[1 : len(b)-1])
}
