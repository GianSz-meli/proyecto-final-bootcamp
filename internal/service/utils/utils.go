package utils

import (
	"reflect"
)

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
				targetField.Set(field.Elem())
				updated = true
			}
		}
	}

	return updated
}
