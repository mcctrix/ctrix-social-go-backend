package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/mcctrix/ctrix-social-go-backend/models"
)

func MergeStructs(s1, s2 interface{}) (map[string]interface{}, error) {
	merged := make(map[string]interface{})

	// Helper to process a single struct's fields
	processStruct := func(val reflect.Value) error {
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		if val.Kind() != reflect.Struct {
			return fmt.Errorf("input is not a struct or pointer to struct: %v", val.Kind())
		}
		typ := val.Type()

		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			value := val.Field(i)

			// Skip unexported fields (first letter is lowercase)
			if !field.IsExported() {
				continue
			}

			// Get JSON tag name, handle omitempty and "-" (ignore)
			jsonTag := field.Tag.Get("json")
			if jsonTag == "-" { // Field explicitly ignored for JSON
				continue
			}

			fieldName := field.Name
			if jsonTag != "" {
				parts := strings.Split(jsonTag, ",")
				if parts[0] != "" { // If the name part of the tag isn't empty
					fieldName = parts[0]
				}
			}

			if value.Kind() == reflect.Slice && value.Len() == 0 {
				continue
			}

			// Only add non-zero values.
			// IsZero() is a simpler way to check for zero value in Go 1.13+
			if value.IsValid() && !value.IsZero() {
				// Handle StringArray: ensure it's converted to []string before map insertion
				// so it serializes correctly as a JSON array.
				if arr, ok := value.Interface().(models.StringArray); ok {
					merged[fieldName] = []string(arr)
				} else {
					merged[fieldName] = value.Interface()
				}
			}
		}
		return nil
	}

	// Process first struct
	if err := processStruct(reflect.ValueOf(s1)); err != nil {
		return nil, fmt.Errorf("error processing first struct: %w", err)
	}

	// Process second struct, overwriting existing fields or adding new ones
	if err := processStruct(reflect.ValueOf(s2)); err != nil {
		return nil, fmt.Errorf("error processing second struct: %w", err)
	}

	return merged, nil
}

/*Make sure to pass a pointer to the struct you want to clear*/
func ClearStruct(formatStruct interface{}, rawBody []byte) ([]byte, error) {

	// 1. Validate that formatStruct is a pointer to a struct
	val := reflect.ValueOf(formatStruct)
	if val.Kind() != reflect.Ptr || val.IsNil() || val.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("ClearStruct: formatStruct must be a non-nil pointer to a struct")
	}

	err := json.Unmarshal(rawBody, &formatStruct)
	if err != nil {
		fmt.Println("Error Unmarshalling the data: ", err)
		return nil, err
	}
	rawForm, err := json.Marshal(formatStruct)
	if err != nil {
		fmt.Println("Error Marshalling the data: ", err)
		return nil, err
	}
	return rawForm, nil
}
