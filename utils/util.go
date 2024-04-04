package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func ConvertMapToStringMap(originalMap map[string]interface{}) (map[string]string, error) {
	convertedMap := make(map[string]string)
	for key, value := range originalMap {
		switch v := value.(type) {
		case string:
			convertedMap[key] = v
		case int, int32, int64, float32, float64:
			convertedMap[key] = fmt.Sprintf("%v", v)
		case bool:
			convertedMap[key] = strconv.FormatBool(v)
		default:
			jsonValue, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("error marshaling value for key '%s': %w", key, err)
			}
			convertedMap[key] = string(jsonValue)
		}
	}
	return convertedMap, nil
}
