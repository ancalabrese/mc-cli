package actions

import (
	"encoding/json"
	"fmt"
)

func unmarshallCustomType(data []byte, result any, labelToValueMap any) error {
	// Type Assertion
	mapInterface, ok := labelToValueMap.(map[string]interface{})
	if !ok {
		return fmt.Errorf("labelToValueMap must be of type map[string]interface{}")
	}
	var label string
	err := json.Unmarshal(data, &label)
	if err != nil {
		return fmt.Errorf("Error unmarshalling custom type: %w", err)
	}

	value, ok := mapInterface[label]
	if !ok {
		return fmt.Errorf("Error unmarshalling %s", label)
	}

	// Ensure 'result' is a pointer and assign directly
	if ptr, ok := result.(*any); ok {
		*ptr = value
	} else {
		return fmt.Errorf("Error unmarshalling custom type: result must be a pointer")
	}
	return nil
}
