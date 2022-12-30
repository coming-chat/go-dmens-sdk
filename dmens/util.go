package dmens

import "encoding/json"

func JsonString(o interface{}) (string, error) {
	data, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
