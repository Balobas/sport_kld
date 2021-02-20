package utils

import "encoding/json"

func UnmarshalFromMap(obj interface{}, objMap map[string]interface{}) error {
	b, err := json.Marshal(objMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, obj)
}
