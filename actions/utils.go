package actions

import (
	"encoding/json"
	"fmt"
)

func unmarshallCustomType[T comparable](data []byte, result *T, labelToValueMap map[string]T, unsupportedValue T) error {
	var label string
	err := json.Unmarshal(data, &label)
	if err != nil {
		return fmt.Errorf("Error unmarshalling custom type: %w", err)
	}

	value, ok := labelToValueMap[label]
	if !ok {
		fmt.Printf("Error unmarshalling %s, type: %T\n", label, result)
		value = unsupportedValue
	}
	result = &value
	return nil
}
