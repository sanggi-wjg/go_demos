package buildpkg

import (
	"encoding/json"
	"errors"
)

// export Add
func Add(a, b int) int {
	return a + b
}

// export JsonDecode
func JsonDecode(target []byte) (map[any]interface{}, error) {
	var data map[any]interface{}
	if err := json.Unmarshal(target, &data); err != nil {
		return nil, errors.New("json decode failed")
	}
	return data, nil
}
