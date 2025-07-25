package utils

import (
	"reflect"
)

// UpdateFields compares the source struct (request body) with the target struct (the existing record).
// For each non-nil pointer field in the source, the corresponding field in the target is updated in memory.
// Returns true if any fields are updated, false otherwise.
func UpdateFields(target interface{}, source interface{}) bool {
	targetReflect := reflect.ValueOf(target)
	sourceReflect := reflect.ValueOf(source)

	if targetReflect.Kind() != reflect.Ptr || targetReflect.Elem().Kind() != reflect.Struct {
		return false
	}
	if sourceReflect.Kind() != reflect.Ptr || sourceReflect.Elem().Kind() != reflect.Struct {
		return false
	}

	targetValue := targetReflect.Elem()
	sourceValue := sourceReflect.Elem()
	sourceType := sourceValue.Type()

	updated := false

	for i := 0; i < sourceValue.NumField(); i++ {
		field := sourceValue.Field(i)
		fieldName := sourceType.Field(i).Name
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			targetField := targetValue.FieldByName(fieldName)
			if targetField.IsValid() && targetField.CanSet() {
				//When target is also a pointer value assign the pointer value
				if targetField.Kind() == reflect.Ptr {
					targetField.Set(field)
					updated = true
				} else if targetField.Type() == field.Elem().Type() {
					// When target is a value, but source is pointer
					targetField.Set(field.Elem())
					updated = true
				}
			}
		}
	}

	return updated
}
